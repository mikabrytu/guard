package scene

import (
	"littlejumbo/guard/config"
	"littlejumbo/guard/internal/objects/player"

	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func Init() {
	drawPlayer()
}

func drawPlayer() {
	rect := utils.RectSpecs{
		PosX:   (config.SCREEN_SIZE.X / 2) - (config.METRICS_PLAYER_SIZE.X / 2),
		PosY:   config.SCREEN_SIZE.Y - config.METRICS_UI_PANEL_HEIGHT - config.METRICS_OBJECT_PLAYER_OFFSET - config.METRICS_PLAYER_SIZE.Y,
		Width:  config.METRICS_PLAYER_SIZE.X,
		Height: config.METRICS_PLAYER_SIZE.Y,
	}

	player.New(config.OBJECT_PLAYER_NAME, rect, render.Green)
}
