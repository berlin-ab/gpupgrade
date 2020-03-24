package greenplum

type gpStart struct {
	cluster *Cluster
	runner  Runner
}

func newGpStart(cluster *Cluster, runner Runner) *gpStart {
	return &gpStart{
		cluster: cluster,
		runner:  runner,
	}
}

func (m *gpStart) Start() error {
	return m.runner.Run("gpstart", "-a", "-d", m.cluster.MasterDataDir())
}

func (m *gpStart) StartMasterOnly() error {
	return m.runner.Run("gpstart", "-m", "-a", "-d", m.cluster.MasterDataDir())
}
