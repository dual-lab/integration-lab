package images

import (
	"github.com/containers/buildah"
	"github.com/dual-lab/integration-lab/containers/common"
)

type helloworldBuilder struct{
  sharedBuilder
}

func NewHelloworldBuilder() common.BuilderFile {
	return &helloworldBuilder{}
}

func (builder *helloworldBuilder) NewBuildOpts() (opts buildah.BuilderOptions, err error) {
  return buildah.BuilderOptions{
    FromImage: "helloworld",
  }, nil
}

func (builder *helloworldBuilder) Compose(b *buildah.Builder) (err error) {
	return
}

func (builder *helloworldBuilder) Name() (name string) {
	return "hello-world-integration"
}
