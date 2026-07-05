package tags

import (
	"bufio"
	"cegla/core"
	"context"
	"fmt"
)

// Fieldset — структурное исключение: <legend>, если есть, должен идти
// первым (проверяется в Render, как Table.Caption/Details.Summary).
type Fieldset []core.FlowContent
type Legend []core.PhrasingContent

// Option — содержимое <select>/<optgroup> напрямую, либо <optgroup>.
type Option []core.PhrasingContent
type Optgroup []core.OptionContent

type Output []core.PhrasingContent
type Progress []core.PhrasingContent // transparent-ish, phrasing fallback
type Meter []core.PhrasingContent

func (Fieldset) Name() string { return "fieldset" }
func (Legend) Name() string   { return "legend" }
func (Option) Name() string   { return "option" }
func (Optgroup) Name() string { return "optgroup" }
func (Output) Name() string   { return "output" }
func (Progress) Name() string { return "progress" }
func (Meter) Name() string    { return "meter" }

func (el Legend) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Fieldset) Render(ctx context.Context, w *bufio.Writer) error {
	for i, child := range el {
		if _, ok := child.(Legend); ok && i != 0 {
			return fmt.Errorf("cegla: <legend> must be the first child of <fieldset>")
		}
	}
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Option) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Optgroup) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Output) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Progress) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Meter) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
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
