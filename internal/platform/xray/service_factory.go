package xray

func NewRuntimeService(
	binaryPath string,
	configPath string,
	workingDir string,
) *Service {

	process := NewProcessManager(
		binaryPath,
		configPath,
		workingDir,
	)

	return NewService(process)
}
