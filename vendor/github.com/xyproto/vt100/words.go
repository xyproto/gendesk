package vt100

import (
	"strings"
)

func Words(line string, colors ...string) string {
	var ok bool
	words := strings.Split(line, " ")
	// Starting out with light gray, then using the last set color for the rest of the words
	color := LightGray
	coloredWords := make([]string, len(words))
	for i, word := range words {
		if i < len(colors) {
			prevColor := color
			color, ok = LightColorMap[colors[i]]
			if !ok {
				// Use the previous color if this color string was not found
				color = prevColor
			}
		}
		coloredWords[i] = color.Get(word)
	}
	return strings.Join(coloredWords, " ")
}
