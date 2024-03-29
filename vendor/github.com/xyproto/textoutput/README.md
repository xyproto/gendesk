# TextOutput

[![GoDoc](https://godoc.org/github.com/xyproto/textoutput?status.svg)](https://godoc.org/github.com/xyproto/textoutput) [![License](https://img.shields.io/badge/license-BSD-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/textoutput/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/textoutput)](https://goreportcard.com/report/github.com/xyproto/textoutput)

Package for controlling text output, with or without colors, using VT100 terminal codes.

## Example use

```go
package main

import (
    "github.com/xyproto/textoutput"
)

func main() {
    // Prepare to output colored text, but respect the NO_COLOR environment variable
    o := textoutput.New()

    o.Print("<blue>hi</blue> ")
    o.Println("<yellow>there</yellow>")
    o.Printf("<green>%s: <red>%d<off>\n", "number", 42)
}
```

## General info

* Version: 1.14.2
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
