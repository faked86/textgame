package game

import (
	"testing"

	"text_game/pkg/game/inventory"
	"text_game/pkg/game/location"
)

func TestFindWaysNoWay(t *testing.T) {
	kitchen := location.NewLocation(
		"кухня",
		"кухня, ничего интересного",
		"ты находишься на кухне, надо собрать рюкзак и идти в универ",
		"здесь есть",
		inventory.Inventory{},
		location.Home,
	)
	game := NewGame(kitchen, &inventory.Inventory{})

	get := len(game.findWays())
	want := 0
	if get != want {
		t.Error("Want empty list, get list with objects")
	}
}

func TestFindWaysOneWay(t *testing.T) {
	kitchen := location.NewLocation(
		"кухня",
		"кухня, ничего интересного",
		"ты находишься на кухне, надо собрать рюкзак и идти в универ",
		"здесь есть",
		inventory.Inventory{},
		location.Home,
	)
	lobby := location.NewLocation(
		"коридор",
		"ничего интересного",
		"ничего интересного",
		"здесь есть",
		inventory.Inventory{},
		location.Home,
	)
	game := NewGame(kitchen, &inventory.Inventory{})
	game.AddLocation(lobby)
	game.AddWay(0, 1)

	get := game.findWays()
	wantLen := 1

	if len(get) != wantLen {
		t.Errorf("Want slice with 1 obj, get %d objects", len(get))
	}

	if &get[len(get)-1] == lobby {
		t.Error("Get wrong object")
	}
}
