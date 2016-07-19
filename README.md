# notifize

a desktop notification package for Go / golang

[![GoDoc](https://godoc.org/github.com/ctcpip/notifize?status.svg)](https://godoc.org/github.com/ctcpip/notifize)
![Development Status](https://img.shields.io/badge/development%20status-pre--alpha-red.svg)
[![Code Climate](https://img.shields.io/codeclimate/github/ctcpip/notifize.svg)](https://codeclimate.com/github/ctcpip/notifize)
[![Issue Count](https://img.shields.io/codeclimate/issues/github/ctcpip/notifize.svg)](https://codeclimate.com/github/ctcpip/notifize/issues)
[![License](https://img.shields.io/badge/license-GNU%20GPLv3-blue.svg)](./LICENSE)

## how it works

send desktop notifications from your Go application

## operating system support

* __GNU/Linux__ using _libnotify_ `notify-send`
* __OS X 10.9+__ using _AppleScript_ `display notification`

_I'd like to add support for other operating systems (primarily Windows and \*BSD)._
* _libnotify_ would probably work for \*BSD but I don't want to implement without testing.
* Windows doesn't seem to have an easy way of doing this, and I want to avoid doing something like bundling a 3rd party binary.  Please let me know of any ideas/suggestions to provide Windows notifications cleanly.

## installation

install and update this Go package with `go get -u github.com/ctcpip/notifize`

## usage example

~~~ go
package main

import "github.com/ctcpip/notifize"

func main() {
  notifize.Display("title", "hello, world", false, "")
}
~~~

## documentation / reference
[godoc.org/github.com/ctcpip/notifize](http://godoc.org/github.com/ctcpip/notifize)

### projects that use notifize

* [timezilla](http://github.com/ctcpip/timezilla)

### stuff nobody cares about

this project uses...

* [Go](http://golang.org)
* [semantic versioning](http://semver.org/)

### license

[GNU GPLv3](http://www.gnu.org/licenses/gpl-3.0.en.html)
