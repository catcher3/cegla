package cegla

import (
	"bufio"
	"context"
)

// --- Фундамент ---
type Container interface {
	Name() string //Возвращает имя тега, например "div", "span", "h1" и т.д.

	Render(ctx context.Context, w *bufio.Writer) error
}

// Composition — контракт comp-уровня (компонентов) (comp.Menu и подобные), в отличие от
// обычного тега из tags не пишет байты сама. Сначала она строит дерево
// целиком из tags+atr (BuildContainer), и только это готовое дерево
// рендерится — сама Composition никогда напрямую не работает с
// *bufio.Writer. Это формализует паттерн Menu.buildNav()/Render(), чтобы
// он был одинаковым для любого будущего компонента, а не придумывался
// каждый раз заново.
//
// Composition дженерик по T — конкретному типу контейнера, который строит
// данная композиция (например tags.Nav для Menu, tags.TR для будущей
// композиции строки таблицы). Это сознательный выбор в пользу негенерик-
// варианта (BuildContainer() Node): T конкретный, поэтому враппер поверх
// композиции (см. ui.MenuWithActions) может сразу append'ить в результат
// BuildContainer(), не делая type assertion обратно к конкретному слайсу —
// в отличие от Attrsibute -> cegla.FlowContent, здесь нетранзитивность
// интерфейсов Go нам не мешает, т.к. T уже конкретен на этапе объявления.
//
// Плата за это: Composition[T] с разными T нельзя хранить в одном общем
// слайсе (`[]Composition[?]` в Go не существует) — если понадобится
// гетерогенная коллекция композиций, храните их как обычный root.Node,
// просто без доступа к BuildContainer у элементов такой коллекции.
type Composition[T Container] interface {
	Container
	BuildContainer() T
}

type Attribute interface {
	Container

	IsAttribute()

	Key() string
	Value() string
}

// Список аттрибутов
type Attrs []Attribute

// Элемент является синглтоном
// Сиглтон - Уникальная сущность в рамка одного контейнера
type Singleton interface {
	Container

	// Путь/критерий по которому отслеживать уникальность сущности
	SingletonPath() string
}

// Определяет сущностей которые предпологается объединять в контейнере
// Важно на одном уровне.
type Aggregator interface {
}
