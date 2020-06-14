package perfume

import (
	"fmt"
	"testing"
)

func TestNewWindow(t *testing.T) {
	window := NewWindow(NewSize(100, 200))
	head := NewHead(NewSize(70, 200), "HeadyHeady")
	body := NewBody(NewSize(30, 200), "BodyBody")
	stack := NewLayout(StackLayoutType)
	input := NewElement(InputElementType, NewRelativeLocation(5, 10))

	err := stack.AddChild(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = body.AddChild(stack)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = window.Add(head)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = window.Add(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	window.PrintTree()
}
