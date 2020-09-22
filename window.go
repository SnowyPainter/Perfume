package perfume

import (
	"sort"

	"github.com/nathan-fiscaletti/consolesize-go"
)

//getFullScreen return fullscreen col/rows
func getFullScreen() Size {
	cols, rows := consolesize.GetConsoleSize()
	return NewSize(uint(rows), uint(cols))
}

//Window is the base of Perfume as Builder
type Window struct {
	isFullScreen bool
	size         Size
	formals      map[FormalElementType]IFormal
}

//NewWindow return new window 0 doesn't work & doesn't work except terminal console
func NewWindow(s Size) (*Window, error) {
	fullscreen := getFullScreen()

	if s.Width == FullSize {
		s.Width = fullscreen.Width
	}
	if s.Height == FullSize {
		s.Height = fullscreen.Height
	}

	if fullscreen.Width < s.Width {
		return nil, ErrOutOfWidth
	} else if fullscreen.Height < s.Height {
		return nil, ErrOutOfHeight
	}
	return &Window{
		size:    s,
		formals: make(map[FormalElementType]IFormal),
	}, nil
}

//Add adds formal element to window
func (w *Window) Add(f IFormal) error {

	if f == nil || w.formals == nil {
		return ErrElementIsNil
	}

	if _, err := w.FindFormal(f.Type()); err == nil {
		return ErrExistFormal
	}

	size := f.Size()
	sumOfHeight := size.Height
	for _, s := range w.GetFormalSizes() {
		sumOfHeight += s.Height
	}
	if size.Width > w.size.Width {
		return ErrOutOfWidth
	}
	if sumOfHeight > w.size.Height {
		return ErrOutOfHeight
	}

	w.formals[f.Type()] = f

	return nil
}

//FindFormal get index of speific formal type.
func (w *Window) FindFormal(targetType FormalElementType) (IFormal, error) {
	if val, found := w.formals[targetType]; found {
		return val, nil
	}

	return nil, ErrNotExist
}

//SetFormal f must be pointer IFormal
func (w *Window) SetFormal(kindof FormalElementType, f IFormal) error {
	if f == nil {
		return ErrElementIsNil
	}

	w.formals[kindof] = f
	return nil
}

func (w *Window) GetFormalByOrder(ascending bool) (formals []FormalElementType) {
	keys := make([]FormalElementType, 0)
	for key := range w.formals {
		keys = append(keys, key)
	}
	sliceFunc := func(i, j int) bool { return keys[i] > keys[j] }
	if ascending {
		sliceFunc = func(i, j int) bool { return keys[i] < keys[j] }
	}
	sort.Slice(keys, sliceFunc)

	return keys
}

//GetFormalSizes return all size of formals that children on window
func (w *Window) GetFormalSizes() (arr []Size) {
	for _, val := range w.formals {
		arr = append(arr, val.Size())
	}
	return
}
