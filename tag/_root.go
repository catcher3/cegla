package tag

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
	return RenderChildren("html", h, ctx, w)
}

type Head []MetadataContent // Строго Metadata

func (h Head) Name() string {
	return "head"
}

func (h Head) Render(ctx context.Context, w *bufio.Writer) error {
	return RenderChildren("head", h, ctx, w)
}
func (h Head) IsDocumentContent() {}

type Body []FlowContent // Теперь это слайс, а не тип Body

func (b Body) Name() string {
	return "body"
}

func (b Body) IsDocumentContent() {}

func (b Body) Render(ctx context.Context, w *bufio.Writer) error {
	return RenderChildren("body", b, ctx, w)
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
