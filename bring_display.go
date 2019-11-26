package main

import (
	"image"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
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

// Custom Widget that represents the remote computer being controlled
type BringDisplay struct {
	widget.BaseWidget
	keyboardHandler
	mouseHandler

	lastUpdate int64

	Display image.Image
	Client  *bring.Client
}

// Creates a new BringDisplay and does all the heavy lifting, setting up all event handlers
func NewBringDisplay(client *bring.Client, width, height int) *BringDisplay {
	empty := image.NewNRGBA(image.Rect(0, 0, width-1, height-1))

	b := &BringDisplay{
		Client: client,
	}
	b.keyboardHandler.display = b
	b.mouseHandler.display = b

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

// This forces a display refresh if there are any pending updates, instead of waiting for the next
// OnSync event. Ex: when moving the mouse we want instant feedback of the new cursor position
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
