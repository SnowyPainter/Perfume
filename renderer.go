package perfume

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func (r *Renderer) isNil(obj interface{}) bool {
	if obj == nil {
		return true
	}
	return false
}

type StackElement func(i int, size Size, originLoc Location, spacing int, dirSum *int) Location

//Renderer render windows and children to terminal
type Renderer struct {
	clear       func()
	printBuffer PrintBuffer
	window      *Window
}

//NewRenderer returns renderer pointer it can be nil
func NewRenderer(w *Window) *Renderer {
	goos := runtime.GOOS
	var cf func()
	if goos == "linux" {
		cf = func() {
			cmd := exec.Command("clear") //Linux example, its tested
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	} else if goos == "windows" {
		cf = func() {
			cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}

	return &Renderer{
		window:      w,
		printBuffer: NewPrintBuffer(w.size),
		clear:       cf,
	}
}

//SetWindow sets window of itself it can be nil
func (r *Renderer) SetWindow(w *Window) {
	r.window = w
}

func callSequence(errors ...error) (int, error) {
	for i, err := range errors {
		if err != nil {
			return i, err
		}
	}
	return -1, nil
}

//checkBorderExist return border opt exists and it vailed
func checkBorderExist(baseElement iBaseElement) bool {
	if opt := baseElement.LoadOption(BorderOption); opt != nil && opt.Get().(string) != "" {
		return true
	}
	return false
}

//bufferBorder draw border in PrintBuffer
func bufferBorder(b *PrintBuffer, border string, start Location, s Size) error {
	if start.X() < 0 || start.Y() < 0 {
		return ErrMinusSize
	}
	x := uint(start.X())
	y := uint(start.Y())

	c, err := callSequence(
		b.SetRow(border, y, x, s.Width),
		b.SetRow(border, y+s.Height-1, x, s.Width),
		b.SetColumn(border, x, y, s.Height+y),
		b.SetColumn(border, x+s.Width-1, y, s.Height+y),
	)
	if err != nil {
		panic(fmt.Sprintf("%d : %s", c, err.Error()))
	}
	return nil
}

//bufferText draw text
func bufferText(b *PrintBuffer, text string, loc Location) error {
	x := uint(loc.X())
	y := uint(loc.Y())
	txtLen := uint(len(text))
	if x+txtLen >= b.size.Width {
		return ErrOutOfWidth
	}
	if y >= b.size.Height {
		return ErrOutOfHeight
	}

	err := b.SetRow(text, y, x, x+txtLen)
	if err != nil {
		return err
	}
	return nil
}

func applyStyleOptions(buffer *PrintBuffer, options map[CommonOption]*Option, elementLoc Location, elementSize Size) {
	for key, option := range options {
		switch key {
		case BorderOption:
			bufferBorder(buffer, option.Get().(string), elementLoc, elementSize)
		}
	}
}

func applyElementProperties(buffer *PrintBuffer, element IElement, location Location, size Size) {
	switch element.Type() {
	case TextElementType:
		txtElement := element.(*Text)
		txt := txtElement.Text()
		width := int(element.Size().Width)
		maxLen := len(txt)
		if len(txt) > width {
			maxLen = width
		}
		bufferText(buffer, txt[:maxLen], location)
	case InputElementType:
		//...
	}
}

//Render render formals, layouts, elements to terminal
func (r *Renderer) Render() {
	window := r.window
	printBufferAddress := &r.printBuffer
	formalStartsByHeight := 0

	for _, formal := range window.GetFormalByOrder(true) {
		borderExist := 0
		formalElement, _ := window.FindFormal(formal)
		formalLoc := NewLocation(0, formalStartsByHeight)
		formalSize := formalElement.Size()

		applyStyleOptions(printBufferAddress, formalElement.LoadAllOption(), formalLoc, formalSize)

		if checkBorderExist(formalElement) {
			borderExist = 1
		}

		for _, layout := range formalElement.GetChildren() {

			layoutSize := layout.Size()
			layoutLoc := NewLocation(borderExist, formalStartsByHeight+borderExist)
			if checkBorderExist(layout) {
				borderExist = 1
			} else {
				borderExist = 0
			}
			if opt := layout.LoadOption(FitParentOption); opt != nil {
				if opt.Get().(bool) {
					layoutSize = formalSize
					layoutLoc = formalLoc
				}
			}

			applyStyleOptions(printBufferAddress, layout.LoadAllOption(), layoutLoc, layoutSize)

			elements := layout.GetChildren()
			if layout.Type() == FreeLayoutType {
				for _, element := range elements {
					elementLoc := SumLocation(element.GetLocation(), layoutLoc.Plus(borderExist))
					elementSize := element.Size()
					applyElementProperties(printBufferAddress, element, elementLoc, elementSize)
					applyStyleOptions(printBufferAddress, element.LoadAllOption(), elementLoc, elementSize)
				}
				continue
			}

			stackLayout := layout.(*StackLayout)
			orientation := stackLayout.Orientation
			elementSpacing := stackLayout.Spacing
			var stackElementFunc StackElement
			if orientation == HorizontalOrientation {
				stackElementFunc = func(i int, size Size, originLoc Location, spacing int, dirSum *int) Location {
					originLoc.SetX(*dirSum + originLoc.X())
					*dirSum += int(size.Width) + spacing
					return originLoc
				}
			} else { //Vertical
				stackElementFunc = func(i int, size Size, originLoc Location, spacing int, dirSum *int) Location {
					originLoc.SetY(*dirSum + originLoc.Y())
					*dirSum += int(size.Height) + spacing
					return originLoc
				}
			}
			elementDir := 0
			for i, element := range elements {
				elementLoc := layoutLoc.Plus(borderExist)
				elementSize := element.Size()
				//Stack up elements..
				elementLoc = stackElementFunc(i, elementSize, elementLoc, elementSpacing, &elementDir)
				applyElementProperties(printBufferAddress, element, elementLoc, elementSize)
				applyStyleOptions(printBufferAddress, element.LoadAllOption(), elementLoc, elementSize)
			}
		}

		// Summing
		formalStartsByHeight += int(formalSize.Height)
	}

	//Later, disunite this snippet to channel
	//Print to console

	fullscreen := getFullScreen()
	end := ""
	for i := uint(0); i < window.size.Height; i++ {
		if fullscreen.Width <= window.size.Width {
			end = ""
		} else {
			end = "\n"
		}
		fmt.Printf("%s%s", r.printBuffer.GetLine(i), end)
	}

}

func (r Renderer) Clear() {
	r.clear()
}

//PrintStruct prints information of window
func (r *Renderer) PrintStruct(depth PrintDepth, form map[PrintLineForm]*Parseable) {
	window := r.window

	fmt.Printf(form[WindowLine].Window(window))

	if depth <= WindowPrintDepth {
		return
	}

	for _, f := range window.formals {
		layouts := f.GetChildren()
		fmt.Printf(form[FormalsLine].Formal(f))
		if depth >= LayoutsPrintDepth {
			for _, l := range layouts {
				elements := l.GetChildren()
				fmt.Printf(form[LayoutsLine].Layout(l))
				if depth >= ElementsPrintDepth {
					for _, e := range elements {
						fmt.Printf(form[ElementsLine].Element(e))
					}
				}
			}
		}

		fmt.Printf("\n")
	}
}
