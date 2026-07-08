package tag

// interfaces_check.go — компилятор проверяет соответствие каждого тега
// нужным интерфейсам прямо при go build. Раньше опечатка вида "isFlow"
// вместо "IsFlow" не давала ошибки компиляции в момент объявления метода
// (в Go interfaces implicit), и всплывала только когда кто-то пытался
// использовать тег в литерале контейнера. Этот файл превращает такие
// ошибки в немедленный build failure. При добавлении нового тега — сразу
// добавляйте сюда строку.
var (
	_ FlowContent = Div{}
	_ FlowContent = Blockquote{}
	_ FlowContent = Main{}
	_ FlowContent = Dialog{}
	_ FlowContent = Header{}
	_ FlowContent = Footer{}
	_ FlowContent = Address{}
	_ FlowContent = Form{}
	_ FlowContent = P{}
	_ FlowContent = Pre{}
	_ FlowContent = Hr{}
	_ FlowContent = Ol{}
	_ FlowContent = Ul{}
	_ FlowContent = LI{}
	_ FlowContent = Dl{}
	_ FlowContent = Figure{}
	_ FlowContent = FigCaption{}
	_ FlowContent = Details{}
	_ FlowContent = Summary{}
	_ FlowContent = Article{}
	_ FlowContent = Section{}
	_ FlowContent = Nav{}
	_ FlowContent = Aside{}
	_ FlowContent = H1{}
	_ FlowContent = H2{}
	_ FlowContent = H3{}
	_ FlowContent = H4{}
	_ FlowContent = H5{}
	_ FlowContent = H6{}
	_ FlowContent = Span{}
	_ FlowContent = Em{}
	_ FlowContent = Strong{}
	_ FlowContent = Small{}
	_ FlowContent = S{}
	_ FlowContent = Cite{}
	_ FlowContent = Q{}
	_ FlowContent = Dfn{}
	_ FlowContent = Abbr{}
	_ FlowContent = Time{}
	_ FlowContent = Code{}
	_ FlowContent = Var{}
	_ FlowContent = Samp{}
	_ FlowContent = Kbd{}
	_ FlowContent = Sub{}
	_ FlowContent = Sup{}
	_ FlowContent = I{}
	_ FlowContent = B{}
	_ FlowContent = U{}
	_ FlowContent = Mark{}
	_ FlowContent = Bdi{}
	_ FlowContent = Bdo{}
	_ FlowContent = Wbr{}
	_ FlowContent = Br{}
	_ FlowContent = Data{}
	_ FlowContent = A{}
	_ FlowContent = Img{}
	_ FlowContent = Embed{}
	_ FlowContent = Iframe{}
	_ FlowContent = Object{}
	_ FlowContent = Video{}
	_ FlowContent = Audio{}
	_ FlowContent = Canvas{}
	_ FlowContent = Svg{}
	_ FlowContent = MathML{}
	_ FlowContent = Button{}
	_ FlowContent = Label{}
	_ FlowContent = Select{}
	_ FlowContent = Textarea{}
	_ FlowContent = Input{}
	_ FlowContent = Fieldset{}
	_ FlowContent = Output{}
	_ FlowContent = Progress{}
	_ FlowContent = Meter{}
	_ FlowContent = NoScript{}
	_ FlowContent = Script{}

	_ PhrasingContent = Span{}
	_ PhrasingContent = A{}
	_ PhrasingContent = Button{}
	_ PhrasingContent = Img{}
	_ PhrasingContent = Data{}

	_ MetadataContent = Title{}
	_ MetadataContent = Script{}
	_ MetadataContent = StyleTag{}
	_ MetadataContent = Base{}
	_ MetadataContent = Link{}
	_ MetadataContent = Meta{}
	_ MetadataContent = NoScript{}

	_ SectioningContent = Article{}
	_ SectioningContent = Section{}
	_ SectioningContent = Nav{}
	_ SectioningContent = Aside{}

	_ HeadingContent = H1{}
	_ HeadingContent = H2{}
	_ HeadingContent = H3{}
	_ HeadingContent = H4{}
	_ HeadingContent = H5{}
	_ HeadingContent = H6{}

	_ EmbeddedContent = Img{}
	_ EmbeddedContent = Iframe{}
	_ EmbeddedContent = Embed{}
	_ EmbeddedContent = Object{}
	_ EmbeddedContent = Video{}
	_ EmbeddedContent = Audio{}
	_ EmbeddedContent = Canvas{}
	_ EmbeddedContent = Svg{}
	_ EmbeddedContent = MathML{}

	_ InteractiveContent = A{}
	_ InteractiveContent = Button{}
	_ InteractiveContent = Details{}
	_ InteractiveContent = Label{}
	_ InteractiveContent = Select{}
	_ InteractiveContent = Textarea{}
	_ InteractiveContent = Input{}

	_ FormAssociatedContent = Button{}
	_ FormAssociatedContent = Label{}
	_ FormAssociatedContent = Select{}
	_ FormAssociatedContent = Textarea{}
	_ FormAssociatedContent = Input{}
	_ FormAssociatedContent = Fieldset{}
	_ FormAssociatedContent = Output{}
	_ FormAssociatedContent = Progress{}

	_ TableChild = Caption{}
	_ TableChild = ColGroup{}
	_ TableChild = THead{}
	_ TableChild = TBody{}
	_ TableChild = TFoot{}
	_ TableChild = TR{}

	_ TableRowContent = TD{}
	_ TableRowContent = TH{}

	_ ListChild = LI{}

	_ DescriptionListContent = Dt{}
	_ DescriptionListContent = Dd{}

	_ SelectContent = Option{}
	_ SelectContent = Optgroup{}

	_ OptionContent = Option{}
)
