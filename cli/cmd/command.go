package cmd

type Command interface {
	Name() string
	OnAction(args []string) error
}
