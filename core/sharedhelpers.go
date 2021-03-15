package core

import (
	"fmt"
	"image/color"
	"math"
	"runtime"
)

func TestRayColor(r Ray, scene *[]Hittable, depth int, noHitFunc func(*Ray) Col3) Col3 {
	//base case
	if depth <= 0 {
		return Vec3{0, 0, 0}
	}
	hit := HitRecord{}
	if GetClosestHit(scene, &r, 0.001, math.Inf(1), &hit) {

		scattered := Ray{}
		attenuation := Col3{}
		emitted := Col3{}
		//man - this syntax... TODO - split up.
		//if we hit a light
		//TODO we may need to add a check that this only gets called from our direct light pass
		//to avoid a double count...
		if light, ok := hit.HitMaterial.(*DiffuseLightMaterial); ok {
			emitted = light.Emit(0, 0, &hit.Hitpoint).Scale(light.IntensityMultipler)
		}
		if (hit.HitMaterial).Scatter(&r, &hit, &attenuation, &scattered) {
			lights := getLights(scene)
			directLight := CalculateDirectLightingForAllLights(&hit, &lights, scene, noHitFunc)
			fullLight := directLight.Add(Multiply(attenuation, TestRayColor(scattered, scene, depth-1, noHitFunc)))

			return fullLight
		} else {
			return emitted
		}

	}
	return noHitFunc(&r)
}

func CalculateDirectLightingForAllLights(hit *HitRecord, lights *[]Hittable, scene *[]Hittable, noHitFunc func(*Ray) Col3) Col3 {
	//foreach light in our list of lights
	outputColor := Col3{0, 0, 0}
	for i := 0; i < len(*lights); i++ {
		//generate random point somewhere on the light
		//TODO make a light interface that might be able to give us a random point on the surface of the light.
		light := (*lights)[i]

		if sphere, ok := light.(*Sphere); ok {

			//instead of actually using this point - we want to use it as a starting point
			//but rather get the closest point along the ray formed by our shading point
			//and this sample point - this will achieve a much more sensical normal
			//if the light is a 3d surface.

			randomPointOnLight := (RandomUnitSphereSample2().Scale(sphere.Radius)).Add(sphere.Center)
			toLightVec := randomPointOnLight.Subtract(hit.Hitpoint)
			toLightVecNormalized := Normalize(toLightVec)

			//if the light vec is facing away from the normal of our hit point
			//then return emit of the hit... likely 0
			if Dot(toLightVecNormalized, hit.Normal) < 0 {
				outputColor = outputColor.Add(Col3{0, 0, 0})
				continue
			}

			//this ray is the ray from our surface to shade to the light - we need to cast it - to see if we
			//can actually make it to the light.
			scatteredRay := NewRay(hit.Hitpoint, toLightVecNormalized)
			tempRec := HitRecord{}
			if sphere.Hit(&scatteredRay, 0.001, math.Inf(1), &tempRec) {
				//now we create our real vector.
				toLightVec = tempRec.Hitpoint.Subtract(hit.Hitpoint)
				toLightVecNormalized = Normalize(toLightVec)
				distSquared := toLightVec.LengthSquared()
				//TODO this is weird- should be 4*pi*r^2
				lightArea := 1 * math.Pi * sphere.Radius * sphere.Radius

				normalLight := tempRec.Normal
				normalLightDotLightToHit := math.Abs(Dot(normalLight, toLightVecNormalized.Negate()))

				solidAngle := (normalLightDotLightToHit * lightArea) / distSquared

				gx0x1 := Dot(hit.Normal, toLightVecNormalized) *
					solidAngle

				val1 := TestRayColor(scatteredRay, scene, 1, noHitFunc)
				val2 := val1.Scale(gx0x1)

				outputColor = outputColor.Add(val2)
				continue
			}

		}
	}
	numLights := float64(len(*lights))
	//don't divide by 0
	if numLights < 1 {
		numLights = 1
	}
	return outputColor.Scale(1.0 / numLights)
}

//only making this sphere to get materials
func getLights(scene *[]Hittable) []Hittable {
	lights := make([]Hittable, 0)

	for i := 0; i < len(*scene); i++ {
		item := ((*scene)[i])
		if _, ok := item.GetMaterial().(*DiffuseLightMaterial); ok {
			lights = append(lights, item)
		}
	}
	return lights
}

func ConvertColor(color Col3, samples int) color.RGBA {
	intermediateCol := color.Scale(1.0 / float64(samples))
	R := Clamp(intermediateCol.X, 0, 1)
	G := Clamp(intermediateCol.Y, 0, 1)
	B := Clamp(intermediateCol.Z, 0, 1)
	final := Col3{
		math.Sqrt(R),
		math.Sqrt(G),
		math.Sqrt(B)}.Scale(255).ToRGBA()

	if final.B == 255 && final.R == 1 {
		fmt.Println(final)
		runtime.Breakpoint()
	}
	return final
}

func AddTrisToScene(scene *[]Hittable, mesh *Mesh, mat Material) {
	for _, face := range mesh.Faces {
		verts := [3]Vec3{}
		uvs := [3]Vec3{}
		norms := [3]Vec3{}
		verts[0] = mesh.Verts[face.VertIndicies[0]-1]
		verts[1] = mesh.Verts[face.VertIndicies[1]-1]
		verts[2] = mesh.Verts[face.VertIndicies[2]-1]

		uvs[0] = mesh.TexCoords[face.TexCoordIndicies[0]-1]
		uvs[1] = mesh.TexCoords[face.TexCoordIndicies[1]-1]
		uvs[2] = mesh.TexCoords[face.TexCoordIndicies[2]-1]

		norms[0] = mesh.Normals[face.NormalIndicies[0]-1]
		norms[1] = mesh.Normals[face.NormalIndicies[1]-1]
		norms[2] = mesh.Normals[face.NormalIndicies[2]-1]

		*scene = append(*scene, &Triangle{Verts: verts[:], UVs: uvs[:], Normals: norms[:], Material: mat})
	}
}
