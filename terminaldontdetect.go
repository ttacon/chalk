// +build !smartmode

package chalk

// DetectTerminal turns off all coloring and styling if the current device
// is not a tty. This function is a no-op if chalk is built without the
// `smartmode` tag.
//
// NOTE: this has not beed tested for windows.
//
// The detection component of this function was taken, without hesitation, from
// golang.org/x/crypto/ssh/terminal.
func DetectTerminal(_ uintptr) {}
