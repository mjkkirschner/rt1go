package core

type Texture interface {
	Value(u float64, v float64, point *Pt3) Col3
}

type SolidTexture struct {
	color Col3
}

func (st *SolidTexture) Value(u float64, v float64, point *Pt3) Col3 {
	return st.color
}

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
	*attenuation = *&mat.Albedo //Vec3{hitRecord.U, hitRecord.V, 0} //*&mat.Albedo
	return true
}
