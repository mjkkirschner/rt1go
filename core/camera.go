package core

import "math"

type Camera struct {
	ViewHeight      float64
	ViewWidth       float64
	FocalLength     float64
	origin          Pt3
	horizontal      Vec3
	vertical        Vec3
	lowerLeftCorner Pt3
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func NewCameraByPoints(lookFrom Pt3, lookAt Pt3, vUp Vec3, fovDegrees float64, aspectRatio float64) Camera {
	theta := degreesToRadians(fovDegrees)
	h := math.Tan(theta / 2.0)
	viewPortH := 2.0 * h
	viewPortW := aspectRatio * viewPortH
	focalLength := 1.0

	w := Normalize(lookFrom.Subtract(lookAt))
	u := Normalize(Cross(vUp, w))
	v := Cross(w, u)
	origin := lookFrom

	cam := Camera{
		viewPortH,
		viewPortW,
		focalLength,
		origin,
		u.Scale(viewPortW),
		v.Scale(viewPortH),
		Vec3{0, 0, 0},
	}
	lowerLeftCorner := cam.origin.Subtract((cam.horizontal.Scale(.5))).Subtract(cam.vertical.Scale(.5)).Subtract(w)
	cam.lowerLeftCorner = lowerLeftCorner

	return cam
}

func NewCamera(viewHeight float64, viewWidth float64, focalLength float64, origin Pt3) Camera {

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
