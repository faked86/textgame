package game

import (
	"errors"

	"text_game/pkg/game/inventory"
	"text_game/pkg/game/location"
)

type Game struct {
	WorldMap        [][]bool
	Locations       []location.Location
	playerInventory *inventory.Inventory
	playerLocation  int
}

func NewGame(start *location.Location, invent *inventory.Inventory) *Game {
	w := &Game{
		WorldMap:        [][]bool{{false}},
		Locations:       []location.Location{*start},
		playerInventory: invent,
		playerLocation:  0,
	}
	return w
}

func (g *Game) AddLocation(l *location.Location) {
	g.Locations = append(g.Locations, *l)
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

func (g *Game) findWays() []location.Location {
	res := make([]location.Location, 0)
	for i, way := range g.WorldMap[g.playerLocation] {
		if way {
			res = append(res, g.Locations[i])
		}
	}
	return res
}

func (g *Game) WaysMessage() {

}

func (g *Game) LookAround() {

}
