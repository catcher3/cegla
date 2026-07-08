package tag

import (
	"bufio"
	"bytes"
	"context"
	"html"

	. "github.com/catcher3/cegla"
	"github.com/catcher3/cegla/render"
)

type DocumentContent interface {
	Container

	// Нет смысла делать через экспортируемый метод, так как это маркер для двух элементов Head и Body.
	IsDocumentContent()
}

// --- Маркеры (Интерфейсы) ---
// --- Базовые категории контента (пересекающиеся, элемент может иметь несколько) ---

type MetadataContent interface {
	Container

	IsMetadata()
}

type FlowContent interface {
	Container

	IsFlow()
}

type SectioningContent interface {
	FlowContent
	IsSectioning()
}

type HeadingContent interface {
	FlowContent
	IsHeading()
}

type PhrasingContent interface {
	FlowContent
	IsPhrasing()
}

type EmbeddedContent interface {
	PhrasingContent

	IsEmbedded()
}

// Interactive и FormAssociated — ортогональные категории,
// элемент типа A или Button получает их ДОПОЛНИТЕЛЬНО к Flow/Phrasing,
// а не вместо. Поэтому они НЕ встраивают Flow/Phrasing в себя —
// каждый элемент декларирует свои роли явно.
type InteractiveContent interface {
	Container
	IsInteractive()
}

type FormAssociatedContent interface {
	Container

	IsFormAssociated()
}

// --- Категории со структурным исключением (единственное законное место для "особых" маркеров) ---

// Валидные прямые дети <table>: caption, colgroup, thead, tbody, tfoot, tr
type TableChild interface {
	Container

	IsTableChild()
}

// Валидные прямые дети <tr>: td, th
type TableRowContent interface {
	Container

	IsTableRowContent()
}

// Валидные прямые дети <select>: option, optgroup
type SelectContent interface {
	Container

	IsSelectContent()
}

// Валидные прямые дети <dl>: dt, dd
type DescriptionListContent interface {
	Container

	IsDescriptionListContent()
}

// Валидные прямые дети <ol>/<ul>: только li
type ListChild interface {
	Container

	IsListChild()
}

// Валидные дети <select>: option напрямую ИЛИ optgroup
// Валидные дети <optgroup>: только option
type OptionContent interface {
	Container

	IsOptionContent()
}

// Контейнера

type Flow []FlowContent
type Phrasing []PhrasingContent
type List []ListChild
type DescriptionList []DescriptionListContent
type Option_ []OptionContent

type SelectChildren []SelectContent

// --- Строгие контейнеры ---
type HTML []DocumentContent

func (h HTML) String() string {
	ctx := context.Background()
	buf := new(bytes.Buffer)

	// Создаем врайтер, который пишет в наш буфер
	w := bufio.NewWriter(buf)

	// Рендерим
	h.Render(ctx, w)

	// ВАЖНО: сбрасываем буфер, чтобы все данные гарантированно попали в buf
	w.Flush()

	return buf.String()
}

func (h HTML) Name() string {
	return "html"
}

func (h HTML) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren("html", h, ctx, w)
}

type Head []MetadataContent // Строго Metadata

func (h Head) Name() string {
	return "head"
}

func (h Head) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren("head", h, ctx, w)
}
func (h Head) IsDocumentContent() {}

type Body []FlowContent // Теперь это слайс, а не тип Body

func (b Body) Name() string {
	return "body"
}

func (b Body) IsDocumentContent() {}

func (b Body) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren("body", b, ctx, w)
}

// --- Реализация Text ---
type Text string

func (t Text) Name() string {
	return "#text"
}

func (t Text) Render(ctx context.Context, w *bufio.Writer) error {
	w.WriteString(html.EscapeString(string(t)))

	return nil
}

// Text — это PhrasingContent
func (t Text) IsPhrasing() {}
func (t Text) IsFlow()     {}
