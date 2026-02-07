package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
	"github.com/Wayne-Francis/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
    if cfg.Next == "" {
	cfg.Next = "https://pokeapi.co/api/v2/location-area"
    }
    res, err := http.Get(cfg.Next)
    if err != nil {
        log.Fatal(err)
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode > 299 {
        log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
    }
    if err != nil {
        log.Fatal(err)
    }

    var m pokeapi.LocationAreaList
    err = json.Unmarshal(body, &m)
    if err != nil {
        fmt.Println(err)
    }

    for _, location := range m.Results {
	fmt.Println(location.Name)
 	}
    if m.Next != nil {
	cfg.Next = *m.Next
	} 
    if m.Next == nil {
	cfg.Next = ""
	}
    if m.Previous != nil {
	cfg.Previous = *m.Previous
	} 
    if m.Previous == nil {
	cfg.Previous = ""
	}	
    return nil
}

func commandMapb(cfg *config) error {
    if cfg.Previous == "" {
      fmt.Println("you're on the first page")
      return nil
    }
    res, err := http.Get(cfg.Previous)
    if err != nil {
        log.Fatal(err)
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode > 299 {
        log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
    }
    if err != nil {
        log.Fatal(err)
    }

    var m pokeapi.LocationAreaList
    err = json.Unmarshal(body, &m)
    if err != nil {
        fmt.Println(err)
    }

    for _, location := range m.Results {
	fmt.Println(location.Name)
 	}
    if m.Previous != nil {
	cfg.Previous = *m.Previous
	} 
    if m.Previous == nil {
	cfg.Previous = ""
	}
    if m.Next != nil {
	cfg.Next = *m.Next
	} 
    if m.Next == nil {
	cfg.Next = ""
	}	
    return nil
}