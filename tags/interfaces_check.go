package tags

import "github.com/catcher3/cegla"

// interfaces_check.go — компилятор проверяет соответствие каждого тега
// нужным интерфейсам прямо при go build. Раньше опечатка вида "isFlow"
// вместо "IsFlow" не давала ошибки компиляции в момент объявления метода
// (в Go interfaces implicit), и всплывала только когда кто-то пытался
// использовать тег в литерале контейнера. Этот файл превращает такие
// ошибки в немедленный build failure. При добавлении нового тега — сразу
// добавляйте сюда строку.
var (
	_ cegla.FlowContent = Div{}
	_ cegla.FlowContent = Blockquote{}
	_ cegla.FlowContent = Main{}
	_ cegla.FlowContent = Dialog{}
	_ cegla.FlowContent = Header{}
	_ cegla.FlowContent = Footer{}
	_ cegla.FlowContent = Address{}
	_ cegla.FlowContent = Form{}
	_ cegla.FlowContent = P{}
	_ cegla.FlowContent = Pre{}
	_ cegla.FlowContent = Hr{}
	_ cegla.FlowContent = Ol{}
	_ cegla.FlowContent = Ul{}
	_ cegla.FlowContent = LI{}
	_ cegla.FlowContent = Dl{}
	_ cegla.FlowContent = Figure{}
	_ cegla.FlowContent = FigCaption{}
	_ cegla.FlowContent = Details{}
	_ cegla.FlowContent = Summary{}
	_ cegla.FlowContent = Article{}
	_ cegla.FlowContent = Section{}
	_ cegla.FlowContent = Nav{}
	_ cegla.FlowContent = Aside{}
	_ cegla.FlowContent = H1{}
	_ cegla.FlowContent = H2{}
	_ cegla.FlowContent = H3{}
	_ cegla.FlowContent = H4{}
	_ cegla.FlowContent = H5{}
	_ cegla.FlowContent = H6{}
	_ cegla.FlowContent = Span{}
	_ cegla.FlowContent = Em{}
	_ cegla.FlowContent = Strong{}
	_ cegla.FlowContent = Small{}
	_ cegla.FlowContent = S{}
	_ cegla.FlowContent = Cite{}
	_ cegla.FlowContent = Q{}
	_ cegla.FlowContent = Dfn{}
	_ cegla.FlowContent = Abbr{}
	_ cegla.FlowContent = Time{}
	_ cegla.FlowContent = Code{}
	_ cegla.FlowContent = Var{}
	_ cegla.FlowContent = Samp{}
	_ cegla.FlowContent = Kbd{}
	_ cegla.FlowContent = Sub{}
	_ cegla.FlowContent = Sup{}
	_ cegla.FlowContent = I{}
	_ cegla.FlowContent = B{}
	_ cegla.FlowContent = U{}
	_ cegla.FlowContent = Mark{}
	_ cegla.FlowContent = Bdi{}
	_ cegla.FlowContent = Bdo{}
	_ cegla.FlowContent = Wbr{}
	_ cegla.FlowContent = Br{}
	_ cegla.FlowContent = Data{}
	_ cegla.FlowContent = A{}
	_ cegla.FlowContent = Img{}
	_ cegla.FlowContent = Embed{}
	_ cegla.FlowContent = Iframe{}
	_ cegla.FlowContent = Object{}
	_ cegla.FlowContent = Video{}
	_ cegla.FlowContent = Audio{}
	_ cegla.FlowContent = Canvas{}
	_ cegla.FlowContent = Svg{}
	_ cegla.FlowContent = MathML{}
	_ cegla.FlowContent = Button{}
	_ cegla.FlowContent = Label{}
	_ cegla.FlowContent = Select{}
	_ cegla.FlowContent = Textarea{}
	_ cegla.FlowContent = Input{}
	_ cegla.FlowContent = Fieldset{}
	_ cegla.FlowContent = Output{}
	_ cegla.FlowContent = Progress{}
	_ cegla.FlowContent = Meter{}
	_ cegla.FlowContent = NoScript{}
	_ cegla.FlowContent = Script{}

	_ cegla.PhrasingContent = Span{}
	_ cegla.PhrasingContent = A{}
	_ cegla.PhrasingContent = Button{}
	_ cegla.PhrasingContent = Img{}
	_ cegla.PhrasingContent = Data{}

	_ cegla.MetadataContent = Title{}
	_ cegla.MetadataContent = Script{}
	_ cegla.MetadataContent = StyleTag{}
	_ cegla.MetadataContent = Base{}
	_ cegla.MetadataContent = Link{}
	_ cegla.MetadataContent = Meta{}
	_ cegla.MetadataContent = NoScript{}

	_ cegla.SectioningContent = Article{}
	_ cegla.SectioningContent = Section{}
	_ cegla.SectioningContent = Nav{}
	_ cegla.SectioningContent = Aside{}

	_ cegla.HeadingContent = H1{}
	_ cegla.HeadingContent = H2{}
	_ cegla.HeadingContent = H3{}
	_ cegla.HeadingContent = H4{}
	_ cegla.HeadingContent = H5{}
	_ cegla.HeadingContent = H6{}

	_ cegla.EmbeddedContent = Img{}
	_ cegla.EmbeddedContent = Iframe{}
	_ cegla.EmbeddedContent = Embed{}
	_ cegla.EmbeddedContent = Object{}
	_ cegla.EmbeddedContent = Video{}
	_ cegla.EmbeddedContent = Audio{}
	_ cegla.EmbeddedContent = Canvas{}
	_ cegla.EmbeddedContent = Svg{}
	_ cegla.EmbeddedContent = MathML{}

	_ cegla.InteractiveContent = A{}
	_ cegla.InteractiveContent = Button{}
	_ cegla.InteractiveContent = Details{}
	_ cegla.InteractiveContent = Label{}
	_ cegla.InteractiveContent = Select{}
	_ cegla.InteractiveContent = Textarea{}
	_ cegla.InteractiveContent = Input{}

	_ cegla.FormAssociatedContent = Button{}
	_ cegla.FormAssociatedContent = Label{}
	_ cegla.FormAssociatedContent = Select{}
	_ cegla.FormAssociatedContent = Textarea{}
	_ cegla.FormAssociatedContent = Input{}
	_ cegla.FormAssociatedContent = Fieldset{}
	_ cegla.FormAssociatedContent = Output{}
	_ cegla.FormAssociatedContent = Progress{}

	_ cegla.TableChild = Caption{}
	_ cegla.TableChild = ColGroup{}
	_ cegla.TableChild = THead{}
	_ cegla.TableChild = TBody{}
	_ cegla.TableChild = TFoot{}
	_ cegla.TableChild = TR{}

	_ cegla.TableRowContent = TD{}
	_ cegla.TableRowContent = TH{}

	_ cegla.ListChild = LI{}

	_ cegla.DescriptionListContent = Dt{}
	_ cegla.DescriptionListContent = Dd{}

	_ cegla.SelectContent = Option{}
	_ cegla.SelectContent = Optgroup{}

	_ cegla.OptionContent = Option{}
)
