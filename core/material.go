package core

type Material interface {
	Scatter(rayIn *Ray, hitRecord *HitRecord, attenuation *Col3, scatteredRay *Ray) bool
}

type DiffuseMaterial struct {
	Albedo Col3
}

func (mat *DiffuseMaterial) Scatter(rayIn *Ray, hitRecord *HitRecord, attenuation *Col3, scatteredRay *Ray) bool {
	norm := hitRecord.Normal

	scatterDirection := norm.Add(Normalize(GetRandomVectorInUnitSphere()))
	if scatterDirection.NearZero() {
		scatterDirection = norm
	}
	*scatteredRay = NewRay(hitRecord.Hitpoint, scatterDirection)
	*attenuation = *&mat.Albedo
	return true
}
