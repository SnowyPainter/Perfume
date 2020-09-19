package perfume

import (
	"fmt"
	"reflect"
)

const (
	_ CommonOption = iota
	BorderOption
	FitParentOption
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
	Type          CommonOption
}

type iLocation interface {
	X() int
	Y() int
}
type locationData struct {
	x int
	y int
}

//RelLocation is location structure which is relative of parent
type RelLocation struct {
	locationData
}

//Location is static location in console
type Location struct {
	locationData
}

//Size is a structure which has Height and Width
type Size struct {
	Height uint
	Width  uint
}

//**********************
// Constructors
//**********************

//NewOption Create designed option, if handlers nil, it is default in/out handler
func NewOption(optType CommonOption, value interface{}) (opt *Option) {

	returnFunc := func(v interface{}) interface{} {
		return v
	}

	settingFunc := func(opt *Option, v interface{}) {
		if opt != nil {
			opt.value = v
			opt.valueType = reflect.TypeOf(v)
		}
	}

	opt = &Option{}
	opt.value = value
	opt.Type = optType
	opt.valueType = reflect.TypeOf(value)
	opt.SetReturnFunc(returnFunc)
	opt.SetSettingFunc(settingFunc)
	return
}

//newLocationData makes new Locationdata
func newLocationData(x, y int) locationData {
	return locationData{
		x: x, y: y,
	}
}

//NewRelativeLocation makes new RelLocation object
func NewRelativeLocation(x, y int) RelLocation {
	return RelLocation{
		locationData: newLocationData(x, y),
	}
}

//NewLocation makes new Location object
func NewLocation(x, y int) Location {
	return Location{
		locationData: newLocationData(x, y),
	}
}

func (l locationData) X() int {
	return l.x
}
func (l locationData) Y() int {
	return l.y
}
func (l *locationData) SetX(n int) {
	l.x = n
}
func (l *locationData) SetY(n int) {
	l.y = n
}

//Location, RelLocation
func (l Location) Plus(number int) Location {
	return NewLocation(
		l.X()+number,
		l.Y()+number,
	)
}

func SumLocation(location1, location2 iLocation) Location {
	return NewLocation(location1.X()+location2.X(), location1.Y()+location2.Y())
}

//NewSize makes new Size object by height and width
func NewSize(height, width uint) Size {
	return Size{
		Height: height, Width: width,
	}
}

//Plus return number added Size
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
