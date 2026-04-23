package geometry

type Shape interface {
	Area() float64
	Perimeter() float64
	BoundingBox() BoundingBox //BoundingBox() returns a value, not pointer
	Intersects(other Shape) (bool, error)
	//Intersects is shape-aware (actual geometry later)
	//Errors are allowed here, not in BoundingBox
}

//Point - atomic geometry unit
//BoundingBox - fast spatial approximation
//Shape - polymorphic geometry contract
