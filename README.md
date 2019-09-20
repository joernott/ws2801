[![GoDoc](https://godoc.org/github.com/joernott/ws2801?status.svg)](https://godoc.org/github.com/joernott/ws2801) [![license](https://img.shields.io/badge/license-BSD%203--clause-blue.svg)](https://github.com/joernott/ws2801/LICENSE) [![cover.run](https://cover.run/go/github.com/joernott/ws2801.svg?style=flat&tag=golang-1.13)](https://cover.run/go?tag=golang-1.13&repo=github.com%2Fjoernott%2Fws2801)
# ws2801 - a library to address ws2801 led strips connected via SPI

This package controls ws2801 LED strips connected via SPI.

The library has been tested on a raspberry pi zero w with two different ws2801
LED strips.

## License
BSD 3-clause license

## Contributions
Contributions / Pull requests are welcome. 

## Documentation
[https://godoc.org/github.com/joernott/ws2801](https://godoc.org/github.com/joernott/ws2801)

## Usage

### Create connection
The LED strip connection is created by using the NewPixels function.
```
	p,err := ws2801.NewPixels(
		32						// number of pixels in the strip
	)
```
### Set a pixel or a range of pixels to a specific color
```
	err := p.SetPixel(0, 255, 0, 0)
	err := p.SetPixels(2, 4, 0, 255, 0)
	err := p.SetPixels(6, 8, 0, 0, 255)
```

The Show function actually transmits the data to the strip. It must be called
whenever you want to show your changes:
```
	err := p.Show()
```

### Clear a pixel or a range of pixels
```
	err := p.ClearPixel(0)
	err := p.ClearPixels(2, 8)
	err := p.Show()
```

## Use the ws2801tool
The ws2801tool is a small commandline tool to set/clear pixels from the command
line.

### Build
```
go get -v github.com/joernott/ws2801tool
```

### Usage
ws2801tool [command]

Available Commands:
  clear       Clears a pixel or a range of pixels
  help        Help about any command
  set         Sets one or more pixels to a specific color

Flags:
  -h, --help                 Help for ws2801tool
  -p, --position int         Position of the pixel.
  -f, --first int            Position of the first pixel when defining a range.
  -l, --last int             Position of the last pixel when defining a range.
  -n, --number_of_leds int   Number of LEDs
  -r, --red uint8            Intensity of the red subpixel
  -g, --green uint8          Intensity of the green subpixel
  -b, --blue uint8           Intensity of the blue subpixel

Use "ws2801tool [command] --help" for more information about a command.
subcommand is required
