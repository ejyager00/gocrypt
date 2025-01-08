/*
Copyright © 2025 Eric Yager
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ejyager00/gocrypt/internal/algorithms"
	"github.com/spf13/cobra"
)

var encrypt bool
var decrypt bool
var inFile string
var outFile string
var key string
var caesar bool
var vigenere bool

var rootCmd = &cobra.Command{
	Use:   "gocrypt",
	Short: "A simple message encryption CLA.",
	Long: `A simple message encryption CLA. Includes simple message ciphers using 
	caesar and vigenere strategies.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !encrypt && !decrypt {
			cmd.Help()
			return
		}

		if !caesar && !vigenere {
			cmd.Help()
			return
		}

		content, err := os.ReadFile(inFile)
		if err != nil {
			fmt.Println("Error reading input file:", err)
			return
		}

		var result string
		if caesar {
			shift := uint8(0)
			fmt.Sscan(key, &shift)
			if encrypt {
				result = algorithms.CaesarEncrypt(string(content), shift)
			} else {
				result = algorithms.CaesarDecrypt(string(content), shift)
			}
		} else {
			if encrypt {
				result = algorithms.VigenereEncrypt(string(content), key)
			} else {
				result = algorithms.VigenereDecrypt(string(content), key)
			}
		}

		err = os.WriteFile(outFile, []byte(result), 0644)
		if err != nil {
			fmt.Println("Error writing output file:", err)
			return
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&encrypt, "encrypt", "e", false, "Encrypt the input")
	rootCmd.Flags().BoolVarP(&decrypt, "decrypt", "d", false, "Decrypt the input")
	rootCmd.Flags().StringVarP(&inFile, "input", "i", "", "Input file path")
	rootCmd.Flags().StringVarP(&outFile, "output", "o", "", "Output file path")
	rootCmd.Flags().StringVarP(&key, "key", "k", "", "Encryption/decryption key")
	rootCmd.Flags().BoolVarP(&caesar, "caesar", "c", false, "Use Caesar cipher")
	rootCmd.Flags().BoolVarP(&vigenere, "vigenere", "v", false, "Use Vigenere cipher")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")
	rootCmd.MarkFlagRequired("key")
}
