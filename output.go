package main

/*
 *
 * Colored text output
 * 
 * Only supports a few selected colors
 *
 */

import (
	"fmt"
	"os"
)

type Output struct {
	color   bool
	enabled bool
}

func NewOutput(color bool, enabled bool) *Output {
	var o *Output = new(Output)
	o.color = color
	o.enabled = enabled
	return o
}

func (o *Output) errText(msg string) {
	if o.enabled {
		fmt.Fprintf(os.Stderr, o.darkRedText(msg)+"\n")
	}
}

func (o *Output) Println(msg string) {
	if o.enabled {
		fmt.Println(msg)
	}
}

func (o *Output) isEnabled() bool {
	return o.enabled
}

func (o *Output) colorOn(num1 int, num2 int) string {
	if o.color {
		return fmt.Sprintf("\033[%d;%dm", num1, num2)
	}
	return ""
}

func (o *Output) colorOff() string {
	if o.color {
		return "\033[0m"
	}
	return ""
}

func (o *Output) darkRedText(s string) string {
	return o.colorOn(0, 31) + s + o.colorOff()
}

func (o *Output) lightGreenText(s string) string {
	return o.colorOn(1, 32) + s + o.colorOff()
}

func (o *Output) darkGreenText(s string) string {
	return o.colorOn(0, 32) + s + o.colorOff()
}

func (o *Output) lightYellowText(s string) string {
	return o.colorOn(1, 33) + s + o.colorOff()
}

func (o *Output) darkYellowText(s string) string {
	return o.colorOn(0, 33) + s + o.colorOff()
}

func (o *Output) lightBlueText(s string) string {
	return o.colorOn(1, 34) + s + o.colorOff()
}

func (o *Output) darkBlueText(s string) string {
	return o.colorOn(0, 34) + s + o.colorOff()
}

func (o *Output) lightCyanText(s string) string {
	return o.colorOn(1, 36) + s + o.colorOff()
}

func (o *Output) lightPurpleText(s string) string {
	return o.colorOn(1, 35) + s + o.colorOff()
}

func (o *Output) darkPurpleText(s string) string {
	return o.colorOn(0, 35) + s + o.colorOff()
}

func (o *Output) darkGrayText(s string) string {
	return o.colorOn(1, 30) + s + o.colorOff()
}
