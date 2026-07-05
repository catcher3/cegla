// composition.go
package core

import (
	"bufio"
	"context"
)

// Composition — контракт UI-уровня (ui.Menu и подобные), в отличие от
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
// в отличие от core.Attribute -> core.FlowContent, здесь нетранзитивность
// интерфейсов Go нам не мешает, т.к. T уже конкретен на этапе объявления.
//
// Плата за это: Composition[T] с разными T нельзя хранить в одном общем
// слайсе (`[]Composition[?]` в Go не существует) — если понадобится
// гетерогенная коллекция композиций, храните их как обычный core.Node,
// просто без доступа к BuildContainer у элементов такой коллекции.
type Composition[T Node] interface {
	Node
	BuildContainer() T
}

// RenderComposition — единая реализация Render для любой Composition[T]:
// построить контейнер, отрендерить его. Тип T выводится компилятором из
// конкретной реализации BuildContainer вызывающего типа — вызывать можно
// без явного указания параметра: core.RenderComposition(m, ctx, w).
func RenderComposition[T Node](c Composition[T], ctx context.Context, w *bufio.Writer) error {
	return c.BuildContainer().Render(ctx, w)
}
