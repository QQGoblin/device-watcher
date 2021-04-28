package server

import "github.com/spf13/cobra"

func NewCmdServerStart() *cobra.Command {

	//var target string
	getCmd := &cobra.Command{
		Use:   "server",
		Short: "start grpc server for this node ",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return getCmd

}
