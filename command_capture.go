package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "github.com/Wayne-Francis/pokedexcli/internal/pokeapi"
    "math/rand"
)

func commandCatch(cfg *config, args []string) error {
    baseurl := "https://pokeapi.co/api/v2/pokemon/"
    if len(args) < 1 {
        fmt.Println("please enter a pokemon to catch")
        return nil
    }
    fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
    pokemon := args[0]
    fullurl := baseurl + pokemon + "/"

    var body []byte
    var err error

    cachedBody, found := cfg.cache.Get(fullurl)
    if found {
        body = cachedBody
    } else {
        res, err := http.Get(fullurl)
        if err != nil {
            return err
        }
        defer res.Body.Close()

        body, err = io.ReadAll(res.Body)
        if err != nil {
            return err
        }

        if res.StatusCode == http.StatusNotFound {
            return fmt.Errorf("no such pokemon: %s", pokemon)
        }
        if res.StatusCode > 299 {
            return fmt.Errorf("catch failed with status code: %d", res.StatusCode)
        }
        cfg.cache.Add(fullurl, body)
    }

    var p pokeapi.Pokemon
    err = json.Unmarshal(body, &p)
    if err != nil {
        return err
    }
    exp := p.BaseExperience
    max_xp := 90
    if exp > max_xp {
        exp = max_xp
    }
    
    roll := rand.Intn(100)

    if roll > exp {
    fmt.Printf("%s was caught!\n", p.Name)
    cfg.pokedex[p.Name] = p
    } else {
    fmt.Printf("%s escaped!\n", p.Name)
    }

    return nil
}