package main

import (
	"image"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/widget"
	"github.com/deluan/bring"
)

type bringDisplayRenderer struct {
	objects []fyne.CanvasObject
	remote  *BringDisplay
}

func (r *bringDisplayRenderer) MinSize() fyne.Size {
	return r.remote.MinSize()
}

func (r *bringDisplayRenderer) Layout(size fyne.Size) {
	if len(r.objects) == 0 {
		return
	}

	r.objects[0].Resize(size)
}

func (r *bringDisplayRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *bringDisplayRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *bringDisplayRenderer) Refresh() {
	if len(r.objects) == 0 {
		raster := canvas.NewImageFromImage(r.remote.Display)
		raster.FillMode = canvas.ImageFillContain
		r.objects = append(r.objects, raster)
	} else {
		r.objects[0].(*canvas.Image).Image = r.remote.Display
	}
	r.Layout(r.remote.Size())
	canvas.Refresh(r.remote)
}

func (r *bringDisplayRenderer) Destroy() {
}

type BringDisplay struct {
	widget.BaseWidget
	lastUpdate      int64
	mouseHandler    *mouseHandler
	keyboardHandler *keyboardHandler

	Display image.Image
	Client  *bring.Client
}

func NewBringDisplay(client *bring.Client, width, height int) *BringDisplay {
	empty := image.NewNRGBA(image.Rect(0, 0, width-1, height-1))

	b := &BringDisplay{
		Client:          client,
		mouseHandler:    newMouseHandler(client),
		keyboardHandler: newKeyboardHandler(client),
	}
	b.SetDisplay(empty)
	b.Client.OnSync(func(img image.Image, ts int64) {
		if ts == b.lastUpdate {
			return
		}
		b.lastUpdate = ts
		b.SetDisplay(img)
	})
	go b.Client.Start()
	return b
}

func (b *BringDisplay) MinSize() fyne.Size {
	b.ExtendBaseWidget(b)
	return fyne.Size{
		Width:  b.Display.Bounds().Dx(),
		Height: b.Display.Bounds().Dy(),
	}
}

func (b *BringDisplay) CreateRenderer() fyne.WidgetRenderer {
	return &bringDisplayRenderer{
		objects: []fyne.CanvasObject{},
		remote:  b,
	}
}

func (b *BringDisplay) FocusGained() {
}

func (b *BringDisplay) FocusLost() {
}

func (b *BringDisplay) Focused() bool {
	return true
}

func (b *BringDisplay) TypedRune(ch rune) {
}

func (b *BringDisplay) TypedKey(ev *fyne.KeyEvent) {
	b.keyboardHandler.TypedKey(ev.Name)
}

func (b *BringDisplay) KeyDown(ev *fyne.KeyEvent) {
	b.keyboardHandler.KeyDown(ev.Name)
}

func (b *BringDisplay) KeyUp(ev *fyne.KeyEvent) {
	b.keyboardHandler.KeyUp(ev.Name)
}

func (b *BringDisplay) updateDisplay() {
	img, ts := b.Client.Screen()
	if ts != b.lastUpdate {
		b.SetDisplay(img)
		b.lastUpdate = ts
	}
}

func (b *BringDisplay) SetDisplay(img image.Image) {
	b.Display = img
	b.Refresh()
}

func (b *BringDisplay) MouseDown(ev *desktop.MouseEvent) {
	b.mouseHandler.MouseDown(ev.Button, ev.Position.X, ev.Position.Y)
	b.updateDisplay()
}

func (b *BringDisplay) MouseUp(ev *desktop.MouseEvent) {
	b.mouseHandler.MouseUp(ev.Button, ev.Position.X, ev.Position.Y)
	b.updateDisplay()
}

func (b *BringDisplay) MouseMoved(ev *desktop.MouseEvent) {
	b.mouseHandler.MouseMove(ev.Position.X, ev.Position.Y)
	b.updateDisplay()
}

func (b *BringDisplay) MouseIn(*desktop.MouseEvent) {
}

func (b *BringDisplay) MouseOut() {
}

// Make sure all necessary interfaces are implemented
var _ desktop.Hoverable = (*BringDisplay)(nil)
var _ desktop.Mouseable = (*BringDisplay)(nil)
var _ desktop.Keyable = (*BringDisplay)(nil)
