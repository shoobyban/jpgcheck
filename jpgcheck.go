package main

import (
	"fmt"
	"image/jpeg"
	"os"
)

func E(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s Error: %v\n", os.Args[1], err.Error())
		os.Exit(-1)
	}
}

func main() {
	imgfile, err := os.Open(os.Args[1])
	E(err)
	defer imgfile.Close()
	img, err := jpeg.Decode(imgfile)
	E(err)
	ib := img.Bounds()
	x := ib.Dx()
	y := ib.Dy()
	px := img.At(x-10, y-10)
	r, g, b, a := px.RGBA()
	if r == 32896 && g == 32896 && b == 32896 && a == 65535 {
		fmt.Fprintf(os.Stderr, "%s Error: broken image (grey bottom)\n", os.Args[1])
	}
}
