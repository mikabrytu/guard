package player

import (
	"fmt"
	"littlejumbo/guard/config"
	"littlejumbo/guard/internal/objects/bullet"
	"time"

	"github.com/mikabrytu/gomes-engine/debug"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/physics"
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
		Physics: player.physics,
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
	p.body = physics.RegisterBody(&p.rect, p.name)
	p.updateBody()
	p.listen()
}

func (p *Player) physics() {
	if p.moving {
		p.rect.PosX += p.speed * p.axis
		p.sprite.UpdateRect(p.rect)
		p.updateBody()
	}
}

func (p *Player) render() {
	if p.isSimple {
		render.DrawRect(p.rect, p.color)
	}

	if debug.IsEnabled() {
		render.DrawRect(*p.body.Rect, config.COLOR_DEBUG)
	}
}

func (p *Player) destroy() {
	p.sprite.ClearSprite()
	physics.RemoveBody(&p.body)
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

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_SPACE, func(params ...any) error {
		p.shoot()
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

// TODO: Make shoot system so a player object doesn't have a bullet dependency
func (p *Player) shoot() {
	id := time.Now().Unix()
	name := fmt.Sprintf("bullet-%v", id)
	rect := utils.RectSpecs{
		PosX:   p.rect.PosX + ((config.METRICS_OBJECT_PLAYER_SIZE.X / 2) - (config.METRICS_OBJECT_BULLET_PLAYER_SIZE.X / 2)),
		PosY:   p.rect.PosY - config.METRICS_OBJECT_BULLET_PLAYER_SIZE.Y,
		Width:  config.METRICS_OBJECT_BULLET_PLAYER_SIZE.X,
		Height: config.METRICS_OBJECT_BULLET_PLAYER_SIZE.Y,
	}

	bullet := bullet.New(name, rect, config.COLOR_OBJECT_PLAYER)
	bullet.SetDirection(-1)
	bullet.SetSpeed(config.OBJECT_BULLET_PLAYER_SPEED)
}

func (p *Player) updateBody() {
	if p.body.Rect != &p.rect {
		p.body.Rect = &p.rect
		p.body.Rect.Height -= 16
		p.body.Rect.PosY += 16
	}
}
