package hub

import "github.com/greenplum-db/gpupgrade/utils/cluster"

type TargetPorts struct {
	MasterPort   uint32
	StandbyPort  uint32
	PrimaryPorts []uint32
}

func DeterminePortsForTarget(desiredPorts []uint32, segmentsByHostname map[string][]cluster.SegConfig) TargetPorts {
	if len(desiredPorts) == 0 {
		desiredPorts = generatePorts(segmentsByHostname)
	}

	sanitizedPorts := sanitize(desiredPorts)
	var masterPort uint32 = 0
	var standbyPort uint32 = 0
	var segmentPorts []uint32

	if containsMaster(segmentsByHostname) {
		masterPort = sanitizedPorts[0]
	}

	if containsStandby(segmentsByHostname) {
		standbyPort = sanitizedPorts[len(sanitizedPorts)-1]
		sanitizedPorts = sanitizedPorts[:len(sanitizedPorts)-1]
	}

	if len(sanitizedPorts) > 0 {
		segmentPorts = sanitizedPorts[1:]
	}

	return TargetPorts{
		MasterPort:   masterPort,
		PrimaryPorts: segmentPorts,
		StandbyPort:  standbyPort,
	}
}

func containsMaster(segmentsByHostname map[string][]cluster.SegConfig) bool {
	return containsSegmentByContentIdAndRole(segmentsByHostname,
		cluster.MASTER_CONTENT_ID, cluster.PrimaryRole)
}

func containsStandby(segmentsByHostname map[string][]cluster.SegConfig) bool {
	return containsSegmentByContentIdAndRole(segmentsByHostname,
		cluster.MASTER_CONTENT_ID, cluster.MirrorRole)
}

func containsSegmentByContentIdAndRole(segmentsByHostname map[string][]cluster.SegConfig, contentId int, role string) bool {
	for _, segments := range segmentsByHostname {
		for _, segment := range segments {
			if segment.ContentID == contentId &&
				segment.Role == role {
				return true
			}
		}
	}
	return false
}

// Create a default port range, starting with the pg_upgrade default of
// 50432. Reserve enough ports to handle the host with the most
// segments.
func generatePorts(segmentsByHost map[string][]cluster.SegConfig) (ports []uint32) {
	var maxSegs int

	for _, segments := range segmentsByHost {
		if len(segments) > maxSegs {
			maxSegs = len(segments)
		}
	}

	for i := 0; i < maxSegs; i++ {
		ports = append(ports, uint32(50432+i))
	}

	return ports
}
