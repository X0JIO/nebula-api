package xray

import (
	"context"
	"errors"
	"os"
	"os/exec"
)

type ProcessManager interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Restart(ctx context.Context) error

	Running() bool
	PID() int
}

func NewProcessManager(
	binaryPath string,
	configPath string,
	workingDir string,
) ProcessManager {

	return &WindowsProcessManager{
		binaryPath: binaryPath,
		configPath: configPath,
		workingDir: workingDir,
	}
}

func (p *WindowsProcessManager) Start(
	ctx context.Context,
) error {

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.cmd != nil {
		return errors.New("xray already running")
	}

	cmd := exec.CommandContext(
		ctx,
		p.binaryPath,
		"run",
		"-config",
		p.configPath,
	)

	if p.workingDir != "" {
		cmd.Dir = p.workingDir
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	p.cmd = cmd

	return nil
}

func (p *WindowsProcessManager) Stop(
	ctx context.Context,
) error {

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.cmd == nil {
		return nil
	}

	err := p.cmd.Process.Kill()

	p.cmd = nil

	return err
}

func (p *WindowsProcessManager) Restart(
	ctx context.Context,
) error {

	if err := p.Stop(ctx); err != nil {
		return err
	}

	return p.Start(ctx)
}

func (p *WindowsProcessManager) Running() bool {

	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.cmd != nil
}

func (p *WindowsProcessManager) PID() int {

	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.cmd == nil || p.cmd.Process == nil {
		return 0
	}

	return p.cmd.Process.Pid
}
