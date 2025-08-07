package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		fmt.Println("Your command was:", cleanedInput[0])
	}

}

func cleanInput(text string) []string {

	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}
