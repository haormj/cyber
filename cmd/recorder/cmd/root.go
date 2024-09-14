package cmd

import (
	"log"

	"github.com/haormj/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "recorder",
	Short:   "recorder toolkit",
	Version: version.FullVersion(),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
