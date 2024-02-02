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
		inv:  Inventory{item.NewBaseItem("рюкзак")},
		want: "рюкзак",
	},
	{
		inv:  Inventory{item.NewBaseItem("рюкзак"), item.NewBaseItem("конспекты")},
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
