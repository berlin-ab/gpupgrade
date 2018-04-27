package integrations_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gp_upgrade/hub/cluster"
	"gp_upgrade/hub/configutils"
	"gp_upgrade/hub/services"
	"gp_upgrade/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"google.golang.org/grpc"
)

// needs the cli and the hub
var _ = Describe("check", func() {
	var (
		dir           string
		hub           *services.HubClient
		commandExecer *testutils.FakeCommandExecer
	)

	BeforeEach(func() {
		var err error
		dir, err = ioutil.TempDir("", "")
		Expect(err).ToNot(HaveOccurred())

		conf := &services.HubConfig{
			CliToHubPort:   7527,
			HubToAgentPort: 6416,
			StateDir:       dir,
		}
		reader := configutils.NewReader()

		commandExecer = &testutils.FakeCommandExecer{}
		commandExecer.SetOutput(&testutils.FakeCommand{})

		hub = services.NewHub(&cluster.Pair{}, &reader, grpc.DialContext, commandExecer.Exec, conf)

		Expect(checkPortIsAvailable(7527)).To(BeTrue())
		go hub.Start()
	})

	AfterEach(func() {
		hub.Stop()
		Expect(checkPortIsAvailable(7527)).To(BeTrue())
		os.RemoveAll(dir)
	})

	// `gp_backup check seginstall` verifies that the user has installed the software on all hosts
	// As a single-node check, this test verifies the mechanics of the check, but would typically succeed.
	// The implementation, however, uses the gp_upgrade_agent binary to verify installation. In real life,
	// all the binaries, gp_upgrade_hub and gp_upgrade_agent included, would be alongside each other.
	// But in our integration tests' context, only the necessary Golang code is compiled, and Ginkgo's default
	// is to compile gp_upgrade_hub and gp_upgrade_agent in separate directories. As such, this test depends on the
	// setup in `integrations_suite_test.go` to replicate the real-world scenario of "install binaries side-by-side".
	//
	// TODO: This test might be interesting to run multi-node; for that, figure out how "installation" should be done
	Describe("seginstall", func() {
		It("updates status PENDING to RUNNING then to COMPLETE if successful", func(done Done) {
			defer close(done)

			config := `[{
  			  "content": 2,
  			  "dbid": 7,
  			  "hostname": "localhost"
  			}]`
			testutils.WriteProvidedConfig(dir, config)

			f, err := os.Create(filepath.Join(dir, "new_cluster_config.json"))
			Expect(err).ToNot(HaveOccurred())
			f.Close()

			Expect(runStatusUpgrade()).To(ContainSubstring("PENDING - Install binaries on segments"))

			trigger := make(chan struct{}, 1)
			commandExecer.SetTrigger(trigger)
			commandExecer.SetOutput(&testutils.FakeCommand{
				Err: nil,
				Out: []byte("some output"),
			})

			wg := &sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer GinkgoRecover()

				Eventually(runStatusUpgrade).Should(ContainSubstring("RUNNING - Install binaries on segments"))
				trigger <- struct{}{}
			}()

			checkSeginstallSession := runCommand("check", "seginstall", "--master-host", "localhost")

			Eventually(checkSeginstallSession).Should(Exit(0))
			wg.Wait()

			Expect(commandExecer.Command()).To(Equal("ssh"))
			Expect(strings.Join(commandExecer.Args(), "")).To(ContainSubstring("ls"))
			Eventually(runStatusUpgrade).Should(ContainSubstring("COMPLETE - Install binaries on segments"))
		})
	})

	It("fails if the --master-host flag is missing", func() {
		checkSeginstallSession := runCommand("check", "seginstall")
		Expect(checkSeginstallSession).Should(Exit(1))
		Expect(string(checkSeginstallSession.Out.Contents())).To(Equal("Required flag(s) \"master-host\" have/has not been set\n"))
	})
})
