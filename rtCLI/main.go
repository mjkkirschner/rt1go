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
)

func testRayColor(r core.Ray, scene *[]core.Hittable) core.Col3 {

	hit := core.HitRecord{}
	if core.GetClosestHit(scene, &r, 0, math.Inf(1), &hit) {
		norm := hit.Normal
		if !hit.FrontFacing {
			norm = norm.Negate()
		}
		return norm.Add(core.Vec3{1.0, 1.0, 1.0}).Scale(255.0).Scale(.5)
	}
	dn := core.Normalize(r.Direction)
	var t = (dn.Y + 1.0) * .5
	return core.Vec3{255, 255, 255}.Scale((1.0 - t)).Add(core.NewVector3(.5*255, .7*255, 1.0*255).Scale((t)))
}

func ConvertColor(color core.Col3, samples int) color.RGBA {
	return color.Scale(1.0 / float64(samples)).ToRGBA()
}

func main() {
	fmt.Println("let's raytrace something")

	fmt.Println("creating a scene with some spheres")

	scene := []core.Hittable{&core.Sphere{core.Vec3{0, 0, -1}, 0.5}, &core.Sphere{core.Vec3{0, -100.5, -1}, 100}}

	for i := 0; i < 1000; i++ {
		newSphere := core.Sphere{core.Vec3{rand.Float64()*100.0 - 50, rand.Float64()*100.0 - 50, rand.Float64()*100.0 - 50}, rand.Float64() * 5.0}
		scene = append(scene, &newSphere)
	}

	fmt.Println("creating camera and image")
	const imageWidth int = 640
	const imageHeight int = 480
	const samplesPerPixel = 2
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	cam := core.NewCamera(2, 2.66666666667, 1, core.NewVector3(0, 0, 0))

	for i := 0; i < imageHeight; i++ {
		for j := 0; j < imageWidth; j++ {
			color := core.Col3{0, 0, 0}
			for s := 0; s < samplesPerPixel; s++ {

				u := (float64(i) + rand.Float64()) / float64(imageHeight-1)
				v := (float64(j) + rand.Float64()) / float64(imageWidth-1)
				r := cam.GetRay(v, u)

				color = color.Add(testRayColor(r, &scene))
			}

			//I think this is a ppm vs goimage discrepancy (is 0,0 top corner or bottom issue)
			img.SetRGBA(j, imageHeight-i, ConvertColor(color, samplesPerPixel))
		}
	}
	outfile, err := os.Create("test.png")
	if err != nil {
		println("some error creating image file")
	}
	png.Encode(outfile, img)
	outfile.Close()
}
