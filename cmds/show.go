package cmds

import (
	"github.com/kooksee/ktask/internal/cnst"
	"github.com/kooksee/ktask/internal/utils"
	"github.com/spf13/cobra"
)

func ShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "show",
		Aliases: []string{"v"},
		Short:   "Show Version Info",
		Run: func(cmd *cobra.Command, args []string) {
			utils.P(cnst.Consumer, cnst.TaskTx, cnst.TaskStatus, cnst.Oss)
		},
	}
}
