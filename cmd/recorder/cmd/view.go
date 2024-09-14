package cmd

import (
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "cyber recorder view",
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
