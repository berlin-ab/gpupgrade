package greenplum

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/greenplum-db/gpupgrade/testutils/exectest"
	"github.com/greenplum-db/gpupgrade/utils"
)

func TestMain(m *testing.M) {
	os.Exit(exectest.Run(m))
}

func StartClusterCmd() {}
func StopClusterCmd()  {}
func PgrepCmd()        {}
func PgrepCmd_Errors() {
	os.Stderr.WriteString("exit status 2")
	os.Exit(2)
}

func init() {
	exectest.RegisterMains(
		StartClusterCmd,
		StopClusterCmd,
		PgrepCmd,
		PgrepCmd_Errors,
	)
}

// TODO: Consolidate with the same function in common_test.go in the
//  hub package. This is tricky due to cycle imports and other issues.
// MustCreateCluster creates a utils.Cluster and calls t.Fatalf() if there is
// any error.
func MustCreateCluster(t *testing.T, segs []SegConfig) *Cluster {
	t.Helper()

	cluster, err := NewCluster(segs)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	return cluster
}

// TODO: Consolidate with the same function in common_test.go in the hub package.
// DevNull implements OutStreams by just discarding all writes.
var DevNull = devNull{}

type devNull struct{}

func (_ devNull) Stdout() io.Writer {
	return ioutil.Discard
}

func (_ devNull) Stderr() io.Writer {
	return ioutil.Discard
}

func TestStartOrStopCluster(t *testing.T) {
	g := NewGomegaWithT(t)

	source := MustCreateCluster(t, []SegConfig{
		{ContentID: -1, DbID: 1, Port: 15432, Hostname: "localhost", DataDir: "basedir/seg-1", Role: "p"},
	})
	source.BinDir = "/source/bindir"

	utils.System.RemoveAll = func(s string) error { return nil }
	utils.System.MkdirAll = func(s string, perm os.FileMode) error { return nil }

	startStopCmd = nil
	pgrepCmd = nil

	defer func() {
		startStopCmd = exec.Command
		pgrepCmd = exec.Command
	}()

	t.Run("isPostmasterRunning succeeds", func(t *testing.T) {
		pgrepCmd = exectest.NewCommandWithVerifier(PgrepCmd,
			func(path string, args ...string) {
				g.Expect(path).To(Equal("bash"))
				g.Expect(args).To(Equal([]string{"-c", "pgrep -F basedir/seg-1/postmaster.pid"}))
			})

		command := pgrepCommand{streams: DevNull}
		err := command.isRunning(source.MasterPidFile())
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("isPostmasterRunning fails", func(t *testing.T) {
		pgrepCmd = exectest.NewCommand(PgrepCmd_Errors)

		command := pgrepCommand{streams: DevNull}
		err := command.isRunning(source.MasterPidFile())
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("stop cluster successfully shuts down cluster", func(t *testing.T) {
		pgrepCmd = exectest.NewCommandWithVerifier(PgrepCmd,
			func(path string, args ...string) {
				g.Expect(path).To(Equal("bash"))
				g.Expect(args).To(Equal([]string{"-c", "pgrep -F basedir/seg-1/postmaster.pid"}))
			})

		startStopCmd = exectest.NewCommandWithVerifier(StopClusterCmd,
			func(path string, args ...string) {
				g.Expect(path).To(Equal("bash"))
				g.Expect(args).To(Equal([]string{"-c", "source /source/bindir/../greenplum_path.sh " +
					"&& /source/bindir/gpstop -a -d basedir/seg-1"}))
			})

		err := source.Stop(DevNull)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("stop cluster detects that cluster is already shutdown", func(t *testing.T) {
		pgrepCmd = exectest.NewCommand(PgrepCmd_Errors)

		var skippedStopClusterCommand = true
		startStopCmd = exectest.NewCommandWithVerifier(PgrepCmd,
			func(path string, args ...string) {
				skippedStopClusterCommand = false
			})

		err := source.Stop(DevNull)
		g.Expect(err).To(HaveOccurred())
		g.Expect(skippedStopClusterCommand).To(Equal(true))
	})

	t.Run("start cluster successfully starts up cluster", func(t *testing.T) {
		startStopCmd = exectest.NewCommandWithVerifier(StartClusterCmd,
			func(path string, args ...string) {
				g.Expect(path).To(Equal("bash"))
				g.Expect(args).To(Equal([]string{"-c", "source /source/bindir/../greenplum_path.sh " +
					"&& /source/bindir/gpstart -a -d basedir/seg-1"}))
			})

		err := source.Start(DevNull)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("start master successfully starts up master only", func(t *testing.T) {
		startStopCmd = exectest.NewCommandWithVerifier(StartClusterCmd,
			func(path string, args ...string) {
				g.Expect(path).To(Equal("bash"))
				g.Expect(args).To(Equal([]string{"-c", "source /source/bindir/../greenplum_path.sh " +
					"&& /source/bindir/gpstart -m -a -d basedir/seg-1"}))
			})

		err := source.StartMasterOnly(DevNull)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("stop master successfully shuts down master only", func(t *testing.T) {
		pgrepCmd = exectest.NewCommandWithVerifier(PgrepCmd,
			func(path string, args ...string) {
				g.Expect(path).To(Equal("bash"))
				g.Expect(args).To(Equal([]string{"-c", "pgrep -F basedir/seg-1/postmaster.pid"}))
			})

		startStopCmd = exectest.NewCommandWithVerifier(StopClusterCmd,
			func(path string, args ...string) {
				g.Expect(path).To(Equal("bash"))
				g.Expect(args).To(Equal([]string{"-c", "source /source/bindir/../greenplum_path.sh " +
					"&& /source/bindir/gpstop -m -a -d basedir/seg-1"}))
			})

		err := source.StopMasterOnly(DevNull)
		g.Expect(err).ToNot(HaveOccurred())
	})
}
