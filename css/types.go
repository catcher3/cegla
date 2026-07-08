package css

import (
	"bufio"
	"context"
	"strings"
)

type CSS interface {
	ToClass() string
}

// Style — набор CSS/tailwind-классов, который ведёт себя как атрибут
// class="...". Реализует Attrsibute напрямую (Style — слайс, поэтому не
// может встроить AttrsMarker как структурное поле — методы продублированы
// вручную), поэтому Style{...} можно класть прямо в литерал любого тега
// рядом с обычными детьми:
//
//	Button{
//	    css.Style{css.Padding("4"), css.Bg("blue-500")},
//	    Text("Click me"),
//	}
//
// cegla.RenderChildren распознаёт Style как Attrsibute с Key()=="class"
// и смёрджит его с любыми другими "class"-атрибутами (например tag.Class)
// через конкатенацию, а не затирает их.
type Style []CSS

func (s Style) Add(css CSS) Style {
	return append(s, css)
}

func (s Style) ToClass() string {
	var res []string
	for _, item := range s {
		res = append(res, item.ToClass())
	}
	return strings.Join(res, " ")
}

// --- Attrsibute ---

func (Style) Name() string    { return "class" }
func (Style) Key() string     { return "class" }
func (s Style) Value() string { return s.ToClass() }
func (Style) IsAttribute()    {}

func (s Style) Render(ctx context.Context, w *bufio.Writer) error {
	return nil
}

func (Style) IsDocumentContent()        {}
func (Style) IsFlow()                   {}
func (Style) IsPhrasing()               {}
func (Style) IsMetadata()               {}
func (Style) IsSectioning()             {}
func (Style) IsHeading()                {}
func (Style) IsEmbedded()               {}
func (Style) IsInteractive()            {}
func (Style) IsFormAssociated()         {}
func (Style) IsTableChild()             {}
func (Style) IsTableRowContent()        {}
func (Style) IsSelectContent()          {}
func (Style) IsDescriptionListContent() {}
func (Style) IsListChild()              {}
func (Style) IsOptionContent()          {}

// 1. Делаем string совместимым с CSS
type Raw string

func (r Raw) ToClass() string { return string(r) }
