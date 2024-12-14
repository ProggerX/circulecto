package draw

import (
	"image/color"
	"math"
)

func (dr *Drawing) DrawLine(x1, y1, x2, y2 float64) {
	gc := dr.gc
	gc.SetStrokeColor(color.RGBA{0, 0, 0, 255})
	gc.SetLineWidth(3)
	gc.MoveTo(x1, y1)
	gc.LineTo(x2, y2)
	gc.Stroke()
}

func (dr *Drawing) DrawCorrectLine(x1, y1, x2, y2 float64) {
	if math.Abs(x2-x1) > math.Abs(y2-y1) {
		dr.DrawLine(x1, y1, x1+math.Abs(x2-x1)/2, y1)
		dr.DrawLine(x1+math.Abs(x2-x1)/2, y1, x1+math.Abs(x2-x1)/2, y2)
		dr.DrawLine(x1+math.Abs(x2-x1)/2, y2, x2, y2)
	} else {
		dr.DrawLine(x1, y1, x1, y1+math.Abs(y2-y1)/2)
		dr.DrawLine(x1, y1+math.Abs(y2-y1)/2, x2, y1+math.Abs(y2-y1)/2)
		dr.DrawLine(x2, y1+math.Abs(y2-y1)/2, x2, y2)
	}
}
