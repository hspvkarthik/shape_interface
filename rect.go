package shape_interface

type Rectangle struct {
	length float64
	width  float64
}

func NewRectangle(length float64, width float64) Rectangle {
	if length < 0 || width < 0 {
		panic("length or width should be positive")
	}
	return Rectangle{length: length, width: width}
}
