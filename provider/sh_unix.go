package provider

import (
	"github.com/solvent-io/zkit/action"
	"context"
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

	return "execin the world bro!", nil
}

