package bullet

import (
	"littlejumbo/guard/config"

	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func New(name string, rect utils.RectSpecs, color render.Color) *Bullet {
	bullet := &Bullet{
		Name:  name,
		rect:  rect,
		color: color,
	}

	bullet.instance = lifecycle.Register(&lifecycle.GameObject{
		Update: bullet.update,
		Render: bullet.render,
	})

	return bullet
}

func (b *Bullet) SetDirection(direction int) {
	b.axis = direction
}

func (b *Bullet) SetSpeed(speed int) {
	b.speed = speed
}

func (b *Bullet) update() {
	b.rect.PosY += b.axis * b.speed

	// TODO: Make a pooling system
	if (b.rect.PosY+b.rect.Height) < 0 || b.rect.PosY > config.SCREEN_SIZE.Y {
		lifecycle.Stop(b.instance)
	}
}

func (b *Bullet) render() {
	render.DrawRect(b.rect, b.color)
}
