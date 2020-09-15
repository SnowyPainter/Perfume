package perfume

//FormalElementType is type of some groupping Layouts
type FormalElementType uint8

//LayoutElementType is type of something that contains only IElement
type LayoutElementType uint8

//ElementType is type of elements which can't be parent
type ElementType uint8

//InputType Types of Input Element
type InputType uint8

//CommonOption can be apply all types
type CommonOption uint8

//ComponentOption Type of styles
type ComponentOption uint8

//LayoutOption options
type LayoutOption uint8

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
