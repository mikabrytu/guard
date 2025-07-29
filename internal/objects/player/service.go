package player

import (
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func New(name string, rect utils.RectSpecs, color render.Color) *Player {
	player := &Player{
		name:     name,
		rect:     rect,
		color:    color,
		isSimple: true,
	}

	lifecycle.Register(&lifecycle.GameObject{
		Start:   player.start,
		Update:  player.update,
		Render:  player.render,
		Destroy: player.destroy,
	})

	return player
}

func (p *Player) SetSprite(path string) {
	p.isSimple = false

	p.sprite = render.NewSprite(p.name, path)
	p.sprite.Init(p.rect)
}

func (p *Player) SetSpeed(speed int) {
	p.speed = speed
}

func (p *Player) start() {
	p.listen()

}

func (p *Player) update() {
	if p.moving {
		p.rect.PosX += p.speed * p.axis
		p.sprite.UpdateRect(p.rect)
	}
}

func (p *Player) render() {
	if p.isSimple {
		render.DrawRect(p.rect, p.color)
	}
}

func (p *Player) destroy() {
	p.sprite.ClearSprite()
}

func (p *Player) listen() {
	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_A, func(params ...any) error {
		if !p.moving {
			p.move(-1)
		}
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_D, func(params ...any) error {
		if !p.moving {
			p.move(1)
		}
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_RELEASED_A, func(params ...any) error {
		if p.moving {
			p.move(0)
		}
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_RELEASED_D, func(params ...any) error {
		if p.moving {
			p.move(0)
		}
		return nil
	})
}

func (p *Player) move(axis int) {
	p.axis = axis

	if axis == 0 {
		p.moving = false
	} else {
		p.moving = true
	}
}
