package main

import (
	"littlejumbo/guard/config"

	gomesengine "github.com/mikabrytu/gomes-engine"
)

func main() {
	gomesengine.Init(config.GAME_TITLE, int32(config.SCREEN_SIZE.X), int32(config.SCREEN_SIZE.Y))
	gomesengine.Run()
}
