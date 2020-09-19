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
	bodyLayout := NewStackLayout("BodyLayout", b.Plus(-2), HorizontalOrientation, 1)
	footLayout := NewFreeLayout("FootLayout", b.Plus(-2))

	t1 := NewText("MyText1", "hello,", NewSize(1, 6))
	t2 := NewText("MyText2", "world!", NewSize(1, 6))
	t3 := NewText("MyText3", "It is foot layout", NewSize(1, foot.Size().Width-2))

	t3Loc := NewRelativeLocation(int(foot.Size().Width/2), 0)
	t3.SetLocation(t3Loc)

	borderOpt := NewOption(BorderOption, "")
	fitOpt := NewOption(FitParentOption, false)

	borderOpt.Set("hd-")
	head.AddOption(borderOpt.Clone())

	borderOpt.Set("body-")
	body.AddOption(borderOpt.Clone())

	borderOpt.Set("ft-")
	foot.AddOption(borderOpt.Clone())

	borderOpt.Set("stcklayut-")
	bodyLayout.AddOption(borderOpt.Clone())
	bodyLayout.AddOption(fitOpt)

	c, err := callSequence(
		bodyLayout.AddChild(t1),
		bodyLayout.AddChild(t2),
		footLayout.AddChild(t3),
		body.AddChild(bodyLayout),
		foot.AddChild(footLayout),
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
		ElementsLine: NewParseable("\t\t└-- element LOC:", RelLocationProperty, "\n"),
	})
}
