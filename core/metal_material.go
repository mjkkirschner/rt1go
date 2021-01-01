package core

type MetalMaterial struct {
	Albedo    Col3
	Roughness float64
}

func (mat *MetalMaterial) Scatter(rayIn *Ray, hitRecord *HitRecord, attenuation *Col3, scatteredRay *Ray) bool {
	norm := hitRecord.Normal
	rayDir := (Normalize(rayIn.Direction))

	reflectedDir := Reflect(&rayDir, &norm)
	*scatteredRay = NewRay(hitRecord.Hitpoint, reflectedDir.Add(RandomUnitSphereSample2().Scale(mat.Roughness)))
	*attenuation = *&mat.Albedo
	return Dot(scatteredRay.Direction, hitRecord.Normal) > 0
}
