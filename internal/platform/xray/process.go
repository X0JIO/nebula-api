package xray

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ProcessManager interface {
	Start(ctx context.Context) error

	Stop(ctx context.Context) error

	Restart(ctx context.Context) error

	Running() bool

	PID() int

	Status() Status

	Version(ctx context.Context) (string, error)
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

	if _, err := os.Stat(p.binaryPath); err != nil {
		return fmt.Errorf(
			"xray binary not found: %s",
			p.binaryPath,
		)
	}

	if p.workingDir != "" {
		cmd.Dir = p.workingDir
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	p.cmd = cmd

	go func() {

		err := cmd.Wait()

		p.mu.Lock()
		defer p.mu.Unlock()

		p.cmd = nil

		if err != nil {
			// TODO: добавить logger
		}

	}()

	return nil
}

func (p *WindowsProcessManager) Stop(
	ctx context.Context,
) error {

	p.mu.Lock()

	if p.cmd == nil {
		p.mu.Unlock()
		return nil
	}

	cmd := p.cmd
	p.cmd = nil

	p.mu.Unlock()

	if err := cmd.Process.Signal(os.Interrupt); err != nil {

		return cmd.Process.Kill()
	}

	done := make(chan error, 1)

	go func() {
		done <- cmd.Wait()
	}()

	select {

	case <-ctx.Done():
		return cmd.Process.Kill()

	case err := <-done:
		return err
	}
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

func (p *WindowsProcessManager) Status() Status {

	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.cmd != nil &&
		p.cmd.Process != nil {

		return Status{
			Running: true,
			PID:     p.cmd.Process.Pid,
		}
	}

	return p.detectExternalProcess()
}

func (p *WindowsProcessManager) Version(
	ctx context.Context,
) (string, error) {

	cmd := exec.CommandContext(
		ctx,
		p.binaryPath,
		"version",
	)

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(
		string(output),
	), nil
}
