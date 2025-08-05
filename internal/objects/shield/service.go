package shield

import (
	"littlejumbo/guard/config"

	"github.com/mikabrytu/gomes-engine/debug"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func New(name string, rect utils.RectSpecs, color render.Color) *Shield {
	shield := &Shield{
		name:     name,
		rect:     rect,
		color:    color,
		isSimple: true,
	}

	lifecycle.Register(&lifecycle.GameObject{
		Start:   shield.start,
		Render:  shield.render,
		Destroy: shield.destroy,
	})

	return shield
}

func (s *Shield) SetSprite(path string) {
	s.isSimple = false

	s.sprite = render.NewSprite(s.name, path)
	s.sprite.Init(s.rect)
}

func (s *Shield) start() {
	rect := s.rect
	rect.Height -= 16 // TODO: Calculate this
	rect.PosY += 16   // TODO: Calculate this

	s.body = physics.RegisterBody(&rect, s.name)
}

func (s *Shield) render() {
	if s.isSimple {
		render.DrawRect(s.rect, s.color)
	}

	if debug.IsEnabled() {
		render.DrawRect(*s.body.Rect, config.COLOR_DEBUG)
	}
}

func (s *Shield) destroy() {
	s.sprite.ClearSprite()
	physics.RemoveBody(&s.body)
}
