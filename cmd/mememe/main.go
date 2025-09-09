package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/suttapak/mememe/internal/cors"
)

// Version is manually updated when creating a new tag
var Version = "v0.1.4"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mememe",
		Short: "mememe is a Golang code generator",
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
			if err := cors.Overwrite(name); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var generateRepositoryCmd = &cobra.Command{
		Use:   "gr [name]",
		Short: "Generate repository",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			name := strings.ToLower(args[0])
			if err := cors.CreateRepository(name); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if err := cors.OverwriteRepository(name); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var cleanCmd = &cobra.Command{
		Use:   "c [name]",
		Short: "Clean service, repository, and controller",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := cors.Clean(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the current version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("mememe version:", Version)
		},
	}

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(generateRepositoryCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}
