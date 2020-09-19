package perfume

import (
	"errors"
)

//ErrElement is an error which is relvant with elements
type ErrElement error

var (
	ErrParentIsNil               ErrElement = errors.New("The Parent Element is nil")
	ErrChildIsNil                           = errors.New("Child Element is nil")
	ErrElementIsNil                         = errors.New("Element is nil")
	ErrMinusSize                            = errors.New("There is minus value in Size")
	ErrMinusLocation                        = errors.New("Location X or Y(or both) is minus")
	ErrExistFormal                          = errors.New("It is exist in window")
	ErrNotExist                             = errors.New("Not Exist")
	ErrOutOfWidth                           = errors.New("Width is out of window")
	ErrOutOfHeight                          = errors.New("Height is out of window")
	ErrElementOptionAlreadyExist            = errors.New("Element has already has that option")
	ErrElementOptionDoesntExist             = errors.New("Element doesn't have that option")
	ErrStartIndexOverEndIndex               = errors.New("The start index is over than end index")
	ErrEndIndexOverMax                      = errors.New("The end index is over than max length(ex width)")
	ErrOverSize                             = errors.New("Row/Col value is over than size")
)
