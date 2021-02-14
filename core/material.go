package core

import (
	"image"
	"os"
)

type Texture interface {
	Value(u float64, v float64, point *Pt3) Col3
}

type SolidTexture struct {
	Color Col3
}

type ImageTexture struct {
	internalimg image.Image
}

func NewImageTexture(path string) *ImageTexture {
	var reader, _ = os.Open(path)
	var image, _, _ = image.Decode(reader)
	return &ImageTexture{image}
}

func (it *ImageTexture) Value(u float64, v float64, point *Pt3) Col3 {
	var u1 = Clamp(u, 0.0, 1.0)
	var v1 = 1.0 - Clamp(v, 0.0, 1.0)
	var u2 = int(u1 * float64(it.internalimg.Bounds().Dx()))
	var v2 = int(v1 * float64(it.internalimg.Bounds().Dy()))
	var r, g, b, _ = it.internalimg.At(u2, v2).RGBA()

	return Col3{float64(uint8(r)) / 255, float64(uint8(g)) / 255, float64(uint8(b)) / 255}
}

func (st *SolidTexture) Value(u float64, v float64, point *Pt3) Col3 {
	return st.Color
}

type Material interface {
	Scatter(rayIn *Ray, hitRecord *HitRecord, attenuation *Col3, scatteredRay *Ray) bool
}

type DiffuseMaterial struct {
	Texture Texture
}

func (mat *DiffuseMaterial) Scatter(rayIn *Ray, hitRecord *HitRecord, attenuation *Col3, scatteredRay *Ray) bool {
	norm := hitRecord.Normal

	scatterDirection := norm.Add(Normalize(GetRandomVectorInUnitSphere()))
	if scatterDirection.NearZero() {
		scatterDirection = norm
	}
	*scatteredRay = NewRay(hitRecord.Hitpoint, scatterDirection)
	*attenuation = mat.Texture.Value(hitRecord.U, hitRecord.V, &hitRecord.Hitpoint)
	return true
}
