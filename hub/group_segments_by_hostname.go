package hub

import (
	"github.com/greenplum-db/gpupgrade/utils"
	"github.com/greenplum-db/gpupgrade/utils/cluster"
)

func GroupSegmentsByHostname(source *utils.Cluster) map[string][]cluster.SegConfig {
	segmentsByHost := make(map[string][]cluster.SegConfig)
	for _, content := range source.ContentIDs {
		segment := source.Primaries[content]
		segmentsByHost[segment.Hostname] = append(segmentsByHost[segment.Hostname], segment)
	}
	return segmentsByHost
}
