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
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
}

var addCmd = &cobra.Command{
	Use:   "add <title> <code|file>",
	Short: "Add a new code snippet (inline or from file)",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.TrimSpace(args[0])
		codeInput := strings.TrimSpace(args[1])
		var code string

		if fileContent, err := os.ReadFile(codeInput); err == nil {
			// If it's a valid file path, use file content
			code = string(fileContent)
		} else {
			// Otherwise, treat it as inline code
			code = codeInput
		}

		snippet := Snippet{
			ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
			Title:     title,
			Code:      code,
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
