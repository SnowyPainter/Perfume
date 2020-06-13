package perfume

import (
	"errors"
)

//ErrElement is an error which is relvant with elements
type ErrElement error

var (
	ErrParentIsNil   ErrElement = errors.New("The Parent Element is nil")
	ErrChildIsNil               = errors.New("Child Element is nil")
	ErrElementIsNil             = errors.New("Element is nil")
	ErrMinusSize                = errors.New("There is minus value in Size")
	ErrMinusLocation            = errors.New("Location X or Y(or both) is minus")
)
