package main

func Encrypt_reverse(s string) string {

	runes := []rune(s)
	var result string

	for i := 0; i < len(s); i++ {

		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] = 'a' + ('z' - runes[i])
		}
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] = 'A' + ('Z' - runes[i])
		}

	}
	result = string(runes)
	return result
}
