package perfume

//IFormalElement is something that contains ILayouts
type IFormalElement interface {
	Size() Size
	Type() FormalElementType
	GetChildren() []ILayout
	AddChild(ILayout) error
}

//ILayoutElement is something whose parent is IFormalElement and it has only IElement children
type ILayoutElement interface {
	Type() LayoutElementType
	GetParent() IFormal
	GetChildren() []IElement
	AddChild(IElement) error
	SetParent(IFormal) error
}

//IBaseElement is a interface that is base of TUI
type IBaseElement interface {
	GetLocation() RelLocation
	Type() ElementType
	GetParent() ILayout
	SetParent(ILayout) error
}

//IFormal is a container of all of Formal objects. It must have IFormalElement
type IFormal interface {
	GetName() string
	SetName(string)
	IFormalElement
}

//ILayout is a container of all of Layout objects. It must have ILayoutElement
type ILayout interface {
	ILayoutElement
}

//IElement is a container. It contains IBaseElement
type IElement interface {
	IBaseElement
}

//FormalElement contains Layout children. it's a structure
type FormalElement struct {
	name     string
	size     Size
	children []ILayout
	kindof   FormalElementType
}

//LayoutElement has IElement children and IFormalElement parent
type LayoutElement struct {
	parent   IFormal
	children []IElement
	kindof   LayoutElementType
}

//Element is structure that is compoent of LayoutElement
type Element struct {
	location RelLocation
	kindof   ElementType
	parent   ILayout
}

//******Formals*******
type Head struct {
	FormalElement
}
type Body struct {
	FormalElement
}
type Footer struct {
	size Size
	FormalElement
}

//******Layouts*******
type FreeLayout struct {
	LayoutElement
}
type StackLayout struct {
	LayoutElement
}

//******Elements*******
type Input struct {
	kind InputType
	Element
}
type Text struct {
	Element
}

//NewHead return new head
func NewHead(s Size, name string) *Head {
	return &Head{
		FormalElement: EmptyFormal(HeadElementType, s, name),
	}
}

//NewBody return new body
func NewBody(s Size, name string) *Body {
	return &Body{
		FormalElement: EmptyFormal(BodyElementType, s, name),
	}
}

//NewElement return empty Element
func NewElement(kindof ElementType, loc RelLocation) *Element {
	return &Element{
		location: loc,
		kindof:   kindof,
		parent:   nil,
	}
}

//NewLayout return LayoutElement by EmptyLayoutElemnt(Pointer)
func NewLayout(kindof LayoutElementType) *LayoutElement {
	return EmptyLayout(kindof)
}

//EmptyFormal returns a FormalElement object whose children init
func EmptyFormal(formal FormalElementType, s Size, name string) FormalElement {
	return FormalElement{
		size:     s,
		kindof:   formal,
		children: make([]ILayout, 0),
		name:     name,
	}
}

//EmptyLayout returns parent-nil layout
func EmptyLayout(layout LayoutElementType) *LayoutElement {
	return &LayoutElement{
		kindof:   layout,
		parent:   nil,
		children: make([]IElement, 0),
	}
}
