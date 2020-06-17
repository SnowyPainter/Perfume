package perfume

import "fmt"

//PrintTree prints all data of window
func (w *Window) PrintTree() {
	fmt.Printf("Window || (%d,%d) Size || %d Formals ||\n\n", w.size.Height, w.size.Width, len(w.formals))
	for _, f := range w.formals {
		children := f.GetChildren()
		fmt.Printf("======= %s =======(%d layouts)\n", f.GetName(), len(children))
		for i, l := range children {
			c := l.GetChildren()
			fmt.Printf("layout\t=(%d)= Type =(%d)= Element =(%d)=\n", i, l.Type(), len(c))
			fmt.Printf("----------------------------------------\n")
			for j, e := range c {
				fmt.Printf("element\t[%d] Its Type =(%d)=\n", j, e.Type())
			}
		}
		fmt.Printf("\n")
	}
}
