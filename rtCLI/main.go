package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"rt1go/core"
)

func testRayColor(r core.Ray) core.Col3 {
	dn := core.Normalize(r.Direction)
	var t = (dn.Y + 1.0) * .5
	return core.Vec3{255, 255, 255}.Scale((1.0 - t)).Add(core.NewVector3(.5*255, .7*255, 1.0*255).Scale((t)))
}

func main() {
	fmt.Println("let's raytrace something")

	fmt.Println("creating camera and image")
	const imageWidth int = 640
	const imageHeight int = 480
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	cam := core.NewCamera(2, 3.5, 1, core.NewVector3(0, 0, 0))
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
