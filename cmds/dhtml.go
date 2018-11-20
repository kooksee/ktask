package cmds

import (
	"github.com/kooksee/ktask/internal/config"
	"github.com/kooksee/ktask/internal/dhtml"
	"github.com/kooksee/ktask/internal/utils"
	"github.com/spf13/cobra"
)

func DHtmlCmd() *cobra.Command {
	var addr = ":8082"
	var sleepTime = 3
	var handle = func(cmd *cobra.Command) *cobra.Command {
		cmd.PersistentFlags().StringVar(&addr, "addr", addr, "addr")
		cmd.PersistentFlags().IntVar(&sleepTime, "sleep_time", sleepTime, "sleep time")
		return cmd
	}

	return handle(&cobra.Command{
		Use:   "dhtml",
		Short: "爬取动态的html",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			config.DefaultConfig().Init()

			cfg := dhtml.NewConfig()
			cfg.SleepTime = sleepTime
			cfg.Init()

			utils.P(cfg)
			return cfg.Run(addr)
		},
	})
}
