//go:build !windows
// +build !windows

package terminal

import (
	"os"
)

func openTTY() (*os.File, error) {
	return os.Open("/dev/tty")
}
