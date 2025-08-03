package alien

import (
	"littlejumbo/guard/config"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/physics"
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

	alien.instance = lifecycle.Register(&lifecycle.GameObject{
		Start:   alien.start,
		Physics: alien.physics,
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

func (a *Alien) SetStep(step math.Vector2) {
	a.step = step
}

func (a *Alien) MoveStep() {
	a.rect.PosX += a.axis.X * a.step.X
	a.sprite.UpdateRect(a.rect)
}

func (a *Alien) InvertX() {
	a.axis.X *= -1
}

func (a *Alien) DescendY() {
	a.rect.PosY += a.step.Y
	a.sprite.UpdateRect(a.rect)
}

func (a *Alien) IsAtScreenEdge() bool {
	return (a.rect.PosX+a.rect.Width) > config.SCREEN_SIZE.X-config.SCREEN_OFFSET.X || a.rect.PosX < config.SCREEN_OFFSET.X
}

func (a *Alien) start() {
	a.body = physics.RegisterBody(&a.rect, a.Name)

	events.Subscribe(config.EVENTS_BULLET_HIT, func(params ...any) error {
		name := params[0].([]any)[0].([]any)[0].(string)
		if name == a.Name {
			lifecycle.Stop(a.instance)
		}

		return nil
	})
}

func (a *Alien) physics() {
	collision := physics.CheckCollision(&a.body)
	if collision.Name != "nil" {
		println("Collision detected")
		lifecycle.Stop(a.instance)
	}
}

func (a *Alien) render() {
	if a.isSimple {
		render.DrawRect(a.rect, a.color)
	}
}

func (a *Alien) destroy() {
	a.sprite.ClearSprite()
	physics.RemoveBody(&a.body)
}
