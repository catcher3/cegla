package tag

import (
	"bufio"
	"context"

	"github.com/catcher3/cegla/render"
)

type Article Flow
type Section Flow
type Nav Flow
type Aside Flow

func (Article) Name() string { return "article" }
func (Section) Name() string { return "section" }
func (Nav) Name() string     { return "nav" }
func (Aside) Name() string   { return "aside" }

func (el Article) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Section) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Nav) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Aside) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (Article) IsFlow()       {}
func (Article) IsSectioning() {}
func (Section) IsFlow()       {}
func (Section) IsSectioning() {}
func (Nav) IsFlow()           {}
func (Nav) IsSectioning()     {}
func (Aside) IsFlow()         {}
func (Aside) IsSectioning()   {}
