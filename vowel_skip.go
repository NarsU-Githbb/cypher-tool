package main

import (
	"fmt"
	"strings"
)

func encrypt_vowel_skip(s string) string {

	vowelMap := map[string]string{

		"a": "e",
		"e": "i",
		"o": "u",
		"u": "y",
		"y": "a",
	}

	length := len(s)
	parts := strings.Split(s, "")
	var result string

	for i := 0; i < length; i++ {
		if val, ok := vowelMap[parts[i]]; ok {
			parts[i] = val

		}
	}
	result = strings.Join(parts, "")

	return result

}

func decrypt_wovel_skip(s string) string {

	vowelMap := map[string]string{

		"e": "a",
		"i": "e",
		"u": "o",
		"y": "u",
		"a": "y",
	}

	parts := strings.Split(s, "")
	var answer string

	for i := 0; i < len(parts); i++ {
		if val, ok := vowelMap[parts[i]]; ok {
			parts[i] = val

		}
	}
	answer = strings.Join(parts, "")

	return answer

}

func main() {

	encrypting := encrypt_vowel_skip("abce")
	decrypting := decrypt_wovel_skip("ebci")
	fmt.Println(encrypting)
	fmt.Println(decrypting)

}
