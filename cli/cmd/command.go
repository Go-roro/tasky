package cmd

import (
	"fmt"
	"os"
)

type Command interface {
	Name() string
	OnAction(args []string) (string, error)
}

var commands = []Command{
	&AddCommand{},
}

func ExecuteCommand(args []string) error {
	if len(os.Args) < 1 {
		return fmt.Errorf("no command provided")
	}

	command := args[0]
	flags := args[1:]

	for _, cmd := range commands {
		if cmd.Name() == command {
			result, err := cmd.OnAction(flags)
			if err != nil {
				return fmt.Errorf("error executing command '%s': %w", command, err)
			}
			fmt.Println(result)
			return nil
		}
	}
	return fmt.Errorf("unknown command: %s", command)
}
