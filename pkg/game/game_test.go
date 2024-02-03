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
	game := NewGame(kitchen, inventory.Inventory{})

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
	game := NewGame(kitchen, inventory.Inventory{})
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

func TestWaysMessageNoWay(t *testing.T) {
	kitchen := location.NewBaseLocation(
		"кухня",
		"кухня, ничего интересного",
		"ты находишься на кухне, надо собрать рюкзак и идти в универ",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)
	game := NewGame(kitchen, inventory.Inventory{})

	got := game.waysMessage()
	want := ""
	if got != want {
		t.Error("Want empty string, got list with objects")
	}
}

func TestWaysMessageOneWay(t *testing.T) {
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
	game := NewGame(kitchen, inventory.Inventory{})
	game.AddLocation(lobby)
	game.AddWay(0, 1)

	got := game.waysMessage()
	want := "можно пройти - коридор"

	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}

func TestWaysMessageMultipleWays(t *testing.T) {
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
	var room location.Location = location.NewBaseLocation(
		"комната",
		"ты в своей комнате",
		"пустая комната",
		"на столе",
		inventory.Inventory{},
		tags.Home,
	)
	game := NewGame(lobby, inventory.Inventory{})

	game.AddLocation(kitchen)
	game.AddWay(0, 1)

	game.AddLocation(room)
	game.AddWay(0, 2)

	got := game.waysMessage()
	want := "можно пройти - кухня, комната"

	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}

func TestWaysMessageDifferentTag1(t *testing.T) {
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
	var room location.Location = location.NewBaseLocation(
		"комната",
		"ты в своей комнате",
		"пустая комната",
		"на столе",
		inventory.Inventory{},
		tags.Outside,
	)
	game := NewGame(lobby, inventory.Inventory{})

	game.AddLocation(kitchen)
	game.AddWay(0, 1)

	game.AddLocation(room)
	game.AddWay(0, 2)

	got := game.waysMessage()
	want := "можно пройти - кухня, улица"

	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}

func TestWaysMessageDifferentTag2(t *testing.T) {
	var lobby location.Location = location.NewBaseLocation(
		"коридор",
		"ничего интересного",
		"ничего интересного",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)
	var outside location.Location = location.NewBaseLocation(
		"улица",
		"на улице весна",
		"ничего интересного",
		"здесь есть",
		inventory.Inventory{},
		tags.Outside,
	)
	game := NewGame(outside, inventory.Inventory{})

	game.AddLocation(lobby)
	game.AddWay(0, 1)

	got := game.waysMessage()
	want := "можно пройти - домой"

	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}
func TestLookAroundNoWay(t *testing.T) {
	kitchen := location.NewBaseLocation(
		"кухня",
		"кухня, ничего интересного",
		"ты находишься на кухне, надо собрать рюкзак и идти в универ",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)
	game := NewGame(kitchen, inventory.Inventory{})

	got := game.LookAround()
	want := "ты находишься на кухне, надо собрать рюкзак и идти в универ."
	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}

func TestLookAroundOneWay(t *testing.T) {
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
	game := NewGame(kitchen, inventory.Inventory{})
	game.AddLocation(lobby)
	game.AddWay(0, 1)

	got := game.LookAround()
	want := "ты находишься на кухне, надо собрать рюкзак и идти в универ. можно пройти - коридор"

	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}

func TestLookAroundMultipleWays(t *testing.T) {
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
	var room location.Location = location.NewBaseLocation(
		"комната",
		"ты в своей комнате",
		"пустая комната",
		"на столе",
		inventory.Inventory{},
		tags.Home,
	)
	game := NewGame(lobby, inventory.Inventory{})

	game.AddLocation(kitchen)
	game.AddWay(0, 1)

	game.AddLocation(room)
	game.AddWay(0, 2)

	got := game.LookAround()
	want := "ничего интересного. можно пройти - кухня, комната"

	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}

func TestLookAroundDifferentTag1(t *testing.T) {
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
	var room location.Location = location.NewBaseLocation(
		"комната",
		"ты в своей комнате",
		"пустая комната",
		"на столе",
		inventory.Inventory{},
		tags.Outside,
	)
	game := NewGame(lobby, inventory.Inventory{})

	game.AddLocation(kitchen)
	game.AddWay(0, 1)

	game.AddLocation(room)
	game.AddWay(0, 2)

	got := game.LookAround()
	want := "ничего интересного. можно пройти - кухня, улица"

	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}

func TestLookAroundDifferentTag2(t *testing.T) {
	var lobby location.Location = location.NewBaseLocation(
		"коридор",
		"ничего интересного",
		"ничего интересного",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)
	var outside location.Location = location.NewBaseLocation(
		"улица",
		"на улице весна",
		"на улице весна",
		"здесь есть",
		inventory.Inventory{},
		tags.Outside,
	)
	game := NewGame(outside, inventory.Inventory{})

	game.AddLocation(lobby)
	game.AddWay(0, 1)

	got := game.LookAround()
	want := "на улице весна. можно пройти - домой"

	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}
