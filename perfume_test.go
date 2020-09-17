package perfume

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRenderer(t *testing.T) {
	testingSize := NewSize(5, 32)

	window, err := NewWindow(testingSize)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//Manbody testing size applied
	body := NewBody(testingSize, "MainBody")

	borderOpt := CreateOption(reflect.TypeOf(""), nil, nil)
	borderOpt.Set("*")
	body.AddOption(
		BorderOption,
		borderOpt,
	)

	err = window.Add(body)
	if err != nil {
		fmt.Println(err, " : ", err.Error())
		return
	}

	renderer := NewRenderer(window)

	renderer.PrintStruct(ElementsPrintDepth, map[PrintLineForm]*Parseable{
		WindowLine:   NewParseable("Window || (", SizeProperty, ") || (", ChildrenLenProperty, ") ||\n\n"),
		FormalsLine:  NewParseable("-- (", NameProperty, ") Formal --\n"),
		LayoutsLine:  NewParseable("\t└--(", TypeProperty, ")layout ", NameProperty, "\n"),
		ElementsLine: NewParseable("\t\t└--(", TypeProperty, ")element LOC:", RelLocationProperty, "\n"),
	})

	fmt.Println("rendering...")

	renderer.Render()
}
