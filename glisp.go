package main

import (
	"bufio"
	"fmt"
	"os"
)

type Atom interface{}

type Symbol string

type Number float64

type String string

type DotPair struct {
	car interface{}
	cdr interface{}
}

type number struct {
	value int
}

func (d *DotPair) Car() interface{} {
	return d.car
}

func (d *DotPair) Cdr() interface{} {
	return d.cdr
}

func cons(car interface{}, cdr interface{}) *DotPair {
	return &DotPair{car, cdr}
}

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

	// 改行文字がない場合も最後のトークンを追加する
	if currentToken != "" {
		result = append(result, currentToken)
	}
	return result, nil
}
