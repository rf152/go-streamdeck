package actionhandlers

import (
	streamdeck "github.com/rf152/go-streamdeck"
	"github.com/rf152/go-streamdeck/buttons"
)

type TextLabelChangeAction struct {
	NewLabel string
}

func (action *TextLabelChangeAction) Pressed(btn streamdeck.Button) {
	mybtn := btn.(*buttons.TextButton)
	mybtn.SetText(action.NewLabel)
}

func NewTextLabelChangeAction(newLabel string) *TextLabelChangeAction {
	return &TextLabelChangeAction{NewLabel: newLabel}
}
