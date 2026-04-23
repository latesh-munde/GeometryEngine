package errors

import "errors"

var (
	ErrInvalidDimension = errors.New("invalid dimensnion: must be > 0") //dimension : radisu, length, etc
	ErrLine             = errors.New("Start and end points should be diffrent")
	ErrInvalidPolygon   = errors.New("invalid polygon: Points to make polygon must be atleast 3")
)
