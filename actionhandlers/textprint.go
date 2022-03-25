package actionhandlers

import (
	"fmt"
	streamdeck "github.com/rf152/go-streamdeck"
)

type TextPrintAction struct {
	Label string
}

func (action *TextPrintAction) Pressed(btn streamdeck.Button) {
	fmt.Println(action.Label)
	fmt.Print("The button pressed is: ")
	fmt.Println(btn)
}

func NewTextPrintAction(label string) *TextPrintAction {
	return &TextPrintAction{Label: label}
}
