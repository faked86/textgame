package location

import (
	"testing"

	"text_game/pkg/game/inventory"
	"text_game/pkg/game/item"
	"text_game/pkg/game/location/tags"
)

type locationCase struct {
	loc  BaseLocation
	want string
}

func TestLookAround(t *testing.T) {
	var kitchen = NewBaseLocation(
		"кухня",
		"кухня, ничего интересного",
		"ты находишься на кухне, надо собрать рюкзак и идти в универ",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)

	var room1 = NewBaseLocation(
		"комната",
		"ты в своей комнате",
		"пустая комната",
		"на столе",
		inventory.Inventory{
			item.NewBaseItem("ключи"),
			item.NewBaseItem("конспекты"),
			item.NewBaseItem("рюкзак"),
		},
		tags.Home,
	)

	var room2 = NewBaseLocation(
		"комната",
		"ты в своей комнате",
		"пустая комната",
		"на столе",
		inventory.Inventory{},
		tags.Home,
	)

	var lookAroundCases = []locationCase{
		{
			loc:  *kitchen,
			want: "ты находишься на кухне, надо собрать рюкзак и идти в универ.",
		},
		{
			loc:  *room1,
			want: "на столе: ключи, конспекты, рюкзак.",
		},
		{
			loc:  *room2,
			want: "пустая комната.",
		},
	}

	for cNum, c := range lookAroundCases {
		get := c.loc.LookAround()
		if get != c.want {
			t.Errorf("Case: %d Want: %s get: %s", cNum, c.want, get)
		}
	}
}

func TestEnter(t *testing.T) {
	var kitchen = NewBaseLocation(
		"кухня",
		"кухня, ничего интересного",
		"ты находишься на кухне, надо собрать рюкзак и идти в универ",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)

	var room = NewBaseLocation(
		"комната",
		"ты в своей комнате",
		"пустая комната",
		"на столе",
		inventory.Inventory{
			item.NewBaseItem("ключи"),
			item.NewBaseItem("конспекты"),
			item.NewBaseItem("рюкзак"),
		},
		tags.Home,
	)

	var lobby = NewBaseLocation(
		"коридор",
		"ничего интересного",
		"ничего интересного",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)

	var outside = NewBaseLocation(
		"улица",
		"на улице весна",
		"ничего интересного",
		"здесь есть",
		inventory.Inventory{},
		tags.Outside,
	)

	var enterCases = []locationCase{
		{
			loc:  *kitchen,
			want: "кухня, ничего интересного.",
		},
		{
			loc:  *room,
			want: "ты в своей комнате.",
		},
		{
			loc:  *lobby,
			want: "ничего интересного.",
		},
		{
			loc:  *outside,
			want: "на улице весна.",
		},
	}

	for cNum, c := range enterCases {
		get := c.loc.Enter()
		if get != c.want {
			t.Errorf("Case: %d Want: %s get: %s", cNum, c.want, get)
		}
	}
}

func TestTakeItemWrong(t *testing.T) {
	keys := item.NewBaseItem("ключи")
	notes := item.NewBaseItem("конспекты")
	backpack := item.NewBaseItem("рюкзак")

	room := NewBaseLocation(
		"комната",
		"ты в своей комнате",
		"пустая комната",
		"на столе",
		inventory.Inventory{
			keys,
			notes,
			backpack,
		},
		tags.Home,
	)

	item, err := room.TakeItem("телефон")

	if err == nil {
		t.Error("want: error, got: nil")
	}

	if err.Error() != "no such item in this location" {
		t.Error("got wrong error")
	}

	if item != nil {
		t.Error("want no item got item")
	}
}

func TestTakeItemRight(t *testing.T) {
	keys := item.NewBaseItem("ключи")
	notes := item.NewBaseItem("конспекты")
	backpack := item.NewBaseItem("рюкзак")

	room := NewBaseLocation(
		"комната",
		"ты в своей комнате",
		"пустая комната",
		"на столе",
		inventory.Inventory{
			keys,
			notes,
			backpack,
		},
		tags.Home,
	)

	item, err := room.TakeItem("ключи")

	if err != nil {
		t.Error("want: no error, got: error")
	}

	if item == nil {
		t.Error("want item got nil")
	}

	if item != keys {
		t.Errorf("want: %p, get: %p", item, &keys)
	}

	for i := range room.Loot {
		if room.Loot[i] == item {
			t.Error("item still in room after it was taken")
		}
	}
}
