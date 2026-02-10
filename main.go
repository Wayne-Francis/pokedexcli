package main

import (
	"bufio"
	"fmt"
    "os"
    "github.com/Wayne-Francis/pokedexcli/internal/pokecache"
    "github.com/Wayne-Francis/pokedexcli/internal/pokeapi"
     "time"
     "math/rand"
)

// main function: program execution starts here
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    rand.Seed(time.Now().UnixNano())
    commands = map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Displays a list of 20 locations",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Displays a list of the last 20 locations",
            callback:    commandMapb,
        },
        "explore": {
            name:        "explore",
            description: "explore a location for pokemon",
            callback:    commandExplore,
        },
        "catch": {
            name:        "catch",
            description: "Try to Catch a Pokemon!",
            callback:    commandCatch,
        },
        "inspect": {
            name:        "inspect",
            description: "Inspect the stats of a Pokemon you have caught",
            callback:    commandInspect,
        },
        "pokedex": {
            name:        "pokedex",
            description: "List all the Pokemon you have caught",
            callback:    commandPokedex,
        },
    }
    cfg := &config{
    cache: pokecache.NewCache(5 * time.Second),
    pokedex: make(map[string]pokeapi.FullPokemon),
    }
    for {
        fmt.Println("Pokedex > ")
        scanner.Scan()         
		line := scanner.Text()     
        tokens := cleanInput(line)
        if len(tokens) == 0 {
        continue
        }
        cmdName := tokens[0]
        args := tokens[1:]

        cmd, ok := commands[cmdName]
        if !ok {
            fmt.Printf("Unknown command: %v\n", cmdName)
        continue
        }
        err := cmd.callback(cfg, args)
        if err != nil {
            fmt.Println(err)
        }
    }
}