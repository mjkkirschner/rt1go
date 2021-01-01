package core

import (
	"math"
	"math/rand"
)

type RefractiveMaterial struct {
	RefractiveIndex float64
}

func (mat *RefractiveMaterial) Scatter(rayIn *Ray, hitRecord *HitRecord, attenuation *Col3, scatteredRay *Ray) bool {

	ri := mat.RefractiveIndex
	if hitRecord.FrontFacing {
		ri = 1.0 / ri
	}
	normalizedInputVector := Normalize(rayIn.Direction)

	cos_theta := math.Min(Dot(normalizedInputVector.Negate(), hitRecord.Normal), 1.0)
	sin_theta := math.Sqrt(1.0 - cos_theta*cos_theta)

	cannot_refract := ri*sin_theta > 1.0

	finalDirection := Vec3{}
	if cannot_refract || (reflectance(cos_theta, ri) > rand.Float64()) {
		//must reflect
		finalDirection = Reflect(&normalizedInputVector, &hitRecord.Normal)
	} else {
		//refract
		finalDirection = Refract(&normalizedInputVector, &hitRecord.Normal, ri)
	}

	*scatteredRay = NewRay(hitRecord.Hitpoint, finalDirection)
	*attenuation = Col3{1, 1, 1}
	return true
}

func reflectance(cos float64, refIndex float64) float64 {
	//Schlick's approximation
	r0 := (1 - refIndex) / (1 + refIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cos), 5)
}
