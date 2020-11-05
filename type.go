package perfume

//FormalElementType is type of some groupping Layouts
type FormalElementType uint8

//LayoutElementType is type of something that contains only IElement
type LayoutElementType uint8

//ElementType is type of elements which can't be parent
type ElementType uint8

//OrientationType for stack layout
type OrientationType uint8

//CommonOption can be apply all types
type CommonOption uint8

//LayoutOption options
type LayoutOption uint8

//---------------------------------------------
//---------------- Elements -------------------
//---------------------------------------------

//Load Formals like this order
const (
	_ FormalElementType = iota
	HeadFormalType
	BodyFormalType
	FooterFormalType
)
const (
	_ LayoutElementType = iota
	FreeLayoutType
	StackLayoutType
)

const (
	_ ElementType = iota
	TextElementType
	InputElementType
)

const (
	FullSize = 0
)

//---------------------------------------------
//---------------- Option, Option Types -------
//---------------------------------------------

const (
	_ OrientationType = iota
	VerticalOrientation
	HorizontalOrientation
)
