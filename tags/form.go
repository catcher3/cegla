package tags

import (
	"bufio"
	"context"
	"fmt"

	"github.com/catcher3/cegla"
)

// Fieldset — структурное исключение: <legend>, если есть, должен идти
// первым (проверяется в Render, как Table.Caption/Details.Summary).
type Fieldset []cegla.FlowContent
type Legend []cegla.PhrasingContent

// Option — содержимое <select>/<optgroup> напрямую, либо <optgroup>.
type Option []cegla.PhrasingContent
type Optgroup []cegla.OptionContent

type Output []cegla.PhrasingContent
type Progress []cegla.PhrasingContent // transparent-ish, phrasing fallback
type Meter []cegla.PhrasingContent

func (Fieldset) Name() string { return "fieldset" }
func (Legend) Name() string   { return "legend" }
func (Option) Name() string   { return "option" }
func (Optgroup) Name() string { return "optgroup" }
func (Output) Name() string   { return "output" }
func (Progress) Name() string { return "progress" }
func (Meter) Name() string    { return "meter" }

func (el Legend) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Fieldset) Render(ctx context.Context, w *bufio.Writer) error {
	for i, child := range el {
		if _, ok := child.(Legend); ok && i != 0 {
			return fmt.Errorf("cegla: <legend> must be the first child of <fieldset>")
		}
	}
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Option) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Optgroup) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Output) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Progress) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Meter) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}

func (Fieldset) IsFlow()           {}
func (Fieldset) IsFormAssociated() {}
func (Legend) IsFlow()             {}

func (Option) IsSelectContent()   {}
func (Option) IsOptionContent()   {}
func (Optgroup) IsSelectContent() {}

func (Output) IsFlow()             {}
func (Output) IsPhrasing()         {}
func (Output) IsFormAssociated()   {}
func (Progress) IsFlow()           {}
func (Progress) IsPhrasing()       {}
func (Progress) IsFormAssociated() {}
func (Meter) IsFlow()              {}
func (Meter) IsPhrasing()          {}
