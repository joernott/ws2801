package ws2801

import (
	"testing"
)

func TestNewPixels(t *testing.T) {
	var err error
	var p *Pixels

	p, err = NewPixels(-1)
	if err == nil {
		t.Fatalf("Could create a LED strip with -1 LEDs")
	}

	p, err = NewPixels(1)
	if err != nil {
		t.Fatalf("Could not create LED strip")
	}
}

func TestSetPixel(t *testing.T) {
	var err error
	var p *Pixels

	p, err = NewPixels(1)
	if err != nil {
		t.Fatalf("Could not create LED strip")
	}

	p.SetPixel(-1, 255, 255, 255)
	if err == nil {
		t.Fatalf("Could change the color of LED -1")
	}

	p.SetPixel(1, 255, 255, 255)
	if err == nil {
		t.Fatalf("Could change the color of LED = NumLEDs")
	}

	p.SetPixel(0, 255, 255, 255)
	if err != nil {
		t.Fatalf("Could not set the color of LED 0")
	}
}

func TestClearPixel(t *testing.T) {
	var err error
	var p *Pixels

	p, err = NewPixels(1)
	if err != nil {
		t.Fatalf("Could not create LED strip")
	}

	p.ClearPixel(-1)
	if err == nil {
		t.Fatalf("Could clear LED -1")
	}

	p.ClearPixel(1)
	if err == nil {
		t.Fatalf("Could clear LED = NumLEDs")
	}

	p.ClearPixel(0)
	if err != nil {
		t.Fatalf("Could not clear LED 0")
	}
}

func TestSetPixels(t *testing.T) {
	var err error
	var p *Pixels

	p, err = NewPixels(4)
	if err != nil {
		t.Fatalf("Could not create LED strip")
	}

	p.SetPixels(-1, 1, 255, 255, 255)
	if err == nil {
		t.Fatalf("Could use illegal start LED -1")
	}

	p.SetPixels(0, 4, 255, 255, 255)
	if err == nil {
		t.Fatalf("Could use end LED = NumLEDs")
	}

	p.SetPixels(3, 0, 255, 255, 255)
	if err == nil {
		t.Fatalf("Could use end > start")
	}

	p.SetPixels(0, 3, 255, 255, 255)
	if err != nil {
		t.Fatalf("Could not set the color of LEDs 0 to 3")
	}
}

func TestClearPixels(t *testing.T) {
	var err error
	var p *Pixels

	p, err = NewPixels(4)
	if err != nil {
		t.Fatalf("Could not create LED strip")
	}

	p.ClearPixels(-1, 1)
	if err == nil {
		t.Fatalf("Could clear illegal start LED -1")
	}

	p.ClearPixels(0, 4)
	if err == nil {
		t.Fatalf("Could clear end LED = NumLEDs")
	}

	p.ClearPixels(3, 0)
	if err == nil {
		t.Fatalf("Could clear end <")
	}

	p.ClearPixels(0, 3)
	if err != nil {
		t.Fatalf("Could not clear LED 0")
	}
}

func TestCount(t *testing.T) {
	var err error
	var p *Pixels

	p, err = NewPixels(4)
	if err != nil {
		t.Fatalf("Could not create LED strip")
	}

	i := p.Count()
	if i != 4 {
		t.Fatalf("Number of LEDs differ from creation")
	}
}
