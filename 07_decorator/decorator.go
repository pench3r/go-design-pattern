package decorator

type IDraw interface {
	Draw() string
}

type Square struct{}

func (s Square) Draw() string {
	return "this is square"
}

type ColorSquare struct {
	color  string
	square IDraw
}

func NewColorSquare(col string, draw IDraw) *ColorSquare {
	return &ColorSquare{color: col, square: draw}
}

func (c ColorSquare) Draw() string {
	return "ColorSquare: " + c.square.Draw() + " with " + c.color
}
