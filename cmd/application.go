package cmd

import (
	"github.com/spf13/cobra"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/config"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch"
)

var runapplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "run application",
	Run: func(cmd *cobra.Command, args []string) {
		runapplication()
	},
}

func runapplication() {

	cfg := config.NewFromEnv()
	application := app.New(cfg, scratch.New(cfg.Config, rootCmd.Version))

	if err := application.Run(); err != nil {
		panic(err)
	}

}
