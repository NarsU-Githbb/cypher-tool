package main

import (
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
