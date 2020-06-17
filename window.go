package perfume

//Window is the base of Perfume as Builder
type Window struct {
	size    Size
	formals map[FormalElementType]IFormal
}

//NewWindow return new window
func NewWindow(s Size) *Window {
	return &Window{
		size:    s,
		formals: make(map[FormalElementType]IFormal),
	}
}

//Add adds formal element to window
func (w *Window) Add(f IFormal) error {
	if f == nil || w.formals == nil {
		return ErrElementIsNil
	}

	if _, err := w.FindFormal(f.Type()); err != nil {
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

	return nil, nil
}

//SetFormal f must be pointer IFormal
func (w *Window) SetFormal(kindof FormalElementType, f IFormal) error {
	if f == nil {
		return ErrElementIsNil
	}

	w.formals[kindof] = f
	return nil
}

//GetFormalSizes return all size of formals that children on window
func (w *Window) GetFormalSizes() (arr []Size) {
	for _, val := range w.formals {
		arr = append(arr, val.Size())
	}
	return
}
