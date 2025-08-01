package player

import (
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Player struct {
	name     string
	rect     utils.RectSpecs
	color    render.Color
	sprite   *render.Sprite
	speed    int
	axis     int
	moving   bool
	isSimple bool
}
