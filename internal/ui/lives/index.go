package lives

import (
	"littlejumbo/guard/config"

	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

const max int = 3

func Init() {
	drawLives()
	drawDivider()
}

func drawLives() {
	for i := range max {
		x := config.METRICS_UI_PANEL_HEIGHT / 4

		if i > 0 {
			x += i * (config.METRICS_OBJECT_PLAYER_SIZE.X + config.METRICS_UI_LIVES_OFFSET)
		}

		rect := utils.RectSpecs{
			PosX:   x,
			PosY:   config.SCREEN_SIZE.Y - config.METRICS_OBJECT_PLAYER_SIZE.Y - (config.METRICS_UI_PANEL_HEIGHT / 4),
			Width:  config.METRICS_OBJECT_PLAYER_SIZE.X,
			Height: config.METRICS_OBJECT_PLAYER_SIZE.Y,
		}

		lifecycle.Register(&lifecycle.GameObject{
			Render: func() {
				render.DrawRect(rect, config.COLOR_UI_PLAYER)
			},
		})
	}
}

func drawDivider() {
	rect := utils.RectSpecs{
		PosX:   0,
		PosY:   config.SCREEN_SIZE.Y - config.METRICS_UI_PANEL_HEIGHT - config.METRICS_UI_DIVIDER_SIZE,
		Width:  config.SCREEN_SIZE.X,
		Height: config.METRICS_UI_DIVIDER_SIZE,
	}

	lifecycle.Register(&lifecycle.GameObject{
		Render: func() {
			render.DrawRect(rect, config.COLOR_UI_DIVIDER)
		},
	})
}
