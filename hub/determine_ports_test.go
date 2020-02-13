package hub_test

import (
	"testing"

	"github.com/greenplum-db/gpupgrade/utils/cluster"

	"github.com/greenplum-db/gpupgrade/hub"
)

func TestDeterminePorts(t *testing.T) {
	t.Run("when no master is given", func(t *testing.T) {
		t.Run("it returns 0 for the Master Port", func(t *testing.T) {
			segmentsByHostname := map[string][]cluster.SegConfig{}
			segmentsByHostname["host1"] = []cluster.SegConfig{
				{ContentID: cluster.MASTER_CONTENT_ID, Role: cluster.MirrorRole},
			}
			targetPorts := hub.DeterminePortsForTarget([]uint32{}, segmentsByHostname)

			var invalidPortValue uint32 = 0
			if targetPorts.MasterPort != invalidPortValue {
				t.Errorf("got master port of %v, expected master to be the invalid value of %v",
					targetPorts.MasterPort,
					invalidPortValue)
			}
		})
	})

	t.Run("when no ports are specified as desired", func(t *testing.T) {
		t.Run("it returns the last assumed port for the standby master port", func(t *testing.T) {
			segmentsByHostname := map[string][]cluster.SegConfig{}

			segmentsByHostname["host1"] = []cluster.SegConfig{
				{ContentID: cluster.MASTER_CONTENT_ID, Role: cluster.MirrorRole},
			}

			targetPorts := hub.DeterminePortsForTarget([]uint32{}, segmentsByHostname)

			var expectedStandbyPort uint32 = 50432
			if targetPorts.StandbyPort != expectedStandbyPort {
				t.Errorf("got standby master port of %v, expected standby master to get the last assumed port of %v",
					targetPorts.StandbyPort,
					expectedStandbyPort)
			}
		})

		t.Run("it returns the first assumed port for the master", func(t *testing.T) {
			segmentsByHostname := map[string][]cluster.SegConfig{}
			segmentsByHostname["host1"] = []cluster.SegConfig{
				{ContentID: cluster.MASTER_CONTENT_ID, Role: cluster.PrimaryRole},
			}
			targetPorts := hub.DeterminePortsForTarget([]uint32{}, segmentsByHostname)

			var expectedMasterPort uint32 = 50432
			if targetPorts.MasterPort != expectedMasterPort {
				t.Errorf("got master port of %v, expected master to get the first assumed port of %v",
					targetPorts.MasterPort,
					expectedMasterPort)
			}
		})
	})
}
