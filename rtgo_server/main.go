package main

import (
	"context"
	"fmt"
	"image"
	"log"
	"math/rand"
	"net"
	"rt1go/core"
	pb "rt1go/protos/rtgo/protos"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRtgoServer
}

const (
	port = ":50051"
)

func (*server) Render(ctx context.Context, in *pb.RenderRequest) (*pb.RenderReply, error) {

	//TODO gotta be something better than this...
	from := ((*pbvec)(in.Camera.Lookfrom)).ConvertToRTVec()
	to := ((*pbvec)(in.Camera.Lookfrom)).ConvertToRTVec()
	vup := ((*pbvec)(in.Camera.Lookfrom)).ConvertToRTVec()

	//we need to convert our render request into our real go objects...

	//first the camera..
	cam := core.NewCameraByPoints(
		from,
		to,
		vup,
		in.Camera.FovDegrees,
		in.Camera.AspectRatio,
		in.Camera.Aperture,
		in.Camera.Focuslength)

	materials := make([]core.Material, 0)
	scene := make([]core.Hittable, 0)

	//next ready all materials.

	for _, pbmat := range in.Material {
		switch pbmat.MaterialType {
		case pb.Material_DiffuseMaterial:
			texColors := pbmat.Texture.Colors
			var Texture core.Texture
			if len(texColors) > 1 {
				//TODO make a textureImage from an array of colors...
				//we would need to add size or stride... etc.
				//core.NewImageTexture()
				//Texture := core.ImageTexture{}
			} else {
				Texture = &core.SolidTexture{Color: (*pbvec)(texColors[0]).ConvertToRTVec()}
			}

			materials = append(materials, &core.DiffuseMaterial{Texture: Texture})

		case pb.Material_LightMaterial:
			materials = append(materials, &core.DiffuseLightMaterial{Albedo: (*pbvec)(pbmat.Albedo).ConvertToRTVec(),
				IntensityMultipler: pbmat.Lightintensity})
		}
	}

	//next parse and generate geometry objects and append them to scene...
	for i, hittable := range in.Hittable {
		switch hittable.HittableType {
		case pb.Hittable_Sphere:
			sphere := core.Sphere{(*pbvec)(hittable.Center).ConvertToRTVec(), hittable.Radius, materials[i]}
			scene = append(scene, &sphere)

		case pb.Hittable_Mesh:

			verti := make([]int, 0)
			normi := make([]int, 0)
			texti := make([]int, 0)
			vertC := make([]core.Vec3, 0)
			normC := make([]core.Vec3, 0)
			texC := make([]core.Vec3, 0)
			faces := make([]core.FaceData, 0)
			//foreach face grab the 3 indicies for each parameter
			for fi := 0; i < len(hittable.Facedata); fi++ {
				faceData := core.FaceData{}
				data := hittable.Facedata[fi]

				verti = append(verti, int(data.Vertindicies[0]))
				verti = append(verti, int(data.Vertindicies[1]))
				verti = append(verti, int(data.Vertindicies[2]))

				normi = append(normi, int(data.Normalindicies[0]))
				normi = append(normi, int(data.Normalindicies[1]))
				normi = append(normi, int(data.Normalindicies[2]))

				texti = append(texti, int(data.Texcoordindicies[0]))
				texti = append(texti, int(data.Texcoordindicies[1]))
				texti = append(texti, int(data.Texcoordindicies[2]))
				faceData.NormalIndicies = normi
				faceData.VertIndicies = verti
				faceData.TexCoordIndicies = texti
				faces = append(faces, faceData)
			}
			//append all actual data
			for vi := 0; i < len(hittable.Verts); vi++ {
				v := (*pbvec)(hittable.Verts[vi]).ConvertToRTVec()
				n := (*pbvec)(hittable.Normals[vi]).ConvertToRTVec()
				t := (*pbvec)(hittable.Verts[vi]).ConvertToRTVec()
				vertC = append(vertC, v)
				normC = append(normC, n)
				texC = append(texC, t)
			}

			tri := core.Triangle{vertC, normC, texC, materials[i]}
			scene = append(scene, &tri)

		}
	}

	//we're ready to perfom the raytrace...
	fmt.Println("performing ray trace, data is all decoded at this point")
	const imageWidth int = 1024
	const imageHeight int = 768
	const samplesPerPixel = 8
	const maxDepth = 5
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	fmt.Println("creating acceleration structure")

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
					colors[s] = core.TestRayColor(r, &scene, maxDepth, testRayNoHitColor)
				}(s)

			}
			wg.Wait()
			color := core.Col3{0, 0, 0}
			for _, curCol := range colors {
				color = color.Add(curCol)
			}
			//I think this is a ppm vs goimage discrepancy (is 0,0 top corner or bottom issue)
			img.SetRGBA(j, imageHeight-i, core.ConvertColor(color, samplesPerPixel))
		}
		fmt.Println("completed line", i, "of ", imageHeight)
	}
	fmt.Println(time.Since(start))

	result := make([]*pb.Vec3, 0)
	for i := 0; i < imageHeight; i++ {
		for j := 0; j < imageWidth; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			result = append(result, &pb.Vec3{X: float64(r), Y: float64(g), Z: float64(b)})
		}
	}
	return &pb.RenderReply{Colors: result}, nil
}

func testRayNoHitColor(ray *core.Ray) core.Col3 {
	//return Vec3{}
	t := (core.Normalize(ray.Direction).Y + 1.0) * .5
	return (core.Col3{1.0, 1.0, 1.0}.Add(core.Col3{.5, .7, 1.0}.Scale(t))).Scale(1.0 - t)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRtgoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type pbvec pb.Vec3

func (vec *pbvec) ConvertToRTVec() core.Vec3 {
	return core.Vec3{vec.X, vec.Y, vec.Y}
}
