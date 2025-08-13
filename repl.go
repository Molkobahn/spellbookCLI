package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(cfg *Config, args ...string)error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"close": {
			name:	"close",
			description:	"Close the spellbook",
			callback:		commandClose,
		},
		"spells": {
			name: 			"spells",
			description:	"Shows list of spells",
			callback:		commandSpells,		
		},
	}
}

func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Spellbook > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}