package perfume

//IFormalElement is something that contains ILayouts
type IFormalElement interface {
	Size() Size
	Type() FormalElementType
	GetChildren() []*ILayout
	AddChild(*ILayout) error
}

//ILayout is something whose parent is IFormalElement and it has only IElement children
type ILayout interface {
	Type() LayoutElementType
	GetParent() *IFormalElement
	GetChildren() *[]IElement
	AddChild(*IElement)
}

//IElement is a interface that is base of TUI
type IElement interface {
	GetLocation() RelLocation
	Type() ElementType
	Parent() *ILayout
}

//FormalElement contains Layout children. it's a structure
type FormalElement struct {
	size     Size
	children []*ILayout
	kindof   FormalElementType
}

//LayoutElement has IElement children and IFormalElement parent
type LayoutElement struct {
	parent   *IFormalElement
	children []*IElement
	kindof   LayoutElementType
}

//Element is structure that is compoent of LayoutElement
type Element struct {
	location RelLocation
	kindof   ElementType
	parent   *ILayout
}

//******Formals*******
type Window struct {
	size    Size
	formals []*FormalElement
}
type Head struct {
	size Size
	FormalElement
}
type Body struct {
	size Size
	FormalElement
}
type Footer struct {
	size Size
	FormalElement
}
type RightSideBar struct {
	FormalElement
}
type LeftSideBar struct {
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
	Element
}
type Text struct {
	Element
}

//NewWindow return new window
func NewWindow(s Size) *Window {
	return &Window{
		size:    s,
		formals: make([]*FormalElement, 0),
	}
}

//NewHead return new head
func NewHead(s Size) *Head {
	return &Head{
		FormalElement: FormalElement{
			size:     s,
			children: make([]*ILayout, 0),
			kindof:   HeadElementType,
		},
	}
}

//NewBody return new body
func NewBody(s Size) *Body {
	return &Body{
		FormalElement: FormalElement{
			size:     s,
			children: make([]*ILayout, 0),
			kindof:   BodyElementType,
		},
	}
}

//******Implements******

//**Except Window**
//AddFormal adds formal element to window
func (w *Window) AddFormal(f *FormalElement) error {
	if f == nil || w.formals == nil {
		return ErrElementIsNil
	}
	w.formals = append(w.formals, f)
	return nil
}

//**FormalELement**

//Size func returns its own size
func (f *FormalElement) Size() Size {
	return f.size
}

//Type func returns its own type
func (f *FormalElement) Type() FormalElementType {
	return f.kindof
}

//GetChildren func returns all of its layout children
func (f *FormalElement) GetChildren() []*ILayout {
	return f.children
}

//AddChild func adds a ILayout to children property
func (f *FormalElement) AddChild(child *ILayout) error {
	if child == nil {
		return ErrChildIsNil
	}

	f.children = append(f.children, child)
	return nil
}

//**LayoutElement**

//Type returns LayoutElementType
func (l *LayoutElement) Type() LayoutElementType {
	return l.kindof
}

//GetParent returns partent(IFormalElement)
func (l *LayoutElement) GetParent() *IFormalElement {
	return l.parent
}

//GetChildren returns children(IElement pointer)
func (l *LayoutElement) GetChildren() []*IElement {
	return l.children
}

//AddChild adds element on children(IElement pointer)
func (l *LayoutElement) AddChild(element *IElement) error {
	if element == nil {
		return ErrChildIsNil
	}
	l.children = append(l.children, element)
	return nil
}

//**Element Component**
//GetLocation return relative location of element
func (e *Element) GetLocation() RelLocation {
	return e.location
}

//Type func return type of element
func (e *Element) Type() ElementType {
	return e.kindof
}

//GetParent func return parent of element
func (e *Element) GetParent() *ILayout {
	return e.parent
}
