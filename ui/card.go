package ui

import (
	"bufio"
	. "cegla/core/atr"
	. "cegla/core/tags"
	"context"
)

type Card struct {
	Id string
	// Title is the title of the card
	Title H1

	Description P

	// Clicks is the number of times the card has been clicked
	Clicks int
}

func (Card) Name() string {
	return "card"
}

func (cd Card) Render(ctx context.Context, w *bufio.Writer) error {
	// Собираем содержимое прямо в Render
	container := Div{
		Class("card"),
		// Передаем поля структуры как детей
		cd.Title,
		cd.Description,
		// Можно добавить еще что-то, например, счетчик кликов
		// Button{Text("Clicks: " + strconv.Itoa(cd.Clicks))},
	}

	return container.Render(ctx, w)
}
