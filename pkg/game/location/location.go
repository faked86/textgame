package location

import (
	"text_game/pkg/game/item"
)

type Location interface {
	LookAround() string
	Enter() string
	TakeItem(itemName string) (*item.BaseItem, error)
	Tag() string
}
