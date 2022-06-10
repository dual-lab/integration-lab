package common

import (
	"context"

	"github.com/containers/buildah"
	"github.com/containers/storage/pkg/unshare"
  is "github.com/containers/image/v5/storage"
)

// Main function that build the images using
// buildah go API
func BuildMain(builderFile BuilderFile) (err error) {
	if buildah.InitReexec() {
		return
	}
	unshare.MaybeReexecUsingUserNamespace(false)

	buildStore, err := builderFile.NewStorage()
	if err != nil {
		return
	}
	defer buildStore.Shutdown(false)

	buildOptions, err := builderFile.NewBuildOpts()
	if err != nil {
		return
	}
  builder, err := buildah.NewBuilder(context.TODO(), buildStore, buildOptions)
  if err != nil {
    return
  }
  defer builder.Delete()

  if err = builderFile.Compose(builder); err !=nil {
    return
  }
  imageRef, err := is.Transport.ParseStoreReference(buildStore, builderFile.Name())
  if err != nil {
    return
  }
  _, _, _, err = builder.Commit(context.TODO(), imageRef, buildah.CommitOptions{})
  if err != nil {
    return
  }
	return
}
