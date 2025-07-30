package alien

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func New(name string, rect utils.RectSpecs, color render.Color) *Alien {
	alien := &Alien{
		name:     name,
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

	a.sprite = render.NewSprite(a.name, path)
	a.sprite.Init(a.rect)
}

func (a *Alien) render() {
	if a.isSimple {
		render.DrawRect(a.rect, a.color)
	}
}

func (a *Alien) destroy() {
	a.sprite.ClearSprite()
}
