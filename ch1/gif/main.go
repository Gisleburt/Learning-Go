package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"os"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
	usage      = "You must provide a frequency (recommend 0 < f <= 5)\n"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usage)
		return
	}
	freq, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dupe: %v\n", err)
		fmt.Fprintf(os.Stderr, usage)
		return
	}
	lissajous(os.Stdout, freq)
}

func lissajous(out io.Writer, freq float64) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 64
		nframes = 64
		delay   = 8
	)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size, 2*size)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size*0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
