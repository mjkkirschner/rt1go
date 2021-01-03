package core

type DiffuseLightMaterial struct {
	Albedo             Col3
	IntensityMultipler float64
}

func (mat *DiffuseLightMaterial) Scatter(rayIn *Ray, hitRecord *HitRecord, attenuation *Col3, scatteredRay *Ray) bool {
	return false
}

func (mat *DiffuseLightMaterial) Emit(u, v float64, pt *Pt3) Col3 {
	return mat.Albedo
}
