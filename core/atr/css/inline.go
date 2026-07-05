// core/atr/css/inline.go
//
// Package css отвечает ТОЛЬКО за инлайн-стили (style="..."), в отличие от
// пакета tw (core/atr/css/tw), который генерирует классы Tailwind
// (class="..."). Раньше здесь был ещё параллельный набор хелперов
// (Style/Atom/CSS/Box/FlexRow/Shadow), дублировавший то, что уже делает tw
// через обычный atr.Attr — он удалён, чтобы не было двух независимых
// способов собрать "class" в одном проекте. Если нужен произвольный
// Tailwind-класс без готового хелпера — tw.Class("любой-класс"); если нужен
// именно инлайновый CSS (style="width: 240px") — хелперы этого файла.
package css

import (
	"cegla/core/atr"
	"fmt"
)

// Helper for generating style attributes
func prop(key, val string) atr.Attr {
	return atr.Custom("style", key+": "+val)
}

// Dimensions
func Width(val string) atr.Attr     { return prop("width", val) }
func WidthPx(px int) atr.Attr       { return prop("width", fmt.Sprintf("%dpx", px)) }
func WidthPct(pct float64) atr.Attr { return prop("width", fmt.Sprintf("%.2f%%", pct)) }
func HeightPx(px int) atr.Attr      { return prop("height", fmt.Sprintf("%dpx", px)) }

// Colors
func Color(val string) atr.Attr           { return prop("color", val) }
func BackgroundColor(val string) atr.Attr { return prop("background-color", val) }

// Spacing
func MarginPx(top, right, bottom, left int) atr.Attr {
	return prop("margin", fmt.Sprintf("%dpx %dpx %dpx %dpx", top, right, bottom, left))
}
func PaddingPx(px int) atr.Attr { return prop("padding", fmt.Sprintf("%dpx", px)) }

// Typography
func FontSizePx(px int) atr.Attr     { return prop("font-size", fmt.Sprintf("%dpx", px)) }
func FontWeight(weight int) atr.Attr { return prop("font-weight", fmt.Sprintf("%d", weight)) }

// Flexbox
func Display(val string) atr.Attr       { return prop("display", val) }
func FlexDirection(val string) atr.Attr { return prop("flex-direction", val) }
func AlignItems(val string) atr.Attr    { return prop("align-items", val) }

// Прочее
func BorderRadiusPx(px int) atr.Attr { return prop("border-radius", fmt.Sprintf("%dpx", px)) }
func Opacity(v float64) atr.Attr     { return prop("opacity", fmt.Sprintf("%.2f", v)) }
func ZIndex(v int) atr.Attr          { return prop("z-index", fmt.Sprintf("%d", v)) }
func Cursor(val string) atr.Attr     { return prop("cursor", val) }
func Overflow(val string) atr.Attr   { return prop("overflow", val) }
func Position(val string) atr.Attr   { return prop("position", val) }
