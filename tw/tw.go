// core/atr/tw/tailwind.go
package tw

import (
	"github.com/catcher3/cegla/tag"
)

// --- Tailwind CSS helpers (class="...") ---

// Class — универсальный хелпер для любого класса. Используется, когда нет
// готового хелпера для конкретного класса.

// Class is a fallback for any arbitrary CSS class
// that doesn't have a dedicated helper function yet.
func Class(v string) tag.Attr {
	return tag.Class(v)
}

// Layout
func Flex() tag.Attr   { return tag.Class("flex") }
func Grid() tag.Attr   { return tag.Class("grid") }
func Hidden() tag.Attr { return tag.Class("hidden") }
func Block() tag.Attr  { return tag.Class("block") }

// Flexbox & Grid
func ItemsCenter() tag.Attr    { return tag.Class("items-center") }
func JustifyCenter() tag.Attr  { return tag.Class("justify-center") }
func JustifyBetween() tag.Attr { return tag.Class("justify-between") }
func FlexCol() tag.Attr        { return tag.Class("flex-col") }
func FlexRow() tag.Attr        { return tag.Class("flex-row") }
func Gap(size string) tag.Attr { return tag.Class("gap-" + size) }

// Spacing (Padding/Margin)
func P(size string) tag.Attr  { return tag.Class("p-" + size) }
func Px(size string) tag.Attr { return tag.Class("px-" + size) }
func Py(size string) tag.Attr { return tag.Class("py-" + size) }
func Pt(size string) tag.Attr { return tag.Class("pt-" + size) }
func Pb(size string) tag.Attr { return tag.Class("pb-" + size) }
func M(size string) tag.Attr  { return tag.Class("m-" + size) }
func Mx(size string) tag.Attr { return tag.Class("mx-" + size) }
func My(size string) tag.Attr { return tag.Class("my-" + size) }
func Mt(size string) tag.Attr { return tag.Class("mt-" + size) }
func Mb(size string) tag.Attr { return tag.Class("mb-" + size) }

// Sizing
func W(size string) tag.Attr    { return tag.Class("w-" + size) }
func H(size string) tag.Attr    { return tag.Class("h-" + size) }
func MinW(size string) tag.Attr { return tag.Class("min-w-" + size) }
func MaxW(size string) tag.Attr { return tag.Class("max-w-" + size) }

// Borders & radius
func Border(size string) tag.Attr       { return tag.Class("border-" + size) }
func BorderColor(color string) tag.Attr { return tag.Class("border-" + color) }
func Rounded(size string) tag.Attr      { return tag.Class("rounded-" + size) }

// Typography
func TextCenter() tag.Attr      { return tag.Class("text-center") }
func Text(size string) tag.Attr { return tag.Class("text-" + size) }
func FontBold() tag.Attr        { return tag.Class("font-bold") }
func Truncate() tag.Attr        { return tag.Class("truncate") }

// Colors (Background, Text, Border)
func Bg(color string) tag.Attr    { return tag.Class("bg-" + color) }
func TextColor(c string) tag.Attr { return tag.Class("text-" + c) }

// Misc
func Shadow(size string) tag.Attr { return tag.Class("shadow-" + size) }
func Transition() tag.Attr        { return tag.Class("transition-all duration-150") }
func Cursor(v string) tag.Attr    { return tag.Class("cursor-" + v) }

// Modifiers (Hover, Focus и т.п.) — оборачивают любой tag.Attr, добавляя
// префикс. Работают для ЛЮБОГО хелпера выше: tw.Hover(tw.Bg("gray-100")),
// tw.Focus(tw.BorderColor("blue-500")).
func Hover(attr tag.Attr) tag.Attr {
	return tag.Class("hover:" + attr.Value())
}
func Focus(attr tag.Attr) tag.Attr {
	return tag.Class("focus:" + attr.Value())
}
func Active(attr tag.Attr) tag.Attr {
	return tag.Class("active:" + attr.Value())
}
func Disabled(attr tag.Attr) tag.Attr {
	return tag.Class("disabled:" + attr.Value())
}

// --- Готовые пресеты (были css.Box/FlexRow/Shadow — перенесены сюда,
// потому что это тоже просто class="...", как и всё остальное в tw;
// отдельный тип-обёртка (css.Style) для этого не нужен) ---

// Box — базовый контейнер: паддинг, рамка, скругление, белый фон.
func Box() tag.Attr { return tag.Class("p-4 border rounded-lg bg-white") }

// Card — то же самое + тень, для карточек.
func Card() tag.Attr { return tag.Class("p-4 border rounded-lg bg-white shadow-md") }

// SoftShadow — готовая тень с плавным увеличением при наведении.
func SoftShadow() tag.Attr { return tag.Class("shadow-md hover:shadow-lg transition-shadow") }
