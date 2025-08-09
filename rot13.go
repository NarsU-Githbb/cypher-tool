package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func encrypt_rot13(s string) string {
	var result strings.Builder
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			result.WriteRune('A' + (r-'A'+13)%26)
		case r >= 'a' && r <= 'z':
			result.WriteRune('a' + (r-'a'+13)%26)
		default:
			result.WriteRune(r)
		}
	}
	return result.String()
}

func decrypt_rot13(s string) string {
	return encrypt_rot13(s)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("ROT13 Encryption/Decryption Tool")
	fmt.Print("Enter your message: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	encrypted := encrypt_rot13(input)
	decrypted := decrypt_rot13(encrypted)

	fmt.Println("Encrypted:", encrypted)
	fmt.Println("Decrypted:", decrypted)
}
