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

	configService := NewDefaultConfigService(
		workingDir,
	)

	return NewService(
		process,
		configService,
	)
}
