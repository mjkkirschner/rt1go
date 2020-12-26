package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"rt1go/core"
)

func testRayColor(r core.Ray) core.Col3 {
	sphereCent := core.Vec3{0, 0, -1}
	hitPoint := hitSphere(sphereCent, .5, &r)
	if hitPoint > 0.0 {
		normal := core.Normalize(r.At((hitPoint)).Subtract(sphereCent))
		return core.Vec3{normal.X + 1, normal.Y + 1, normal.Z + 1}.Scale(255).Scale(.5)
	}
	dn := core.Normalize(r.Direction)
	var t = (dn.Y + 1.0) * .5
	return core.Vec3{255, 255, 255}.Scale((1.0 - t)).Add(core.NewVector3(.5*255, .7*255, 1.0*255).Scale((t)))
}

func hitSphere(center core.Pt3, radius float64, ray *core.Ray) float64 {
	//TODO - solve this geometrically for funzies
	oc := ray.Origin.Subtract(center)
	a := core.Dot(ray.Direction, ray.Direction)
	b := 2.0 * core.Dot(oc, ray.Direction)
	c := core.Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	//negative is behind us
	if discriminant < 0 {
		return -1.0
	} else {
		//only return one solution
		return (-b - math.Sqrt(discriminant)) / (2.0 * a)
	}
}

func main() {
	fmt.Println("let's raytrace something")

	fmt.Println("creating camera and image")
	const imageWidth int = 640
	const imageHeight int = 480
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	cam := core.NewCamera(2, 2.66666666667, 1, core.NewVector3(0, 0, 0))
	lowerLeftCorner := cam.Origin.Subtract((cam.Horizontal.Scale(.5))).Subtract(cam.Vertical.Scale(.5)).Subtract(core.NewVector3(0, 0, float64(cam.FocalLength)))
	fmt.Printf("lower left corner is:", lowerLeftCorner)

	for i := 0; i < imageHeight; i++ {
		for j := 0; j < imageWidth; j++ {
			u := float64(i) / float64(imageHeight-1)
			v := float64(j) / float64(imageWidth-1)
			r := core.NewRay(core.NewVector3(0, 0, 0), lowerLeftCorner.Add(cam.Horizontal.Scale(v)).Add(cam.Vertical.Scale((u))).Subtract(cam.Origin))
			col := testRayColor(r)
			//I think this is a ppm vs goimage discrepancy (is 0,0 top corner or bottom issue)
			img.SetRGBA(j, imageHeight-i, col.ToRGBA())
		}
	}
	outfile, err := os.Create("test.png")
	if err != nil {
		println("some error creating image file")
	}
	png.Encode(outfile, img)
	outfile.Close()
}
