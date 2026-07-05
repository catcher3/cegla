package tags

import (
	"bufio"
	"cegla/core"
	"context"
)

type Span []core.PhrasingContent
type Em []core.PhrasingContent
type Strong []core.PhrasingContent
type Small []core.PhrasingContent
type S []core.PhrasingContent
type Cite []core.PhrasingContent
type Q []core.PhrasingContent
type Dfn []core.PhrasingContent
type Abbr []core.PhrasingContent
type Time []core.PhrasingContent
type Code []core.PhrasingContent
type Var []core.PhrasingContent
type Samp []core.PhrasingContent
type Kbd []core.PhrasingContent
type Sub []core.PhrasingContent
type Sup []core.PhrasingContent
type I []core.PhrasingContent
type B []core.PhrasingContent
type U []core.PhrasingContent
type Mark []core.PhrasingContent
type Bdi []core.PhrasingContent
type Bdo []core.PhrasingContent
type Wbr []core.Attribute // void

// A — транзитивный (transparent) элемент: реальные допустимые дети зависят
// от контекста использования (см. ограничение системы типов, embedded.go).
// Берём самое широкое разрешение — PhrasingContent.
type A []core.PhrasingContent

func (Span) Name() string   { return "span" }
func (Em) Name() string     { return "em" }
func (Strong) Name() string { return "strong" }
func (Small) Name() string  { return "small" }
func (S) Name() string      { return "s" }
func (Cite) Name() string   { return "cite" }
func (Q) Name() string      { return "q" }
func (Dfn) Name() string    { return "dfn" }
func (Abbr) Name() string   { return "abbr" }
func (Time) Name() string   { return "time" }
func (Code) Name() string   { return "code" }
func (Var) Name() string    { return "var" }
func (Samp) Name() string   { return "samp" }
func (Kbd) Name() string    { return "kbd" }
func (Sub) Name() string    { return "sub" }
func (Sup) Name() string    { return "sup" }
func (I) Name() string      { return "i" }
func (B) Name() string      { return "b" }
func (U) Name() string      { return "u" }
func (Mark) Name() string   { return "mark" }
func (Bdi) Name() string    { return "bdi" }
func (Bdo) Name() string    { return "bdo" }
func (Wbr) Name() string    { return "wbr" }
func (A) Name() string      { return "a" }

func (el Span) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Em) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Strong) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Small) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el S) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Cite) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Q) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Dfn) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Abbr) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Time) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Code) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Var) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Samp) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Kbd) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Sub) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Sup) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el I) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el B) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el U) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Mark) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Bdi) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Bdo) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (el Wbr) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderVoid(el.Name(), el, ctx, w)
}
func (el A) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}

func (Span) IsFlow()       {}
func (Span) IsPhrasing()   {}
func (Em) IsFlow()         {}
func (Em) IsPhrasing()     {}
func (Strong) IsFlow()     {}
func (Strong) IsPhrasing() {}
func (Small) IsFlow()      {}
func (Small) IsPhrasing()  {}
func (S) IsFlow()          {}
func (S) IsPhrasing()      {}
func (Cite) IsFlow()       {}
func (Cite) IsPhrasing()   {}
func (Q) IsFlow()          {}
func (Q) IsPhrasing()      {}
func (Dfn) IsFlow()        {}
func (Dfn) IsPhrasing()    {}
func (Abbr) IsFlow()       {}
func (Abbr) IsPhrasing()   {}
func (Time) IsFlow()       {}
func (Time) IsPhrasing()   {}
func (Code) IsFlow()       {}
func (Code) IsPhrasing()   {}
func (Var) IsFlow()        {}
func (Var) IsPhrasing()    {}
func (Samp) IsFlow()       {}
func (Samp) IsPhrasing()   {}
func (Kbd) IsFlow()        {}
func (Kbd) IsPhrasing()    {}
func (Sub) IsFlow()        {}
func (Sub) IsPhrasing()    {}
func (Sup) IsFlow()        {}
func (Sup) IsPhrasing()    {}
func (I) IsFlow()          {}
func (I) IsPhrasing()      {}
func (B) IsFlow()          {}
func (B) IsPhrasing()      {}
func (U) IsFlow()          {}
func (U) IsPhrasing()      {}
func (Mark) IsFlow()       {}
func (Mark) IsPhrasing()   {}
func (Bdi) IsFlow()        {}
func (Bdi) IsPhrasing()    {}
func (Bdo) IsFlow()        {}
func (Bdo) IsPhrasing()    {}
func (Wbr) IsFlow()        {}
func (Wbr) IsPhrasing()    {}

func (A) IsFlow()        {}
func (A) IsPhrasing()    {}
func (A) IsInteractive() {}

// --- Br — void ---

type Br []core.Attribute

func (Br) Name() string { return "br" }
func (el Br) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderVoid(el.Name(), el, ctx, w)
}
func (Br) IsFlow()     {}
func (Br) IsPhrasing() {}

// --- Data ---

type Data []core.PhrasingContent // атрибут value + текстовое содержимое

func (Data) Name() string { return "data" }
func (el Data) Render(ctx context.Context, w *bufio.Writer) error {
	return core.RenderChildren(el.Name(), el, ctx, w)
}
func (Data) IsFlow()     {}
func (Data) IsPhrasing() {}
