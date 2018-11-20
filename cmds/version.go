package cmds

import (
	"fmt"

	"github.com/kooksee/kask/version"
	"github.com/spf13/cobra"
)

// VersionCmd ...
func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show Version Info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ktask version", version.Version)
			fmt.Println("ktask commit version", version.CommitVersion)
			fmt.Println("ktask build version", version.BuildVersion)
		},
	}
}
