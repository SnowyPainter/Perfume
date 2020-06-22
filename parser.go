package perfume

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//PrintDepth options
type PrintDepth int

//PrintLineForm has parsable string value
type PrintLineForm int

//PrintPropertyType ppt
type PrintPropertyType string

const (
	_ PrintDepth = iota
	WindowPrintDepth
	FormalsPrintDepth
	LayoutsPrintDepth
	ElementsPrintDepth
)

const (
	_ PrintLineForm = iota
	WindowLine
	FormalsLine
	LayoutsLine
	ElementsLine
)

const (
	_                   PrintPropertyType = ""
	NameProperty                          = "%Name%"
	TypeProperty                          = "%Type%"
	SizeProperty                          = "%Size%"
	ParentNameProperty                    = "%ParentName%"
	ChildrenLenProperty                   = "%ChildrenLen%"
	RelLocationProperty                   = "%RelLocation%"
)

const (
	//PropertyChar is the base parsing character
	PropertyChar = "%" //%[^%%]*%
)

//Parseable is string that can be parsed with values
type Parseable struct {
	content string
}

//NewParseable return parseable(content) pointer
func NewParseable(c string) *Parseable {
	return &Parseable{content: c}
}

func (p *Parseable) findProperties() []string {
	r, _ := regexp.Compile("%[^%%]*%")
	return r.FindAllString(p.content, -1)
}

//Window returns parsed string for line formatting
func (p *Parseable) Window(w *Window) string {
	parsed := p.content
	for _, k := range p.findProperties() {
		switch k {
		case SizeProperty:
			sizeStr := fmt.Sprintf("%d,%d", w.size.Width, w.size.Height)
			parsed = strings.Replace(parsed, SizeProperty, sizeStr, -1)
		case ChildrenLenProperty:
			parsed = strings.Replace(parsed, ChildrenLenProperty, strconv.Itoa(len(w.formals)), -1)
		}
	}
	return parsed
}

//Formal returns parsed string for line formatting
func (p *Parseable) Formal(f IFormal) string {
	parsed := p.content
	for _, k := range p.findProperties() {
		switch k {
		case TypeProperty:
			replace(&parsed, TypeProperty, fmt.Sprintf("%v", f.Type()))
		case SizeProperty:
			size := f.Size()
			sizeStr := fmt.Sprintf("%d,%d", size.Width, size.Height)
			replace(&parsed, SizeProperty, sizeStr)
		case ChildrenLenProperty:
			replace(&parsed, ChildrenLenProperty, strconv.Itoa(f.ChildrenCount()))
		case NameProperty:
			replace(&parsed, NameProperty, f.GetName())
		}
	}
	return parsed
}

//Layout returns parsed string for line formatting
func (p *Parseable) Layout(l ILayout) string {
	parsed := p.content
	for _, k := range p.findProperties() {
		switch k {
		case TypeProperty:
			replace(&parsed, TypeProperty, fmt.Sprintf("%v", l.Type()))
		case ChildrenLenProperty:
			replace(&parsed, ChildrenLenProperty, strconv.Itoa(l.ChildrenCount()))
		case NameProperty:
			replace(&parsed, NameProperty, l.GetName())
		case ParentNameProperty:
			replace(&parsed, ParentNameProperty, l.GetParent().GetName())
		}
	}
	return parsed
}

//Element returns parsed string for line formatting
func (p *Parseable) Element(e IElement) string {
	parsed := p.content
	for _, k := range p.findProperties() {
		switch k {
		case TypeProperty:
			replace(&parsed, TypeProperty, fmt.Sprintf("%v", e.Type()))
		case NameProperty:
			replace(&parsed, NameProperty, e.GetName())
		case ParentNameProperty:
			replace(&parsed, ParentNameProperty, e.GetParent().GetName())
		case RelLocationProperty:
			relLoc := e.GetLocation()
			locStr := fmt.Sprintf("%d, %d", relLoc.X, relLoc.Y)
			replace(&parsed, RelLocationProperty, locStr)
		}
	}
	return parsed
}

func replace(target *string, old string, new string) {
	*target = strings.Replace(*target, old, new, -1)
}
