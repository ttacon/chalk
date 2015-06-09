// +build smartmode
// +build darwin dragonfly freebsd linux,!appengine netbsd openbsd

package chalk

import (
	"os"
	"syscall"
	"unsafe"
)

// DetectTerminal turns off all coloring and styling if the current device
// is not a tty. This function is a no-op if chalk is built without the
// `smartmode` tag.
//
// NOTE: this has not beed tested for windows.
//
// The detection component of this function was taken, without hesitation, from
// golang.org/x/crypto/ssh/terminal.
func DetectTerminal() {

	// TODO(ttacon): fix syscall.TIOCGETA = ioctlReadTermios
	// on linux it's a different value
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, os.Stdout.Fd(), syscall.TIOCGETA, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)

	if err == 0 {
		return
	}

	Black = noColor
	Red = noColor
	Green = noColor
	Yellow = noColor
	Blue = noColor
	Magenta = noColor
	Cyan = noColor
	White = noColor
	ResetColor = noColor

	Bold = emptyTextStyle
	Dim = emptyTextStyle
	Italic = emptyTextStyle
	Underline = emptyTextStyle
	Inverse = emptyTextStyle
	Hidden = emptyTextStyle
	Strikethrough = emptyTextStyle
}
