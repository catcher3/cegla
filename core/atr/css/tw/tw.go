// core/atr/tw/tailwind.go
package tw

import "cegla/core/atr"

// Class is a fallback for any arbitrary CSS class
// that doesn't have a dedicated helper function yet.
func Class(v string) atr.Attr {
	return atr.Class(v)
}

// Layout
func Flex() atr.Attr   { return atr.Class("flex") }
func Grid() atr.Attr   { return atr.Class("grid") }
func Hidden() atr.Attr { return atr.Class("hidden") }
func Block() atr.Attr  { return atr.Class("block") }

// Flexbox & Grid
func ItemsCenter() atr.Attr    { return atr.Class("items-center") }
func JustifyCenter() atr.Attr  { return atr.Class("justify-center") }
func JustifyBetween() atr.Attr { return atr.Class("justify-between") }
func FlexCol() atr.Attr        { return atr.Class("flex-col") }
func FlexRow() atr.Attr        { return atr.Class("flex-row") }
func Gap(size string) atr.Attr { return atr.Class("gap-" + size) }

// Spacing (Padding/Margin)
func P(size string) atr.Attr  { return atr.Class("p-" + size) }
func Px(size string) atr.Attr { return atr.Class("px-" + size) }
func Py(size string) atr.Attr { return atr.Class("py-" + size) }
func Pt(size string) atr.Attr { return atr.Class("pt-" + size) }
func Pb(size string) atr.Attr { return atr.Class("pb-" + size) }
func M(size string) atr.Attr  { return atr.Class("m-" + size) }
func Mx(size string) atr.Attr { return atr.Class("mx-" + size) }
func My(size string) atr.Attr { return atr.Class("my-" + size) }
func Mt(size string) atr.Attr { return atr.Class("mt-" + size) }
func Mb(size string) atr.Attr { return atr.Class("mb-" + size) }

// Sizing
func W(size string) atr.Attr    { return atr.Class("w-" + size) }
func H(size string) atr.Attr    { return atr.Class("h-" + size) }
func MinW(size string) atr.Attr { return atr.Class("min-w-" + size) }
func MaxW(size string) atr.Attr { return atr.Class("max-w-" + size) }

// Borders & radius
func Border(size string) atr.Attr       { return atr.Class("border-" + size) }
func BorderColor(color string) atr.Attr { return atr.Class("border-" + color) }
func Rounded(size string) atr.Attr      { return atr.Class("rounded-" + size) }

// Typography
func TextCenter() atr.Attr      { return atr.Class("text-center") }
func Text(size string) atr.Attr { return atr.Class("text-" + size) }
func FontBold() atr.Attr        { return atr.Class("font-bold") }
func Truncate() atr.Attr        { return atr.Class("truncate") }

// Colors (Background, Text, Border)
func Bg(color string) atr.Attr    { return atr.Class("bg-" + color) }
func TextColor(c string) atr.Attr { return atr.Class("text-" + c) }

// Misc
func Shadow(size string) atr.Attr { return atr.Class("shadow-" + size) }
func Transition() atr.Attr        { return atr.Class("transition-all duration-150") }
func Cursor(v string) atr.Attr    { return atr.Class("cursor-" + v) }

// Modifiers (Hover, Focus и т.п.) — оборачивают любой atr.Attr, добавляя
// префикс. Работают для ЛЮБОГО хелпера выше: tw.Hover(tw.Bg("gray-100")),
// tw.Focus(tw.BorderColor("blue-500")).
func Hover(attr atr.Attr) atr.Attr {
	return atr.Class("hover:" + attr.Value())
}
func Focus(attr atr.Attr) atr.Attr {
	return atr.Class("focus:" + attr.Value())
}
func Active(attr atr.Attr) atr.Attr {
	return atr.Class("active:" + attr.Value())
}
func Disabled(attr atr.Attr) atr.Attr {
	return atr.Class("disabled:" + attr.Value())
}

// --- Готовые пресеты (были css.Box/FlexRow/Shadow — перенесены сюда,
// потому что это тоже просто class="...", как и всё остальное в tw;
// отдельный тип-обёртка (css.Style) для этого не нужен) ---

// Box — базовый контейнер: паддинг, рамка, скругление, белый фон.
func Box() atr.Attr { return atr.Class("p-4 border rounded-lg bg-white") }

// Card — то же самое + тень, для карточек.
func Card() atr.Attr { return atr.Class("p-4 border rounded-lg bg-white shadow-md") }

// SoftShadow — готовая тень с плавным увеличением при наведении.
func SoftShadow() atr.Attr { return atr.Class("shadow-md hover:shadow-lg transition-shadow") }
