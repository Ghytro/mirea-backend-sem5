package admin

import (
	"bytes"
	"context"
	"os"
	"os/exec"
)

type Service struct {
	shellPath string
}

func NewService() *Service {
	return &Service{
		shellPath: os.Getenv("SHELL_PATH"),
	}
}

func (s *Service) ExecCommand(ctx context.Context, command string) (string, error) {
	cmd := exec.CommandContext(ctx, s.shellPath, "-c", command)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err != nil {
		return stderr.String(), err
	}
	return stdout.String(), nil
}
