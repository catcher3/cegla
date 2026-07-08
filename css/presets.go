package css

var (
	// Box — базовый контейнер для элементов
	Box = Style{
		Raw("p-4"),
		Raw("border"),
		Raw("rounded-lg"),
		Raw("bg-white"),
	}

	// FlexRow — контейнер для горизонтального выравнивания
	FlexRow = Style{
		Raw("flex"),
		Raw("flex-row"),
		Raw("items-center"),
		Raw("gap-4"),
	}

	// Shadow — готовая тень
	Shadow = Style{
		Raw("shadow-md"),
		Raw("hover:shadow-lg"),
		Raw("transition-shadow"),
	}
)
