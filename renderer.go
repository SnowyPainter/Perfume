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
	window *Window
}

//NewRenderer returns renderer pointer it can be nil
func NewRenderer(w *Window) *Renderer {
	return &Renderer{window: w}
}

//SetWindow sets window of itself it can be nil
func (r *Renderer) SetWindow(w *Window) {
	r.window = w
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
