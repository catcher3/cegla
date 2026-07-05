package main

import (
	"cegla/core"
	. "cegla/core"
	. "cegla/core/atr"
	"cegla/core/atr/css"
	"cegla/core/atr/css/tw"
	"cegla/core/atr/htmx"
	"cegla/core/tags"
	. "cegla/core/tags"
	"fmt"
)

func PrimaryButton(label string) tags.Button {
	return tags.Button{
		// Tailwind classes (merged via space)
		tw.Bg("blue-600"),
		tw.Hover(tw.Bg("blue-700")),
		tw.Text("white"),
		tw.FontBold(),
		tw.Py("2"),
		tw.Px("4"),
		tw.Class("rounded"), // standard class fallback

		// Inline styles (merged via semicolon)
		css.WidthPct(100),
		css.MarginPx(10, 0, 10, 0),

		core.Text(label),
	}
}

func main() {
	pageHello := HTML{
		Lang("en"),
		Head{
			Title{core.Text("Hello World")},
		},
		Body{
			H1{
				Class("main-title"),
				ID("HelloWorld"),
				Text("Hello World"),
			},
			Button{
				tw.P("4"),
				tw.Bg("blue-500"),
				Text("Click me"),
				htmx.Post("/api/v1/postbutton"),
			},

			PrimaryButton("Hello wrold"),
		},
	}

	fmt.Println(pageHello)
}
