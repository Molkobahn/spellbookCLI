package main

import(
	"fmt"
)

const baseURL = "https://www.dnd5eapi.co/api/2014"

type Config struct {
	args string
}

func main() {
	cfg := Config{}
	fmt.Println("Welcome to your spellbook!")
	startRepl(&cfg)
}