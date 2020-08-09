package model

import (
	"time"
)

type Image struct {
	ID      string     `json:"id" bson:"_id"`
	Path    string     `json:"path" bson:"path"`
	Type    string     `json:"extension" bson:"extension"`
	Size    int        `json:"size" bson:"size"`
	Ts      time.Time `json:"ts" bson:"ts"`
	Variety []Resized  `json:"variety" bson:"variety"`
}

type Resized struct {
	Path   string `json:"path" bson:"path"`
	Width  int    `json:"width" bson:"width"`
	Height int    `json:"height" bson:"height"`
}

type SizeInput struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (si SizeInput) Valid() bool {
	return si.Height > 0 && si.Width > 0
}
