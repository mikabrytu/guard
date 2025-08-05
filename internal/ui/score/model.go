package score

import "time"

type Score struct {
	Score     int       `json:"score"`
	TimeStamp time.Time `json:"timestamp"`
}
