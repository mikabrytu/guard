package bullet

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Bullet struct {
	Name     string
	instance *lifecycle.GameObject
	rect     utils.RectSpecs
	color    render.Color
	axis     int
	speed    int
	body     physics.RigidBody
}
