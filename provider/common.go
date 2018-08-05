package provider

import (
	"context"
	"github.com/solvent-io/zkit/action"
	"github.com/chuckpreslar/emission"
)

type Provider interface {
	Realize(ctx context.Context) error
}

type ProviderOptions struct {
	OutputPath string
	TargetPath string
	WorkPath string

	Debug bool
	Verbose bool
}

// Need to add provider switching
// for now defaults will work on all OSs we care about
func Get(ac action.Action, emitter *emission.Emitter) Provider {
	switch ac.Type() {
	case "Sh":
		return NewShUnix(ac, emitter)
	}

	return nil
}

func Phase(ctx context.Context) string {
	return ctx.Value("phase").(string)
}

func Options(ctx context.Context) *ProviderOptions {
	return ctx.Value("options").(*ProviderOptions)
}