package main

import (
	"github.com/kooksee/ktask/cmds"
)

func main() {

	rootCmd := cmds.RootCmd()
	rootCmd.AddCommand(
		cmds.VersionCmd(),
		cmds.HttpProxyExampleCmd(),
		cmds.LogTaskCmd(),
		cmds.ShowCmd(),
		cmds.DHtmlCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}
