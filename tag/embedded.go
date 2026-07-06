package tag

import (
	"bufio"
	"context"

	"github.com/catcher3/cegla"
)

// --- Void embedded-элементы (только атрибуты, без детей) ---

type Img []cegla.Attribute
type Embed []cegla.Attribute
type Source []cegla.Attribute // только внутри picture/video/audio
type Track []cegla.Attribute  // только внутри video/audio

func (Img) Name() string    { return "img" }
func (Embed) Name() string  { return "embed" }
func (Source) Name() string { return "source" }
func (Track) Name() string  { return "track" }

func (el Img) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderVoid(el.Name(), el, ctx, w)
}
func (el Embed) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderVoid(el.Name(), el, ctx, w)
}
func (el Source) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderVoid(el.Name(), el, ctx, w)
}
func (el Track) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderVoid(el.Name(), el, ctx, w)
}

func (Img) IsFlow()       {}
func (Img) IsPhrasing()   {}
func (Img) IsEmbedded()   {}
func (Embed) IsFlow()     {}
func (Embed) IsPhrasing() {}
func (Embed) IsEmbedded() {}

// --- Iframe: не void по спеке, но fallback-контент почти не используется —
// упрощаем до атрибутов, как у большинства современных DSL.

type Iframe []cegla.Attribute

func (Iframe) Name() string { return "iframe" }
func (el Iframe) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderVoid(el.Name(), el, ctx, w)
}
func (Iframe) IsFlow()     {}
func (Iframe) IsPhrasing() {}
func (Iframe) IsEmbedded() {}

// --- Transparent content model (Object, Video, Audio, Canvas) ---
//
// По спеке их реальные допустимые дети зависят от контекста использования
// (Phrasing внутри Phrasing-родителя, Flow внутри Flow-родителя). Система
// маркерных интерфейсов такое не выражает без дублирования типов на каждый
// контекст, поэтому здесь сознательно взято самое широкое разрешение
// ([]cegla.FlowContent) — см. обсуждение архитектуры. Это единственное
// известное ограничение системы типов, не баг.

type Object []cegla.FlowContent
type Video []cegla.FlowContent // + Source/Track внутри
type Audio []cegla.FlowContent // + Source/Track внутри
type Canvas []cegla.FlowContent

func (Object) Name() string { return "object" }
func (Video) Name() string  { return "video" }
func (Audio) Name() string  { return "audio" }
func (Canvas) Name() string { return "canvas" }

func (el Object) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Video) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Audio) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}
func (el Canvas) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderChildren(el.Name(), el, ctx, w)
}

func (Object) IsFlow()     {}
func (Object) IsPhrasing() {}
func (Object) IsEmbedded() {}
func (Video) IsFlow()      {}
func (Video) IsPhrasing()  {}
func (Video) IsEmbedded()  {}
func (Audio) IsFlow()      {}
func (Audio) IsPhrasing()  {}
func (Audio) IsEmbedded()  {}
func (Canvas) IsFlow()     {}
func (Canvas) IsPhrasing() {}
func (Canvas) IsEmbedded() {}

// --- Svg / MathML — заглушки: свой мир атрибутов/контента, не переиспользуют
// HTML-категории. Полноценная поддержка — отдельная задача на будущее.

type Svg []cegla.Attribute
type MathML []cegla.Attribute

func (Svg) Name() string    { return "svg" }
func (MathML) Name() string { return "math" }

func (el Svg) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderVoid(el.Name(), el, ctx, w)
}
func (el MathML) Render(ctx context.Context, w *bufio.Writer) error {
	return cegla.RenderVoid(el.Name(), el, ctx, w)
}

func (Svg) IsFlow()        {}
func (Svg) IsPhrasing()    {}
func (Svg) IsEmbedded()    {}
func (MathML) IsFlow()     {}
func (MathML) IsPhrasing() {}
func (MathML) IsEmbedded() {}
