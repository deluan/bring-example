package main

import (
	"image"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/widget"
)

type bringRemoteRenderer struct {
	objects []fyne.CanvasObject
	remote  *BringRemote
}

// MinSize calculates the minimum size of a button.
// This is based on the contained text, any icon that is set and a standard
// amount of padding added.
func (r *bringRemoteRenderer) MinSize() fyne.Size {
	return r.remote.Size()
}

// Layout the components of the button widget
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
	Display image.Image
}

func NewBringRemote(width, height int) *BringRemote {
	empty := image.NewNRGBA(image.Rect(0, 0, width-1, height-1))

	w := &BringRemote{}
	w.SetImage(empty)
	return w
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

func (b *BringRemote) SetImage(img image.Image) {
	b.Display = img
	b.Refresh()
}

func (b *BringRemote) MouseMoved(ev *desktop.MouseEvent) {
	//fmt.Printf("%#v\n", ev)
}

func (b *BringRemote) MouseIn(*desktop.MouseEvent) {
	//panic("implement me")
}

func (b *BringRemote) MouseOut() {
	//panic("implement me")
}

var _ desktop.Hoverable = (*BringRemote)(nil)
