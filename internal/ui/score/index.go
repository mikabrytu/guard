package score

import (
	"fmt"
	"littlejumbo/guard/config"

	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/ui"
)

var score *ui.Font
var high *ui.Font

func Init() {
	load()
	draw()
}

func load() {
	specs := ui.FontSpecs{
		Name: config.UI_FONT_NAME,
		Path: config.PATH_FONT,
		Size: config.METRICS_UI_FONT_SIZE,
	}

	score = ui.NewFont(specs, config.SCREEN_SIZE)
	high = ui.NewFont(specs, config.SCREEN_SIZE)
}

func draw() {
	sText := fmt.Sprintf(config.UI_SCORE, 0)
	hText := fmt.Sprintf(config.UI_HIGH_SCORE, 0)
	offset := math.Vector2{
		X: config.METRICS_UI_FONT_OFFSET,
		Y: config.METRICS_UI_FONT_OFFSET,
	}

	score.Init(sText, render.White, config.VEC2_ZERO)
	score.AlignText(ui.TopLeft, offset)

	high.Init(hText, render.White, config.VEC2_ZERO)
	high.AlignText(ui.TopRight, offset)
}
