package perfume

//iFormalElement is something that contains ILayouts
type iFormalElement interface {
	Size() Size
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
	Type() ElementType
}

//iBaseElement is the base interface of all elements(layout,formals ...)
type iBaseElement interface {
	GetName() string
	SetName(string)
	LoadAllOption() map[CommonOption]*Option
	LoadOption(CommonOption) *Option
	AddOption(CommonOption, *Option) error
	SetOptionItself(CommonOption, *Option) error
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
	size     Size
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
	kindof   ElementType
	parent   ILayout
}

//ElementBase is structure that is base of all of elements
type ElementBase struct {
	name          string
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

//NewLayout return LayoutElement by EmptyLayoutElemnt(Pointer)
func NewLayout(kindof LayoutElementType, name string) *LayoutElement {
	return EmptyLayout(kindof, name)
}

//NewElement return empty Element
func NewElement(kindof ElementType, name string, loc RelLocation) *Element {
	return EmptyElement(kindof, loc, name)
}

//NewBase return baseelement -> Root for all ofvs
func NewBase(name string) ElementBase {
	return ElementBase{
		name:          name,
		publicOptions: make(map[CommonOption]*Option),
	}
}

//EmptyFormal returns a FormalElement object whose children init
func EmptyFormal(formal FormalElementType, s Size, name string) FormalElement {
	return FormalElement{
		size:        s,
		kindof:      formal,
		children:    make([]ILayout, 0),
		ElementBase: NewBase(name),
	}
}

//EmptyLayout returns parent-nil layout
func EmptyLayout(layout LayoutElementType, name string) *LayoutElement {
	return &LayoutElement{
		ElementBase: NewBase(name),
		kindof:      layout,
		parent:      nil,
		children:    make([]IElement, 0),
	}
}

//EmptyElement returns layoutless element(dependenced)
func EmptyElement(element ElementType, loc RelLocation, name string) *Element {
	return &Element{
		ElementBase: NewBase(name),
		kindof:      element,
		parent:      nil,
		location:    loc,
	}
}
