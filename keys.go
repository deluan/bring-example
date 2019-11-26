package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
	"github.com/deluan/bring"
)

type keyboardHandler struct {
	client *bring.Client
	shift  bool
	caps   bool
}

func (ks *keyboardHandler) sendKey(key bring.KeyCode, pressed bool) {
	k := int(key)
	if k >= int('A') && k <= int('Z') {
		if !ks.shift && !ks.caps {
			k = k + 32
		}
	}
	_ = ks.client.SendKey(bring.KeyCode(k), pressed)
}

func mapDesktopKey(keyName fyne.KeyName) bring.KeyCode {
	if k, ok := desktopKeyMap[keyName]; ok {
		return k
	}
	return -1
}

func (ks *keyboardHandler) KeyUp(keyName fyne.KeyName) {
	if k := mapDesktopKey(keyName); k >= 0 {
		if k == bring.KeyLeftShift || k == bring.KeyRightShift {
			ks.shift = false
		}
		if k == bring.KeyCapsLock {
			ks.caps = false
		}
		ks.sendKey(k, false)
	}
}

func (ks *keyboardHandler) KeyDown(keyName fyne.KeyName) {
	if k := mapDesktopKey(keyName); k >= 0 {
		if k == bring.KeyLeftShift || k == bring.KeyRightShift {
			ks.shift = true
		}
		if k == bring.KeyCapsLock {
			ks.caps = true
		}
		ks.sendKey(k, true)
	}
}

func (ks *keyboardHandler) TypedKey(keyName fyne.KeyName) {
	k, ok := keyMap[keyName]
	if !ok && len(keyName) == 1 && keyName[0] < 128 {
		k = bring.KeyCode(keyName[0])
	}
	if k > 0 {
		ks.sendKey(k, true)
		ks.sendKey(k, false)
	}
}

var (
	keyMap        map[fyne.KeyName]bring.KeyCode
	desktopKeyMap map[fyne.KeyName]bring.KeyCode
)

func init() {
	desktopKeyMap = map[fyne.KeyName]bring.KeyCode{
		desktop.KeyAltLeft:      bring.KeyLeftAlt,
		desktop.KeyAltRight:     bring.KeyRightAlt,
		desktop.KeyControlLeft:  bring.KeyLeftControl,
		desktop.KeyControlRight: bring.KeyRightControl,
		desktop.KeyShiftLeft:    bring.KeyLeftShift,
		desktop.KeyShiftRight:   bring.KeyRightShift,
		desktop.KeySuperLeft:    bring.KeySuper,
		desktop.KeySuperRight:   bring.KeySuper,
		//:     bring.KeyCapsLock,
	}

	keyMap = map[fyne.KeyName]bring.KeyCode{
		fyne.KeyBackspace: bring.KeyBackspace,
		fyne.KeyDelete:    bring.KeyDelete,
		fyne.KeyDown:      bring.KeyDown,
		fyne.KeyEnd:       bring.KeyEnd,
		fyne.KeyEnter:     bring.KeyEnter,
		fyne.KeyReturn:    bring.KeyEnter,
		fyne.KeyEscape:    bring.KeyEscape,
		fyne.KeyF1:        bring.KeyF1,
		fyne.KeyF2:        bring.KeyF2,
		fyne.KeyF3:        bring.KeyF3,
		fyne.KeyF4:        bring.KeyF4,
		fyne.KeyF5:        bring.KeyF5,
		fyne.KeyF6:        bring.KeyF6,
		fyne.KeyF7:        bring.KeyF7,
		fyne.KeyF8:        bring.KeyF8,
		fyne.KeyF9:        bring.KeyF9,
		fyne.KeyF10:       bring.KeyF10,
		fyne.KeyF11:       bring.KeyF11,
		fyne.KeyF12:       bring.KeyF12,
		fyne.KeyHome:      bring.KeyHome,
		fyne.KeyInsert:    bring.KeyInsert,
		fyne.KeyLeft:      bring.KeyLeft,
		fyne.KeyPageDown:  bring.KeyPageDown,
		fyne.KeyPageUp:    bring.KeyPageUp,
		fyne.KeyRight:     bring.KeyRight,
		fyne.KeyTab:       bring.KeyTab,
		fyne.KeyUp:        bring.KeyUp,
		//:        bring.KeyPause,
		//:  bring.KeyPrintScreen,
		//:      bring.KeyNumLock,
		//:         bring.KeyMeta,
		//:          bring.KeyWin,
	}
}
