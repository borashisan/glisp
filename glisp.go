package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	str, _ := reader.ReadString('\n')
	array, err := parser(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(array)
}

func parser(str string) ([]string, error) {
	var result []string
	var currentToken string

	runeSpace := rune(' ')
	runeNewline := rune('\n')
	runeSingleQuote := rune('\'')

	for _, c := range str {
		if c == runeSpace || c == runeNewline {
			if currentToken != "" {
				result = append(result, currentToken)
				currentToken = ""
			}
			continue
		}
		if c == runeSingleQuote {
			if currentToken != "" {
				return nil, fmt.Errorf("*** - SYSTEM::READ-EVAL-PRINT: variable %s has no value", currentToken)
			}
			result = append(result, string(c))
			continue
		}
		currentToken += string(c)
	}

	if currentToken != "" {
		result = append(result, currentToken)
	}
	return result, nil
}
