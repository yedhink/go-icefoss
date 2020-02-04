package main
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{ color.Black ,
color.RGBA{0x10,0x0ef,0x10,0xff},
color.RGBA{0x09,0x0f6,0x09,0xff},
color.RGBA{0x00,0xff,0x00,0xff},
color.RGBA{0x01,0x0fe,0x01,0xff},
color.RGBA{0x02,0x0fd,0x02,0xff},
color.RGBA{0x03,0x0fc,0x03,0xff},
color.RGBA{0x04,0x0fb,0x04,0xff},
color.RGBA{0x05,0x0fa,0x05,0xff},
color.RGBA{0x06,0x0f9,0x06,0xff},
color.RGBA{0x07,0x0f8,0x07,0xff},
color.RGBA{0x08,0x0f7,0x08,0xff},
color.RGBA{0x09,0x0f6,0x09,0xff},
color.RGBA{0x0a,0x0f5,0x0a,0xff},
color.RGBA{0x0b,0x0f4,0x0b,0xff},
color.RGBA{0x0c,0x0f3,0x0c,0xff},
color.RGBA{0x0d,0x0f2,0x0d,0xff},
color.RGBA{0x0e,0x0f1,0x0e,0xff},
color.RGBA{0x0f,0x0f0,0x0f,0xff},
color.RGBA{0x10,0x0ef,0x10,0xff},
color.RGBA{0x11,0x0ee,0x11,0xff},
color.RGBA{0x12,0x0ed,0x12,0xff},
color.RGBA{0x13,0x0ec,0x13,0xff},
color.RGBA{0x14,0x0eb,0x14,0xff},
}

const (
	whiteIndex = 0 // First colour in palette
	blackIndex = 1 // Next colour in palette
)

func main () {
	lissajous (os.Stdout )
}

func lissajous ( out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 400
		nframes = 100
		delay = 8
	)

freq := rand.Float64() // Relative frequency of y oscillator
anim := gif.GIF{ LoopCount : nframes }
phase := 0.5 // phase difference
for i := 0; i < nframes ; i ++ { // 64 Animation frames
	rect := image.Rect(0 , 0 , 2* size +1 , 2* size +1)
	// Create a new black animation frame
	img := image.NewPaletted( rect , palette )
	// " cycles " times revolutions
	for t := 0.0; t < cycles *2* math.Pi ; t += res {
		x := math.Sin( t )
		y := math.Tan( t * freq + phase )
		img.SetColorIndex( size + int ( x * size +0.5) , size + int ( y * size +0.5) , blackIndex )
	}
	phase += 2.0
	anim.Delay = append(anim.Delay , delay )
	anim.Image = append(anim.Image , img )
}
gif . EncodeAll ( out , & anim )
}

