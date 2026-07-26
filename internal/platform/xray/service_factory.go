package xray

func NewRuntimeService(
	binaryPath string,
	configPath string,
	workingDir string,
	config *ConfigService,
	client Client,
) *Service {

	process := NewProcessManager(
		binaryPath,
		configPath,
		workingDir,
	)

	return NewService(
		process,
		client,
		config,
	)
}
