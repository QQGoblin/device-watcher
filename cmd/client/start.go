package client

import (
	"github.com/QQGoblin/device-watcher/pkg/grpc"
	"github.com/spf13/cobra"
)

var clienCmd = &cobra.Command{
	Use:   "client",
	Short: "start grpc server for this node ",
	Run: func(cmd *cobra.Command, args []string) {
		grpc.Start()
	},
}

func NewCmdClientStart() *cobra.Command {
	return clienCmd
}
