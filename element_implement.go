package perfume

//******Implements******

//ElementBase
//GetName return name(string)
func (eb *ElementBase) GetName() string {
	return eb.name
}

//SetName set name(string)
func (eb *ElementBase) SetName(name string) {
	eb.name = name
}

//**FormalELement**
//Type func returns its own type

//ChildrenCount return count of children it has. (time complex)
func (f *FormalElement) ChildrenCount() int {
	return len(f.children)
}

func (f *FormalElement) Type() FormalElementType {
	return f.kindof
}

//Size func returns its own size
func (f *FormalElement) Size() Size {
	return f.size
}

//GetChildren func returns all of its layout children
func (f *FormalElement) GetChildren() []ILayout {
	return f.children
}

//AddChild func adds a ILayout to children property
func (f *FormalElement) AddChild(child ILayout) error {
	if child == nil {
		return ErrChildIsNil
	}
	err := child.SetParent(f)
	if err != nil {
		return err
	}
	f.children = append(f.children, child)
	return nil
}

//**LayoutElement**

//ChildrenCount return count of children it has. (time complex)
func (l *LayoutElement) ChildrenCount() int {
	return len(l.children)
}

//Type returns LayoutElementType
func (l *LayoutElement) Type() LayoutElementType {
	return l.kindof
}

//GetParent returns partent(IFormalElement)
func (l *LayoutElement) GetParent() IFormal {
	return l.parent
}

//GetChildren returns children(IElement pointer)
func (l *LayoutElement) GetChildren() []IElement {
	return l.children
}

//AddChild adds element on children(IElement pointer)
func (l *LayoutElement) AddChild(element IElement) error {
	if element == nil {
		return ErrChildIsNil
	}
	element.SetParent(l)
	l.children = append(l.children, element)
	return nil
}

//SetParent set its parent(pointer)
func (l *LayoutElement) SetParent(formal IFormal) error {
	if formal == nil {
		return ErrParentIsNil
	}
	l.parent = formal
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
func (e *Element) GetParent() ILayout {
	return e.parent
}

//SetParent set its parent(pointer)
func (e *Element) SetParent(formal ILayout) error {
	if e.parent == nil || formal == nil {
		return ErrParentIsNil
	}
	e.parent = formal
	return nil
}
