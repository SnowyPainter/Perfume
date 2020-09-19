package perfume

import "fmt"

func (r *Renderer) isNil(obj interface{}) bool {
	if obj == nil {
		return true
	}
	return false
}

//Renderer render windows and children to terminal
type Renderer struct {
	printBuffer PrintBuffer
	window      *Window
}

//NewRenderer returns renderer pointer it can be nil
func NewRenderer(w *Window) *Renderer {
	return &Renderer{window: w, printBuffer: NewPrintBuffer(w.size)}
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

func bufferBorder(b *PrintBuffer, border string, start RelLocation, s Size) error {
	if start.X < 0 || start.Y < 0 {
		return ErrMinusSize
	}
	x := uint(start.X)
	y := uint(start.Y)

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

func applyStyleOptions(buffer *PrintBuffer, structLevel int, options map[CommonOption]*Option, elementLoc RelLocation, elementSize Size) {
	for key, option := range options {
		switch key {
		case BorderOption:
			bufferBorder(buffer, option.Get().(string), elementLoc, elementSize)
		}
	}
}

//Render render formals, layouts, elements to terminal
func (r *Renderer) Render() {
	window := r.window
	elements := make([]string, 0)

	formalStartsByHeight := 0
	elementLevel := 0

	for _, formal := range window.GetFormalByOrder(true) {
		elementLevel = 0 //Must be at top

		formalElement, _ := window.FindFormal(formal)
		formalLoc := NewRelativeLocation(elementLevel, formalStartsByHeight)
		formalSize := formalElement.Size()

		applyStyleOptions(&r.printBuffer, elementLevel,
			formalElement.LoadAllOption(), formalLoc, formalSize)

		for _, layout := range formalElement.GetChildren() {
			elementLevel = 1 //Must be at top

			layoutSize := layout.Size()
			layoutLoc := NewRelativeLocation(elementLevel, formalStartsByHeight+elementLevel)

			if opt := layout.LoadOption(FitParentOption); opt != nil {
				if opt.Get().(bool) {
					layoutSize = formalSize
					layoutLoc = formalLoc
				}
			}

			applyStyleOptions(&r.printBuffer, elementLevel,
				layout.LoadAllOption(), layoutLoc, layoutSize)

			for _, element := range layout.GetChildren() {
				elementLevel = 2 //Must be at top
				//Style Edit

				elements = append(elements, element.GetName())
			}
		}

		// Summing
		formalStartsByHeight += int(formalSize.Height)
	}

	//Print to console
	for i := uint(0); i < window.size.Height; i++ {
		fmt.Println(r.printBuffer.GetLine(i))
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
