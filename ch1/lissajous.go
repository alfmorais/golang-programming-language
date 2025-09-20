// Lissajous gera animacões GIF de figuras Lissajous aleatórias.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var pallete = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // primeira cor da paleta
	blackIndex = 1 // próxima cor da paleta
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	lissajous(os.Stdout, r)
}

func lissajous(out io.Writer, r *rand.Rand) {
	const (
		cycles  = 5     // número de revoluções completas da oscilador x
		res     = 0.001 // resolução angular
		size    = 100   // canvas de imagem cobre de [-size..+size]
		nframes = 64    // número de quadros da animação
		delay   = 8     // atraso entre quadros em unidades de 10ms
	)

	freq := r.Float64() * 3.0 // frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferença de fase

	for index := 0; index < nframes; index++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)

		for k := range img.Pix {
			img.Pix[k] = whiteIndex
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
