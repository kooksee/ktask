package cmds

import (
	"github.com/kooksee/ktask/internal/config"
	"github.com/kooksee/ktask/internal/kts"
	"github.com/kooksee/ktask/internal/utils"
	"github.com/spf13/cobra"
)

func LogTaskCmd() *cobra.Command {
	var count = 100
	var logHttpCmd = func(cmd *cobra.Command) *cobra.Command {
		cmd.PersistentFlags().IntVar(&count, "count", count, "log task num")
		return cmd
	}

	return logHttpCmd(&cobra.Command{
		Use:   "log",
		Short: "log task generator",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			cfg := config.DefaultConfig()
			cfg.Init()

			task := &kts.Task{}
			for i := 0; i < count; i++ {
				m := task.Mock()
				utils.P(m)
				utils.MustNotError(utils.Retry(3, func() error {
					_, err := cfg.TaskPost(m.Encode())
					return err
				}))
			}

			return nil
		},
	})
}
