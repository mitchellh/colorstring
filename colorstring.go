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
	// Colors maps a color string to the integer code for that color.
	Colors map[string]int

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
		if len(m) > 0 {
			result.WriteString(v[m[1]:nm[0]])
		}
		m = nm

		var replace string
		if code, ok := c.Colors[v[m[0]+1:m[1]-1]]; ok {
			colored = true
			replace = fmt.Sprintf("\033[0;%dm", code)
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
var DefaultColors map[string]int

func init() {
	DefaultColors = map[string]int{
		"blue": 34,
	}

	def = Colorize{
		Colors: DefaultColors,
		Reset:  true,
	}
}

var def Colorize
var parseRe = regexp.MustCompile(`(?i)\[[a-z0-9_-]+\]`)
