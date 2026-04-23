package geometry

type BoundingBox struct {
	MinX, MinY, MaxX, MaxY float64
}

func NewBoundingBox(minX, minY, maxX, maxY float64) BoundingBox {
	return BoundingBox{
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
	}
}

func (bbox BoundingBox) Intersects(other BoundingBox) bool {
	// fmt.Println("code Will be written")
	return true
}
