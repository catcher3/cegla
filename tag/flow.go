package tag

import (
	"bufio"
	"context"
	"fmt"

	. "github.com/catcher3/cegla"
	"github.com/catcher3/cegla/render"
)

// --- Простые flow-контейнеры ---
type Div Flow
type Blockquote Flow
type Main Flow
type Dialog Flow
type Header Flow
type Footer Flow
type Address Flow
type Form Flow // нельзя вложенный <form> — не выражается категорией, только runtime/линтом

func (Div) Name() string        { return "div" }
func (Blockquote) Name() string { return "blockquote" }
func (Main) Name() string       { return "main" }
func (Dialog) Name() string     { return "dialog" }
func (Header) Name() string     { return "header" }
func (Footer) Name() string     { return "footer" }
func (Address) Name() string    { return "address" }
func (Form) Name() string       { return "form" }

func (el Div) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Blockquote) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Main) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Dialog) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Header) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Footer) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Address) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Form) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (Div) IsFlow()        {}
func (Blockquote) IsFlow() {}
func (Main) IsFlow()       {}
func (Dialog) IsFlow()     {}
func (Header) IsFlow()     {}
func (Footer) IsFlow()     {}
func (Address) IsFlow()    {}
func (Form) IsFlow()       {}

// --- P / Pre — только Phrasing ---

type P Phrasing
type Pre Phrasing

func (P) Name() string   { return "p" }
func (Pre) Name() string { return "pre" }

func (el P) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Pre) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (P) IsFlow()   {}
func (Pre) IsFlow() {}

// --- Hr — void-элемент, только атрибуты ---

type Hr Attrs

func (Hr) Name() string { return "hr" }
func (el Hr) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderVoid(el.Name(), el, ctx, w)
}
func (Hr) IsFlow() {}

// --- Списки: Ol/Ul -> только LI ---

type Ol List
type Ul List
type LI Flow // содержимое <li> — обычный Flow

func (Ol) Name() string { return "ol" }
func (Ul) Name() string { return "ul" }
func (LI) Name() string { return "li" }

func (el Ol) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Ul) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el LI) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (Ol) IsFlow()      {}
func (Ul) IsFlow()      {}
func (LI) IsListChild() {}
func (LI) IsFlow()      {} // сам LI валиден и как FlowContent (напр. внутри TD/Div)

// --- Dl -> только Dt/Dd ---

type Dl DescriptionList
type Dt Phrasing
type Dd Flow

func (Dl) Name() string { return "dl" }
func (Dt) Name() string { return "dt" }
func (Dd) Name() string { return "dd" }

func (el Dl) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Dt) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Dd) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (Dl) IsFlow()                   {}
func (Dt) IsDescriptionListContent() {}
func (Dd) IsDescriptionListContent() {}

// --- Figure -> максимум один FigCaption (проверяется в Render) ---

type Figure Flow
type FigCaption Flow

func (Figure) Name() string     { return "figure" }
func (FigCaption) Name() string { return "figcaption" }

func (el FigCaption) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (el Figure) Render(ctx context.Context, w *bufio.Writer) error {
	captions := 0
	for _, child := range el {
		if _, ok := child.(FigCaption); ok {
			captions++
		}
	}
	if captions > 1 {
		return fmt.Errorf("cegla: <figure> allows at most one <figcaption>, got %d", captions)
	}
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (Figure) IsFlow()     {}
func (FigCaption) IsFlow() {}

// --- Details -> Summary (если есть) должен идти первым (проверяется в Render) ---

type Details Flow
type Summary Phrasing

func (Details) Name() string { return "details" }
func (Summary) Name() string { return "summary" }

func (el Summary) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (el Details) Render(ctx context.Context, w *bufio.Writer) error {
	for i, child := range el {
		if _, ok := child.(Summary); ok && i != 0 {
			return fmt.Errorf("cegla: <summary> must be the first child of <details>")
		}
	}
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (Details) IsFlow()        {}
func (Summary) IsFlow()        {}
func (Details) IsInteractive() {}
