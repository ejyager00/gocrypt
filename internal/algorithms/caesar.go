/*
Copyright © 2025 Eric Yager
*/
package algorithms

import "github.com/ejyager00/gocrypt/internal/util"

func CaesarEncrypt(cleartext string, shift uint8) string {
	b := []byte(cleartext)
	for i, x := range b {
		if util.IsUpperCase(x) {
			b[i] = ((x-65)+shift)%26 + 65
		} else if util.IsLowerCase(x) {
			b[i] = ((x-97)+shift)%26 + 97
		} else {
			b[i] = x
		}
	}
	return string(b)
}

func CaesarDecrypt(ciphertext string, shift uint8) string {
	shift = 26 - (shift % 26)
	return CaesarEncrypt(ciphertext, shift)
}
