package internal

import (
	"github.com/janeczku/go-spinner"
	"time"
)

var spinnerCharSet = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
var spinnerSpeed = time.Millisecond * 50

type Terminal struct{}

func (*Terminal) ShowBusy(message string) *spinner.Spinner {
	s := spinner.StartNew(message)
	s.SetSpeed(spinnerSpeed)
	s.SetCharset(spinnerCharSet)
	return s
}

func NewTerminal() Terminal {
	return Terminal{}
}
