package model

import "time"

type HoroTimer struct {
	Id       int
	Context  string
	Length   int
	InsertAt time.Time
	DoneAt   time.Time
}
