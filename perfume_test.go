package perfume

import (
	"fmt"
	"testing"

	"github.com/nathan-fiscaletti/consolesize-go"
)

//percent return percent that max 100
func percent(v uint, p float32) float32 {
	return float32(v) * (p / 100)
}

func createWindow(row, col uint) (*Window, error) {

	//If use FullSize value, then you must get this 'real' value after NewWindow and check Error
	fullSize := NewSize(FullSize, FullSize)
	testingSize := fullSize //NewSize(row, col)

	window, err := NewWindow(testingSize)
	if err != nil {
		return nil, err
	}
	winSize := window.size
	h := NewSize(uint(percent(winSize.Height, 15)), winSize.Width)
	b := NewSize(uint(percent(winSize.Height, 60)), winSize.Width)
	f := NewSize(uint(percent(winSize.Height, 25)), winSize.Width)

	//Manbody testing size applied
	head := NewHead(h, "Head")
	body := NewBody(b, "Body")
	foot := NewFooter(f, "Footer")
	bodyLayout := NewStackLayout("BodyLayout", b.Plus(-2), VerticalOrientation, 1)
	footLayout := NewFreeLayout("FootLayout", b.Plus(-2))

	t1 := NewText("MyText1", "hello,", NewSize(1, 6))
	t2 := NewText("MyText2", "world!", NewSize(1, 6))
	t3 := NewText("MyText3", "It is foot layout", NewSize(1, foot.Size().Width-2))

	t3Loc := NewRelativeLocation(int(foot.Size().Width/2), 1)
	t3.SetLocation(t3Loc)

	borderOpt := NewOption(BorderOption, "")
	fitOpt := NewOption(FitParentOption, true)

	borderOpt.Set("*")
	head.AddOption(borderOpt.Clone())

	borderOpt.Set("*")
	body.AddOption(borderOpt.Clone())

	borderOpt.Set("*")
	foot.AddOption(borderOpt.Clone())

	borderOpt.Set("=")
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

	cols, rows := consolesize.GetConsoleSize()
	window, err := createWindow(uint(rows), uint(cols))
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
