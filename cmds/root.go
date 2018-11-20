package cmds

import (
	"github.com/kooksee/ktask/internal/config"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	var root = func(cmd *cobra.Command) *cobra.Command {
		cfg := config.DefaultConfig()
		cmd.PersistentFlags().BoolVar(&cfg.Debug, "debug", cfg.Debug, "debug")
		return cmd
	}

	return root(&cobra.Command{
		Use:   "ktask",
		Short: "分布式任务处理器",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			return nil
		},
	})
}
