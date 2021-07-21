package jumper_transform

import (
	"github.com/jumperzq86/jumper_transform/impl/transform"
	"github.com/jumperzq86/jumper_transform/interf"
)

func Newtransform() interf.Transform {
	var tf transform.Transform
	return &tf
}
