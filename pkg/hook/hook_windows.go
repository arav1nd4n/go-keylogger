package hook

import (
	"syscall"
	"unsafe"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx        = user32.NewProc("SetWindowsHookExW")
	callNextHookEx          = user32.NewProc("CallNextHookEx")
	unhookWindowsHookEx     = user32.NewProc("UnhookWindowsHookEx")
	getMessageW             = user32.NewProc("GetMessageW")
	getForegroundWindow     = user32.NewProc("GetForegroundWindow")
	getWindowTextW          = user32.NewProc("GetWindowTextW")
	getWindowTextLengthW    = user32.NewProc("GetWindowTextLengthW")
)

const (
	WH_KEYBOARD_LL = 13
	WM_KEYDOWN     = 0x0100
	WM_SYSKEYDOWN  = 0x0104
)

type (
	HOOKPROC  func(int, uintptr, uintptr) uintptr
	KBDLLHOOKSTRUCT struct {
		VkCode      uint32
		ScanCode    uint32
		Flags       uint32
		Time        uint32
		DwExtraInfo uintptr
	}
)

var (
	keyboardHook uintptr
	keyCallback  func(string, string)
)

func Start(callback func(string, string)) error {
	keyCallback = callback
	hookProc := syscall.NewCallback(keyboardHookProc)

	r, _, err := setWindowsHookEx.Call(
		WH_KEYBOARD_LL,
		hookProc,
		0,
		0,
	)
	if r == 0 {
		return err
	}
	keyboardHook = r

	go messageLoop()

	return nil
}

func Stop() {
	if keyboardHook != 0 {
		unhookWindowsHookEx.Call(keyboardHook)
		keyboardHook = 0
	}
}

func keyboardHookProc(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 && (wParam == WM_KEYDOWN || wParam == WM_SYSKEYDOWN) {
		kbdStruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		keyName := getKeyName(kbdStruct.VkCode)
		windowTitle := getActiveWindowTitle()
		keyCallback(keyName, windowTitle)
	}

	r, _, _ := callNextHookEx.Call(keyboardHook, uintptr(nCode), wParam, lParam)
	return r
}

func messageLoop() {
	var msg struct {
		HWnd    uintptr
		Message uint32
		WParam  uintptr
		LParam  uintptr
		Time    uint32
		Pt      struct{ X, Y int32 }
	}
	
	for {
		ret, _, _ := getMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if ret == 0 {
			break
		}
	}
}

func getKeyName(vkCode uint32) string {
	return "Key"
}

func getActiveWindowTitle() string {
	return "Window"
}
