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

//Render render formals, layouts, elements to terminal
func (r *Renderer) Render() {
	window := r.window
	elements := make([]string, 0)

	for _, formal := range window.formals {

		//Style Edit - TEST - TEMPORARY

		if borderOpt := formal.LoadOption(BorderOption); borderOpt != nil {
			border := borderOpt.Get().(string)
			size := formal.Size()
			_ = r.printBuffer.SetRow(border, 0, 0, size.Width)
			_ = r.printBuffer.SetRow(border, size.Height-1, 0, size.Width)
			_ = r.printBuffer.SetColumn(border, 0, 0, size.Height)
			_ = r.printBuffer.SetColumn(border, size.Width-1, 0, size.Height)
		}

		for _, layout := range formal.GetChildren() {

			//Style Edit

			for _, element := range layout.GetChildren() {

				//Style Edit

				elements = append(elements, element.GetName())
			}
		}
	}

	for i := uint(0); i < window.size.Height; i++ {
		fmt.Println(r.printBuffer.GetLine(i))
	}

	/*for i, e := range elements {
		fmt.Println(i, " : ", e)
	}*/
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
