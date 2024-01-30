package inventory

import (
	"testing"

	"text_game/pkg/game/item"
)

type stringCase struct {
	inv  Inventory
	want string
}

var stringCases = []stringCase{
	{
		inv:  Inventory{},
		want: "",
	},
	{
		inv:  Inventory{&item.Item{Name: "рюкзак"}},
		want: "рюкзак",
	},
	{
		inv:  Inventory{&item.Item{Name: "рюкзак"}, &item.Item{Name: "конспекты"}},
		want: "рюкзак, конспекты",
	},
}

func TestString(t *testing.T) {
	for cNum, c := range stringCases {
		if c.want != c.inv.String() {
			t.Errorf("Case: %d Want: %s get: %s", cNum, c.want, &c.inv)
		}
	}
}
