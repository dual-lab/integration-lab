package common

import (
	"context"

	"github.com/containers/buildah"
	is "github.com/containers/image/v5/storage"
	"github.com/sirupsen/logrus"
)

// Main function that build the images using
// buildah go API
func BuildMain(builderFile BuilderFile) (err error) {
  
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
  logWriter := logrus.StandardLogger().Writer()
  _, _, _, err = builder.Commit(context.TODO(), imageRef, buildah.CommitOptions{
    ReportWriter: logWriter,
  })

  defer logWriter.Close()
  if err != nil {
    return
  }
	return
}
