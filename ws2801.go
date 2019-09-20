// Package ws2801 uses periph to access ws2801 LED strips via SPI
//
// This package uses the SPI interface provioded by periph to control a ws2801
// LED strip.
package ws2801

import (
	"errors"

	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

// The Pixels object is the interface to the LED strip.
type Pixels struct {
	port    spi.PortCloser
	conn    spi.Conn
	numleds int
	buffer  []byte
}

// NewPixels creates a new Pixels object.
//
// NumLEDs is the total number of LEDs in the strip.
func NewPixels(NumLEDs int) (*Pixels, error) {
	if NumLEDs <= 0 {
		return nil, errors.New("Illegal number of LEDs")
	}
	if _, err := host.Init(); err != nil {
		return nil, err
	}
	p := new(Pixels)

	port, err := spireg.Open("")
	if err != nil {
		return nil, err
	}
	p.port = port
	c, err := port.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		port.Close()
		return nil, err
	}
	p.conn = c
	p.numleds = NumLEDs
	p.buffer = make([]byte, NumLEDs*3)
	return p, nil
}

// Close closes the Pixels object and the SPI connection.
func (p *Pixels) Close() {
	p.port.Close()
}

// Show must be called after any change of pixel colors.
func (p *Pixels) Show() error {
	r := make([]byte, p.numleds*3)
	err := p.conn.Tx(p.buffer, r)
	return err
}

// Count returns the number of LEDs of the Pixels object.
func (p *Pixels) Count() int {
	return p.numleds
}

// SetPixel sets the color of a pixel.
//
// Position is the number of the pixel starting with 0.
// Red specifies the intensity of the red subpixel. It ranges between 0 and 255.
// Green specifies the intensity of the green subpixel. It ranges between 0 and 255.
// Blue specifies the intensity of the blue subpixel. It ranges between 0 and 255.
func (p *Pixels) SetPixel(Position int, Red byte, Green byte, Blue byte) error {
	if (Position < 0) || (Position > p.numleds) {
		return errors.New("Position out of range.")
	}
	p.buffer[Position*3] = Red
	p.buffer[Position*3+1] = Green
	p.buffer[Position*3+2] = Blue
	return nil
}

// ClearPixel sets the color of a pixel to black.
//
// Position is the number of the pixel starting with 0.
func (p *Pixels) ClearPixel(Position int) error {
	if (Position < 0) || (Position > p.numleds) {
		return errors.New("Position out of range.")
	}
	p.buffer[Position*3] = 0
	p.buffer[Position*3+1] = 0
	p.buffer[Position*3+2] = 0
	return nil
}

// SetPixels sets the color of a range of pixels.
//
// Start is the position of the first pixel to set.
// End is the position of the last pixel to set.
// Red specifies the intensity of the red subpixels. It ranges between 0 and 255.
// Green specifies the intensity of the green subpixels. It ranges between 0 and 255.
// Blue specifies the intensity of the blue subpixels. It ranges between 0 and 255.
func (p *Pixels) SetPixels(Start int, End int, Red byte, Green byte, Blue byte) error {
	if Start < 0 {
		return errors.New("Illegal start value")
	}
	if End >= p.numleds {
		return errors.New("Illegal end value")
	}
	if Start > End {
		return errors.New("Start must be smaller than End")
	}
	for i := Start; i <= End; i++ {
		p.buffer[i*3] = Red
		p.buffer[i*3+1] = Green
		p.buffer[i*3+2] = Blue
	}
	return nil
}

// ClearPixels sets the color of a range of pixels.
//
// Start is the position of the first pixel to set.
// End is the position of the last pixel to set.
// Red specifies the intensity of the red subpixels. It ranges between 0 and 255.
// Green specifies the intensity of the green subpixels. It ranges between 0 and 255.
// Blue specifies the intensity of the blue subpixels. It ranges between 0 and 255.
func (p *Pixels) ClearPixels(Start int, End int) error {
	if Start < 0 {
		return errors.New("Illegal start value")
	}
	if End >= p.numleds {
		return errors.New("Illegal end value")
	}
	if Start > End {
		return errors.New("Start must be smaller than End")
	}
	for i := Start * 3; i <= End*3+2; i++ {
		p.buffer[i] = 0
	}
	return nil
}
