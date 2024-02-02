package location

import (
	"errors"
	"fmt"

	"text_game/pkg/game/inventory"
	"text_game/pkg/game/item"
	"text_game/pkg/game/location/tags"
)

type BaseLocation struct {
	Name                string
	Loot                inventory.Inventory
	enterText           string
	lookAroundTextEmpty string
	lookAroundTextLoot  string
	tag                 tags.Tag
}

func NewBaseLocation(name, enterTxt, lkArndTxtEmpty, lkArndTxtLoot string, loot inventory.Inventory, tag tags.Tag) *BaseLocation {
	return &BaseLocation{
		Name:                name,
		Loot:                loot,
		enterText:           enterTxt,
		lookAroundTextEmpty: lkArndTxtEmpty,
		lookAroundTextLoot:  lkArndTxtLoot,
		tag:                 tag,
	}
}

func (l *BaseLocation) LookAround() string {
	if len(l.Loot) == 0 {
		return fmt.Sprintf("%s.", l.lookAroundTextEmpty)
	}

	return fmt.Sprintf("%s: %s.", l.lookAroundTextLoot, &l.Loot)
}

func (l *BaseLocation) Enter() string {
	return l.enterText + "."
}

func (l *BaseLocation) TakeItem(itemName string) (*item.BaseItem, error) {
	for i, item := range l.Loot {
		if item.Name() == itemName {
			l.Loot = append(l.Loot[:i], l.Loot[i+1:]...)
			return item, nil
		}
	}
	return nil, errors.New("no such item in this location")
}

func (l *BaseLocation) Tag() tags.Tag {
	return l.tag
}
