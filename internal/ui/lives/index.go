package lives

import (
	"littlejumbo/guard/config"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

var lives [3]*lifecycle.GameObject
var current int

const max int = 3

func Init() {
	current = max

	drawLives()
	drawDivider()
	listen()
}

func listen() {
	events.Subscribe(config.EVENTS_PLAYER_HIT, func(params ...any) error {
		if current <= 0 {
			return nil
		}

		current--

		if current == 0 {
			events.Emit(config.EVENTS_GAME_OVER)
		} else {
			lifecycle.Disable(lives[current])
		}

		return nil
	})
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

		lives[i] = lifecycle.Register(&lifecycle.GameObject{
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
