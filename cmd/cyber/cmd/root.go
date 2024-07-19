package cmd

import (
	"log"

	"github.com/haormj/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "cyber",
	Short:   "cyber toolkit",
	Version: version.FullVersion(),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
