package tags

import (
	"bufio"
	"cegla/core"
	"context"
)

type H1 []core.PhrasingContent
type H2 []core.PhrasingContent
type H3 []core.PhrasingContent
type H4 []core.PhrasingContent
type H5 []core.PhrasingContent
type H6 []core.PhrasingContent

func (H1) Name() string { return "h1" }
func (H2) Name() string { return "h2" }
func (H3) Name() string { return "h3" }
func (H4) Name() string { return "h4" }
func (H5) Name() string { return "h5" }
func (H6) Name() string { return "h6" }

func (el H1) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el H2) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el H3) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el H4) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el H5) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el H6) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
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
// поэтому типизирован через core.HeadingContent, а не FlowContent.
type Hgroup []core.HeadingContent

func (Hgroup) Name() string { return "hgroup" }
func (el Hgroup) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (Hgroup) IsFlow() {}
