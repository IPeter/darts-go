package model

import "github.com/google/uuid"
import "time"

type Throw struct {
	ID       string    `json:"id"`
	Score    int       `json:"score"`
	Modifier int       `json:"modifier"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
	Cam1Img  string    `json:"cam1img"`
	Cam2Img  string    `json:"cam2img"`
	Time     time.Time `json:"time"`
}

func NewThrow(score int, modifier int, x int, y int, cam1img string, cam2img string) *Throw {
	return &Throw{
		ID:       uuid.New().String(),
		Score:    score,
		Modifier: modifier,
		X:       x,
		Y:       y,
		Cam1Img: cam1img,
		Cam2Img: cam2img,
		Time:     time.Now().UTC(),
	}
}
