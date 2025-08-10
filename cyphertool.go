package main

import "fmt"

func main() {

	toEncrypt, encoding, message := getInput()
	var result string

	switch encoding {
	case "ROT13":
		result = encrypt_rot13(message)

	case "reverse":
		result = Encrypt_reverse(message)

	case "vowelskip":
		if toEncrypt == true {
			result = encrypt_vowel_skip(message)
		} else if toEncrypt == false {
			result = decrypt_vowel_skip(message)
		}
	}

	fmt.Print("Decrypted message using ", encoding)
	fmt.Print(":\n", result)
	fmt.Print("\n")
}
