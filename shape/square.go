package shape

type Square struct {
	k float32
}

func (s Square) GetLength() float32    { return s.k }
func (s Square) GetArea() float32      { return s.k * s.k }
func (s Square) GetPerimeter() float32 { return s.k * 4 }

func (s *Square) SetLength(l float32) { s.k = l }
