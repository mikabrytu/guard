package alien

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func New(name string, rect utils.RectSpecs, color render.Color) *Alien {
	alien := &Alien{
		name:  name,
		rect:  rect,
		color: color,
	}

	lifecycle.Register(&lifecycle.GameObject{
		Render: alien.render,
	})

	return alien
}

func (a *Alien) render() {
	render.DrawRect(a.rect, a.color)
}
