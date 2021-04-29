package server

import (
	clientset "github.com/QQGoblin/device-watcher/pkg/client/clientset/versioned"
	"github.com/QQGoblin/device-watcher/pkg/client/informers/externalversions"
	"github.com/QQGoblin/device-watcher/pkg/controller"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
	"time"
)

//var target string
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start grpc server for this node ",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

var (
	kubeconfig string
)

func NewCmdServerStart() *cobra.Command {
	return serverCmd
}

func init() {
	serverCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "~/.kube/config", "kubernetes config")
}

func start() {
	stopCh := signals.SetupSignalHandler()
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		klog.Fatal("Error building kubeconfig clientset: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatal("Error create kubernetes client %s", err.Error())
	}

	deviceClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatal("Error create device.lqingcloud.cn client %s", err.Error())
	}

	deviceInformerFactory := externalversions.NewSharedInformerFactory(deviceClient, time.Second*30)

	nicController := controller.NewNicController(kubeClient, deviceClient, deviceInformerFactory)

	go deviceInformerFactory.Start(stopCh)

	if err = nicController.Run(stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}
