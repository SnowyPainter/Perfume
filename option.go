package perfume

import (
	"fmt"
	"reflect"
)

const (
	_ CommonOption = iota
	BorderOption
)

const (
	_ ComponentOption = iota
	MarginOption
	PaddingOption
	WidthOption
	HeightOption
)
const (
	_ LayoutOption = iota
	SpacingOption
	OrientationOption
)

//OptionGetHandler handle value when calls Get
type OptionGetHandler func(val interface{}) interface{}

//OptionSetHandler handle value when calls Set
type OptionSetHandler func(opt *Option, val interface{})

type Option struct {
	value         interface{}
	getHandleFunc OptionGetHandler
	setHandleFunc OptionSetHandler
	valueType     reflect.Type
}

//RelLocation is location structure which is relative of parent
type RelLocation struct {
	X int
	Y int
}

//Size is a structure which has Height and Width
type Size struct {
	Height uint
	Width  uint
}

//**********************
// Constructors
//**********************

//CreateOption Create designed option, if handlers nil, it is default in/out handler
func CreateOption(valType reflect.Type, returnFunc OptionGetHandler, settingFunc OptionSetHandler) (opt *Option) {

	if returnFunc == nil {
		//Default
		returnFunc = func(v interface{}) interface{} {
			return v
		}
	}
	if settingFunc == nil {
		//Default
		settingFunc = func(opt *Option, v interface{}) {
			if opt != nil {
				opt.value = v
				opt.valueType = reflect.TypeOf(v)
			}
		}
	}

	opt = &Option{}
	opt.valueType = valType
	opt.SetReturnFunc(returnFunc)
	opt.SetSettingFunc(settingFunc)
	return
}

//NewRelativeLocation makes new RelLocation object
func NewRelativeLocation(x int, y int) RelLocation {
	return RelLocation{
		X: x, Y: y,
	}
}

//NewSize makes new Size object by height and width
func NewSize(height uint, width uint) Size {
	return Size{
		Height: height, Width: width,
	}
}

func (s Size) Plus(number int) Size {
	n := uint(number)
	if s.Height-n < 1 || s.Width-n < 1 {
		return Size{}
	}
	size := Size{
		Height: s.Height + n,
		Width:  s.Width + n,
	}
	return size
}

//SetSettingFunc set Set func property func
func (o *Option) SetSettingFunc(f OptionSetHandler) {
	o.setHandleFunc = f
}

//SetReturnFunc set Get func property func
func (o *Option) SetReturnFunc(f OptionGetHandler) {
	o.getHandleFunc = f
}

func (o Option) Get() interface{} {
	return o.getHandleFunc(o.value)
}
func (o *Option) Set(val interface{}) {
	o.setHandleFunc(o, val)
}

func (o Option) Clone() *Option {
	return &o
}

func (size Size) String() string {
	return fmt.Sprintf("Height : %d, Width : %d", size.Height, size.Width)
}
