package tags

import (
	"bufio"
	"context"

	"github.com/catcher3/cegla"
)

type Article []cegla.FlowContent
type Section []cegla.FlowContent
type Nav []cegla.FlowContent
type Aside []cegla.FlowContent

func (Article) Name() string { return "article" }
func (Section) Name() string { return "section" }
func (Nav) Name() string     { return "nav" }
func (Aside) Name() string   { return "aside" }

func (el Article) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Section) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Nav) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Aside) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}

func (Article) IsFlow()       {}
func (Article) IsSectioning() {}
func (Section) IsFlow()       {}
func (Section) IsSectioning() {}
func (Nav) IsFlow()           {}
func (Nav) IsSectioning()     {}
func (Aside) IsFlow()         {}
func (Aside) IsSectioning()   {}
