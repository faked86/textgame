package location

import (
	"text_game/pkg/game/item"
	"text_game/pkg/game/location/tags"
)

type Location interface {
	Name() string
	LookAround() string
	Enter() string
	TakeItem(itemName string) (item.Item, error)
	Tag() tags.Tag
}
