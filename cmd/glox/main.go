package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var hasError bool = false

type Token struct {
	tokenType string
	lexeme    string
	literal   string
	line      int
}

func newToken(tokenType string, lexeme string, literal string, line int) *Token {
	t := Token{tokenType: tokenType, lexeme: lexeme, literal: literal, line: line}
	return &t
}

func (t Token) toString() string {
	return t.tokenType + " " + t.lexeme + " " + t.literal
}

var hadError bool = false

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

	run(string(data))

	if hadError {
		os.Exit(65)
	}
}

func runPrompt() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		run(line)
		hadError = false
	}
}

func run(source string) {
	var tokens []string
	var errors []string
	line := 1

	for _, char := range source {

		charString := string(char)

		switch charString {
		case "(":
			tokens = append(tokens, "LEFT_PAREN"+" "+charString+" "+"null")
			break
		case ")":
			tokens = append(tokens, "RIGHT_PAREN"+" "+charString+" "+"null")
			break
		case "{":
			tokens = append(tokens, "LEFT_BRACE"+" "+charString+" "+"null")
			break
		case "}":
			tokens = append(tokens, "RIGHT_BRACE"+" "+charString+" "+"null")
			break
		case ",":
			tokens = append(tokens, "COMMA"+" "+charString+" "+"null")
			break
		case ".":
			tokens = append(tokens, "DOT"+" "+charString+" "+"null")
			break
		case "-":
			tokens = append(tokens, "MINUS"+" "+charString+" "+"null")
			break
		case "+":
			tokens = append(tokens, "PLUS"+" "+charString+" "+"null")
			break
		case ";":
			tokens = append(tokens, "SEMICOLON"+" "+charString+" "+"null")
			break
		case "/":
			tokens = append(tokens, "SLASH"+" "+charString+" "+"null")
			break
		case "*":
			tokens = append(tokens, "STAR"+" "+charString+" "+"null")
			break
		case "\n":
			line++
			break
		default:
			hasError = true
			errors = append(errors, "[line "+strconv.Itoa(line)+"] Error: Unexpected character: "+charString)
		}
	}

	for _, err := range errors {
		fmt.Fprint(os.Stderr, "\n", err)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

	fmt.Println("EOF  null")

	if hasError {
		os.Exit(65)
	}
}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d ] Error %s : %s", line, where, message)
	hadError = true
}
