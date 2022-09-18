package admin

import (
	"bytes"
	"context"
	"os/exec"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ExecCommand(ctx context.Context, command string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, command, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	return stdout.String(), err
}
