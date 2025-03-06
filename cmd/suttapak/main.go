package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/suttapak/mememe/internal/cors"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "suttapak",
		Short: "Suttapak is a Golang code generator",
	}

	var generateCmd = &cobra.Command{
		Use:   "g [name]",
		Short: "Generate service, repository, and controller",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := strings.ToLower(args[0])
			if err := cors.Create(name); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(generateCmd)
	rootCmd.Execute()
}
