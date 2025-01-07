/*
Copyright © 2025 Eric Yager
*/
package util

func IsUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func IsLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func AllUpper(text string) string {
	b := []byte(text)
	for i, x := range b {
		if x >= 'a' && x <= 'z' {
			b[i] = x - 32
		} else {
			b[i] = x
		}
	}
	return string(b)
}

func AllLower(text string) string {
	b := []byte(text)
	for i, x := range b {
		if x >= 'A' && x <= 'Z' {
			b[i] = x + 32
		} else {
			b[i] = x
		}
	}
	return string(b)
}
