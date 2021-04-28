package controller

import (
	"fmt"
	"github.com/QQGoblin/device-watcher/pkg/client/clientset/versioned"
	"github.com/QQGoblin/device-watcher/pkg/client/informers/externalversions"
	v1beta12 "github.com/QQGoblin/device-watcher/pkg/client/listers/device/v1beta1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"time"
)

type nicController struct {
	// k8s 标准资源的client
	kubeClient kubernetes.Interface
	// device group 对应的 client
	deviceClient versioned.Interface
	// 通过Key获取Cache的对象
	nicLister v1beta12.NicLister
	// Controller启动时这个接口用来等待Informer的缓存同步
	nicSynced cache.InformerSynced
	// 工作队列，存储对象为 string(namespace/name)
	workqueue workqueue.RateLimitingInterface
	recorder  record.EventRecorder
}

func NewNicController(kubeClient kubernetes.Interface, deviceClient versioned.Interface, f externalversions.SharedInformerFactory) *nicController {

	// 构建对应的Informer
	nicInformer := f.Device().V1beta1().Nics()
	nicInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    nil,
		UpdateFunc: nil,
		DeleteFunc: nil,
	})

	return &nicController{
		kubeClient:   kubeClient,
		deviceClient: deviceClient,
		nicLister:    nicInformer.Lister(),
		nicSynced:    nicInformer.Informer().HasSynced, // 这里返回的是函数
		workqueue:    workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "nic"),
	}
}

func (c *nicController) Run(stopCh <-chan struct{}) error {
	defer func() {
		utilruntime.HandleCrash()
		c.workqueue.ShutDown()
		klog.V(0).Info("shutting down nic controller")
	}()

	klog.V(0).Info("starting nic controller")

	// 等待 Informer 同步缓存
	if !cache.WaitForCacheSync(stopCh, c.nicSynced) {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	// 启动一个工作协程，这里其实可以启动多个
	go wait.Until(c.runWorker, time.Second*5, stopCh)
	<-stopCh
	return nil
}

func (c *nicController) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *nicController) processNextWorkItem() bool {
	// workqueue 为空时阻塞
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		defer c.workqueue.Done(obj)
		// 工作队列中存储的Key为 namespace/name 的字符串
		var key string
		var ok bool
		if key, ok = obj.(string); !ok {
			// 工作队列中的Key不是string时，调用Forget删除该key
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// 执行调协函数，该函数中包含业务逻辑
		if err := c.reconcile(key); err != nil {
			// 调协处理异常时，重新将key入队
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		// 处理Key无异常，从工作队列中移除
		c.workqueue.Forget(obj)
		klog.Infof("Successfully synced %s:%s", "key", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

func (c *nicController) reconcile(key string) error {
	klog.Infoln(c.nicLister.Get(key))
	return nil
}
