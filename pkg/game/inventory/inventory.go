package inventory

import (
	"fmt"
	"strings"

	"text_game/pkg/game/item"
)

type Inventory []*item.BaseItem

func (i *Inventory) String() string {
	var sb strings.Builder

	for _, item := range *i {
		sb.WriteString(fmt.Sprintf("%s, ", item.Name()))
	}
	res := sb.String()
	res = strings.TrimRight(res, ", ")

	return res
}
