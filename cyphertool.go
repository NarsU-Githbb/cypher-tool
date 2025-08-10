package main

import "fmt"

func main() {

	toEncrypt, encoding, message := getInput()
	var result string

	switch encoding {
	case "ROT13":
		result = encrypt_rot13(message)

	case "Reverse":
		result = Encrypt_reverse(message)

	case "Vowelskip":
		if toEncrypt == true {
			result = encrypt_vowel_skip(message)
		} else if toEncrypt == false {
			result = decrypt_wovel_skip(message)
		}
	}

	fmt.Println("Decrypted message using reverse:\n", result)
}
