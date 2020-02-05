package agent

import (
	"github.com/pkg/errors"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/upgrade"
)

//
// Functions to perform an upgrade of a single segment:
//

func upgradeSegment(segment Segment, request *idl.UpgradePrimariesRequest, host string) error {
	err := restoreBackup(request, segment)

	if err != nil {
		return err
	}

	err = performUpgrade(segment, request)

	if err != nil {
		failedAction := "upgrade"
		if request.CheckOnly {
			failedAction = "check"
		}
		return errors.Wrapf(err, "failed to %s primary on host %s with content %d", failedAction, host, segment.Content)
	}

	return nil
}

func performUpgrade(segment Segment, request *idl.UpgradePrimariesRequest) error {
	dbid := int(segment.DBID)
	segmentPair := upgrade.SegmentPair{
		Source: &upgrade.Segment{request.SourceBinDir, segment.SourceDataDir, dbid, int(segment.SourcePort)},
		Target: &upgrade.Segment{request.TargetBinDir, segment.TargetDataDir, dbid, int(segment.TargetPort)},
	}

	options := []upgrade.Option{
		upgrade.WithExecCommand(execCommand),
		upgrade.WithWorkDir(segment.WorkDir),
		upgrade.WithSegmentMode(),
	}

	if request.CheckOnly {
		options = append(options, upgrade.WithCheckOnly())
	}

	if request.UseLinkMode {
		options = append(options, upgrade.WithLinkMode())
	}

	return upgrade.Run(segmentPair, options...)
}

func restoreBackup(request *idl.UpgradePrimariesRequest, segment Segment) error {
	if request.CheckOnly {
		return nil
	}

	return Rsync(request.MasterBackupDir, segment.TargetDataDir, []string{
		"internal.auto.conf",
		"postgresql.conf",
		"pg_hba.conf",
		"postmaster.opts",
		"gp_dbid",
		"gpssh.conf",
		"gpperfmon",
	})
}
