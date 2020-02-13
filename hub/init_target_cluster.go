package hub

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/greenplum-db/gp-common-go-libs/dbconn"
	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"golang.org/x/xerrors"

	"github.com/greenplum-db/gpupgrade/db"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/step"
	"github.com/greenplum-db/gpupgrade/utils"
	"github.com/greenplum-db/gpupgrade/utils/cluster"
)

func (s *Server) GenerateInitsystemConfig(ports []uint32) (int, error) {
	sourceDBConn := db.NewDBConn("localhost", int(s.Source.MasterPort()), "template1")
	return s.writeConf(sourceDBConn, ports)
}

func (s *Server) initsystemConfPath() string {
	return filepath.Join(s.StateDir, "gpinitsystem_config")
}

func (s *Server) writeConf(sourceDBConn *dbconn.DBConn, ports []uint32) (int, error) {
	err := sourceDBConn.Connect(1)
	if err != nil {
		return 0, errors.Wrap(err, "could not connect to database")
	}
	defer sourceDBConn.Close()

	gpinitsystemConfig, err := CreateInitialInitsystemConfig(s.Source.MasterDataDir())
	if err != nil {
		return 0, err
	}

	gpinitsystemConfig, err = GetCheckpointSegmentsAndEncoding(gpinitsystemConfig, sourceDBConn)
	if err != nil {
		return 0, err
	}

	gpinitsystemConfig, masterPort, err := WriteSegmentArray(gpinitsystemConfig, s.Source, ports)
	if err != nil {
		return 0, xerrors.Errorf("generating segment array: %w", err)
	}

	return masterPort, WriteInitsystemFile(gpinitsystemConfig, s.initsystemConfPath())
}

func (s *Server) CreateTargetCluster(stream step.OutStreams, masterPort int, targetPorts TargetPorts) error {
	err := s.InitTargetCluster(stream, targetPorts)
	if err != nil {
		return err
	}

	conn := db.NewDBConn("localhost", masterPort, "template1")
	defer conn.Close()

	s.Target, err = utils.ClusterFromDB(conn, s.Target.BinDir)
	if err != nil {
		return errors.Wrap(err, "could not retrieve target configuration")
	}

	if err := s.SaveConfig(); err != nil {
		return err
	}

	return nil
}

func (s *Server) InitTargetCluster(stream step.OutStreams, targetPorts TargetPorts) error {
	agentConns, err := s.AgentConns()
	if err != nil {
		return errors.Wrap(err, "Could not get/create agents")
	}

	err = CreateAllDataDirectories(agentConns, s.Source)
	if err != nil {
		return err
	}

	err = RunInitsystemForTargetCluster(stream, s.Target, s.initsystemConfPath())
	if err != nil {
		return err
	}

	return UpgradeStandby(StandbyConfig{
		GreenplumEnv: &greenplumEnv{
			binDir:              s.Target.BinDir,
			masterPort:          int(targetPorts.MasterPort),
			masterDataDirectory: s.Target.MasterDataDir(),
		},
		Port:          int(targetPorts.StandbyPort),
		Hostname:      s.Source.StandbyHostname(),
		DataDirectory: s.Source.StandbyDataDirectory() + "_upgrade",
	})
}

func GetCheckpointSegmentsAndEncoding(gpinitsystemConfig []string, dbConnector *dbconn.DBConn) ([]string, error) {
	checkpointSegments, err := dbconn.SelectString(dbConnector, "SELECT current_setting('checkpoint_segments') AS string")
	if err != nil {
		return gpinitsystemConfig, errors.Wrap(err, "Could not retrieve checkpoint segments")
	}
	encoding, err := dbconn.SelectString(dbConnector, "SELECT current_setting('server_encoding') AS string")
	if err != nil {
		return gpinitsystemConfig, errors.Wrap(err, "Could not retrieve server encoding")
	}
	gpinitsystemConfig = append(gpinitsystemConfig,
		fmt.Sprintf("CHECK_POINT_SEGMENTS=%s", checkpointSegments),
		fmt.Sprintf("ENCODING=%s", encoding))
	return gpinitsystemConfig, nil
}

func CreateInitialInitsystemConfig(sourceMasterDataDir string) ([]string, error) {
	gpinitsystemConfig := []string{`ARRAY_NAME="gp_upgrade cluster"`}

	segPrefix, err := GetMasterSegPrefix(sourceMasterDataDir)
	if err != nil {
		return gpinitsystemConfig, errors.Wrap(err, "Could not get master segment prefix")
	}

	gplog.Info("Data Dir: %s", sourceMasterDataDir)
	gplog.Info("segPrefix: %v", segPrefix)
	gpinitsystemConfig = append(gpinitsystemConfig, "SEG_PREFIX="+segPrefix, "TRUSTED_SHELL=ssh")

	return gpinitsystemConfig, nil
}

func WriteInitsystemFile(gpinitsystemConfig []string, gpinitsystemFilepath string) error {
	gpinitsystemContents := []byte(strings.Join(gpinitsystemConfig, "\n"))

	err := ioutil.WriteFile(gpinitsystemFilepath, gpinitsystemContents, 0644)
	if err != nil {
		return errors.Wrap(err, "Could not write gpinitsystem_config file")
	}
	return nil
}

func upgradeDataDir(path string) string {
	// e.g.
	//   /data/primary/seg1
	// becomes
	//   /data/primary_upgrade/seg1
	path = filepath.Clean(path)
	parent := fmt.Sprintf("%s_upgrade", filepath.Dir(path))
	return filepath.Join(parent, filepath.Base(path))
}

// sanitize sorts and deduplicates a slice of port numbers.
func sanitize(ports []uint32) []uint32 {
	sort.Slice(ports, func(i, j int) bool { return ports[i] < ports[j] })

	dedupe := ports[:0] // point at the same backing array

	var last uint32
	for i, port := range ports {
		if i == 0 || port != last {
			dedupe = append(dedupe, port)
		}
		last = port
	}

	return dedupe
}

func WriteSegmentArray(config []string, source *utils.Cluster, ports []uint32) ([]string, int, error) {
	master, segments, err := getTargetConfig(source, ports)

	var primaries []*cluster.SegConfig

	for _, segment := range segments {
		segment := segment
		if segment.Role == cluster.PrimaryRole && segment.ContentID != cluster.MASTER_CONTENT_ID {
			primaries = append(primaries, segment)
		}
	}

	if err != nil {
		return nil, 0, err
	}

	config = append(config,
		fmt.Sprintf("QD_PRIMARY_ARRAY=%s~%d~%s~%d~%d~0",
			master.Hostname,
			master.Port,
			master.DataDir,
			master.DbID,
			master.ContentID,
		),
	)

	config = append(config, "declare -a PRIMARY_ARRAY=(")
	for _, primary := range primaries {
		config = append(config,
			fmt.Sprintf("\t%s~%d~%s~%d~%d~0",
				primary.Hostname,
				primary.Port,
				primary.DataDir,
				primary.DbID,
				primary.ContentID,
			),
		)
	}
	config = append(config, ")")

	return config, master.Port, nil
}

func getTargetConfig(source *utils.Cluster, desiredPorts []uint32) (*cluster.SegConfig, []*cluster.SegConfig, error) {
	segmentsByHost := GroupSegmentsByHostname(source)
	targetPorts := DeterminePortsForTarget(desiredPorts, segmentsByHost)

	// Use a copy of the source cluster's segment configs rather than modifying
	// the source cluster. This keeps the in-memory representation of source
	// cluster consistent with its on-disk representation.
	copyOrPrimaries := make(map[int]cluster.SegConfig)

	for _, segments := range segmentsByHost {
		numPrimariesFound := 0
		for _, segment := range segments {
			if segment.ContentID == cluster.MASTER_CONTENT_ID || segment.Role == cluster.MirrorRole {
				continue
			}

			if len(targetPorts.PrimaryPorts) > numPrimariesFound {
				segment.Port = int(targetPorts.PrimaryPorts[numPrimariesFound])
				segment.DataDir = upgradeDataDir(segment.DataDir)
				copyOrPrimaries[segment.ContentID] = segment
				numPrimariesFound++
			}
		}
	}

	master, ok := source.Primaries[cluster.MASTER_CONTENT_ID]
	master.Port = int(targetPorts.MasterPort)
	master.DataDir = upgradeDataDir(master.DataDir)
	if !ok {
		return nil, nil, errors.New("old cluster contains no master segment")
	}

	return &master, flattenSegments(copyOrPrimaries), nil
}

//
// Transform map of ContentId -> SegConfig to a list of
// SegConfigs
//
// - enforces a sorted order to stabilize test that assumes an order
//
func flattenSegments(copySegments map[int]cluster.SegConfig) []*cluster.SegConfig {
	// Determine ordering based on ContentId
	keys := []int{}
	for key, _ := range copySegments {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Build up list to return
	returnSegments := []*cluster.SegConfig{}
	for _, key := range keys {
		segment := copySegments[key]
		returnSegments = append(returnSegments, &segment)
	}
	return returnSegments
}

func CreateAllDataDirectories(agentConns []*Connection, source *utils.Cluster) error {
	targetDataDir := path.Dir(source.MasterDataDir()) + "_upgrade"
	err := utils.CreateDataDirectory(targetDataDir)
	if err != nil {
		return err
	}
	err = CreateSegmentDataDirectories(agentConns, source)
	if err != nil {
		return err
	}

	return nil
}

func RunInitsystemForTargetCluster(stream step.OutStreams, target *utils.Cluster, gpinitsystemFilepath string) error {
	gphome := filepath.Dir(path.Clean(target.BinDir)) //works around https://github.com/golang/go/issues/4837 in go10.4

	args := "-a -I " + gpinitsystemFilepath
	if target.Version.SemVer.Major < 7 {
		// For 6X we add --ignore-warnings to gpinitsystem to return 0 on
		// warnings and 1 on errors. 7X and later does this by default.
		args += " --ignore-warnings"
	}

	script := fmt.Sprintf("source %[1]s/greenplum_path.sh && %[1]s/bin/gpinitsystem %[2]s",
		gphome,
		args,
	)
	cmd := execCommand("bash", "-c", script)

	cmd.Stdout = stream.Stdout()
	cmd.Stderr = stream.Stderr()

	err := cmd.Run()
	if err != nil {
		return xerrors.Errorf("gpinitsystem: %w", err)
	}

	return nil
}

func GetMasterSegPrefix(datadir string) (string, error) {
	const masterContentID = "-1"

	base := path.Base(datadir)
	if !strings.HasSuffix(base, masterContentID) {
		return "", fmt.Errorf("path requires a master content identifier: '%s'", datadir)
	}

	segPrefix := strings.TrimSuffix(base, masterContentID)
	if segPrefix == "" {
		return "", fmt.Errorf("path has no segment prefix: '%s'", datadir)
	}
	return segPrefix, nil
}

func CreateSegmentDataDirectories(agentConns []*Connection, cluster *utils.Cluster) error {
	wg := sync.WaitGroup{}
	errChan := make(chan error, len(agentConns))

	for _, conn := range agentConns {
		wg.Add(1)

		go func(c *Connection) {
			defer wg.Done()

			segments, err := cluster.SegmentsOn(c.Hostname)
			if err != nil {
				errChan <- err
				return
			}

			req := new(idl.CreateSegmentDataDirRequest)
			for _, seg := range segments {
				// gpinitsystem needs the *parent* directories of the new
				// segment data directories to exist.
				datadir := filepath.Dir(upgradeDataDir(seg.DataDir))
				req.Datadirs = append(req.Datadirs, datadir)
			}

			_, err = c.AgentClient.CreateSegmentDataDirectories(context.Background(), req)
			if err != nil {
				gplog.Error("Error creating segment data directories on host %s: %s",
					c.Hostname, err.Error())
				errChan <- err
			}
		}(conn)
	}

	wg.Wait()
	close(errChan)

	// TODO: Use a multierror to differentiate errors between hosts.
	for err := range errChan {
		if err != nil {
			return xerrors.Errorf("segment data directories: %w", err)
		}
	}
	return nil
}
