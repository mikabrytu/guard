package shield

import (
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Shield struct {
	name     string
	rect     utils.RectSpecs
	color    render.Color
	sprite   *render.Sprite
	isSimple bool
}
