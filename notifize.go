// Package notifize - a desktop notification package for Go / golang
package notifize

import (
	"os/exec"
	"runtime"
)

// Display a notification
func Display(summary string, body string, isUrgent bool, iconPath string) {

	switch runtime.GOOS {

	case "darwin":

		exec.Command("osascript", "-e", "display notification \""+body+"\" with title \""+summary+"\"").Run()

	case "linux":

		if isUrgent {
			exec.Command("notify-send", "-i", iconPath, summary, body, "-u", "critical").Run()
		} else {
			exec.Command("notify-send", "-i", iconPath, summary, body).Run()
		}

	case "android", "dragonfly", "freebsd", "nacl", "netbsd", "openbsd", "plan9", "solaris", "windows":
		// not implemented

	}

}
