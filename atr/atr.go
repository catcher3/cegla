// atr.go
// Аттрибуты — это отдельные сущности, которые могут быть применены к элементам HTML. Они представляют собой пары ключ-значение, которые определяют свойства и поведение элементов.
package atr

import (
	"bufio"
	"context"

	"github.com/catcher3/cegla"
)

// Attr — простой тип для одного атрибута.
//
// Встраивает cegla.AttrMarker, поэтому автоматически реализует все
// маркерные интерфейсы контента (FlowContent, PhrasingContent,
// MetadataContent, DocumentContent, TableChild, TableRowContent,
// SelectContent, ListChild, OptionContent и т.д.) — атрибут можно положить
// в литерал любого тега, независимо от его категории.
type Attr struct {
	cegla.AttrMarker
	K, V string
}

func (a Attr) Name() string { return a.K }

func (a Attr) IsAttribute()  {}
func (a Attr) Key() string   { return a.K }
func (a Attr) Value() string { return a.V }

// Render для атрибута ничего не пишет напрямую — атрибуты рендерятся внутри
// открывающего тега через cegla.RenderChildren/cegla.RenderVoid. Метод нужен
// только для удовлетворения cegla.Node.
func (a Attr) Render(ctx context.Context, w *bufio.Writer) error {
	return nil
}

// --- Хелперы для удобного создания ---

func Class(v string) Attr       { return Attr{K: "class", V: v} }
func ID(v string) Attr          { return Attr{K: "id", V: v} }
func StyleAtr(v string) Attr    { return Attr{K: "style", V: v} }
func Href(v string) Attr        { return Attr{K: "href", V: v} }
func DataAtr(k, v string) Attr  { return Attr{K: "data-" + k, V: v} }
func Lang(code string) Attr     { return Attr{K: "lang", V: code} }
func Src(v string) Attr         { return Attr{K: "src", V: v} }
func Srcset(v string) Attr      { return Attr{K: "srcset", V: v} }
func Sizes(v string) Attr       { return Attr{K: "sizes", V: v} }
func Alt(v string) Attr         { return Attr{K: "alt", V: v} }
func Rel(v string) Attr         { return Attr{K: "rel", V: v} }
func TypeAttr(v string) Attr    { return Attr{K: "type", V: v} }
func NameAttr(v string) Attr    { return Attr{K: "name", V: v} }
func Value(v string) Attr       { return Attr{K: "value", V: v} }
func Placeholder(v string) Attr { return Attr{K: "placeholder", V: v} }
func Content(v string) Attr     { return Attr{K: "content", V: v} }
func Charset(v string) Attr     { return Attr{K: "charset", V: v} }
func Target(v string) Attr      { return Attr{K: "target", V: v} }
func Scope(v string) Attr       { return Attr{K: "scope", V: v} }
func For(v string) Attr         { return Attr{K: "for", V: v} }    // <label for="...">
func Method(v string) Attr      { return Attr{K: "method", V: v} } // <form method="post">
func Action(v string) Attr      { return Attr{K: "action", V: v} } // <form action="...">
func Enctype(v string) Attr     { return Attr{K: "enctype", V: v} }
func Min(v string) Attr         { return Attr{K: "min", V: v} }
func Max(v string) Attr         { return Attr{K: "max", V: v} }
func Step(v string) Attr        { return Attr{K: "step", V: v} }
func MinLength(v string) Attr   { return Attr{K: "minlength", V: v} }
func MaxLength(v string) Attr   { return Attr{K: "maxlength", V: v} }
func Pattern(v string) Attr     { return Attr{K: "pattern", V: v} }
func Rows(v string) Attr        { return Attr{K: "rows", V: v} }
func Cols(v string) Attr        { return Attr{K: "cols", V: v} }
func ColSpan(v string) Attr     { return Attr{K: "colspan", V: v} }
func RowSpan(v string) Attr     { return Attr{K: "rowspan", V: v} }
func Download(v string) Attr    { return Attr{K: "download", V: v} }
func Media(v string) Attr       { return Attr{K: "media", V: v} }
func Loading(v string) Attr     { return Attr{K: "loading", V: v} } // "lazy"/"eager"
func TabIndex(v string) Attr    { return Attr{K: "tabindex", V: v} }
func Role(v string) Attr        { return Attr{K: "role", V: v} }
func AriaLabel(v string) Attr   { return Attr{K: "aria-label", V: v} }
func AriaHidden(v string) Attr  { return Attr{K: "aria-hidden", V: v} }

// --- Булевы атрибуты ---
//
// В HTML присутствие атрибута = true независимо от его значения:
// <input disabled="false"> ВСЁ РАВНО отключает поле — браузер не смотрит
// на значение таких атрибутов, только на факт наличия. Поэтому у Disabled/
// Checked/Required и т.д. нет параметра bool — их нужно просто НЕ вызывать,
// когда условие ложно, а не вызывать с "false":
//
//	Input{
//		atr.TypeAttr("text"),
//		condIf(user.Locked, atr.Disabled()), // включать в слайс только когда true
//	}
//
// Собрать такой слайс с условием проще всего обычным append в коде
// компонента (как это уже делает Menu), а не пытаться выразить условие
// прямо в литерале.
func Disabled() Attr  { return Attr{K: "disabled", V: "disabled"} }
func Checked() Attr   { return Attr{K: "checked", V: "checked"} }
func Selected() Attr  { return Attr{K: "selected", V: "selected"} }
func ReadOnly() Attr  { return Attr{K: "readonly", V: "readonly"} }
func Required() Attr  { return Attr{K: "required", V: "required"} }
func Multiple() Attr  { return Attr{K: "multiple", V: "multiple"} }
func Autofocus() Attr { return Attr{K: "autofocus", V: "autofocus"} }

// Custom — универсальный хелпер для произвольного атрибута,
// когда именованного хелпера ещё нет.
func Custom(key, value string) Attr { return Attr{K: key, V: value} }
