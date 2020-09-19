package perfume

//iFormalElement is something that contains ILayouts
type iFormalElement interface {
	GetChildren() []ILayout
	ChildrenCount() int
	AddChild(ILayout) error
	Type() FormalElementType
}

//iLayoutElement is something whose parent is iFormalElement and it has only IElement children
type iLayoutElement interface {
	GetParent() IFormal
	GetChildren() []IElement
	ChildrenCount() int
	AddChild(IElement) error
	SetParent(IFormal) error
	Type() LayoutElementType
}

//iElement is a interface that is base of TUI
type iElement interface {
	GetLocation() RelLocation
	GetParent() ILayout
	SetParent(ILayout) error
	SetLocation(RelLocation)
	Type() ElementType
}

//iBaseElement is the base interface of all elements(layout,formals ...)
type iBaseElement interface {
	Size() Size
	GetName() string
	SetName(string)
	LoadAllOption() map[CommonOption]*Option
	LoadOption(CommonOption) *Option
	AddOption(*Option) error
	SetOption(*Option) error
}

//IFormal is a container of all of Formal objects. It must have iFormalElement
type IFormal interface {
	iFormalElement
	iBaseElement
}

//ILayout is a container of all of Layout objects. It must have iLayoutElement
type ILayout interface {
	iLayoutElement
	iBaseElement
}

//IElement is a container. It contains iElement
type IElement interface {
	iElement
	iBaseElement
}

//FormalElement contains Layout children. it's a structure
type FormalElement struct {
	ElementBase
	children []ILayout
	kindof   FormalElementType
}

//LayoutElement has IElement children and iFormalElement parent
type LayoutElement struct {
	ElementBase
	parent   IFormal
	children []IElement
	kindof   LayoutElementType
}

//Element is structure that is compoent of LayoutElement
type Element struct {
	ElementBase
	location RelLocation
	parent   ILayout
	kindof   ElementType
}

//ElementBase is structure that is base of all of elements
type ElementBase struct {
	name          string
	size          Size
	publicOptions map[CommonOption]*Option
}

//******Formals*******
type Head struct {
	FormalElement
}
type Body struct {
	FormalElement
}
type Footer struct {
	FormalElement
}

//******Layouts*******
type FreeLayout struct {
	LayoutElement
}
type StackLayout struct {
	Orientation OrientationType
	Spacing     int
	LayoutElement
}

//******Elements*******
type Input struct {
	kind InputType
	Element
}
type Text struct {
	value string
	Element
}
