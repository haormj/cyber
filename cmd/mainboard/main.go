package main

import (
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/mainboard"
	"github.com/haormj/cyber/state"
	"github.com/haormj/version"
	"github.com/spf13/cobra"
)

var mainboardCmd = &cobra.Command{
	Use:     "mainboard",
	Short:   "mainboard",
	Version: version.FullVersion(),
	Run: func(cmd *cobra.Command, args []string) {
		dagConfs, err := cmd.Flags().GetStringSlice("dag_conf")
		if err != nil {
			log.Logger.Error("get dag_conf error", "err", err)
			return
		}

		moduleController := mainboard.NewModuleController(dagConfs)
		if !moduleController.Init() {
			moduleController.Clear()
			log.Logger.Error("module start error")
			return
		}

		state.WaitForShutdown()
		moduleController.Clear()
		log.Logger.Info("exit mainboard")
	},
}

func init() {
	mainboardCmd.Flags().StringSliceP("dag_conf", "d", nil, "module dag config file")
}

func main() {
	if err := mainboardCmd.Execute(); err != nil {
		log.Logger.Error("mainboardCmd.Execute error", "err", err)
	}
}
