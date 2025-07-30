package shield

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
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

func (s *Shield) render() {
	if s.isSimple {
		render.DrawRect(s.rect, s.color)
	}
}

func (s *Shield) destroy() {
	s.sprite.ClearSprite()
}
