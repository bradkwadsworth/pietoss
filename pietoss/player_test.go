package pietoss

import (
	"strconv"
	"testing"
)

var players = []struct {
	in  map[string]string
	out *Player
}{
	{in: map[string]string{"name": "Shemp", "height": "5.85"}, out: &Player{Name: "Shemp", Height: 5.85}},
	{in: map[string]string{"name": "", "height": "-1.0"}, out: &Player{Name: "random", Height: 5.75}},
}

func TestNewPlayer(t *testing.T) {
	t.Log("Create a new Player")
	for _, i := range players {
		if s, err := strconv.ParseFloat(i.in["height"], 64); err == nil {
			if x := NewPlayer(i.in["name"], s); *x != *i.out {
				t.Log(i.out)
				t.Log(x)
				t.Errorf("%v does not equal %v", i.in, x)
			}
		}
	}
}
