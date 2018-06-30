package action

import (
	"fmt"
)

type Sh struct {
	Name string `hcl:"name,label"`
	Cmd []string `hcl:"cmd"`
	Output bool  `hcl:"output,optional"`
}

func (s *Sh) Key() string {
	return s.Name
}

func (s *Sh) Id() string {
	return fmt.Sprint(s.Type(), ".", s.Key())
}

func (s *Sh) Type() string {
	return "Sh"
}
