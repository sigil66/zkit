package provider

import (
	"github.com/solvent-io/zkit/action"
	"context"
	"os/exec"
	"bufio"
	"strings"
	"github.com/chuckpreslar/emission"
	"github.com/solvent-io/zkit/phase"
)

type ShUnix struct {
	*emission.Emitter
	sh *action.Sh

	phaseMap map[string]string
}

func NewShUnix(sh action.Action, phaseMap map[string]string, emitter *emission.Emitter) *ShUnix {
	return &ShUnix{emitter,  sh.(*action.Sh), phaseMap}
}

func (s *ShUnix) Realize(ctx context.Context) error {
	switch s.phaseMap[Phase(ctx)] {
	case "exec":
		return s.exec(ctx)
	default:
		return nil
	}
}

func (s *ShUnix) exec(ctx context.Context) error {
	var err error
	var shell string
	options := Opts(ctx)

	if s.sh.Shell != "" {
		shell = s.sh.Shell
	} else {
		shell = "bash"
	}

	cmd := exec.Command(shell, "-c", strings.Join(s.sh.Cmd, " "))

	if s.sh.Env != nil {
		cmd.Env = s.envFromMap(s.sh.Env)
	}

	if options.Verbose || s.sh.Output {
		cmdReader, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				s.Emit("action.verbose.content", scanner.Text())
			}
		}()

		s.Emit("action.verbose.header", strings.Join(s.sh.Cmd, " "))
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (s *ShUnix) envFromMap(env map[string]string) []string {
	var result []string

	for k, v := range env {
		result = append(result, strings.Join([]string{k,v}, "="))
	}

	return result
}
