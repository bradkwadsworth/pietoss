package pietoss

// Player ...
type Player struct {
	Name   string
	Height float64
}

type Players []*Player

func randomName() string {
	return "random"
}

func randomHeight() float64 {
	return 5.75
}

// NewPlayer ...
func NewPlayer(n string, h float64) *Player {
	if len(n) == 0 {
		n = randomName()
	}
	if h == -1.0 {
		h = randomHeight()
	}
	return &Player{Name: n, Height: h}
}
