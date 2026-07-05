package ui

import (
	"bufio"
	"cegla/core"
	"cegla/core/atr"
	"cegla/core/atr/css/tw"
	"cegla/core/tags"
	"context"
)

// MenuItem — не core.Node, а чистые данные: то, из чего Render строит
// реальные теги. Специально не Node, чтобы пункты меню было удобно собирать
// из БД/конфига/JSON, не думая про интерфейсы рендера вообще.
type MenuItem struct {
	Label  string
	Href   string
	Active bool
}

// Orientation определяет, как меню располагается: сверху (горизонтальный
// navbar) или сбоку (вертикальный sidebar). Влияет и на flex-direction
// контейнера, и на его размеры/позиционирование.
type Orientation int

const (
	Horizontal Orientation = iota // верхнее меню (navbar)
	Vertical                      // боковое меню (sidebar)
)

// Menu — композиция (struct), а не типизированный контейнер. Реализует
// core.Node вручную и внутри Render собирает дерево целиком из tags + atr +
// tw, ничего не рендерит "напрямую".
type Menu struct {
	Brand       string
	Items       []MenuItem
	Orientation Orientation      // Horizontal (по умолчанию) или Vertical
	Attrs       []core.Attribute // доп. атрибуты на корневой <nav>
}

// Name() возвращает имя фактически рендерящегося корневого тега (<nav>),
// а не название компонента — иначе node.Name() разойдётся с тем, что
// реально пишется в Render.
func (Menu) Name() string { return "nav" }

// BuildContainer собирает дерево <nav> целиком, но НЕ рендерит его —
// возвращает tags.Nav как есть. Реализует core.Composition. Это точка
// расширения: обёртки над Menu (см. MenuWithActions ниже) вызывают
// BuildContainer(), дописывают своё в готовое дерево и уже сами вызывают
// Render — не дублируя сборку брэнда/списка/классов заново.
func (m Menu) BuildContainer() tags.Nav {
	nav := make(tags.Nav, 0, len(m.Attrs)+4)

	// core.Attribute не гарантирует core.FlowContent на уровне статических
	// интерфейсов (в Go совместимость интерфейсов не транзитивна), даже
	// если каждая реальная реализация её имеет через core.AttrMarker —
	// поэтому явный мост через any(...).(core.FlowContent).
	for _, a := range m.Attrs {
		if fc, ok := any(a).(core.FlowContent); ok {
			nav = append(nav, fc)
		}
	}

	list := make(tags.Ul, 0, len(m.Items)+3)

	switch m.Orientation {
	case Vertical:
		nav = append(nav,
			tw.Flex(), tw.FlexCol(),
			tw.Class("w-64 h-screen p-4 bg-white border-r border-gray-200"),
		)
		list = append(list, tw.Flex(), tw.FlexCol(), tw.Gap("1"))
	default: // Horizontal
		nav = append(nav,
			tw.Flex(), tw.ItemsCenter(), tw.JustifyBetween(),
			tw.Class("px-6 py-4 bg-white shadow-sm"),
		)
		list = append(list, tw.Flex(), tw.ItemsCenter(), tw.Gap("6"))
	}

	if m.Brand != "" {
		nav = append(nav, tags.Span{
			tw.FontBold(),
			tw.Text("xl"),
			core.Text(m.Brand),
		})
	}

	for _, item := range m.Items {
		list = append(list, tags.LI{menuLink(item)})
	}
	nav = append(nav, list)

	return nav
}

// Render — тонкая обёртка над core.RenderComposition: собрать контейнер,
// отрендерить его. Сама Menu никогда не пишет в *bufio.Writer напрямую.
func (m Menu) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderComposition(m, ctx, w)
}

// IsFlow — делает Menu вкладываемым напрямую в Body{}/Div{} наравне с
// обычными тегами: Body{ Menu{...}, MainApp(...) }.
func (Menu) IsFlow() {}

// menuLink собирает одну ссылку пункта меню. Активный пункт получает
// постоянную подсветку и НЕ получает hover-классов — наведение мышкой
// имеет смысл только для того, что ещё не выбрано. Обычный пункт наоборот:
// подсветка (фон + цвет текста) появляется только при hover — это и есть
// "выделение при наведении на мышку", через tw.Hover(...) (готовый хелпер
// из tw.go, который добавляет префикс "hover:" перед именем класса).
func menuLink(item MenuItem) tags.A {
	link := tags.A{
		atr.Href(item.Href),
		tw.Class("px-3 py-2 rounded transition-colors"),
	}

	if item.Active {
		link = append(link, tw.Class("bg-blue-50 text-blue-600 font-semibold"))
	} else {
		link = append(link,
			tw.TextColor("gray-600"),
			tw.Hover(tw.Bg("gray-100")),
			tw.Hover(tw.TextColor("gray-900")),
		)
	}

	link = append(link, core.Text(item.Label))
	return link
}

// MenuWithActions — обёртка (враппер) над Menu: добавляет произвольный
// блок в конец <nav> (кнопки, поиск, аватар пользователя), не трогая
// логику сборки бренда/списка/классов — она переиспользуется из Menu через
// buildNav(). Встраивание (embedding) Menu в MenuWithActions даёт "бесплатно"
// Name() и IsFlow() — их не нужно объявлять заново, метод-резолюшн Go сам
// поднимет их из встроенного поля. Явно переопределяется только Render.
type MenuWithActions struct {
	Menu
	Actions []core.FlowContent // например кнопка "Войти" или поле поиска
}

func (m MenuWithActions) BuildContainer() tags.Nav {
	// BuildContainer теперь дженерик по T=tags.Nav (Composition[tags.Nav]),
	// поэтому Menu.BuildContainer() уже возвращает конкретный tags.Nav —
	// в отличие от предыдущей негенерик-версии, тут не нужен ни type
	// assertion, ни проверка ok.
	nav := m.Menu.BuildContainer()
	nav = append(nav, m.Actions...)
	return nav
}

func (m MenuWithActions) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderComposition(m, ctx, w)
}
