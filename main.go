package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// The program assumes that the input is well-formed, i.e. that either an argument
// is provided that specifies the path of a file to read in, or that the
// input is through STDIN.
func main() {
	var input io.Reader
	var err error

	if len(os.Args) > 1 { // Input is name of a file to read in
		inputPath := os.Args[1]
		input, err = os.Open(inputPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to open input file", err)
			return
		}
	} else { // assume that otherwise, input is from STDIN
		input = bufio.NewReader(os.Stdin)
	}
	output := processInput(input)
	if output != nil {
		sort.Strings(output)
		for _, line := range output {
			fmt.Println(line)
		}
	}
	return
}

// processInput takes in an input file and loops through the lines, processing
// each command. It stores the accounts in a map where keys are the names associated
// with the account (e.g. Bob: Account{Balance: 400, Limit: 500, Error: false}) and
// calls Credit and Charge on those accounts as needed.
// It assumes each line is well-formed, i.e. no malformed commands, space delimited
// arguments, and that the first argument is a command, followed by a name and then
// any extra args needed to execute the command.
func processInput(input io.Reader) []string {
	output := []string{}
	scanner := bufio.NewScanner(input)

	var accounts = map[string]*Account{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		inputStrings := strings.Split(line, " ")

		// assume well-formed input
		command := inputStrings[0]
		name := inputStrings[1]
		acc, accountExists := accounts[name]

		switch command {
		case "Add":
			if accountExists {
				fmt.Printf("Error: account already exists under name %s\n", name)
				continue
			}
			cardNum := inputStrings[2]
			limit := inputStrings[3]
			acc := createAccount(cardNum, limit)
			accounts[name] = acc
		case "Credit":
			if !accountExists {
				fmt.Printf("Error: no account exists under name %s\n", name)
				continue
			}
			amount := inputStrings[2]
			acc.Credit(amount)
		case "Charge":
			if !accountExists {
				fmt.Printf("Error: no account exists under name %s\n", name)
				continue
			}
			amount := inputStrings[2]
			acc.Charge(amount)
		default:
			// this shouldn't happen since input is guaranteed to be well-formed
			fmt.Printf("Error: unknown command %s\n", command)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error processing input", err)
		return nil
	}

	for name, acc := range accounts {
		if acc.Error {
			output = append(output, fmt.Sprintf("%s: error", name))
		} else {
			output = append(output, fmt.Sprintf("%s: %d", name, acc.Balance))
		}
	}
	return output
}
