package draw

import (
	"image"
	"image/draw"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Drawing struct {
	img draw.Image
	gc  draw2d.GraphicContext
}

func (self *Drawing) Init(w int, h int) {
	self.img = image.NewRGBA(image.Rect(0, 0, w, h))
	self.gc = draw2dimg.NewGraphicContext(self.img)
	self.gc.BeginPath()
	self.gc.MoveTo(0, 0)
	self.gc.LineTo(float64(w), 0)
	self.gc.LineTo(float64(w), float64(h))
	self.gc.LineTo(0, float64(h))
	self.gc.Close()
	self.gc.Fill()
}

func (self *Drawing) Save() {
	draw2dimg.SaveToPngFile("output.png", self.img)
}
