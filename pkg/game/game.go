package game

import (
	"errors"
	"fmt"
	"strings"

	"text_game/pkg/game/inventory"
	"text_game/pkg/game/location"
)

type Game struct {
	WorldMap        [][]bool
	Locations       []location.Location
	playerInventory inventory.Inventory
	playerLocation  int
}

func NewGame(start location.Location, invent inventory.Inventory) *Game {
	w := &Game{
		WorldMap:        [][]bool{{false}},
		Locations:       []location.Location{start},
		playerInventory: invent,
		playerLocation:  0,
	}
	return w
}

func (g *Game) AddLocation(l location.Location) {
	g.Locations = append(g.Locations, l)
	locNum := len(g.WorldMap)
	newRow := make([]bool, locNum+1)

	for i := range g.WorldMap {
		g.WorldMap[i] = append(g.WorldMap[i], false)
		newRow[i] = false
	}

	g.WorldMap = append(g.WorldMap, newRow)
}

func (g *Game) AddWay(l1, l2 int) error {
	numLoc := len(g.Locations)
	if l1 >= numLoc || l2 >= numLoc {
		return errors.New("trying to add way to uninitialized location")
	}
	g.WorldMap[l1][l2] = true
	g.WorldMap[l2][l1] = true
	return nil
}

func (g *Game) findWays() []int {
	res := make([]int, 0)
	for i, way := range g.WorldMap[g.playerLocation] {
		if way {
			res = append(res, i)
		}
	}
	return res
}

func (g *Game) waysMessage() string {
	ways := g.findWays()

	if len(ways) == 0 {
		return ""
	}

	var sb strings.Builder
	curLocTag := g.Locations[g.playerLocation].Tag()

	for _, wayNum := range ways {
		wayTag := g.Locations[wayNum].Tag()

		if curLocTag != wayTag {
			sb.WriteString(fmt.Sprintf("%s, ", wayTag))
		} else {
			sb.WriteString(fmt.Sprintf("%s, ", g.Locations[wayNum].Name()))
		}
	}
	res := sb.String()
	if len(res) > 0 {
		res = res[:len(res)-2] // Remove the last ", " from the string
	}

	return "можно пройти - " + res
}

func (g *Game) LookAround() string {
	curLocMessage := g.Locations[g.playerLocation].LookAround()
	waysMessage := g.waysMessage()

	if waysMessage == "" {
		return curLocMessage
	}
	return curLocMessage + " " + waysMessage
}

func (g *Game) Walk(dest string) string {
	destNum := -1
	for i, loc := range g.Locations {
		if loc.Name() == dest {
			destNum = i
		}
	}
	if destNum == -1 {
		return fmt.Sprintf("локация %s не существует", dest)
	}

	ways := g.findWays()
	possible := false
	for _, w := range ways {
		if destNum == w {
			possible = true
		}
	}
	if !possible {
		return fmt.Sprintf("нет пути в %s", g.Locations[destNum].Name())
	}

	g.playerLocation = destNum
	return g.Locations[g.playerLocation].Enter() + " " + g.waysMessage()
}

func (g *Game) TakeItem(itemName string) string {
	item, err := g.Locations[g.playerLocation].TakeItem(itemName)
	if err != nil {
		switch err.Error() {
		case "no such item in this location":
			return "нет такого"
		default:
			panic(err)
		}
	}

	g.playerInventory = append(g.playerInventory, item)
	return "предмет добавлен в инвентарь: " + itemName
}
