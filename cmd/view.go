package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view <id>",
	Short: "View a snippet by ID",
	Long:  "View a snippet by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		snippets := loadSnippets()

		if len(snippets) == 0 {
			ascii := loadASCIISticker("stickers/orphmoji_scared.txt")
			fmt.Println("❌ No snippets found.")
			fmt.Println(ascii)
			return
		}

		for _, snip := range snippets {
			if snip.ID == id {
				printSnippet(snip.ID, snip.Title, snip.Code, snip.CreatedAt)
				fmt.Println("Copied code to your clipboard!")
				clipboard.WriteAll(snip.Code)
				return
			}
		}

		ascii := loadASCIISticker("stickers/orphmoji_confused.txt")
		fmt.Println("❌ Snippet not found.")
		fmt.Println(ascii)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
