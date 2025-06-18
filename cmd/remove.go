package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <id>",
	Short: "Remove a snippet by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		snippets := loadSnippets()

		index := -1
		for i, s := range snippets {
			if s.ID == id {
				index = i
				break
			}
		}

		if index == -1 {
			fmt.Println("❌ Snippet not found.")
			return
		}

		snippets = append(snippets[:index], snippets[index+1:]...)
		saveAllSnippets(snippets)

		fmt.Println("🗑️ Snippet removed.")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func saveAllSnippets(snippets []Snippet) {
	home, _ := os.UserHomeDir()
	file := home + "/.hackabin/snippets.json"

	newData, _ := json.MarshalIndent(snippets, "", "  ")
	_ = os.WriteFile(file, newData, 0644)
}
