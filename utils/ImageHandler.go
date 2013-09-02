package utils

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

const (
	Center = iota
	LeftTop
	LeftDown
	RithtTop
	RightDown
)

func MakeMark(side int, path, savepath string) {
	imgb, _ := os.Open(path)
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	b := img.Bounds()
	wmb, _ := os.Open("/upload/watermark/watermark.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()
	a := watermark.Bounds()
	var offset image.Point
	switch side {
	case 0:
		offset = image.Pt(b.Dx()/2, b.Dy()/2)
	case 1:
		offset = image.Pt(0, 0)
	case 2:
		offset = image.Pt(0, b.Dx())
	case 3:
		offset = image.Pt(b.Dx()-a.Dx(), 0)
	case 4:
		offset = image.Pt(b.Dx()-a.Dx(), b.Dy()-a.Dy())
	default:
		offset = image.Pt(b.Dx()/2, b.Dy()/2)
	}
	m := image.NewRGBA(b)
	draw.Draw(m, b, img, image.ZP, draw.Over)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	imgw, _ := os.Create("water.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	defer imgw.Close()
}
