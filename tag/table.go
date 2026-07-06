package tags

import (
	"bufio"
	"context"
	"fmt"

	"github.com/catcher3/cegla"
)

// Table типизирован как []cegla.TableChild, а не []cegla.FlowContent —
// это сохраняет категорийную защиту: внутрь можно положить только
// Caption/ColGroup/THead/TBody/TFoot/TR, а не произвольный Div/P. Порядок
// (caption — первым и максимум один раз) — runtime-проверка в Render,
// т.к. маркерные интерфейсы порядок выразить не могут.
type Table []cegla.TableChild

type Caption []cegla.FlowContent
type ColGroup []Col
type Col []cegla.Attribute // void-элемент (только атрибут span), без детей
type THead []TR
type TBody []TR
type TFoot []TR
type TR []cegla.TableRowContent
type TD []cegla.FlowContent // содержимое ячейки — обычный Flow
type TH []cegla.FlowContent // + атрибут scope, содержимое — обычный Flow

func (Table) Name() string    { return "table" }
func (Caption) Name() string  { return "caption" }
func (ColGroup) Name() string { return "colgroup" }
func (Col) Name() string      { return "col" }
func (THead) Name() string    { return "thead" }
func (TBody) Name() string    { return "tbody" }
func (TFoot) Name() string    { return "tfoot" }
func (TR) Name() string       { return "tr" }
func (TD) Name() string       { return "td" }
func (TH) Name() string       { return "th" }

func (el Table) Render(ctx context.Context, w *bufio.Writer) error {
	captions := 0
	for i, child := range el {
		if _, ok := child.(Caption); ok {
			captions++
			if i != 0 {
				return fmt.Errorf("cegla: <caption> must be the first child of <table>")
			}
		}
	}
	if captions > 1 {
		return fmt.Errorf("cegla: <table> allows only one <caption>, got %d", captions)
	}
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Caption) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el ColGroup) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Col) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderVoid(el.Name(), el, ctx, w)
}
func (el THead) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el TBody) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el TFoot) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el TR) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el TD) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el TH) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}

func (Caption) IsTableChild()  {}
func (ColGroup) IsTableChild() {}
func (THead) IsTableChild()    {}
func (TBody) IsTableChild()    {}
func (TFoot) IsTableChild()    {}
func (TR) IsTableChild()       {} // <tr> валиден и напрямую в <table>, и внутри thead/tbody/tfoot

func (TD) IsTableRowContent() {}
func (TH) IsTableRowContent() {}

// TD/TH также валидны как обычный Flow, если понадобится (напр. составные
// компоненты, которые хотят работать generically с Flow) — сознательно НЕ
// делаем этого: TD не должен просачиваться в Body{} (см. предыдущее
// обсуждение архитектуры), поэтому IsFlow() здесь намеренно отсутствует.
