package perfume

import "fmt"

//NewHead return new head
func NewHead(s Size, name string) *Head {
	return &Head{
		FormalElement: EmptyFormal(HeadFormalType, s, name),
	}
}

//NewBody return new body
func NewBody(s Size, name string) *Body {
	return &Body{
		FormalElement: EmptyFormal(BodyFormalType, s, name),
	}
}

//NewFooter return new footer
func NewFooter(s Size, name string) *Footer {
	return &Footer{
		FormalElement: EmptyFormal(FooterFormalType, s, name),
	}
}

//NewBase return baseelement -> Root for all ofvs
func NewBase(name string, size Size) ElementBase {
	return ElementBase{
		name:          name,
		size:          size,
		publicOptions: make(map[CommonOption]*Option),
	}
}

//EmptyFormal returns a FormalElement object whose children init
func EmptyFormal(formal FormalElementType, s Size, name string) FormalElement {
	return FormalElement{
		kindof:      formal,
		children:    make([]ILayout, 0),
		ElementBase: NewBase(name, s),
	}
}

//EmptyLayout returns parent-nil layout
func EmptyLayout(layout LayoutElementType, s Size, name string) LayoutElement {
	return LayoutElement{
		ElementBase: NewBase(name, s),
		kindof:      layout,
		parent:      nil,
		children:    make([]IElement, 0),
	}
}

//EmptyElement returns layoutless element(dependenced)
func EmptyElement(element ElementType, loc RelLocation, s Size, name string) Element {
	return Element{
		ElementBase: NewBase(name, s),
		kindof:      element,
		parent:      nil,
		location:    loc,
	}
}

//Layouts ....
func NewStackLayout(name string, size Size, oriType OrientationType, spacing int) *StackLayout {
	return &StackLayout{
		LayoutElement: EmptyLayout(StackLayoutType, size, name),
		Orientation:   oriType,
		Spacing:       spacing,
	}
}
func NewFreeLayout(name string, size Size) *FreeLayout {
	return &FreeLayout{
		LayoutElement: EmptyLayout(FreeLayoutType, size, name),
	}
}

//Elements ...
func NewText(name string, text string, size Size) *Text {
	if text != "" && size.Height == 0 {
		fmt.Println("Warning, size.Height is 0")
		return nil
	}

	return &Text{
		value:   text,
		Element: EmptyElement(TextElementType, NewRelativeLocation(0, 0), size, name),
	}
}
