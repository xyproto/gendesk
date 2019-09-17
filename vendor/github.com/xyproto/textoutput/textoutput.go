// Package textoutput offers a simple way to use vt100 and output colored text
package textoutput

import (
	"fmt"
	"github.com/xyproto/vt100"
	"os"
)

type TextOutput struct {
	color   bool
	enabled bool
}

func NewTextOutput(color, enabled bool) *TextOutput {
	return &TextOutput{color, enabled}
}

// Write an error message in red to stderr if output is enabled
func (o *TextOutput) Err(msg string) {
	if o.enabled {
		vt100.Red.Error(msg)
	}
}

// Write an error message to stderr and quit with exit code 1
func (o *TextOutput) ErrExit(msg string) {
	o.Err(msg)
	os.Exit(1)
}

// Checks if textual output is enabled
func (o *TextOutput) IsEnabled() bool {
	return o.enabled
}

func (o *TextOutput) DarkRed(s string) string {
	if o.color {
		return vt100.Red.Get(s)
	}
	return s
}

func (o *TextOutput) DarkGreen(s string) string {
	if o.color {
		return vt100.Green.Get(s)
	}
	return s
}

func (o *TextOutput) DarkYellow(s string) string {
	if o.color {
		return vt100.Yellow.Get(s)
	}
	return s
}

func (o *TextOutput) DarkBlue(s string) string {
	if o.color {
		return vt100.Blue.Get(s)
	}
	return s
}

func (o *TextOutput) DarkPurple(s string) string {
	if o.color {
		return vt100.Magenta.Get(s)
	}
	return s
}

func (o *TextOutput) DarkCyan(s string) string {
	if o.color {
		return vt100.Cyan.Get(s)
	}
	return s
}

func (o *TextOutput) DarkGray(s string) string {
	if o.color {
		return vt100.DarkGray.Get(s)
	}
	return s
}

func (o *TextOutput) LightRed(s string) string {
	if o.color {
		return vt100.LightRed.Get(s)
	}
	return s
}

func (o *TextOutput) LightGreen(s string) string {
	if o.color {
		return vt100.LightGreen.Get(s)
	}
	return s
}

func (o *TextOutput) LightYellow(s string) string {
	if o.color {
		return vt100.LightYellow.Get(s)
	}
	return s
}

func (o *TextOutput) LightBlue(s string) string {
	if o.color {
		return vt100.LightBlue.Get(s)
	}
	return s
}

func (o *TextOutput) LightPurple(s string) string {
	if o.color {
		return vt100.LightMagenta.Get(s)
	}
	return s
}

func (o *TextOutput) LightCyan(s string) string {
	if o.color {
		return vt100.LightCyan.Get(s)
	}
	return s
}

func (o *TextOutput) White(s string) string {
	if o.color {
		return vt100.White.Get(s)
	}
	return s
}

// Given a line with words and several color strings, color the words
// in the order of the colors. The last color will color the rest of the
// words.
func (o *TextOutput) Words(line string, colors ...string) string {
	if o.color {
		return vt100.Words(line, colors...)
	}
	return line
}

// Write a message to stdout if output is enabled
func (o *TextOutput) Println(msg ...interface{}) {
	if o.enabled {
		fmt.Println(msg...)
	}
}

// Change the color state in the terminal emulator
func (o *TextOutput) ColorOn(attribute1, attribute2 int) string {
	if !o.color {
		return ""
	}
	return fmt.Sprintf("\033[%d;%dm", attribute1, attribute2)
}

// Change the color state in the terminal emulator
func (o *TextOutput) ColorOff() string {
	if !o.color {
		return ""
	}
	return "\033[0m"
}
