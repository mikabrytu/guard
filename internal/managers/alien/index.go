package alien

import (
	"fmt"
	"littlejumbo/guard/config"
	"littlejumbo/guard/internal/objects/alien"
	"time"

	"github.com/mikabrytu/gomes-engine/utils"
)

var aliens [][]*alien.Alien
var moving bool
var skip bool

const MOVE_DELAY int = 1000
const ROWS int = 5
const COLS int = 11

func Init() {
	moving = true
	skip = false

	draw()
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

			alien := alien.New(name, rect, config.COLOR_OBJECT_ALIEN)
			alien.SetSprite(path)
			alien.SetStep(config.OBJECT_ALIEN_STEP)
			aliens[i][j] = alien
		}
	}
}

func move() {
	time.AfterFunc(time.Duration(MOVE_DELAY)*time.Millisecond, func() {
		if !skip {
			a := aliens[0][0]
			b := aliens[0][COLS-1]
			if a.IsAtScreenEdge() || b.IsAtScreenEdge() {
				callFunc(func(a *alien.Alien) {
					a.InvertX()
					a.DescendY()
				})
				skip = true
				move()

				return
			}
		}

		callFunc(func(a *alien.Alien) {
			a.MoveStep()
		})

		if moving {
			skip = false
			move()
		}
	})
}

func callFunc(callback func(*alien.Alien)) {
	for _, row := range aliens {
		for _, a := range row {
			callback(a)
		}
	}
}
