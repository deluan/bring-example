package main

import (
	"strconv"

	"fyne.io/fyne/app"
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
	bringApp := app.New()
	client := createBringClient("vnc", "10.1.0.11", "5901")
	bringDisplay := NewBringDisplay(client, defaultWidth, defaultHeight)

	w := bringApp.NewWindow("Bring it Fyne")
	w.SetContent(widget.NewVBox(
		widget.NewHBox(
			widget.NewButton("Quit", func() {
				bringApp.Quit()
			}),
		),
		bringDisplay,
	))

	w.ShowAndRun()
}
