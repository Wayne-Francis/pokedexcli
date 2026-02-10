package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
        fmt.Println("please enter a pokemon to inspect")
        return nil
    }
    pokemon_name := args[0]
    if pokemon, caught := cfg.pokedex[pokemon_name]; !caught {
        fmt.Println("Pokemon has not been caught")
    } else {
        fmt.Printf("Name: %v\n", pokemon.Name)
        fmt.Printf("Height:  %v\n", pokemon.Height)
        fmt.Printf("Weight:  %v\n", pokemon.Weight)
        fmt.Printf("Stats:\n")
        for _, stat := range pokemon.Stats {
        fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
    }
        fmt.Printf("Types:\n")   
        for _, t := range pokemon.Types {
        fmt.Printf("  - %s\n", t.Type.Name)
    }
    }
    return nil
}
