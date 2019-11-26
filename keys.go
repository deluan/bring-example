package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
	"github.com/deluan/bring"
)

var (
	keys map[fyne.KeyName]bring.KeyCode
	specialKeys map[fyne.KeyName]bool
)

func init() {
	specialKeys = map[fyne.KeyName]bool{
		desktop.KeyAltLeft: true,
		desktop.KeyAltRight: true,
		desktop.KeyControlLeft: true,
		desktop.KeyControlRight: true,
		desktop.KeyShiftLeft: true,
		desktop.KeyShiftRight: true,
	}

	keys = map[fyne.KeyName]bring.KeyCode{
		desktop.KeyAltLeft:      bring.KeyLeftAlt,
		desktop.KeyAltRight:     bring.KeyRightAlt,
		desktop.KeyControlLeft:  bring.KeyLeftControl,
		desktop.KeyControlRight: bring.KeyRightControl,
		desktop.KeyShiftLeft:    bring.KeyLeftShift,
		desktop.KeyShiftRight:   bring.KeyRightShift,
		fyne.KeyBackspace:    bring.KeyBackspace,
		//:     bring.KeyCapsLock,
		fyne.KeyDelete:       bring.KeyDelete,
		fyne.KeyDown:         bring.KeyDown,
		fyne.KeyEnd:          bring.KeyEnd,
		fyne.KeyEnter:        bring.KeyEnter,
		fyne.KeyReturn:        bring.KeyEnter,
		fyne.KeyEscape:       bring.KeyEscape,
		fyne.KeyF1:           bring.KeyF1,
		fyne.KeyF2:           bring.KeyF2,
		fyne.KeyF3:           bring.KeyF3,
		fyne.KeyF4:           bring.KeyF4,
		fyne.KeyF5:           bring.KeyF5,
		fyne.KeyF6:           bring.KeyF6,
		fyne.KeyF7:           bring.KeyF7,
		fyne.KeyF8:           bring.KeyF8,
		fyne.KeyF9:           bring.KeyF9,
		fyne.KeyF10:          bring.KeyF10,
		fyne.KeyF11:          bring.KeyF11,
		fyne.KeyF12:          bring.KeyF12,
		fyne.KeyHome:         bring.KeyHome,
		fyne.KeyInsert:       bring.KeyInsert,
		fyne.KeyLeft:         bring.KeyLeft,
		//:      bring.KeyNumLock,
		fyne.KeyPageDown:     bring.KeyPageDown,
		fyne.KeyPageUp:       bring.KeyPageUp,
		//:        bring.KeyPause,
		//:  bring.KeyPrintScreen,
		fyne.KeyRight:        bring.KeyRight,
		fyne.KeyTab:          bring.KeyTab,
		fyne.KeyUp:           bring.KeyUp,
		//:         bring.KeyMeta,
		//:        bring.KeySuper,
		//:          bring.KeyWin,
	}
}
