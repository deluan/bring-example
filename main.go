package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	if len(os.Args) < 4 {
		cmd := filepath.Base(os.Args[0])
		fmt.Printf("Usage: %s <vnc|rdp> address port", cmd)
		os.Exit(1)
	}
	client := createBringClient(os.Args[1], os.Args[2], os.Args[3])

	bringApp := app.New()
	title := fmt.Sprintf("%s (%s:%s)", strings.ToUpper(os.Args[1]), os.Args[2], os.Args[3])
	w := bringApp.NewWindow(title)

	bringDisplay := NewBringDisplay(client, defaultWidth, defaultHeight)
	w.CenterOnScreen()
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
