package config

import "github.com/mikabrytu/gomes-engine/math"

var SCREEN_SIZE = math.Vector2{X: 624, Y: 608}

const METRICS_UI_PANEL_HEIGHT int = 64
const METRICS_UI_FONT_SIZE = 18
const METRICS_UI_FONT_OFFSET = 16
const METRICS_UI_LIVES_OFFSET = 8
const METRICS_UI_DIVIDER_SIZE = 2

const METRICS_OBJECT_PLAYER_OFFSET int = 32
const METRICS_OBJECT_ALIEN_OFFSET int = 8
const METRICS_OBJECT_SHIELD_OFFSET int = 64

var METRICS_PLAYER_SIZE = math.Vector2{X: 32, Y: 32}
var METRICS_OBJECT_ALIEN_SIZE = math.Vector2{X: 32, Y: 32}
var METRICS_OBJECT_SHIELD_SIZE = math.Vector2{X: 64, Y: 64}

var VEC2_ZERO = math.Vector2{X: 0, Y: 0}
