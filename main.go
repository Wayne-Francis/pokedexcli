package main

import (
	"bufio"
	"fmt"
    "os"
)

// main function: program execution starts here
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("Pokedex > ")
        scanner.Scan()         // moves to the next line
		line := scanner.Text()     // gets the current line
        clean_input := cleanInput(line)
        input := clean_input[0]
        cmd, ok := commands[input]
        if !ok {
            fmt.Printf("Unknown command: %v\n", input)
            continue
        }
        err := cmd.callback()
        if err != nil {
            fmt.Println(err)
        }
    }
}