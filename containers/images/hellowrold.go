package images

import "github.com/dual-lab/integration-lab/containers/common/BuilderFile"

type helloworldBuilder struct{}

func NewHelloworldBuilder() *BuilderFile {
  return &helloworldBuilder{}
}
