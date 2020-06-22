package perfume

import (
	"testing"
)

func TestRenderer(t *testing.T) {

	window := NewWindow(NewSize(32, 80))
	body := NewBody(NewSize(32, 80), "MainBody")
	stack := NewLayout(StackLayoutType, "MyLayout")
	input := NewElement(InputElementType, "MyInput", NewRelativeLocation(5, 10))

	_ = body.AddChild(stack)
	_ = stack.AddChild(input)
	_ = window.Add(body)

	renderer := NewRenderer(window)

	renderer.PrintStruct(ElementsPrintDepth, map[PrintLineForm]*Parseable{
		WindowLine:   NewParseable("Window || (%Size%) || (%ChildrenLen%) ||\n\n"),
		FormalsLine:  NewParseable("-- (%Name%) Formal --\n"),
		LayoutsLine:  NewParseable("\t└--(%Type%)layout %Name%\n"),
		ElementsLine: NewParseable("\t\t└--(%Type%)element LOC:%RelLocation%\n"),
	})
}
