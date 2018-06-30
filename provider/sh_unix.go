package provider

import (
	"github.com/solvent-io/zkit/action"
	"context"
	"os/exec"
	"bufio"
	"fmt"
	"strings"
	"os"
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
	var err error
	options := ctx.Value("options").(*Options)
	cmd := exec.Command("bash", "-c", strings.Join(s.sh.Cmd, " "))

	if options.Verbose || s.sh.Output {
		cmdReader, err := cmd.StdoutPipe()
		if err != nil {
			return "", err
		}

		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				os.Stdout.WriteString(fmt.Sprintf("  %s\n", scanner.Text()))
			}
		}()

		os.Stdout.WriteString(fmt.Sprintf("  > %s\n", strings.Join(s.sh.Cmd, " ")))
	}

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

