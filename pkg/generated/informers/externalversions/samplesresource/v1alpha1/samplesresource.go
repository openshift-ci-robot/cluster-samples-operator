// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	samplesresourcev1alpha1 "github.com/openshift/cluster-samples-operator/pkg/apis/samplesresource/v1alpha1"
	versioned "github.com/openshift/cluster-samples-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift/cluster-samples-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift/cluster-samples-operator/pkg/generated/listers/samplesresource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SamplesResourceInformer provides access to a shared informer and lister for
// SamplesResources.
type SamplesResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SamplesResourceLister
}

type samplesResourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSamplesResourceInformer constructs a new informer for SamplesResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSamplesResourceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSamplesResourceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSamplesResourceInformer constructs a new informer for SamplesResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSamplesResourceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SamplesresourceV1alpha1().SamplesResources().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SamplesresourceV1alpha1().SamplesResources().Watch(options)
			},
		},
		&samplesresourcev1alpha1.SamplesResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *samplesResourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSamplesResourceInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *samplesResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&samplesresourcev1alpha1.SamplesResource{}, f.defaultInformer)
}

func (f *samplesResourceInformer) Lister() v1alpha1.SamplesResourceLister {
	return v1alpha1.NewSamplesResourceLister(f.Informer().GetIndexer())
}