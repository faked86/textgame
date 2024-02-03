package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"text_game/pkg/game"
	"text_game/pkg/game/inventory"
	"text_game/pkg/game/item"
	"text_game/pkg/game/location"
	"text_game/pkg/game/location/tags"
)

var g game.Game

func main() {
	var input string
	var err error

	initGame()

	fmt.Println("тебе надо собраться в универ")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(handleCommand(input))
	}
}

func initGame() {
	kitchen := location.NewBaseLocation(
		"кухня",
		"кухня, ничего интересного",
		"ты находишься на кухне, надо собрать рюкзак и идти в универ",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)
	lobby := location.NewBaseLocation(
		"коридор",
		"ничего интересного",
		"ничего интересного",
		"здесь есть",
		inventory.Inventory{},
		tags.Home,
	)
	room := location.NewBaseLocation(
		"комната",
		"ты в своей комнате",
		"пустая комната",
		"на столе",
		inventory.Inventory{
			item.NewBaseItem("ключи"),
			item.NewBaseItem("конспекты"),
			item.NewBaseItem("рюкзак")},
		tags.Home,
	)
	outside := location.NewBaseLocation(
		"улица",
		"на улице весна",
		"ничего интересного",
		"здесь есть",
		inventory.Inventory{},
		tags.Outside,
	)

	g = *game.NewGame(kitchen, inventory.Inventory{})

	g.AddLocation(lobby)
	g.AddWay(0, 1)

	g.AddLocation(room)
	g.AddWay(1, 2)

	g.AddLocation(outside)
	g.AddWay(1, 3)
}

func handleCommand(command string) string {
	args := strings.Split(command, " ")
	if len(args) == 0 {
		return "введите команду"
	}
	switch args[0] {
	case "осмотреться":
		if len(args) > 1 {
			return "команда осмотреться не принимает аргументов"
		}
		return g.LookAround()
	case "идти":
		if len(args) != 2 {
			return "команда идти принимает 1 аргумент"
		}
		return g.Walk(args[1])
	case "взять":
		if len(args) != 2 {
			return "команда идти принимает 1 аргумент"
		}
		return g.TakeItem(args[1])
	}
	return "неизвестная команда"
}
