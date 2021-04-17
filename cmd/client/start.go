package client

import (
	"github.com/QQGoblin/device-watcher/pkg/grpc"
	"github.com/spf13/cobra"
)

func NewCmdStart() *cobra.Command {

	//var target string
	getCmd := &cobra.Command{
		Use:   "client",
		Short: "start grpc server for this node ",
		Run: func(cmd *cobra.Command, args []string) {
			grpc.Start()
		},
	}
	return getCmd

}
