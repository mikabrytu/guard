package scene

import (
	"fmt"
	"littlejumbo/guard/config"
	"littlejumbo/guard/internal/objects/alien"
	"littlejumbo/guard/internal/objects/player"

	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func Init() {
	drawPlayer()
	drawAliens()
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

func drawAliens() {
	rows := 5
	cols := 11

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := 96 // TODO: Calculate this to ensure responsiveness
			y := config.METRICS_UI_PANEL_HEIGHT

			if j > 0 {
				x += j * (config.METRICS_OBJECT_ALIEN_SIZE.X + config.METRICS_OBJECT_ALIEN_OFFSET)
			}

			if i > 0 {
				y += i * (config.METRICS_OBJECT_ALIEN_SIZE.Y + config.METRICS_OBJECT_ALIEN_OFFSET)
			}

			name := fmt.Sprintf(config.OBJECT_ALIEN_NAME, i, j)
			rect := utils.RectSpecs{
				PosX:   x,
				PosY:   y,
				Width:  config.METRICS_OBJECT_ALIEN_SIZE.X,
				Height: config.METRICS_OBJECT_ALIEN_SIZE.Y,
			}

			alien.New(name, rect, render.White)
		}
	}
}
