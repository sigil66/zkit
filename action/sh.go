package action

import (
	"fmt"
)

type Sh struct {
	Name string `hcl:"name,label"`
	Cmd []string `hcl:"cmd"`
	Env map[string]string `hcl:"env,optional"`
	Output bool  `hcl:"output,optional"`
	Shell string `hcl:"shell,optional"`

	OnCondition *bool `hcl:"on_condition,optional"`
	AllowFailure bool `hcl:"allow_failure,optional"`
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

func (s *Sh) Condition() *bool {
	return s.OnCondition
}

func (s *Sh) MayFail() bool {
	return s.AllowFailure
}
