/*
Copyright © 2025 Eric Yager
*/
package algorithms

import "github.com/ejyager00/gocrypt/internal/util"

const (
	upperOffset = byte(65)
	lowerOffset = byte(97)
	alphabetLen = byte(26)
)

func VigenereEncrypt(cleartext string, keyword string) string {
	return vigenereTransform(cleartext, keyword, subtractKey)
}
func VigenereDecrypt(ciphertext string, keyword string) string {
	return vigenereTransform(ciphertext, keyword, addKey)
}

func vigenereTransform(text string, keyword string, operation func(byte, byte, byte) byte) string {
	b := []byte(text)
	ku, kl, kLen := getKeys(keyword)

	for i, x := range b {
		if util.IsUpper(x) {
			b[i] = operation(x, ku[i%kLen], upperOffset)
		} else if util.IsLower(x) {
			b[i] = operation(x, kl[i%kLen], lowerOffset)
		}
	}
	return string(b)
}

func subtractKey(char byte, key byte, offset byte) byte {
	return ((char-offset)-(key-offset)+alphabetLen)%alphabetLen + offset
}

func addKey(char byte, key byte, offset byte) byte {
	return ((char-offset)+(key-offset))%alphabetLen + offset
}

func getKeys(k string) ([]byte, []byte, int) {
	ku := []byte(util.AllUpper(k))
	return ku, []byte(util.AllLower(k)), len(ku)
}
