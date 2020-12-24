package core

type Camera struct {
	ViewHeight  float64
	ViewWidth   float64
	FocalLength int
	Origin      Pt3
	Horizontal  Vec3
	Vertical    Vec3
}

func NewCamera(viewHeight float64, viewWidth float64, focalLength int, origin Pt3) Camera {
	return Camera{
		viewHeight,
		viewWidth,
		focalLength,
		origin,
		Vec3{viewWidth, 0, 0},
		Vec3{0, viewHeight, 0}}
}
