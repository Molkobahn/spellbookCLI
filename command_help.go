package main

import(
	"fmt"
)

func commandHelp(cfg *Config, args ...string) error {
	commands := getCommands()
	fmt.Println("Available commands:")
	for _, command := range commands {
		fmt.Println("")
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}