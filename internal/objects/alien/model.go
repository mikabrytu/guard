package alien

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Alien struct {
	Name     string
	instance *lifecycle.GameObject
	rect     utils.RectSpecs
	color    render.Color
	sprite   *render.Sprite
	axis     math.Vector2
	step     math.Vector2
	Body     physics.RigidBody
	score    int
	isSimple bool
}
