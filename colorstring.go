// colorstring provides functions for colorizing strings for terminal
// output.
package colorstring

import (
	"bytes"
	"fmt"
	"regexp"
)

// Color colorizes your strings using the default settings.
//
// If you want to customize, use the Colorize struct.
func Color(v string) string {
	return def.Color(v)
}

// Colorize colorizes your strings, giving you the ability to customize
// some of the colorization process.
//
// The options in Colorize can be set to customize colorization. If you're
// only interested in the defaults, just use the top Color function directly,
// which creates a default Colorize.
type Colorize struct {
	// Colors maps a color string to the code for that color. The code
	// is a string so that you can use more complex colors to set foreground,
	// background, attributes, etc. For example, "boldblue" might be
	// "1;34"
	Colors map[string]string

	// If true, color attributes will be ignored. This is useful if you're
	// outputting to a location that doesn't support colors and you just
	// want the strings returned.
	Disable bool

	// Reset, if true, will reset the color after each colorization by
	// adding a reset code at the end.
	Reset bool
}

func (c *Colorize) Color(v string) string {
	matches := parseRe.FindAllStringIndex(v, -1)
	if len(matches) == 0 {
		return v
	}

	result := new(bytes.Buffer)
	if matches[0][0] > 0 {
		if _, err := result.WriteString(v[:matches[0][0]]); err != nil {
			panic(err)
		}
	}

	colored := false
	var m []int
	for _, nm := range matches {
		// Write the text in between this match and the last
		if len(m) > 0 {
			result.WriteString(v[m[1]:nm[0]])
		}
		m = nm

		// If we're disabled, just ignore the color code information
		if c.Disable {
			continue
		}

		var replace string
		if code, ok := c.Colors[v[m[0]+1:m[1]-1]]; ok {
			colored = true
			replace = fmt.Sprintf("\033[%sm", code)
		} else {
			replace = v[m[0]:m[1]]
		}

		result.WriteString(replace)
	}
	result.WriteString(v[m[1]:])

	if colored && c.Reset {
		// Write the clear byte at the end
		result.WriteString("\033[0m")
	}

	return result.String()
}

// DefaultColors are the default colors used when colorizing.
var DefaultColors map[string]string

func init() {
	DefaultColors = map[string]string{
		"blue": "34",
	}

	def = Colorize{
		Colors: DefaultColors,
		Reset:  true,
	}
}

var def Colorize
var parseRe = regexp.MustCompile(`(?i)\[[a-z0-9_-]+\]`)
