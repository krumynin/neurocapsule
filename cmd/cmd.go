package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	Version = "" // Will be set by -ldflags during compile time
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "neurocapsule",
	Short: "Нейрокапсула - сервис для подбора базового гардероба в один клик",
	Long:  `A service for brewing magical eerie enchanted mysterious otherwordly spellbinding potions.`,
}

func init() {
	cfgFile := os.Getenv("NEUROCAPSULE_CONF")
	if cfgFile != "" {
		err := godotenv.Load(cfgFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	rootCmd.AddCommand(runapplicationCmd)
	rootCmd.Version = Version
}

func Execute() {
	if len(os.Args[1:]) == 0 {
		// run server by default
		os.Args = append(os.Args, runapplicationCmd.Use)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
