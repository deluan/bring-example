package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
	"github.com/deluan/bring"
)

// Handles keyboard events mapping between Bring and Fyne
type keyboardHandler struct {
	display *BringDisplay
	shift   bool
	caps    bool
}

func (ks *keyboardHandler) TypedKey(ev *fyne.KeyEvent) {
	keyName := ev.Name
	k, ok := keyMap[keyName]
	if !ok && len(keyName) == 1 && keyName[0] < 128 {
		k = bring.KeyCode(keyName[0])
	}
	if k > 0 {
		ks.sendKey(k, true)
		ks.sendKey(k, false)
	}
}

func (ks *keyboardHandler) KeyDown(ev *fyne.KeyEvent) {
	ks.handleDesktopKey(ev.Name, true)
}

func (ks *keyboardHandler) KeyUp(ev *fyne.KeyEvent) {
	ks.handleDesktopKey(ev.Name, false)
}

func (ks *keyboardHandler) handleDesktopKey(keyName fyne.KeyName, pressed bool) {
	if keyCode, ok := desktopKeyMap[keyName]; ok {
		if keyCode == bring.KeyLeftShift || keyCode == bring.KeyRightShift {
			ks.shift = pressed
		}
		if keyCode == bring.KeyCapsLock {
			ks.caps = pressed
		}
		ks.sendKey(keyCode, pressed)
	}
}

func (ks *keyboardHandler) sendKey(key bring.KeyCode, pressed bool) {
	if key >= 'A' && key <= 'Z' {
		if !ks.shift && !ks.caps {
			key = key + 32
		}
	}
	_ = ks.display.Client.SendKey(key, pressed)
}

func (ks *keyboardHandler) Focused() bool {
	return true
}

func (ks *keyboardHandler) FocusGained() {
}

func (ks *keyboardHandler) FocusLost() {
}

func (ks *keyboardHandler) TypedRune(ch rune) {
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
		fyne.KeySpace:     bring.KeyCode(32),
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

// Make sure all necessary interfaces are implemented
var _ desktop.Keyable = (*keyboardHandler)(nil)
