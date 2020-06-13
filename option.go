package perfume

//RelLocation is location structure which is relative of parent
type RelLocation struct {
	X int
	Y int
}

//Size is a structure which has Height and Width
type Size struct {
	Height int
	Width  int
}

//**********************
// Constructors
//**********************

//NewRelativeLocation makes new RelLocation object
func NewRelativeLocation(x int, y int) RelLocation {
	return RelLocation{
		X: x, Y: y,
	}
}

//NewSize makes new Size object by height and width
func NewSize(height int, width int) Size {
	return Size{
		Height: height, Width: width,
	}
}
