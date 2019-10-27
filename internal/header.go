package internal

import (
	"github.com/fatih/color"
)

type Header struct {
	Logo string
	Tag  string
}

func (h *Header) String() string {
	return h.Logo + "\n" + h.Tag + "\n"
}

func (h *Header) Print() {
	logoColor := color.New(color.FgHiBlack)
	_, _ = logoColor.Println(h.Logo)

	tagColor := color.New(color.FgBlue)
	_, _ = tagColor.Println(h.Tag + "\n")

}

func NewHeader() *Header {
	var header Header
	header.Logo =
		`
	   ▄██████▄     ▄████████ ███▄▄▄▄   ███▄▄▄▄   ▄██   ▄   
	  ███    ███   ███    ███ ███▀▀▀██▄ ███▀▀▀██▄ ███   ██▄ 
	  ███    █▀    ███    █▀  ███   ███ ███   ███ ███▄▄▄███ 
	 ▄███         ▄███▄▄▄     ███   ███ ███   ███ ▀▀▀▀▀▀███ 
	▀▀███ ████▄  ▀▀███▀▀▀     ███   ███ ███   ███ ▄██   ███ 
	  ███    ███   ███    █▄  ███   ███ ███   ███ ███   ███ 
	  ███    ███   ███    ███ ███   ███ ███   ███ ███   ███ 
	  ████████▀    ██████████  ▀█   █▀   ▀█   █▀   ▀█████▀  
`
	header.Tag = "		   *** gRPC Generator for Go ***"
	return &header
}
