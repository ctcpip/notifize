/*

notifize - a desktop notification package for Go / golang
Copyright (C) 2016  Chris de Almeida

http://github.com/ctcpip/notifize

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

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
