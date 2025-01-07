/*
Copyright © 2025 Eric Yager
*/
package util

func IsUpperCase(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func IsLowerCase(b byte) bool {
	return b >= 'a' && b <= 'z'
}
