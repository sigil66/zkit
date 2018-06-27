package action

type Action interface {
	Id() string
	Key() string
	Type() string
}

type Actions []Action
