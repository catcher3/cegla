package tags

import "cegla/core"

// interfaces_check.go — компилятор проверяет соответствие каждого тега
// нужным интерфейсам прямо при go build. Раньше опечатка вида "isFlow"
// вместо "IsFlow" не давала ошибки компиляции в момент объявления метода
// (в Go interfaces implicit), и всплывала только когда кто-то пытался
// использовать тег в литерале контейнера. Этот файл превращает такие
// ошибки в немедленный build failure. При добавлении нового тега — сразу
// добавляйте сюда строку.
var (
	_ core.FlowContent = Div{}
	_ core.FlowContent = Blockquote{}
	_ core.FlowContent = Main{}
	_ core.FlowContent = Dialog{}
	_ core.FlowContent = Header{}
	_ core.FlowContent = Footer{}
	_ core.FlowContent = Address{}
	_ core.FlowContent = Form{}
	_ core.FlowContent = P{}
	_ core.FlowContent = Pre{}
	_ core.FlowContent = Hr{}
	_ core.FlowContent = Ol{}
	_ core.FlowContent = Ul{}
	_ core.FlowContent = LI{}
	_ core.FlowContent = Dl{}
	_ core.FlowContent = Figure{}
	_ core.FlowContent = FigCaption{}
	_ core.FlowContent = Details{}
	_ core.FlowContent = Summary{}
	_ core.FlowContent = Article{}
	_ core.FlowContent = Section{}
	_ core.FlowContent = Nav{}
	_ core.FlowContent = Aside{}
	_ core.FlowContent = H1{}
	_ core.FlowContent = H2{}
	_ core.FlowContent = H3{}
	_ core.FlowContent = H4{}
	_ core.FlowContent = H5{}
	_ core.FlowContent = H6{}
	_ core.FlowContent = Span{}
	_ core.FlowContent = Em{}
	_ core.FlowContent = Strong{}
	_ core.FlowContent = Small{}
	_ core.FlowContent = S{}
	_ core.FlowContent = Cite{}
	_ core.FlowContent = Q{}
	_ core.FlowContent = Dfn{}
	_ core.FlowContent = Abbr{}
	_ core.FlowContent = Time{}
	_ core.FlowContent = Code{}
	_ core.FlowContent = Var{}
	_ core.FlowContent = Samp{}
	_ core.FlowContent = Kbd{}
	_ core.FlowContent = Sub{}
	_ core.FlowContent = Sup{}
	_ core.FlowContent = I{}
	_ core.FlowContent = B{}
	_ core.FlowContent = U{}
	_ core.FlowContent = Mark{}
	_ core.FlowContent = Bdi{}
	_ core.FlowContent = Bdo{}
	_ core.FlowContent = Wbr{}
	_ core.FlowContent = Br{}
	_ core.FlowContent = Data{}
	_ core.FlowContent = A{}
	_ core.FlowContent = Img{}
	_ core.FlowContent = Embed{}
	_ core.FlowContent = Iframe{}
	_ core.FlowContent = Object{}
	_ core.FlowContent = Video{}
	_ core.FlowContent = Audio{}
	_ core.FlowContent = Canvas{}
	_ core.FlowContent = Svg{}
	_ core.FlowContent = MathML{}
	_ core.FlowContent = Button{}
	_ core.FlowContent = Label{}
	_ core.FlowContent = Select{}
	_ core.FlowContent = Textarea{}
	_ core.FlowContent = Input{}
	_ core.FlowContent = Fieldset{}
	_ core.FlowContent = Output{}
	_ core.FlowContent = Progress{}
	_ core.FlowContent = Meter{}
	_ core.FlowContent = NoScript{}
	_ core.FlowContent = Script{}

	_ core.PhrasingContent = Span{}
	_ core.PhrasingContent = A{}
	_ core.PhrasingContent = Button{}
	_ core.PhrasingContent = Img{}
	_ core.PhrasingContent = Data{}

	_ core.MetadataContent = Title{}
	_ core.MetadataContent = Script{}
	_ core.MetadataContent = StyleTag{}
	_ core.MetadataContent = Base{}
	_ core.MetadataContent = Link{}
	_ core.MetadataContent = Meta{}
	_ core.MetadataContent = NoScript{}

	_ core.SectioningContent = Article{}
	_ core.SectioningContent = Section{}
	_ core.SectioningContent = Nav{}
	_ core.SectioningContent = Aside{}

	_ core.HeadingContent = H1{}
	_ core.HeadingContent = H2{}
	_ core.HeadingContent = H3{}
	_ core.HeadingContent = H4{}
	_ core.HeadingContent = H5{}
	_ core.HeadingContent = H6{}

	_ core.EmbeddedContent = Img{}
	_ core.EmbeddedContent = Iframe{}
	_ core.EmbeddedContent = Embed{}
	_ core.EmbeddedContent = Object{}
	_ core.EmbeddedContent = Video{}
	_ core.EmbeddedContent = Audio{}
	_ core.EmbeddedContent = Canvas{}
	_ core.EmbeddedContent = Svg{}
	_ core.EmbeddedContent = MathML{}

	_ core.InteractiveContent = A{}
	_ core.InteractiveContent = Button{}
	_ core.InteractiveContent = Details{}
	_ core.InteractiveContent = Label{}
	_ core.InteractiveContent = Select{}
	_ core.InteractiveContent = Textarea{}
	_ core.InteractiveContent = Input{}

	_ core.FormAssociatedContent = Button{}
	_ core.FormAssociatedContent = Label{}
	_ core.FormAssociatedContent = Select{}
	_ core.FormAssociatedContent = Textarea{}
	_ core.FormAssociatedContent = Input{}
	_ core.FormAssociatedContent = Fieldset{}
	_ core.FormAssociatedContent = Output{}
	_ core.FormAssociatedContent = Progress{}

	_ core.TableChild = Caption{}
	_ core.TableChild = ColGroup{}
	_ core.TableChild = THead{}
	_ core.TableChild = TBody{}
	_ core.TableChild = TFoot{}
	_ core.TableChild = TR{}

	_ core.TableRowContent = TD{}
	_ core.TableRowContent = TH{}

	_ core.ListChild = LI{}

	_ core.DescriptionListContent = Dt{}
	_ core.DescriptionListContent = Dd{}

	_ core.SelectContent = Option{}
	_ core.SelectContent = Optgroup{}

	_ core.OptionContent = Option{}
)
