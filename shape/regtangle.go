package shape

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return 2 * (r.width + r.height)
}
