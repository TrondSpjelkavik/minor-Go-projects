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

var palette = []color.Color{color.RGBA{0xDD, 0xEE, 0xBB, 0xFF}, color.RGBA{0xCC, 0xBB, 0xBB, 0xDD} }

const (
	whiteIndex = 0
	blackIndex = 1
)


func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {

	const (
		cycles = 10
		res = 0.004
		size = 100
		nframes = 24
		delay = 8
	)

	freq := rand.Float64() * 4.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*4*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.8), size+int(y*size+0.8), blackIndex)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	} 

	gif.EncodeAll(out, &anim)

}

  
