package perfume

import (
	"fmt"
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

	borderOpt := NewOption(BorderOption, "*", nil, nil)
	fitOpt := NewOption(FitParentOption, true, nil, nil)

	head.AddOption(borderOpt.Clone())

	borderOpt.Set("=")
	body.AddOption(borderOpt.Clone())

	borderOpt.Set("-")
	foot.AddOption(borderOpt.Clone())

	borderOpt.Set("0")
	hlayout.AddOption(borderOpt.Clone())
	hlayout.AddOption(fitOpt)
	c, err := callSequence(
		body.AddChild(hlayout),
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
