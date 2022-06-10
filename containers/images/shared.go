package images

import (
	"github.com/containers/storage"
	"github.com/containers/storage/pkg/unshare"
)

type sharedBuilder struct{}

func (builder *sharedBuilder) NewStorage() (store storage.Store, err error) {
	buildStoreOptions, err := storage.DefaultStoreOptions(unshare.IsRootless(), unshare.GetRootlessUID())
	return storage.GetStore(buildStoreOptions)
}
