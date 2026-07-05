package css

type Atom string

func (a Atom) ToClass() string { return string(a) }

// Padding создает класс отступа: p-4, m-2 и т.д.
func Padding(size string) Atom { return Atom("p-" + size) }

// Bg создает класс фона: bg-blue-500
func Bg(color string) Atom { return Atom("bg-" + color) }

// Text создает класс текста: text-sm, text-bold
func Text(style string) Atom { return Atom("text-" + style) }

// Border создает класс рамки
func Border(size string) Atom { return Atom("border-" + size) }
