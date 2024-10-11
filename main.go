package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		cmd := scanner.Text()
		fmt.Println(cmd)
		if scanner.Err() != nil {
			// Handle error.
		}
	}
}
