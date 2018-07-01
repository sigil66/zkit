package action

type Action interface {
	Id() string
	Key() string
	Type() string
	Condition() *bool
}

type Actions []Action
