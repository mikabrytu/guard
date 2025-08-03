package alien

import (
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Alien struct {
	Name     string
	rect     utils.RectSpecs
	color    render.Color
	sprite   *render.Sprite
	axis     math.Vector2
	step     math.Vector2
	isSimple bool
}
