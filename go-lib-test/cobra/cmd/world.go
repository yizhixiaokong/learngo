/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// worldCmd represents the world command
var worldCmd = &cobra.Command{
	Use:   "world",
	Args:  cobra.MaximumNArgs(1), // accept at most 1 argument
	Short: "a demo sub command for cobra",
	Long:  `a demo sub command for cobra, accept a string parameter, print hello xxx to console, default print hello world.`,
	Run: func(cmd *cobra.Command, args []string) {
		upper, _ := cmd.Flags().GetBool("upper")
		world := "world"
		if len(args) > 0 {
			world = args[0]
		}

		if upper {
			world = strings.ToUpper(world)
		}

		cmd.Printf("hello %s!\n", world)
	},
}

func init() {
	helloCmd.AddCommand(worldCmd)

	// -U, --upper, a flag to print world in upper case
	worldCmd.Flags().BoolP("upper", "U", false, "print world in upper case")
}
