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
        fmt.Print("Pokedex > ")
        scanner.Scan()         // moves to the next line
		line := scanner.Text()     // gets the current line
        clean_input := cleanInput(line)
        first_word := clean_input[0]
        fmt.Printf("Your command was: %s\n", first_word)
    }
}