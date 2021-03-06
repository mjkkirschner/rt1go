package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"rt1go/core"
	"runtime"
	"sync"
	"time"
)

func calculateDirectLightingForAllLights(hit *core.HitRecord, lights *[]core.Hittable, scene *[]core.Hittable) core.Col3 {
	//foreach light in our list of lights
	outputColor := core.Col3{0, 0, 0}
	for i := 0; i < len(*lights); i++ {
		//generate random point somewhere on the light
		//TODO make a light interface that might be able to give us a random point on the surface of the light.
		light := (*lights)[i]

		if sphere, ok := light.(*core.Sphere); ok {

			//instead of actually using this point - we want to use it as a starting point
			//but rather get the closest point along the ray formed by our shading point
			//and this sample point - this will achieve a much more sensical normal
			//if the light is a 3d surface.

			randomPointOnLight := (core.RandomUnitSphereSample2().Scale(sphere.Radius)).Add(sphere.Center)
			toLightVec := randomPointOnLight.Subtract(hit.Hitpoint)
			toLightVecNormalized := core.Normalize(toLightVec)

			//if the light vec is facing away from the normal of our hit point
			//then return emit of the hit... likely 0
			if core.Dot(toLightVecNormalized, hit.Normal) < 0 {
				outputColor = outputColor.Add(core.Col3{0, 0, 0})
				continue
			}

			//this ray is the ray from our surface to shade to the light - we need to cast it - to see if we
			//can actually make it to the light.
			scatteredRay := core.NewRay(hit.Hitpoint, toLightVecNormalized)
			tempRec := core.HitRecord{}
			if sphere.Hit(&scatteredRay, 0.001, math.Inf(1), &tempRec) {
				//now we create our real vector.
				toLightVec = tempRec.Hitpoint.Subtract(hit.Hitpoint)
				toLightVecNormalized = core.Normalize(toLightVec)
				distSquared := toLightVec.LengthSquared()
				//TODO this is weird- should be 4*pi*r^2
				lightArea := 1 * math.Pi * sphere.Radius * sphere.Radius

				normalLight := tempRec.Normal
				normalLightDotLightToHit := math.Abs(core.Dot(normalLight, toLightVecNormalized.Negate()))

				solidAngle := (normalLightDotLightToHit * lightArea) / distSquared

				gx0x1 := core.Dot(hit.Normal, toLightVecNormalized) *
					solidAngle

				val1 := testRayColor(scatteredRay, scene, 1)
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

func testRayColor(r core.Ray, scene *[]core.Hittable, depth int) core.Col3 {
	//base case
	if depth <= 0 {
		return core.Vec3{0, 0, 0}
	}
	hit := core.HitRecord{}
	if core.GetClosestHit(scene, &r, 0.001, math.Inf(1), &hit) {

		scattered := core.Ray{}
		attenuation := core.Col3{}
		emitted := core.Col3{}
		//man - this syntax... TODO - split up.
		//if we hit a light
		//TODO we may need to add a check that this only gets called from our direct light pass
		//to avoid a double count...
		if light, ok := hit.HitMaterial.(*core.DiffuseLightMaterial); ok {
			emitted = light.Emit(0, 0, &hit.Hitpoint).Scale(light.IntensityMultipler)
		}
		if (hit.HitMaterial).Scatter(&r, &hit, &attenuation, &scattered) {
			lights := getLights(scene)
			directLight := calculateDirectLightingForAllLights(&hit, &lights, scene)
			fullLight := directLight.Add(core.Multiply(attenuation, testRayColor(scattered, scene, depth-1)))

			return fullLight
		} else {
			return emitted
		}

	}
	return testRayNoHitColor(&r)
}

//only making this sphere to get materials
func getLights(scene *[]core.Hittable) []core.Hittable {
	lights := make([]core.Hittable, 0)

	for i := 0; i < len(*scene); i++ {
		item := ((*scene)[i])
		if _, ok := item.GetMaterial().(*core.DiffuseLightMaterial); ok {
			lights = append(lights, item)
		}
	}
	return lights
}

func testRayNoHitColor(ray *core.Ray) core.Col3 {
	//return core.Vec3{}
	t := (core.Normalize(ray.Direction).Y + 1.0) * .5
	return (core.Col3{1.0, 1.0, 1.0}.Add(core.Col3{.5, .7, 1.0}.Scale(t))).Scale(1.0 - t)
}

func ConvertColor(color core.Col3, samples int) color.RGBA {
	intermediateCol := color.Scale(1.0 / float64(samples))
	R := core.Clamp(intermediateCol.X, 0, 1)
	G := core.Clamp(intermediateCol.Y, 0, 1)
	B := core.Clamp(intermediateCol.Z, 0, 1)
	final := core.Col3{
		math.Sqrt(R),
		math.Sqrt(G),
		math.Sqrt(B)}.Scale(255).ToRGBA()

	if final.B == 255 && final.R == 1 {
		fmt.Println(final)
		runtime.Breakpoint()
	}
	return final
}

func AddTrisToScene(scene *[]core.Hittable, mesh *core.Mesh, mat core.Material) {
	for _, face := range mesh.Faces {
		verts := [3]core.Vec3{}
		uvs := [3]core.Vec3{}
		norms := [3]core.Vec3{}
		verts[0] = mesh.Verts[face.VertIndicies[0]-1]
		verts[1] = mesh.Verts[face.VertIndicies[1]-1]
		verts[2] = mesh.Verts[face.VertIndicies[2]-1]

		uvs[0] = mesh.TexCoords[face.TexCoordIndicies[0]-1]
		uvs[1] = mesh.TexCoords[face.TexCoordIndicies[1]-1]
		uvs[2] = mesh.TexCoords[face.TexCoordIndicies[2]-1]

		norms[0] = mesh.Normals[face.NormalIndicies[0]-1]
		norms[1] = mesh.Normals[face.NormalIndicies[1]-1]
		norms[2] = mesh.Normals[face.NormalIndicies[2]-1]

		*scene = append(*scene, &core.Triangle{Verts: verts[:], UVs: uvs[:], Normals: norms[:], Material: mat})
	}
}

func main() {
	fmt.Println("let's raytrace something")

	fmt.Println("creating a scene with some spheres")

	scene := []core.Hittable{
		//		&core.Sphere{core.Vec3{.7, 1, 0}, 0.2, &core.DiffuseMaterial{core.Col3{.8, .6, .6}}},
		//&core.Sphere{core.Vec3{0, 1.2, 0}, 0.5, &core.DiffuseLightMaterial{core.Col3{1, .8, .1}, 3.5}},
		//	&core.Sphere{core.Vec3{0, -100.5, 0}, 100, &core.MetalMaterial{core.Col3{1, 1, 1}, .2}},
	}
	/*
		for i := 0; i < 300; i++ {
			newSphere := core.Sphere{core.Vec3{rand.Float64()*10.0 - 5, 0, rand.Float64()*10.0 - 5}, .2,
				&core.DiffuseMaterial{&core.SolidTexture{core.Vec3{.8, .6, .6}}}}
			scene = append(scene, &newSphere)
		}
	*/
	fmt.Println("creating camera and image")
	const imageWidth int = 1024
	const imageHeight int = 768
	const samplesPerPixel = 8
	const maxDepth = 5
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	cam := core.NewCameraByPoints(core.Pt3{0, -2, -5}, core.Pt3{0, 0, 0}, core.Vec3{0, 1, 0}, 45.0, 4.0/3.0, .0000001, 4)
	meshbox := core.LoadMeshFromOBJAtPath("./static/cone_tex.obj")

	var boxMat = core.DiffuseMaterial{core.NewImageTexture("./static/cone_tex.png")}
	AddTrisToScene(&scene, &meshbox, &boxMat)

	//meshcubes := core.LoadMeshFromOBJAtPath("./static/glasscubes.obj")

	//AddTrisToScene(&scene, &meshcubes, &core.RefractiveMaterial{1.5})

	bvhForScene := core.NewBVHNode(&scene, 0, len(scene))
	scene = []core.Hittable{&bvhForScene}

	start := time.Now()
	var wg = &sync.WaitGroup{}
	for i := 0; i < imageHeight; i++ {
		for j := 0; j < imageWidth; j++ {
			colors := [samplesPerPixel]core.Col3{}
			for s := 0; s < samplesPerPixel; s++ {
				wg.Add(1)

				go func(s int) {
					defer wg.Done()
					u := (float64(i) + rand.Float64()) / float64(imageHeight-1)
					v := (float64(j) + rand.Float64()) / float64(imageWidth-1)
					r := cam.GetRay(v, u)
					colors[s] = testRayColor(r, &scene, maxDepth)
				}(s)

			}
			wg.Wait()
			color := core.Col3{0, 0, 0}
			for _, curCol := range colors {
				color = color.Add(curCol)
			}
			//I think this is a ppm vs goimage discrepancy (is 0,0 top corner or bottom issue)
			img.SetRGBA(j, imageHeight-i, ConvertColor(color, samplesPerPixel))
		}
		fmt.Println("completed line", i, "of ", imageHeight)
	}
	fmt.Println(time.Since(start))

	outfile, err := os.Create("test.png")
	if err != nil {
		println("some error creating image file")
	}
	png.Encode(outfile, img)
	outfile.Close()
}
