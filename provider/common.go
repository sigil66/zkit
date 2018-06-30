package provider

import (
	"context"
	"github.com/solvent-io/zkit/action"
)

type Provider interface {
	Realize(phase string, ctx context.Context) (string, error)
}

type Options struct {
	OutputPath string
	TargetPath string
	WorkPath string

	Debug bool
	Verbose bool
}

// Need to add provider switching
// for now defaults will work on all OSs we care about
func Get(ac action.Action) Provider {
	switch ac.Type() {
	case "Sh":
		return NewShUnix(ac)
	}

	return nil
}
