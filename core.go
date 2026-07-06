package cegla

import (
	"bufio"
	"context"
)

// --- Фундамент ---

type Node interface {
	Name() string //Возвращает имя тега, например "div", "span", "h1" и т.д.

	Render(ctx context.Context, w *bufio.Writer) error
}

type Attribute interface {
	Node
	IsAttribute()

	Key() string
	Value() string
}

type DocumentContent interface {
	Node

	// Нет смысла делать через экспортируемый метод, так как это маркер для двух элементов Head и Body.
	IsDocumentContent()
}

// --- Маркеры (Интерфейсы) ---
// --- Базовые категории контента (пересекающиеся, элемент может иметь несколько) ---

type MetadataContent interface {
	Node
	IsMetadata()
}

type FlowContent interface {
	Node
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
	Node
	IsInteractive()
}

type FormAssociatedContent interface {
	Node
	IsFormAssociated()
}

// --- Категории со структурным исключением (единственное законное место для "особых" маркеров) ---

// Валидные прямые дети <table>: caption, colgroup, thead, tbody, tfoot, tr
type TableChild interface {
	Node
	IsTableChild()
}

// Валидные прямые дети <tr>: td, th
type TableRowContent interface {
	Node
	IsTableRowContent()
}

// Валидные прямые дети <select>: option, optgroup
type SelectContent interface {
	Node
	IsSelectContent()
}

// Валидные прямые дети <dl>: dt, dd
type DescriptionListContent interface {
	Node
	IsDescriptionListContent()
}

// Валидные прямые дети <ol>/<ul>: только li
type ListChild interface {
	Node
	IsListChild()
}

// Валидные дети <select>: option напрямую ИЛИ optgroup
// Валидные дети <optgroup>: только option
type OptionContent interface {
	Node
	IsOptionContent()
}

// --- Универсальный маркер для атрибутов ---
//
// Attribute — единственный тип, которому разрешено находиться сразу во ВСЕХ
// категориях контента (это осознанное исключение из принципа "никакого any",
// иначе на каждый новый пакет атрибутов пришлось бы вручную реализовывать
// десяток IsX() методов под каждую категорию, где ожидается атрибут).
//
// Любой тип атрибута (atr.Attr, htmx.Attr, css.Style, ...) должен встроить
// AttrMarker, чтобы автоматически удовлетворять DocumentContent, FlowContent,
// PhrasingContent, MetadataContent и всем "структурным исключениям"
// (TableChild, TableRowContent, SelectContent, DescriptionListContent,
// ListChild, OptionContent), а также Sectioning/Heading/Embedded/
// Interactive/FormAssociated. Сам по себе AttrMarker не реализует Node —
// Name()/Render()/Key()/Value()/IsAttribute() всё равно нужно писать в
// конкретном типе атрибута.
type AttrMarker struct{}

func (AttrMarker) IsDocumentContent()        {}
func (AttrMarker) IsFlow()                   {}
func (AttrMarker) IsPhrasing()               {}
func (AttrMarker) IsMetadata()               {}
func (AttrMarker) IsSectioning()             {}
func (AttrMarker) IsHeading()                {}
func (AttrMarker) IsEmbedded()               {}
func (AttrMarker) IsInteractive()            {}
func (AttrMarker) IsFormAssociated()         {}
func (AttrMarker) IsTableChild()             {}
func (AttrMarker) IsTableRowContent()        {}
func (AttrMarker) IsSelectContent()          {}
func (AttrMarker) IsDescriptionListContent() {}
func (AttrMarker) IsListChild()              {}
func (AttrMarker) IsOptionContent()          {}
