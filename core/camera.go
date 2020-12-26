package core

type Camera struct {
	ViewHeight      float64
	ViewWidth       float64
	FocalLength     int
	origin          Pt3
	horizontal      Vec3
	vertical        Vec3
	lowerLeftCorner Pt3
}

func NewCamera(viewHeight float64, viewWidth float64, focalLength int, origin Pt3) Camera {

	cam := Camera{
		viewHeight,
		viewWidth,
		focalLength,
		origin,
		Vec3{viewWidth, 0, 0},
		Vec3{0, viewHeight, 0},
		Vec3{0, 0, 0}}

	lowerLeftCorner := cam.origin.Subtract((cam.horizontal.Scale(.5))).Subtract(cam.vertical.Scale(.5)).Subtract(NewVector3(0, 0, float64(cam.FocalLength)))
	cam.lowerLeftCorner = lowerLeftCorner
	return cam
}

func (cam *Camera) GetRay(u float64, v float64) Ray {
	return NewRay(cam.origin, cam.lowerLeftCorner.Add(cam.horizontal.Scale(u)).Add(cam.vertical.Scale(v)).Subtract(cam.origin))
}
