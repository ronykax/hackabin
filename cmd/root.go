package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var note string

var rootCmd = &cobra.Command{
	Use:   "hackabin [file]",
	Short: "Upload a code file with context",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}

		fmt.Println("Note:", note)
		fmt.Println("Contents of", file+":\n")
		fmt.Println(string(content))
	},
}

func init() {
	rootCmd.Flags().StringVar(&note, "note", "", "Optional note about the code snippet")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
