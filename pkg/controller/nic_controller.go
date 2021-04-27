package controller

//
//import (
//	"fmt"
//	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
//	"k8s.io/apimachinery/pkg/util/wait"
//	"k8s.io/client-go/tools/cache"
//	"k8s.io/client-go/util/workqueue"
//	"k8s.io/klog"
//	clientset "lqingcloud.cn/device-watcher/pkg/client/clientset/versioned"
//	devClient "lqingcloud.cn/device-watcher/pkg/client/clientset/versioned/typed/device/v1beta1"
//	"lqingcloud.cn/device-watcher/pkg/client/informers/externalversions"
//	"lqingcloud.cn/device-watcher/pkg/client/informers/externalversions/device/v1beta1"
//	"time"
//)
//
//type nicController struct {
//	nicClient   devClient.NicInterface
//	nicInformer v1beta1.NicInformer
//	// workqueue is a rate limited work queue. This is used to queue work to be
//	// processed instead of performing it as soon as a change happens. This
//	// means we can ensure we only process a fixed amount of resources at a
//	// time, and makes it easy to ensure we are never processing the same item
//	// simultaneously in two different workers.
//	workqueue workqueue.RateLimitingInterface
//}
//
//func NewNicController(c clientset.Clientset, f externalversions.SharedInformerFactory) *nicController {
//
//	f.Device().V1beta1().Nics().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
//		AddFunc:    nil,
//		UpdateFunc: nil,
//		DeleteFunc: nil,
//	})
//
//	return &nicController{
//		nicClient:   c.DeviceV1beta1().Nics(),
//		nicInformer: f.Device().V1beta1().Nics(),
//		workqueue:   workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "nic"),
//	}
//}
//
//func (c *nicController) Run(stopCh <-chan struct{}) error {
//	defer func() {
//		utilruntime.HandleCrash()
//		c.workqueue.ShutDown()
//		klog.V(0).Info("shutting down nic controller")
//	}()
//
//	klog.V(0).Info("starting nic controller")
//
//	if !cache.WaitForCacheSync(stopCh, c.nicInformer.Informer().GetController().HasSynced) {
//		return fmt.Errorf("failed to wait for caches to sync")
//	}
//
//	// 启动一个工作协程，这里其实可以启动多个
//	go wait.Until(c.runWorker, time.Second*5, stopCh)
//
//	<-stopCh
//	return nil
//}
//
//func (c *nicController) runWorker() {
//	for c.processNextWorkItem() {
//
//	}
//}
//
//func (c *nicController) processNextWorkItem() bool {
//	obj, shutdown := c.workqueue.Get()
//
//	if shutdown {
//		return false
//	}
//
//	// We wrap this block in a func so we can defer c.workqueue.Done.
//	err := func(obj interface{}) error {
//		// We call Done here so the workqueue knows we have finished
//		// processing this item. We also must remember to call Forget if we
//		// do not want this work item being re-queued. For example, we do
//		// not call Forget if a transient error occurs, instead the item is
//		// put back on the workqueue and attempted again after a back-off
//		// period.
//		defer c.workqueue.Done(obj)
//		var key string
//		var ok bool
//		// We expect strings to come off the workqueue. These are of the
//		// form namespace/name. We do this as the delayed nature of the
//		// workqueue means the items in the informer cache may actually be
//		// more up to date that when the item was initially put onto the
//		// workqueue.
//		if key, ok = obj.(string); !ok {
//			// As the item in the workqueue is actually invalid, we call
//			// Forget here else we'd go into a loop of attempting to
//			// process a work item that is invalid.
//			c.workqueue.Forget(obj)
//			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
//			return nil
//		}
//		// Run the reconcile, passing it the namespace/name string of the
//		// Foo resource to be synced.
//		if err := c.reconcile(key); err != nil {
//			// Put the item back on the workqueue to handle any transient errors.
//			c.workqueue.AddRateLimited(key)
//			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
//		}
//		// Finally, if no error occurs we Forget this item so it does not
//		// get queued again until another change happens.
//		c.workqueue.Forget(obj)
//		klog.Infof("Successfully synced %s:%s", "key", key)
//		return nil
//	}(obj)
//
//	if err != nil {
//		utilruntime.HandleError(err)
//		return true
//	}
//
//	return true
//}
//
//func (c *nicController) reconcile(key string) error {
//
//}
