package inventory

import (
	"fmt"
	"strings"
	"text_game/pkg/game/item"
)

type Inventory []item.Item

func (i Inventory) String() string {
	var sb strings.Builder

	for _, item := range i {
		sb.WriteString(fmt.Sprintf("%s, ", item.Name()))
	}
	res := sb.String()
	if len(res) > 0 {
		res = res[:len(res)-2] // Remove the last ", " from the string
	}

	return res
}
