/*
Copyright 2020 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	devicev1beta1 "github.com/QQGoblin/device-watcher/pkg/apis/device/v1beta1"
	versioned "github.com/QQGoblin/device-watcher/pkg/client/clientset/versioned"
	internalinterfaces "github.com/QQGoblin/device-watcher/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/QQGoblin/device-watcher/pkg/client/listers/device/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NicInformer provides access to a shared informer and lister for
// Nics.
type NicInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.NicLister
}

type nicInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNicInformer constructs a new informer for Nic type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNicInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNicInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNicInformer constructs a new informer for Nic type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNicInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DeviceV1beta1().Nics().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DeviceV1beta1().Nics().Watch(options)
			},
		},
		&devicev1beta1.Nic{},
		resyncPeriod,
		indexers,
	)
}

func (f *nicInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNicInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nicInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&devicev1beta1.Nic{}, f.defaultInformer)
}

func (f *nicInformer) Lister() v1beta1.NicLister {
	return v1beta1.NewNicLister(f.Informer().GetIndexer())
}
