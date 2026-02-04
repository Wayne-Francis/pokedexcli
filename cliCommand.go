package main

type cliCommand struct {
    name        string
    description string
    callback    func(*config) error
}

type config struct {
    Next        string
    Previous    string
}

var commands = map[string]cliCommand{
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
        description: "Displays a list the last 20 locations",
        callback:    commandMapb,
    },
}