package main

import(
	"os"
	"fmt"
)

func commandClose(cfg *Config, args ...string) error {
	fmt.Println("Closing the Spellbook!")
	os.Exit(0)
	return nil
}