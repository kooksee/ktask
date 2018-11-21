package main

import (
	"github.com/kooksee/ktask/cmds"
	"github.com/kooksee/ktask/internal/utils"
)

func main() {

	rootCmd := cmds.RootCmd()
	rootCmd.AddCommand(
		cmds.VersionCmd(),
		cmds.HttpProxyExampleCmd(),
		cmds.ShowCmd(),
		cmds.DHtmlCmd(),
	)

	utils.MustNotError(rootCmd.Execute())
}
