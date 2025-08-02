package alien

import (
	"fmt"
	"littlejumbo/guard/config"
	"littlejumbo/guard/internal/objects/alien"
	"time"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

var aliens [][]*alien.Alien
var moving bool
var eventFired bool

const MOVE_DELAY int = 1000
const ROWS int = 5
const COLS int = 11

func Init() {
	moving = true
	eventFired = false

	draw()
	listen()
	move()
}

func draw() {
	aliens = make([][]*alien.Alien, ROWS)
	for i := range ROWS {
		aliens[i] = make([]*alien.Alien, COLS)

		for j := range COLS {
			x := 96 // TODO: Calculate start X
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

			path := ""
			if i == 0 {
				path = config.PATH_SPRITE_ALIEN_C
			} else if i < 3 {
				path = config.PATH_SPRITE_ALIEN_B
			} else {
				path = config.PATH_SPRITE_ALIEN_A
			}

			alien := alien.New(name, rect, render.White)
			alien.SetSprite(path)
			alien.SetStep(config.OBJECT_ALIEN_STEP)
			aliens[i][j] = alien
		}
	}
}

func listen() {
	events.Subscribe(config.EVENT_ALIEN_AT_SCREEN_BOUNDARY, func(params ...any) error {
		if eventFired {
			return nil
		}

		for _, row := range aliens {
			for _, a := range row {
				a.InvertDirectionX()
			}
		}

		eventFired = true
		return nil
	})
}

func move() {
	if eventFired {
		eventFired = false
	}

	time.AfterFunc(time.Duration(MOVE_DELAY)*time.Millisecond, func() {
		for _, row := range aliens {
			for _, a := range row {
				a.MoveStep()
			}
		}

		if moving {
			move()
		}
	})
}
