package perfume

import (
	"fmt"
	"reflect"
	"testing"
)

func createWindow() (*Window, error) {
	testingSize := NewSize(20, 128)
	h := NewSize(4, testingSize.Width)
	b := NewSize(12, testingSize.Width)
	f := NewSize(4, testingSize.Width)

	window, err := NewWindow(testingSize)
	if err != nil {
		return nil, err
	}
	//Manbody testing size applied
	head := NewHead(h, "Head")
	body := NewBody(b, "Body")
	foot := NewFooter(f, "Footer")
	hlayout := NewLayout(StackLayoutType, h.Plus(-2), "HeadLayout")

	borderOpt := CreateOption(reflect.TypeOf(""), nil, nil)
	borderOpt.Set("*")
	head.AddOption(BorderOption, borderOpt.Clone())

	borderOpt.Set("=")
	body.AddOption(BorderOption, borderOpt.Clone())

	borderOpt.Set("-")
	foot.AddOption(BorderOption, borderOpt.Clone())

	borderOpt.Set("0")
	hlayout.AddOption(BorderOption, borderOpt.Clone())

	c, err := callSequence(
		head.AddChild(hlayout),
		window.Add(body),
		window.Add(head),
		window.Add(foot),
	)
	if err != nil {
		fmt.Println(c, " : ", err.Error())
		return nil, err
	}
	return window, nil
}

func TestRenderer(t *testing.T) {
	window, err := createWindow()
	if err != nil {
		fmt.Println(err)
		return
	}
	renderer := NewRenderer(window)
	printRendererStruct(renderer)
	renderer.Render()
}

func printRendererStruct(r *Renderer) {
	r.PrintStruct(ElementsPrintDepth, map[PrintLineForm]*Parseable{
		WindowLine:   NewParseable("Window || (", SizeProperty, ") || (", ChildrenLenProperty, ") ||\n\n"),
		FormalsLine:  NewParseable("-- (", NameProperty, ") Formal --\n"),
		LayoutsLine:  NewParseable("\t└--(", TypeProperty, ")layout ", NameProperty, "\n"),
		ElementsLine: NewParseable("\t\t└--(", TypeProperty, ")element LOC:", RelLocationProperty, "\n"),
	})
}
