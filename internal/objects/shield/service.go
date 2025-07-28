package shield

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func New(name string, rect utils.RectSpecs, color render.Color) *Shield {
	shield := &Shield{
		name:  name,
		rect:  rect,
		color: color,
	}

	lifecycle.Register(&lifecycle.GameObject{
		Render: shield.render,
	})

	return shield
}

func (p *Shield) render() {
	render.DrawRect(p.rect, p.color)
}
