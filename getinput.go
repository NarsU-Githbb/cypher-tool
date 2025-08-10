package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() (toEncrypt bool, encoding string, message string) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Print("Select operation (1/2): \n 1. Encrypt. \n 2. Decrypt. \n")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			toEncrypt = true
		case "2":
			toEncrypt = false
		default:
			fmt.Print("Input is invalid. Please try again with patience and courage!\n")
			continue
		}
		break
	}


	fmt.Print("Select operation (2/2): \n 1. ROT13. \n 2. Reverse. \n 3.Vowelskip. \n")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			encoding = "ROT13"
		case "2":
			encoding = "Reverse"
		case "3":
			encoding = "Vowelskip"
		default:
			fmt.Print("Input is invalid. Please try again with patience and courage!\n")
			continue
		}
		break
	}


	fmt.Print("Enter the message: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	

	message = input
	return toEncrypt, encoding, message
}
