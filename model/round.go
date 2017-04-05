package model

type Round struct {
	Throws  map[string]*Throw `json:"throws"`
	current int
	max     int
}

func NewRound() *Round {
	return &Round{
		Throws:  make(map[string]*Throw, 3),
		current: 0,
		max:     3,
	}
}
