package lister

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	workv1 "open-cluster-management.io/api/work/v1"
	"open-cluster-management.io/sdk-go/pkg/cloudevents/generic/types"
	"open-cluster-management.io/sdk-go/pkg/cloudevents/work/store"
)

// WatcherStoreLister list the ManifestWorks from the WorkClientWatcherStore.
type WatcherStoreLister struct {
	store store.WorkClientWatcherStore
}

func NewWatcherStoreLister(store store.WorkClientWatcherStore) *WatcherStoreLister {
	return &WatcherStoreLister{
		store: store,
	}
}

// List returns the ManifestWorks from the WorkClientWatcherCache with list options.
func (l *WatcherStoreLister) List(options types.ListOptions) ([]*workv1.ManifestWork, error) {
	return l.store.List(options.ClusterName, metav1.ListOptions{})
}
