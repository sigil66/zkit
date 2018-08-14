package provider

import (
	"context"
	"github.com/solvent-io/zkit/action"
	"github.com/chuckpreslar/emission"
)

type Provider interface {
	Realize(ctx context.Context) error
}

type Options struct {
	OutputPath string
	TargetPath string
	WorkPath string
	CachePath string

	Debug bool
	Verbose bool
}

type Factory struct {
	phaseMap map[string]map[string]string
	emitter *emission.Emitter
}

func New(phaseMap map[string]map[string]string, emitter *emission.Emitter) *Factory {
	return &Factory{phaseMap, emitter}
}

// Need to add provider switching
// for now defaults will work on all OSs we care about
func (f *Factory) Get(ac action.Action) Provider {
	switch ac.Type() {
	case "Sh":
		return NewShUnix(ac, f.phaseMap["Sh"], f.emitter)
	}

	return nil
}

func Phase(ctx context.Context) string {
	return ctx.Value("phase").(string)
}

func Opts(ctx context.Context) *Options {
	return ctx.Value("options").(*Options)
}