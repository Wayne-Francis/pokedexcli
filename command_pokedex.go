package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")
    if len(cfg.pokedex) == 0 {
      fmt.Printf("You have not captured any Pokemon!\n")
    }
    for _, pokemon := range cfg.pokedex {
        fmt.Printf(" - %v\n", pokemon.Name)
    }
    return nil
}