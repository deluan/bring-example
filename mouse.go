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

type mouseHandler struct {
	display *BringDisplay
	buttons map[desktop.MouseButton]bool
	x, y    int
}

func (ms *mouseHandler) pressedButtons() []bring.MouseButton {
	var buttons []bring.MouseButton
	for b, pressed := range ms.buttons {
		bb := mouseBtnMap[b]
		if pressed {
			buttons = append(buttons, bb)
		}
	}
	return buttons
}

func (ms *mouseHandler) sendMouse(x, y int) {
	ms.x, ms.y = x, y
	if err := ms.display.Client.SendMouse(image.Pt(x, y), ms.pressedButtons()...); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func (ms *mouseHandler) MouseDown(ev *desktop.MouseEvent) {
	ms.buttons[ev.Button] = true
	ms.sendMouse(ev.Position.X, ev.Position.Y)
	ms.display.updateDisplay()
}

func (ms *mouseHandler) MouseUp(ev *desktop.MouseEvent) {
	ms.buttons[ev.Button] = false
	ms.sendMouse(ev.Position.X, ev.Position.Y)
	ms.display.updateDisplay()
}

func (ms *mouseHandler) MouseMoved(ev *desktop.MouseEvent) {
	x, y := ev.Position.X, ev.Position.Y
	if ms.x == x && ms.y == y {
		return
	}
	ms.sendMouse(x, y)
	ms.display.updateDisplay()
}

func (ms *mouseHandler) MouseIn(*desktop.MouseEvent) {
	if ms.buttons == nil {
		ms.buttons = make(map[desktop.MouseButton]bool)
	}
}

func (ms *mouseHandler) MouseOut() {
}

// Make sure all necessary interfaces are implemented
var _ desktop.Hoverable = (*mouseHandler)(nil)
var _ desktop.Mouseable = (*mouseHandler)(nil)
