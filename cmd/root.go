/*
Copyright © 2025 Eric Yager
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocrypt",
	Short: "A simple message encryption CLA.",
	Long: `A simple message encryption CLA. Includes simple message ciphers using 
	caesar and vigenere strategies.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
