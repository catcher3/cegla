package ui

import (
	"bufio"
	"context"

	"github.com/catcher3/cegla"
	"github.com/catcher3/cegla/atr"
	"github.com/catcher3/cegla/atr/css/tw"
	"github.com/catcher3/cegla/tag"
)

// Avatar — это композиция. Она хранит состояние и знает, как собрать себя в tag.Div.
type Avatar struct {
	AvatarClass      string
	ContainerClass   string
	Source           string
	Placeholder      string
	PlaceholderClass string
}

func (Avatar) Name() string { return "div" }

// BuildContainer реализует логику построения дерева.
func (a Avatar) BuildContainer() tag.Div {
	// Внутренний контейнер с изображением
	container := tag.Div{
		tw.Class(a.ContainerClass),
		tw.Rounded("full"),          // <--- Принудительно делаем круглым
		tw.Class("overflow-hidden"), // <--- Обязательно для обрезки фото
		tag.Img{
			atr.Src(a.Source),
			tw.Class("object-cover w-full h-full"), // Чтобы фото заполняло круг
		},
	}

	// Условный рендеринг placeholder
	if a.Placeholder != "" {
		container = append(container, tag.Span{
			tw.Class(a.PlaceholderClass),
			tag.Text(a.Placeholder),
		})
	}

	// Корневой элемент аватара
	return tag.Div{
		tw.Class("avatar"),
		tw.Class(a.AvatarClass),
		container,
	}
}

func (a Avatar) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderComposition(a, ctx, w)
}

func (Avatar) IsFlow()     {}
func (Avatar) IsPhrasing() {}

// --- AvatarGroup ---

type AvatarGroup struct {
	Class    string
	Children []cegla.FlowContent // В cegla дети передаются как слайс в структуру
}

func (AvatarGroup) Name() string { return "div" }

func (g AvatarGroup) BuildContainer() tag.Div {
	group := tag.Div{
		tw.Class("avatar-group rtl:space-x-reverse"),
		tw.Class(g.Class),
	}

	// Добавляем детей (обычно это список Avatar)
	group = append(group, g.Children...)

	return group
}

func (g AvatarGroup) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderComposition(g, ctx, w)
}

func (AvatarGroup) IsFlow() {}
