package main

import (
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"lqingcloud.cn/device-watcher/cmd/client"
	dwlogger "lqingcloud.cn/device-watcher/pkg/logs"
	"lqingcloud.cn/device-watcher/pkg/version"
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
