package perfume

import (
	"testing"
)

func TestNewWindow(t *testing.T) {
	window := NewWindow(NewSize(100, 200))
	head := NewHead(NewSize(30, 200)).FormalElement
	body := NewBody(NewSize(30, 200)).FormalElement
	window.AddFormal(&head)
	window.AddFormal(&body)
}
