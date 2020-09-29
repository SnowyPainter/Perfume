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

//GetLocationFromStackElement return absolute location. it must be called continuously with dirLocSum(location direction value(x,y) sum)
type GetLocationFromStackElement func(i int, size Size, originLoc Location, spacing int, dirLocSum *int) Location

var getLocationFromHorizontalStackElement GetLocationFromStackElement = func(i int, size Size, originLoc Location, spacing int, dirLocSum *int) Location {
	originLoc.SetX(*dirLocSum + originLoc.X())
	*dirLocSum += int(size.Width) + spacing
	return originLoc
}

var getLocationFromVerticalStackElement GetLocationFromStackElement = func(i int, size Size, originLoc Location, spacing int, dirLocSum *int) Location {
	originLoc.SetY(*dirLocSum + originLoc.Y())
	*dirLocSum += int(size.Height) + spacing
	return originLoc
}

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

//checkChanges check exists changes from printbuffer
func checkChanges(pb PrintBuffer) bool {
	if c, r := pb.GetChanges(); len(c) <= 0 && len(r) <= 0 {
		return false
	}
	return true
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

//applyElementProperties only apply properties for Element
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

//Buffer buffer strings to print buffer, so, it make non-applied buffers
func (r *Renderer) Buffer() {
	window := r.window
	printBufferAddress := &r.printBuffer
	formalStartsByHeight := 0

	for _, formal := range window.GetFormalByOrder(true) {
		borderExist := 0
		formalElement, _ := window.FindFormal(formal)
		formalLoc := NewLocation(0, formalStartsByHeight)
		formalSize := formalElement.Size()

		//apply formal options
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

			//apply layout options
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
			elementDirectionSum := 0
			elementSpacing := stackLayout.Spacing
			stackElementFunc := getLocationFromHorizontalStackElement
			if orientation == VerticalOrientation {
				stackElementFunc = getLocationFromVerticalStackElement
			}

			for i, element := range elements {
				elementSize := element.Size()

				elementLoc := layoutLoc.Plus(borderExist)
				elementLoc = stackElementFunc(i, elementSize, elementLoc, elementSpacing, &elementDirectionSum)

				applyElementProperties(printBufferAddress, element, elementLoc, elementSize)
				applyStyleOptions(printBufferAddress, element.LoadAllOption(), elementLoc, elementSize)
			}
		}

		formalStartsByHeight += int(formalSize.Height)
	}
}

//Render render formals, layouts, elements to terminal & make buffer apply to terminal
func (r *Renderer) Render() {
	if !checkChanges(r.printBuffer) {
		return
	}
	window := r.window

	//Later, disunite this snippet to channel
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
	r.printBuffer.ApplyChanges()
}

//Clear clear terminal not even clear the terminal, check changes and clear organically
func (r Renderer) Clear() {
	if checkChanges(r.printBuffer) {
		r.clear()
	}
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
