package scene

import (
	"fmt"
	"littlejumbo/guard/config"
	"littlejumbo/guard/internal/objects/player"
	"littlejumbo/guard/internal/objects/shield"
	"littlejumbo/guard/internal/ui/lives"
	"littlejumbo/guard/internal/ui/score"

	alien_manager "littlejumbo/guard/internal/managers/alien"

	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func Init() {
	score.Init()
	lives.Init()

	drawPlayer()
	drawShields()

	alien_manager.Init()
}

func drawPlayer() {
	rect := utils.RectSpecs{
		PosX:   (config.SCREEN_SIZE.X / 2) - (config.METRICS_PLAYER_SIZE.X / 2),
		PosY:   config.SCREEN_SIZE.Y - config.METRICS_UI_PANEL_HEIGHT - config.METRICS_OBJECT_PLAYER_OFFSET - config.METRICS_PLAYER_SIZE.Y,
		Width:  config.METRICS_PLAYER_SIZE.X,
		Height: config.METRICS_PLAYER_SIZE.Y,
	}

	player := player.New(config.OBJECT_PLAYER_NAME, rect, render.Green)
	player.SetSprite(config.PATH_SPRITE_PLAYER)
	player.SetSpeed(config.OBJECT_PLAYER_SPEED)
}

func drawShields() {
	max := 4

	for i := 0; i < max; i++ {
		x := 88 // TODO: Calculate start X

		if x > 0 {
			x += i * (config.METRICS_OBJECT_SHIELD_SIZE.X + config.METRICS_OBJECT_SHIELD_OFFSET)
		}

		name := fmt.Sprintf(config.OBJECT_SHIELD_NAME, 0)
		rect := utils.RectSpecs{
			PosX:   x,
			PosY:   config.SCREEN_SIZE.Y - config.METRICS_UI_PANEL_HEIGHT - (2 * config.METRICS_OBJECT_PLAYER_OFFSET) - config.METRICS_PLAYER_SIZE.Y - config.METRICS_OBJECT_SHIELD_SIZE.Y,
			Width:  config.METRICS_OBJECT_SHIELD_SIZE.X,
			Height: config.METRICS_OBJECT_SHIELD_SIZE.Y,
		}

		shield := shield.New(name, rect, render.Green)
		shield.SetSprite(config.PATH_SPRITE_SHIELD)
	}
}
