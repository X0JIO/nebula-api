package xray

import (
	"os/exec"
	"sync"
)

type WindowsProcessManager struct {
	binaryPath string
	configPath string
	workingDir string

	cmd *exec.Cmd
	mu  sync.RWMutex
}
