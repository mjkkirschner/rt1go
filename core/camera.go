package core

import (
	"math"
)

type Camera struct {
	ViewHeight      float64
	ViewWidth       float64
	FocalLength     float64
	origin          Pt3
	horizontal      Vec3
	vertical        Vec3
	lowerLeftCorner Pt3
	lensRadius      float64
	u               Vec3
	v               Vec3
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func NewCameraByPoints(lookFrom Pt3, lookAt Pt3, vUp Vec3, fovDegrees float64, aspectRatio float64, aperture float64, focus_len float64) Camera {
	theta := degreesToRadians(fovDegrees)
	h := math.Tan(theta / 2.0)
	viewPortH := 2.0 * h
	viewPortW := aspectRatio * viewPortH
	focalLength := focus_len

	w := Normalize(lookFrom.Subtract(lookAt))
	u := Normalize(Cross(vUp, w))
	v := Cross(w, u)
	origin := lookFrom

	cam := Camera{
		viewPortH,
		viewPortW,
		focalLength,
		origin,
		u.Scale(viewPortW * focalLength),
		v.Scale(viewPortH * focalLength),
		Vec3{0, 0, 0},
		aperture / 2.0,
		u,
		v,
	}
	lowerLeftCorner := cam.origin.Subtract((cam.horizontal.Scale(.5))).Subtract(cam.vertical.Scale(.5)).Subtract(w.Scale(focalLength))
	cam.lowerLeftCorner = lowerLeftCorner

	return cam
}

func (cam *Camera) GetRay(u float64, v float64) Ray {
	var rd = randomVecInUnitDisk().Scale(cam.lensRadius)
	var offset = cam.u.Scale(rd.X).Add(cam.v.Scale(rd.Y))

	return NewRay(cam.origin.Add(offset), cam.lowerLeftCorner.Add(cam.horizontal.Scale(u)).Add(cam.vertical.Scale(v)).Subtract(cam.origin).Subtract(offset))
}

func randomVecInUnitDisk() Vec3 {
	for true {
		var p = RandomVectorByRange(-1.0, 1.0)
		p.Z = 0
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
	return Vec3{}
}
