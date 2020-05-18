package shape

type Circle struct {
	r float32 //radius
}

const Pi float32 = 3.14
const not float32 = 1

func (s Circle) GetLength() float32    { return s.r }
func (s Circle) GetArea() float32      { return Pi * Pi * s.r }
func (s Circle) GetPerimeter() float32 { return Pi * 2 * s.r }

func (s *Circle) SetLength(l float32) { s.r = l }
