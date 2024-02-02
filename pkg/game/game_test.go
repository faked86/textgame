package game

import (
	"testing"

	"text_game/pkg/game/inventory"
	"text_game/pkg/game/location"
	"text_game/pkg/game/location/tags"
)

func TestFindWaysNoWay(t *testing.T) {
	kitchen := location.NewBaseLocation(
		"кухня",
		"кухня, ничего интересного",
		"ты находишься на кухне, надо собрать рюкзак и идти в универ",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)
	game := NewGame(kitchen, &inventory.Inventory{})

	got := len(game.findWays())
	want := 0
	if got != want {
		t.Error("Want empty list, get list with objects")
	}
}

func TestFindWaysOneWay(t *testing.T) {
	var kitchen location.Location = location.NewBaseLocation(
		"кухня",
		"кухня, ничего интересного",
		"ты находишься на кухне, надо собрать рюкзак и идти в универ",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)
	var lobby location.Location = location.NewBaseLocation(
		"коридор",
		"ничего интересного",
		"ничего интересного",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)
	game := NewGame(kitchen, &inventory.Inventory{})
	game.AddLocation(lobby)
	game.AddWay(0, 1)

	got := game.findWays()
	wantLen := 1
	wantElem := 1

	if len(got) != wantLen {
		t.Errorf("Want slice with 1 obj, get %d objects", len(got))
	}

	if got[len(got)-1] != wantElem {
		t.Error("Get wrong object")
	}
}
