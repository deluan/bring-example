package main

import (
	"fmt"
	"image"

	"fyne.io/fyne/driver/desktop"
	"github.com/deluan/bring"
)

var mouseBtnMap = map[desktop.MouseButton]bring.MouseButton{
	desktop.LeftMouseButton:  bring.MouseLeft,
	desktop.RightMouseButton: bring.MouseRight,
}

type mouseState struct {
	client  *bring.Client
	buttons map[desktop.MouseButton]bool
	x, y    int
}

func (ms *mouseState) pressedButtons() []bring.MouseButton {
	var buttons []bring.MouseButton
	for b, pressed := range ms.buttons {
		bb := mouseBtnMap[b]
		if pressed {
			buttons = append(buttons, bb)
		}
	}
	return buttons
}

func (ms *mouseState) sendMouse(x, y int) {
	ms.x, ms.y = x, y
	if err := ms.client.SendMouse(image.Pt(x, y), ms.pressedButtons()...); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func (ms *mouseState) MouseDown(button desktop.MouseButton, x, y int) {
	ms.buttons[button] = true
	ms.sendMouse(x, y)
}

func (ms *mouseState) MouseUp(button desktop.MouseButton, x, y int) {
	ms.buttons[button] = false
	ms.sendMouse(x, y)
}

func (ms *mouseState) MouseMove(x, y int) {
	if ms.x == x && ms.y == y {
		return
	}
	ms.sendMouse(x, y)
}
