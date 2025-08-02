package alien

import (
	"littlejumbo/guard/config"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func New(name string, rect utils.RectSpecs, color render.Color) *Alien {
	alien := &Alien{
		Name:     name,
		rect:     rect,
		color:    color,
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
	if a.axis.X == 0 {
		a.axis.X = 1
	}

	a.rect.PosX += a.axis.X * a.step
	a.sprite.UpdateRect(a.rect)

	if (a.rect.PosX + a.rect.Width) >= config.SCREEN_SIZE.X-config.SCREEN_OFFSET.X {
		events.Emit(config.EVENT_ALIEN_AT_SCREEN_BOUNDARY)
	}
}

func (a *Alien) InvertDirectionX() {
	a.axis.X *= -1
}

func (a *Alien) render() {
	if a.isSimple {
		render.DrawRect(a.rect, a.color)
	}
}

func (a *Alien) destroy() {
	a.sprite.ClearSprite()
}
