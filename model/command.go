package model

const (
	GameEndsWithWinner = -2
	HandsOnBoard       = -1
	ThrowOutOfBoard    = 0
	ThrowNormal        = 1
	ThrowDouble        = 2
	ThrowTriple        = 3
)

type CamCommand struct {
	Score    int
	Modifier int
	X int
	Y int
	Cam1Img string
	Cam2Img string
}

func (cc *CamCommand) IsValid() bool {
	if cc.Modifier >= -1 && cc.Modifier <= 3 {
		return true
	}

	return false
}

// where num is a number from: [1 .. 20, 25]  a modifier is in : [-1, 0, 1, 2, 3]. 1: simple, 2: double, 3: triple, 0: out of bounds, -1: darts are removed before the third throw.
