package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved snippets",
	Run: func(cmd *cobra.Command, args []string) {
		snippets := loadSnippets()
		if len(snippets) == 0 {
			fmt.Println("No snippets found.")
			return
		}

		for _, s := range snippets {
			createdAt, _ := time.Parse(time.RFC3339, s.CreatedAt)
			fmt.Printf("🆔 %s\n📌 %s [%s]\n📅 %s\n\n",
				s.ID, s.Title, s.Language, createdAt.Format("Jan 2, 2006 3:04PM"))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func loadSnippets() []Snippet {
	home, _ := os.UserHomeDir()
	file := home + "/.hackabin/snippets.json"

	data, err := os.ReadFile(file)
	if err != nil {
		return []Snippet{}
	}

	var snippets []Snippet
	_ = json.Unmarshal(data, &snippets)
	return snippets
}
