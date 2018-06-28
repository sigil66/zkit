package provider

import (
	"github.com/solvent-io/zkit/action"
	"context"
	"os/exec"
	"bufio"
	"fmt"
)

type ShUnix struct {
	sh *action.Sh
}

func NewShUnix(sh action.Action) *ShUnix {
	return &ShUnix{sh.(*action.Sh)}
}

func (s *ShUnix) Realize(phase string, ctx context.Context) (string, error) {
	switch phase {
	case "build":
		return s.exec(ctx)
	default:
		return "", nil
	}
}

func (s *ShUnix) exec(ctx context.Context) (string, error) {
	cmd := exec.Command(s.sh.Cmd[0], s.sh.Cmd[1:]...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("    %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	return "", nil
}

