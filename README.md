# colorstring [![Build Status](https://travis-ci.org/mitchellh/colorstring.svg)](https://travis-ci.org/mitchellh/colorstring)

colorstring is a [Go](http://www.golang.org) library for outputting colored
strings to a console using a simple inline syntax in your string to specify
the color to print as.

For example, the string `[blue]hello [red]world` would output the text
"hello world" in two colors. The API of colorstring allows for easily disabling
colors, adding aliases, etc.

## Installation

Standard `go get`:

```
$ go get github.com/mitchellh/colorstring
```

## Usage & Example

For usage and examples see the [Godoc](http://godoc.org/github.com/mitchellh/colorstring).

Usage is easy enough:

```go
fmt.Println(colorstring.Color("[blue]Hello [red]World!"))
```

Additionally, the `Colorize` struct can be used to set options such as
custom colors, color disabling, etc.

## 256 color terminal support

Many modern terminals support 256 colors, so this library offers this with two special kinds of colors:

* `[rgb]` where each letter is a number 0 to 5; 0 being darkest and 5 being brightest.
* `[grayN]` where N is a number 0 to 15; 0 being darkest (black) and 15 being brightest.
* Background colors are also supported like the rest of the library `[_rgb_]`
  would be background color for that rgb level.
