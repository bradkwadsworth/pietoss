package pietoss

import "testing"

var games = []struct {
	in  Players
	out *Game
}{
	{in: Players{&Player{Name: "Larry"}, &Player{Name: "Shemp"}}, out: &Game{Players: Players{&Player{Name: "Larry"}, &Player{Name: "Shemp"}}}},
}

func TestNewGame(t *testing.T) {
	t.Log("Create a new game")
	for _, i := range games {
		g := NewGame(i.in)
		for y, x := range g.Players {
			if *x != *i.out.Players[y] {
				t.Errorf("%v does not equal %v", *x, *i.out.Players[y])
			}
		}
	}
}
