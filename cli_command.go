package main

import (
    "github.com/Wayne-Francis/pokedexcli/internal/pokecache"
    "github.com/Wayne-Francis/pokedexcli/internal/pokeapi"
)

var commands map[string]cliCommand

type cliCommand struct {
    name        string
    description string
    callback    func(*config, []string) error
}

type config struct {
    Next        string
    Previous    string
    cache 	*pokecache.Cache
    pokedex  map[string]pokeapi.Pokemon
}

//var commands = map[string]cliCommand{
   // "exit": {
    //    name:        "exit",
   //     description: "Exit the Pokedex",
   //     callback:    commandExit,
   // },
	//"help": {
    //    name:        "help",
    //    description: "Displays a help message",
     //   callback:    commandHelp,
    //},
	//"map": {
   //     name:        "map",
     //   description: "Displays a list of 20 locations",
    //    callback:    commandMap,
   // },
    //"mapb": {
     //   name:        "mapb",
    //   description: "Displays a list the last 20 locations",
    //    callback:    commandMapb,
   // },
    //"explore": {
      //  name:        "explore",
      //  description: "Displays a list of pokemon within an area",
      //  callback:    commandExplore,
  //  },
//}