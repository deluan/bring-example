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

type bringRemoteRenderer struct {
	objects []fyne.CanvasObject
	remote  *BringRemote
}

func (r *bringRemoteRenderer) MinSize() fyne.Size {
	return r.remote.MinSize()
}

func (r *bringRemoteRenderer) Layout(size fyne.Size) {
	if len(r.objects) == 0 {
		return
	}

	r.objects[0].Resize(size)
}

func (r *bringRemoteRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *bringRemoteRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *bringRemoteRenderer) Refresh() {
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

func (r *bringRemoteRenderer) Destroy() {
}

type BringRemote struct {
	widget.BaseWidget
	lastUpdate    int64
	mouseState    *mouseState
	keyboardState *keyboardState

	Display image.Image
	Client  *bring.Client
}

func NewBringRemote(client *bring.Client, width, height int) *BringRemote {
	empty := image.NewNRGBA(image.Rect(0, 0, width-1, height-1))

	b := &BringRemote{
		Client:        client,
		mouseState:    &mouseState{client: client, buttons: make(map[desktop.MouseButton]bool)},
		keyboardState: &keyboardState{client: client},
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

func (b *BringRemote) MinSize() fyne.Size {
	b.ExtendBaseWidget(b)
	return fyne.Size{
		Width:  b.Display.Bounds().Dx(),
		Height: b.Display.Bounds().Dy(),
	}
}

func (b *BringRemote) CreateRenderer() fyne.WidgetRenderer {
	return &bringRemoteRenderer{
		objects: []fyne.CanvasObject{},
		remote:  b,
	}
}

func (b *BringRemote) FocusGained() {
}

func (b *BringRemote) FocusLost() {
}

func (b *BringRemote) Focused() bool {
	return true
}

func (b *BringRemote) TypedRune(ch rune) {
}

func (b *BringRemote) TypedKey(ev *fyne.KeyEvent) {
	b.keyboardState.TypedKey(ev.Name)
}

func (b *BringRemote) KeyDown(ev *fyne.KeyEvent) {
	b.keyboardState.KeyDown(ev.Name)
}

func (b *BringRemote) KeyUp(ev *fyne.KeyEvent) {
	b.keyboardState.KeyUp(ev.Name)
}

func (b *BringRemote) updateDisplay() {
	img, ts := b.Client.Screen()
	if ts != b.lastUpdate {
		b.SetDisplay(img)
		b.lastUpdate = ts
	}
}

func (b *BringRemote) SetDisplay(img image.Image) {
	b.Display = img
	b.Refresh()
}

func (b *BringRemote) MouseDown(ev *desktop.MouseEvent) {
	b.mouseState.MouseDown(ev.Button, ev.Position.X, ev.Position.Y)
	b.updateDisplay()
}

func (b *BringRemote) MouseUp(ev *desktop.MouseEvent) {
	b.mouseState.MouseUp(ev.Button, ev.Position.X, ev.Position.Y)
	b.updateDisplay()
}

func (b *BringRemote) MouseMoved(ev *desktop.MouseEvent) {
	b.mouseState.MouseMove(ev.Position.X, ev.Position.Y)
	b.updateDisplay()
}

func (b *BringRemote) MouseIn(*desktop.MouseEvent) {
}

func (b *BringRemote) MouseOut() {
}

// Make sure all necessary interfaces are implemented
var _ desktop.Hoverable = (*BringRemote)(nil)
var _ desktop.Mouseable = (*BringRemote)(nil)
var _ desktop.Keyable = (*BringRemote)(nil)
