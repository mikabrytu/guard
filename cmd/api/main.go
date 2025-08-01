package main

import (
	"littlejumbo/guard/config"
	scene "littlejumbo/guard/internal/scenes/main"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
)

func main() {
	gomesengine.Init(config.GAME_TITLE, int32(config.SCREEN_SIZE.X), int32(config.SCREEN_SIZE.Y))
	lifecycle.SetSmoothStep(0.9)

	settings()
	scene.Init()

	gomesengine.Run()
}

func settings() {
	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_ESCAPE, func(params ...any) error {
		lifecycle.Kill()
		return nil
	})
}
