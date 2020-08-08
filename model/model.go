package model

import "time"

type Image struct {
	ID        string     `json:"id"`
	Path      string     `json:"path"`
	Extension string     `json:"extension"`
	Size      int        `json:"size"`
	Ts        *time.Time `json:"ts"`
	Variety   []*Resized `json:"variety"`
}

type Resized struct {
	Path   string `json:"path"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type SizeInput struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (si SizeInput) Valid() bool {
	return si.Height > 0 && si.Width > 0
}
