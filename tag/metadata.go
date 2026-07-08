package tag

import (
	"bufio"
	"context"

	. "github.com/catcher3/cegla"
	"github.com/catcher3/cegla/render"
)

// Title/Script/StyleTag — единственное сознательное исключение из правила
// "теги — это слайсы": они, как и cegla.Text, могут содержать ТОЛЬКО текст
// (по спеке <title>/<script>/<style> не принимают HTML-элементы вообще), и
// именно так их использует ваш целевой пример: Title("Hello World"). Если
// нужны атрибуты (например src/type/async у <script>), это отдельный тег
// ScriptSrc/аналог — не смешиваем текстовое содержимое и атрибуты в одном
// string-типе.
type Title Phrasing
type Script Phrasing
type StyleTag Phrasing

func (Title) Name() string    { return "title" }
func (Script) Name() string   { return "script" }
func (StyleTag) Name() string { return "style" }

func (el Title) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (el Script) Render(ctx context.Context, w *bufio.Writer) error {
	// Используем RenderChildren (а не RenderVoid), потому что <script> обязан иметь закрывающий тег </script>
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (el StyleTag) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}

func (Title) IsMetadata()    {}
func (Script) IsMetadata()   {}
func (Script) IsFlow()       {}
func (Script) IsPhrasing()   {} // <script> одновременно Metadata, Flow и Phrasing
func (StyleTag) IsMetadata() {}

// --- Void metadata-элементы (только атрибуты) ---

type Base Attrs
type Link Attrs
type Meta Attrs

func (Base) Name() string { return "base" }
func (Link) Name() string { return "link" }
func (Meta) Name() string { return "meta" }

func (el Base) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderVoid(el.Name(), el, ctx, w)
}
func (el Link) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderVoid(el.Name(), el, ctx, w)
}
func (el Meta) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderVoid(el.Name(), el, ctx, w)
}

func (Base) IsMetadata() {}
func (Link) IsMetadata() {}
func (Meta) IsMetadata() {}

// NoScript — двойного назначения (Metadata + Flow), содержимое — обычный Flow.
type NoScript Flow

func (NoScript) Name() string { return "noscript" }
func (el NoScript) Render(ctx context.Context, w *bufio.Writer) error {
	return render.RenderChildren(el.Name(), el, ctx, w)
}
func (NoScript) IsMetadata() {}
func (NoScript) IsFlow()     {}
