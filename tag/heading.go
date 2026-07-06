package tags

import (
	"bufio"
	"context"

	"github.com/catcher3/cegla"
)

type H1 []cegla.PhrasingContent
type H2 []cegla.PhrasingContent
type H3 []cegla.PhrasingContent
type H4 []cegla.PhrasingContent
type H5 []cegla.PhrasingContent
type H6 []cegla.PhrasingContent

func (H1) Name() string { return "h1" }
func (H2) Name() string { return "h2" }
func (H3) Name() string { return "h3" }
func (H4) Name() string { return "h4" }
func (H5) Name() string { return "h5" }
func (H6) Name() string { return "h6" }

func (el H1) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el H2) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el H3) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el H4) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el H5) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el H6) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}

func (H1) IsFlow()    {}
func (H1) IsHeading() {}
func (H2) IsFlow()    {}
func (H2) IsHeading() {}
func (H3) IsFlow()    {}
func (H3) IsHeading() {}
func (H4) IsFlow()    {}
func (H4) IsHeading() {}
func (H5) IsFlow()    {}
func (H5) IsHeading() {}
func (H6) IsFlow()    {}
func (H6) IsHeading() {}

// Hgroup — структурное исключение: по спеке допускает только H1-H6 (+ P),
// поэтому типизирован через cegla.HeadingContent, а не FlowContent.
type Hgroup []cegla.HeadingContent

func (Hgroup) Name() string { return "hgroup" }
func (el Hgroup) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (Hgroup) IsFlow() {}
