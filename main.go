package main

import (
	"image"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"github.com/deluan/bring"
	"github.com/sirupsen/logrus"
)

const (
	guacdAddress  = "localhost:4822"
	defaultWidth  = 1024
	defaultHeight = 768
)

// Creates and initialize Bring's Session and Client
func createBringClient(protocol, hostname, port string) *bring.Client {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true, ForceColors: true})
	logger.SetLevel(logrus.DebugLevel)

	client, err := bring.NewClient(guacdAddress, protocol, map[string]string{
		"hostname": hostname,
		"port":     port,
		"password": "vncpassword",
		"width":    strconv.Itoa(defaultWidth),
		"height":   strconv.Itoa(defaultHeight),
	}, logger)
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	app := app.New()
	remote := NewBringRemote(defaultWidth, defaultHeight)

	w := app.NewWindow("Bring it Fyne")
	w.SetContent(widget.NewVBox(
		widget.NewHBox(
			widget.NewButton("Quit", func() {
				app.Quit()
			}),
		),
		remote,
	))
	client := createBringClient("vnc", "10.0.0.11", "5901")

	w.Canvas().SetOnTypedRune(func(ch rune) {
		_ = client.SendText(string(ch))
	})
	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		if k, ok := keys[ev.Name]; ok && !specialKeys[ev.Name] {
			_ = client.SendKey(k, true)
			_ = client.SendKey(k, false)
		}
	})

	var lastUpdate int64
	client.OnSync(func(img image.Image, ts int64) {
		if ts == lastUpdate {
			return
		}
		lastUpdate = ts
		remote.SetImage(img)
		canvas.Refresh(remote)
	})
	go client.Start()

	w.ShowAndRun()
}
