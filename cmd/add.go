package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type Snippet struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Language  string `json:"language"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
}

var addCmd = &cobra.Command{
	Use:   "add <title> <language> <code>",
	Short: "Add a new code snippet",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		snippet := Snippet{
			ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
			Title:     strings.TrimSpace(args[0]),
			Language:  strings.TrimSpace(args[1]),
			Code:      strings.TrimSpace(args[2]),
			CreatedAt: time.Now().Format(time.RFC3339),
		}

		saveSnippet(snippet)
		fmt.Println("✅ Snippet saved!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func saveSnippet(s Snippet) {
	home, _ := os.UserHomeDir()
	dir := home + "/.hackabin"
	file := dir + "/snippets.json"

	_ = os.MkdirAll(dir, os.ModePerm)

	var snippets []Snippet

	data, err := os.ReadFile(file)
	if err == nil {
		_ = json.Unmarshal(data, &snippets)
	}

	snippets = append(snippets, s)

	newData, _ := json.MarshalIndent(snippets, "", "  ")
	_ = os.WriteFile(file, newData, 0644)
}
