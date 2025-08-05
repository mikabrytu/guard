package alien

import (
	"container/list"
	"fmt"
	"littlejumbo/guard/config"
	"littlejumbo/guard/internal/objects/alien"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	gomesdebug "github.com/mikabrytu/gomes-engine/debug"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

var aliens [][]*alien.Alien
var aggressive *list.List
var moving bool
var skip bool

const ROWS int = 5
const COLS int = 11
const AGGRESSIVE_MIN = 7

func Init() {
	moving = true
	skip = false

	listen()
	draw()
	move()
	shoot()

	if gomesdebug.IsEnabled() {
		debug()
	}
}

func listen() {
	events.Subscribe(config.EVENTS_ALIEN_DESTROYED, func(params ...any) error {
		name := params[0].([]any)[0].([]any)[0].(string)

		re := regexp.MustCompile("[0-9]+")
		search := re.FindAllString(name, -1)

		if len(search) != 2 {
			fmt.Printf("Name %v is invalid.\n", name)
			return nil
		}

		var err error
		var index []int = make([]int, 2)
		index[0], err = strconv.Atoi(search[0])
		index[1], err = strconv.Atoi(search[1])

		if err != nil {
			panic(err)
		}

		tryRemoveAlien(index)
		tryRemoveAggresive(name)

		return nil
	})
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
			score := 0
			if i == 0 {
				path = config.PATH_SPRITE_ALIEN_C
				score = config.SCORE_VALUE_ALIEN_C
			} else if i < 3 {
				path = config.PATH_SPRITE_ALIEN_B
				score = config.SCORE_VALUE_ALIEN_B
			} else {
				path = config.PATH_SPRITE_ALIEN_A
				score = config.SCORE_VALUE_ALIEN_A
			}

			alien := alien.New(name, rect, config.COLOR_OBJECT_ALIEN)
			alien.SetSprite(path)
			alien.SetStep(config.OBJECT_ALIEN_STEP)
			alien.SetScore(score)
			aliens[i][j] = alien
		}
	}
}

func move() {
	time.AfterFunc(time.Duration(config.DELAY_ALIEN_MOVEMENT)*time.Millisecond, func() {
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

func shoot() {
	time.AfterFunc(config.DELAY_ALIEN_SHOOT*time.Millisecond, func() {
		updateAggressive()

		id := rand.Intn(aggressive.Len())
		index := 0
		for e := aggressive.Front(); e != nil; e = e.Next() {
			if index == id {
				a := e.Value.(*alien.Alien)
				a.Shoot()

				shoot()
				return
			}
			index++
		}
	})
}

func debug() {
	lifecycle.Register(&lifecycle.GameObject{
		Render: func() {
			if aggressive == nil {
				return
			}

			for e := aggressive.Front(); e != nil; e = e.Next() {
				alien := e.Value.(*alien.Alien)
				render.DrawRect(*alien.Body.Rect, render.Orange)
			}
		},
	})
}

func tryRemoveAlien(index []int) {
	alien := aliens[index[0]][index[1]]
	if alien == nil {
		fmt.Printf("Alien at coord [%v, %v] not found\n", index[0], index[1])
	} else {
		aliens[index[0]][index[1]] = nil
	}
}

func tryRemoveAggresive(name string) {
	if aggressive == nil {
		return
	}

	for e := aggressive.Front(); e != nil; e = e.Next() {
		a := e.Value.(*alien.Alien)
		if a.Name == name {
			aggressive.Remove(e)
			break
		}
	}
}

func callFunc(callback func(*alien.Alien)) {
	for _, row := range aliens {
		for _, a := range row {
			if a == nil {
				continue
			}

			callback(a)
		}
	}
}

func updateAggressive() {
	aggressive = list.New()

	if aggressive.Len() >= COLS {
		return
	}

	bottom := ROWS - 1

	for bottom >= 0 && aggressive.Len() < COLS {
		for i, a := range aliens[bottom] {
			if a == nil {
				continue
			}

			if bottom < ROWS-1 && aliens[bottom+1][i] != nil {
				continue
			}

			aggressive.PushBack(a)
		}

		bottom--
	}
}
