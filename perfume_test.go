package perfume

import (
	"fmt"
	"testing"
)

func TestRenderer(t *testing.T) {

	window, err := NewWindow(NewSize(32, 80))
	if err != nil {
		fmt.Println(err.Error())
	}
	body := NewBody(NewSize(32, 80), "MainBody")
	stack := NewLayout(StackLayoutType, "MyLayout")
	input := NewElement(InputElementType, "MyInput", NewRelativeLocation(5, 10))

	_ = body.AddChild(stack)
	_ = stack.AddChild(input)
	_ = window.Add(body)

	renderer := NewRenderer(window)

	renderer.PrintStruct(ElementsPrintDepth, map[PrintLineForm]*Parseable{
		WindowLine:   NewParseable("Window || (", SizeProperty, ") || (", ChildrenLenProperty, ") ||\n\n"),
		FormalsLine:  NewParseable("-- (", NameProperty, ") Formal --\n"),
		LayoutsLine:  NewParseable("\t└--(", TypeProperty, ")layout ", NameProperty, "\n"),
		ElementsLine: NewParseable("\t\t└--(", TypeProperty, ")element LOC:", RelLocationProperty, "\n"),
	})
}
