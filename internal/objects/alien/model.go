package alien

import (
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Alien struct {
	name  string
	rect  utils.RectSpecs
	color render.Color
}
