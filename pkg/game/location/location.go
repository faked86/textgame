package location

import (
	"errors"
	"fmt"
	"text_game/pkg/game/inventory"
	"text_game/pkg/game/item"
)

type Location struct {
	Name                string
	Loot                inventory.Inventory
	enterText           string
	lookAroundTextEmpty string
	lookAroundTextLoot  string
	Tag                 Tag
}

func NewLocation(name, enterTxt, lkArndTxtEmpty, lkArndTxtLoot string, loot inventory.Inventory, tag Tag) *Location {
	return &Location{
		Name:                name,
		Loot:                loot,
		enterText:           enterTxt,
		lookAroundTextEmpty: lkArndTxtEmpty,
		lookAroundTextLoot:  lkArndTxtLoot,
		Tag:                 tag,
	}
}

func (l *Location) LookAround() string {
	if len(l.Loot) == 0 {
		return fmt.Sprintf("%s.", l.lookAroundTextEmpty)
	}

	return fmt.Sprintf("%s: %s.", l.lookAroundTextLoot, &l.Loot)
}

func (l *Location) Enter() string {
	return l.enterText + "."
}

func (l *Location) TakeItem(itemName string) (*item.BaseItem, error) {
	for i, item := range l.Loot {
		if item.Name() == itemName {
			l.Loot = append(l.Loot[:i], l.Loot[i+1:]...)
			return item, nil
		}
	}
	return nil, errors.New("no such item in this location")
}
