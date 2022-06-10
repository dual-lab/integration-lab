package common

import (
	"github.com/containers/buildah"
	"github.com/containers/storage"
)

// Interface that wraps all the methods needed to build an OCI image
type BuilderFile interface {
	// NewStorage function used to obtain the buildah storage, requied to
	// build a OCI image
	NewStorage() (store storage.Store, err error)
	// NewBuildOpts function used to obtain the build options for the specific
	// image
	NewBuildOpts() (opts buildah.BuilderOptions, err error)
	// Compose function used to run all the instruction tha compose the whole
	// images, like all the instructions in a dockerfile
	Compose(b *buildah.Builder) (err error)
  // Image name
  Name()(name string)
}
