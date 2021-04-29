package main

import (
	"github.com/QQGoblin/device-watcher/cmd/client"
	"github.com/QQGoblin/device-watcher/cmd/server"
	"github.com/QQGoblin/device-watcher/pkg/version"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"os"
)

func main() {

	if err := run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	//klog.InitFlags(nil)
	//dwlogger.InitLogs()
	//defer dwlogger.FlushLogs()

	cmd := &cobra.Command{
		Use:   "dw",
		Short: "watcher device on this node ",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			welcome()
		},
	}
	cmd.AddCommand(
		client.NewCmdClientStart(),
		server.NewCmdServerStart(),
	)

	return cmd.Execute()
}

func welcome() {
	klog.Infof("Starting device watcher ...")
	klog.Infof("Version Tag : %s", version.GetVersion())
	klog.Infof("GitCommint : %s", version.GetGitCommit())
}
