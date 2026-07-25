package xray

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"sync"
)

type ProcessManager struct {
	binary string
	config string

	cmd *exec.Cmd
	mu  sync.Mutex
}

func NewProcessManager(
	binary string,
	config string,
) *ProcessManager {

	return &ProcessManager{
		binary: binary,
		config: config,
	}
}

func (p *ProcessManager) Start(
	ctx context.Context,
) error {

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.cmd != nil {
		return errors.New("xray already running")
	}

	cmd := exec.CommandContext(
		ctx,
		p.binary,
		"run",
		"-config",
		p.config,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	p.cmd = cmd

	return nil
}

func (p *ProcessManager) Stop() error {

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.cmd == nil {
		return nil
	}

	err := p.cmd.Process.Kill()

	p.cmd = nil

	return err
}

func (p *ProcessManager) Restart(
	ctx context.Context,
) error {

	if err := p.Stop(); err != nil {
		return err
	}

	return p.Start(ctx)
}

func (p *ProcessManager) Running() bool {

	p.mu.Lock()
	defer p.mu.Unlock()

	return p.cmd != nil
}
