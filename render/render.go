// render.go
package render

import (
	"bufio"
	"context"
	"html"

	. "github.com/catcher3/cegla"
)

// Вспомогательная функция для тегов-слайсов, чтобы не дублировать логику
func getLastID[T Container](content []T) string {
	id := ""
	for _, item := range content {
		// Проверяем, является ли элемент атрибутом "id"
		if attr, ok := any(item).(Attribute); ok && attr.Key() == "id" {
			id = attr.Value()
		}
	}
	return id
}

type attrEntry struct {
	key   string
	value string
}

// mergeAttr добавляет атрибут в накопленный список.
// Правило "последний верный" — для обычных атрибутов.
// Правило конкатенации — для class (через пробел) и style (через "; ").
func mergeAttr(entries []attrEntry, key, value string) []attrEntry {
	for i := range entries {
		if entries[i].key != key {
			continue
		}
		switch key {
		case "class":
			entries[i].value += " " + value
		case "style":
			entries[i].value += "; " + value
		default:
			entries[i].value = value // последний верный
		}
		return entries
	}
	return append(entries, attrEntry{key: key, value: value})
}

// RenderChildren — единая реализация рендера для любого контейнерного тега:
// FlowContent, PhrasingContent, MetadataContent, TableChild, TableRowContent,
// ListChild, SelectContent, DescriptionListContent, OptionContent и т.д.
// Атрибуты (Attrsibute) и обычные дети распознаются по типу прямо внутри
// одного слайса, поэтому Div{Class("a"), P{...}} работает без отдельного
// поля Attrs. Функция generic, чтобы не плодить RenderFlow/RenderPhrasing/...
// по одной на каждую категорию контента.
func RenderChildren[T Container](tagName string, content []T, ctx context.Context, w *bufio.Writer) error {
	w.WriteByte('<')
	w.WriteString(tagName)

	var entries []attrEntry // nil, если атрибутов нет — лишней аллокации не будет
	for _, item := range content {
		if attr, ok := any(item).(Attribute); ok {
			entries = mergeAttr(entries, attr.Key(), attr.Value())
		}
	}
	for _, e := range entries {
		w.WriteByte(' ')
		w.WriteString(e.key)
		w.WriteString(`="`)
		w.WriteString(html.EscapeString(e.value))
		w.WriteByte('"')
	}
	w.WriteByte('>')

	for _, item := range content {
		if _, ok := any(item).(Attribute); ok {
			continue
		}
		if err := item.Render(ctx, w); err != nil {
			return err
		}
	}

	w.WriteByte('<')
	w.WriteByte('/')
	w.WriteString(tagName)
	w.WriteByte('>')
	return nil
}

// RenderVoid — рендер void-элементов (br, hr, img, input, meta, link, base,
// col, source, track, embed, wbr). У них нет и не может быть детей — только
// атрибуты, поэтому вызывающий тип должен быть слайсом Attribute
// (например type Img []Attrsibute), а не FlowContent/PhrasingContent.
// Это гарантирует на этапе компиляции, что внутрь void-тега нельзя случайно
// положить дочерний элемент.
func RenderVoid(tagName string, attrs []Attribute, ctx context.Context, w *bufio.Writer) error {
	w.WriteByte('<')
	w.WriteString(tagName)

	var entries []attrEntry
	for _, attr := range attrs {
		entries = mergeAttr(entries, attr.Key(), attr.Value())
	}
	for _, e := range entries {
		w.WriteByte(' ')
		w.WriteString(e.key)
		w.WriteString(`="`)
		w.WriteString(html.EscapeString(e.value))
		w.WriteByte('"')
	}
	w.WriteString(" />")
	return nil
}

// RenderComposition — единая реализация Render для любой Composition[T]:
// построить контейнер, отрендерить его. Тип T выводится компилятором из
// конкретной реализации BuildContainer вызывающего типа — вызывать можно
// без явного указания параметра: cegla.RenderComposition(m, ctx, w).
func RenderComposition[T Container](c Composition[T], ctx context.Context, w *bufio.Writer) error {
	return c.BuildContainer().Render(ctx, w)
}
