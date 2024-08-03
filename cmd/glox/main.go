package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	COMMA       = ","
	DOT         = "."
	MINUS       = "-"
	PLUS        = "+"
	SEMICOLON   = ";"
	SLASH       = "/"
	STAR        = "*"
)

func main() {

	input := os.Args

	if len(input) > 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./glox tokenize <filename>")
		fmt.Fprintln(os.Stderr, "OR")
		fmt.Fprintln(os.Stderr, "Usage: ./glox [script]")
		os.Exit(65)
	} else if len(input) == 3 {
		command := input[1]
		if command != "tokenize" {
			fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
			os.Exit(1)
		}
		runFile(input[2])
	} else {
		runPrompt()
	}
}

func runFile(filePath string) {

	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file from path %s\n", filePath)
		os.Exit(1)
	}

	var errors []string
	var lexemes []string

	for _, char := range string(data) {

		charString := string(char)

		switch charString {
		case LEFT_PAREN:
			lexemes = append(lexemes, charString)
		case RIGHT_PAREN:
			lexemes = append(lexemes, charString)
		case LEFT_BRACE:
			lexemes = append(lexemes, charString)
		case RIGHT_BRACE:
			lexemes = append(lexemes, charString)
		case COMMA:
			lexemes = append(lexemes, charString)
		case DOT:
			lexemes = append(lexemes, charString)
		case MINUS:
			lexemes = append(lexemes, charString)
		case PLUS:
			lexemes = append(lexemes, charString)
		case SEMICOLON:
			lexemes = append(lexemes, charString)
		case SLASH:
			lexemes = append(lexemes, charString)
		case STAR:
			lexemes = append(lexemes, charString)
		default:
			errors = append(errors, charString)
		}
	}

	lexemes = append(lexemes, "EOF")

	for _, error := range errors {
		fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %s\n", error)
	}

	for _, lexeme := range lexemes {
		fmt.Println(lexeme)
	}
}

func runPrompt() {
	fmt.Println("Running prompt")
}
