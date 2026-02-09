package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "github.com/Wayne-Francis/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args []string) error {
    baseurl := "https://pokeapi.co/api/v2/location-area/"
    if len(args) < 1 {
        fmt.Println("please enter an area to explore")
        return nil
    }
    areaname := args[0]
    fullurl := baseurl + areaname + "/"

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
            return fmt.Errorf("no such area: %s", areaname)
        }
        if res.StatusCode > 299 {
            return fmt.Errorf("explore failed with status code: %d", res.StatusCode)
        }

        cfg.cache.Add(fullurl, body)
    }

    var a pokeapi.LocationArea
    err = json.Unmarshal(body, &a)
    if err != nil {
        return err
    }

    fmt.Printf("Exploring %s...\n", areaname)
    fmt.Println("Found Pokemon:")
    for _, enc := range a.PokemonEncounters {
        fmt.Println(enc.Pokemon.Name)
    }

    return nil
}