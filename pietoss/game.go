package pietoss

// Game ...
type Game struct {
	Players Players
}

// NewGame ...
func NewGame(p Players) *Game {
	return &Game{Players: p}
}
