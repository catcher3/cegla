package tags

import (
	"bufio"
	"cegla/core"
	"context"
)

// Примечание: A, Details, Embed, Iframe уже определены как полноценные
// элементы в других файлах (prasing.go, flow.go, embedded.go) — здесь
// только теги, у которых Interactive/FormAssociated — основная роль.
// Img с usemap интерактивен условно (зависит от наличия атрибута), это не
// выражается системой типов и потому не выделено в отдельный тип.

type Button []core.PhrasingContent
type Label []core.PhrasingContent
type Select []core.SelectContent
type Textarea []core.PhrasingContent // содержимое — text/значение по умолчанию
type Input []core.Attribute          // void-элемент, полностью управляется атрибутами

func (Button) Name() string   { return "button" }
func (Label) Name() string    { return "label" }
func (Select) Name() string   { return "select" }
func (Textarea) Name() string { return "textarea" }
func (Input) Name() string    { return "input" }

func (bn Button) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(bn.Name(), bn, ctx, w)
}
func (el Label) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Select) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Textarea) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Input) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderVoid(el.Name(), el, ctx, w)
}

func (Button) IsFlow()           {}
func (Button) IsPhrasing()       {}
func (Button) IsInteractive()    {}
func (Button) IsFormAssociated() {}

func (Label) IsFlow()           {}
func (Label) IsPhrasing()       {}
func (Label) IsInteractive()    {}
func (Label) IsFormAssociated() {}

func (Select) IsFlow()           {}
func (Select) IsPhrasing()       {}
func (Select) IsInteractive()    {}
func (Select) IsFormAssociated() {}

func (Textarea) IsFlow()           {}
func (Textarea) IsPhrasing()       {}
func (Textarea) IsInteractive()    {}
func (Textarea) IsFormAssociated() {}

func (Input) IsFlow()           {}
func (Input) IsPhrasing()       {}
func (Input) IsInteractive()    {}
func (Input) IsFormAssociated() {}
