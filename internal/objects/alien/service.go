package alien

import (
	"littlejumbo/guard/config"

	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func New(name string, rect utils.RectSpecs, color render.Color) *Alien {
	alien := &Alien{
		Name:     name,
		rect:     rect,
		color:    color,
		axis:     math.Vector2{X: -1, Y: 0},
		isSimple: true,
	}

	lifecycle.Register(&lifecycle.GameObject{
		Render:  alien.render,
		Destroy: alien.destroy,
	})

	return alien
}

func (a *Alien) SetSprite(path string) {
	a.isSimple = false

	a.sprite = render.NewSprite(a.Name, path)
	a.sprite.Init(a.rect)
}

func (a *Alien) SetStep(step int) {
	a.step = step
}

func (a *Alien) MoveStep() {
	a.rect.PosX += a.axis.X * a.step
	a.sprite.UpdateRect(a.rect)
}

func (a *Alien) InvertDirectionX() {
	a.axis.X *= -1
}

func (a *Alien) IsAtScreenEdge() bool {
	return (a.rect.PosX+a.rect.Width) > config.SCREEN_SIZE.X-config.SCREEN_OFFSET.X || a.rect.PosX < config.SCREEN_OFFSET.X
}

func (a *Alien) render() {
	if a.isSimple {
		render.DrawRect(a.rect, a.color)
	}
}

func (a *Alien) destroy() {
	a.sprite.ClearSprite()
}
