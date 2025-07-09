package cmd

import (
	"fmt"
	"os"
)

type Command interface {
	Name() string
	OnAction(args []string) (string, error)
	CommandDescription() string
	CommandHelp() string
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

	executeCommand, err := findCommand(command)
	if err != nil {
		return err
	}

	if flags != nil && (flags[0] == "--help" || flags[0] == "-h") {
		fmt.Println(executeCommand.CommandHelp())
		return nil
	}

	result, err := executeCommand.OnAction(flags)
	if err != nil {
		return fmt.Errorf("error executing command '%s': %w", command, err)
	}
	fmt.Println(result)
	return nil
}

func findCommand(command string) (Command, error) {
	for _, cmd := range commands {
		if cmd.Name() == command {
			return cmd, nil
		}
	}
	return nil, fmt.Errorf("unknown command: %s", command)
}

func PrintStartManual() {
	fmt.Println("📝 Tasky - Interactive Task Manager")
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		fmt.Printf("  - %s: %s \n", cmd.Name(), cmd.CommandDescription())
	}
	fmt.Println("\nType 'exit' or 'quit' to exit the application.")
	fmt.Println("\nUse '<command> --help' for more information on a specific command.")
	fmt.Println("Example: 'add --help' for help on the 'add' command.")
}
