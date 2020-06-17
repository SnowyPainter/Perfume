package perfume

import (
	"fmt"
	"testing"
)

func myStackPlease() *LayoutElement {
	stack := NewLayout(StackLayoutType, "MyLayout")
	input := NewElement(InputElementType, "MyInput", NewRelativeLocation(5, 10))

	err := stack.AddChild(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	return stack
}

func TestNewWindow(t *testing.T) {
	window := NewWindow(NewSize(32, 80))
	body := NewBody(NewSize(32, 80), "MainBody")

	stack := myStackPlease()

	err := body.AddChild(stack)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = window.Add(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	window.PrintTree()
}
