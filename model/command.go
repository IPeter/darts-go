package model

const (
	ThrowOutOfBoard = 0
	ThrowNormal     = 1
	ThrowDouble     = 2
	ThrowTriple     = 3
)

type CamCommand struct {
	Score    int
	Modifier int
}

func (cc *CamCommand) IsValid() bool {
	if cc.Modifier >= -1 && cc.Modifier <= 3 {
		return true
	}

	return false
}

// ahol num a szám (1 .. 20, 25)  a modifier pedig: 0, 1, 2, 3. 0: pályán kívüli dobás, 1: sima, 2: dupla, 3: meglepő módon trippla
