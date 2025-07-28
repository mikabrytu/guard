package player

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func New(name string, rect utils.RectSpecs, color render.Color) *Player {
	player := &Player{
		name:  name,
		rect:  rect,
		color: color,
	}

	lifecycle.Register(&lifecycle.GameObject{
		Render: player.render,
	})

	return player
}

func (p *Player) render() {
	render.DrawRect(p.rect, p.color)
}
