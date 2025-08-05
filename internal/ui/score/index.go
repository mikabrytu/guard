package score

import (
	"fmt"
	"littlejumbo/guard/config"
	"os"
	"time"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	savesystem "github.com/mikabrytu/gomes-engine/systems/save"
	"github.com/mikabrytu/gomes-engine/ui"
)

var sFont *ui.Font
var hFont *ui.Font
var data Score
var current int
var high int

func Init() {
	listen()
	loadFont()
	loadData()
	draw()
}

func listen() {
	events.Subscribe(config.EVENTS_ALIEN_DESTROYED, func(params ...any) error {
		points := params[0].([]any)[0].([]any)[1].(int)

		current += points
		text := fmt.Sprintf(config.UI_SCORE, current)
		sFont.UpdateText(text)

		return nil
	})

	events.Subscribe(config.EVENTS_GAME_OVER, func(params ...any) error {
		saveData()

		return nil
	})
}

func loadFont() {
	specs := ui.FontSpecs{
		Name: config.UI_FONT_NAME,
		Path: config.PATH_FONT,
		Size: config.METRICS_UI_FONT_SIZE,
	}

	sFont = ui.NewFont(specs, config.SCREEN_SIZE)
	hFont = ui.NewFont(specs, config.SCREEN_SIZE)
}

func loadData() {
	if !fileExists(config.PATH_SAVE_FILE) {
		createFile(config.PATH_SAVE_FILE)
	}

	err := savesystem.Load(config.PATH_SAVE_FILE, &data)
	if err != nil {
		panic(err)
	}

	current = 0
	high = data.Score
}

func saveData() {
	if current <= high {
		return
	}

	data.Score = current
	data.TimeStamp = time.Now()
	savesystem.Save(data, config.PATH_SAVE_FILE)
}

func draw() {
	sText := fmt.Sprintf(config.UI_SCORE, current)
	hText := fmt.Sprintf(config.UI_HIGH_SCORE, high)
	offset := math.Vector2{
		X: config.METRICS_UI_FONT_OFFSET,
		Y: config.METRICS_UI_FONT_OFFSET,
	}

	sFont.Init(sText, render.White, config.VEC2_ZERO)
	sFont.AlignText(ui.TopLeft, offset)

	hFont.Init(hText, render.White, config.VEC2_ZERO)
	hFont.AlignText(ui.TopRight, offset)
}

// TODO: Add this to the engine as an util
func createFile(path string) {
	var empty Score
	savesystem.Save(empty, path)
}

// TODO: Add this to the engine as an util
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
