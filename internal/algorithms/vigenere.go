/*
Copyright © 2025 Eric Yager
*/
package algorithms

import "github.com/ejyager00/gocrypt/internal/util"

func VigenereEncrypt(cleartext string, keyword string) string {
	b := []byte(cleartext)
	ku, kl, k_len := getKeys(keyword)
	for i, x := range b {
		if util.IsUpper(x) {
			b[i] = ((x-65)-(ku[i%k_len]-65)+26)%26 + 65
		} else if util.IsLower(x) {
			b[i] = ((x-97)-(kl[i%k_len]-97)+26)%26 + 97
		} else {
			b[i] = x
		}
	}
	return string(b)
}

func VigenereDecrypt(ciphertext string, keyword string) string {
	b := []byte(ciphertext)
	ku, kl, k_len := getKeys(keyword)
	for i, x := range b {
		if util.IsUpper(x) {
			b[i] = ((x-65)+(ku[i%k_len]-65))%26 + 65
		} else if util.IsLower(x) {
			b[i] = ((x-97)+(kl[i%k_len]-97))%26 + 97
		} else {
			b[i] = x
		}
	}
	return string(b)
}

func getKeys(k string) ([]byte, []byte, int) {
	ku := []byte(util.AllUpper(k))
	return ku, []byte(util.AllLower(k)), len(ku)
}
