package main

import (
	"fmt"
)

func commandHelp(cfg *config, arg []string) error {
	fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:\n")

    for _, cmd := range commands {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }

    return nil
}
