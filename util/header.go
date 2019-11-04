package util

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
	logoColor := color.New(color.FgMagenta)
	_, _ = logoColor.Println(h.Logo)
}

func NewHeader() *Header {
	var header Header
	header.Logo = `
                            ,e,         
    e88 888  ,e e,  888 8e   "   ,e e,  
   d888 888 d88 88b 888 88b 888 d88 88b 
   Y888 888 888   , 888 888 888 888   , 
    "88 888  "YeeP" 888 888 888  "YeeP" 
     ,  88P                             
    "8",P"     Go Service Generator                           
`
	return &header
}
