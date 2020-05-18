package shape

type Shape interface {
	GetLength() float32
	SetLength(float32)

	GetPerimeter() float32
	GetArea() float32
}

func CreateShapeFromNumber(num int) Shape {
	if num == 0 {
		return new(Square)
	} else {
		return new(Circle)
	}
}
