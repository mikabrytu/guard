package player

import (
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Player struct {
	Name   string
	Rect   utils.RectSpecs
	Color  render.Color
	speed  int
	axis   int
	moving bool
}
