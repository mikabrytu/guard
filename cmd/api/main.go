package main

import (
	"littlejumbo/guard/config"
	scene "littlejumbo/guard/internal/scenes/main"
	"time"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
)

func main() {
	gomesengine.Init(config.GAME_TITLE, int32(config.SCREEN_SIZE.X), int32(config.SCREEN_SIZE.Y))

	settings()
	scene.Init()
	listen()

	gomesengine.Run()
}

func settings() {
	lifecycle.SetSmoothStep(0.9)

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_ESCAPE, func(params ...any) error {
		lifecycle.Kill()
		return nil
	})
}

func listen() {
	events.Subscribe(config.EVENTS_GAME_OVER, func(params ...any) error {
		time.AfterFunc(500*time.Millisecond, func() {
			lifecycle.StopAll()
		})

		time.AfterFunc(1000*time.Millisecond, func() {
			scene.Init()
		})

		return nil
	})
}
