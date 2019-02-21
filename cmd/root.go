package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/zloy2005/webshop/cmd/server"
)

var RootCmd = &cobra.Command{
	Use: "webshop",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	RootCmd.AddCommand(server.Cmd)
}
