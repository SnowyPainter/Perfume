package perfume

//FormalElementType is type of some groupping Layouts
type FormalElementType int

//LayoutElementType is type of something that contains only IElement
type LayoutElementType int

//ElementType is type of elements which can't be parent
type ElementType int

//InputType Types of Input Element
type InputType int

//ComponentOption Type of styles
type ComponentOption int

//LayoutOption options
type LayoutOption int

const (
	_ FormalElementType = iota
	HeadElementType
	BodyElementType
	FooterElementType
)

const (
	_ LayoutElementType = iota
	FreeLayoutType
	StackLayoutType
)

const (
	_ ElementType = iota
	InputElementType
	TextElementType
)

const (
	_ InputType = iota
	TextInputType
	ChooseInputType
	CheckInputType
	NumericInputType
)

const (
	_ ComponentOption = iota
	MarginOption
	PaddingOption
	WidthOption
	HeightOption
)
const (
	_ LayoutOption = iota
	SpacingOption
	OrientationOption
)
