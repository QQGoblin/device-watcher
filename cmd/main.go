package main

import (
	"github.com/QQGoblin/device-watcher/cmd/client"
	dwlogger "github.com/QQGoblin/device-watcher/pkg/logs"
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
	dwlogger.InitLogs()
	defer dwlogger.FlushLogs()

	cmd := &cobra.Command{
		Use:   "dw",
		Short: "watcher device on this node ",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			welcome()
		},
	}
	cmd.AddCommand(
		client.NewCmdStart(),
	)

	return cmd.Execute()
}

func welcome() {
	klog.Info("Starting device watcher client...")
	klog.Info("Version Tag : %s", version.GetVersion())
	klog.Info("GitCommint : %s", version.GetGitCommit())
}
