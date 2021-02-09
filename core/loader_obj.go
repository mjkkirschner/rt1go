package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type FaceData struct {
	VertIndicies     []int
	NormalIndicies   []int
	TexCoordIndicies []int
}

type Mesh struct {
	Faces     []FaceData
	Verts     []Vec3
	TexCoords []Vec3
	Normals   []Vec3
}

func LoadMeshFromOBJAtPath(path string) Mesh {
	faces := LoadTrisFromOBJatPath(path)
	mesh := Mesh{Faces: faces}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chunk := strings.Split(line, " ")
		if len(chunk) < 1 {
			continue
		}
		if chunk[0] == "v" {
			x, _ := strconv.ParseFloat(chunk[1], 0)
			y, _ := strconv.ParseFloat(chunk[2], 0)
			z, _ := strconv.ParseFloat(chunk[3], 0)
			mesh.Verts = append(mesh.Verts, Vec3{x, y, z})
		} else if chunk[0] == "vt" {
			u, _ := strconv.ParseFloat(chunk[1], 0)
			v, _ := strconv.ParseFloat(chunk[2], 0)
			mesh.TexCoords = append(mesh.TexCoords, Vec3{u, v, 0})
		} else if chunk[0] == "vn" {
			x, _ := strconv.ParseFloat(chunk[1], 0)
			y, _ := strconv.ParseFloat(chunk[2], 0)
			z, _ := strconv.ParseFloat(chunk[3], 0)
			mesh.Normals = append(mesh.Normals, Vec3{x, y, z})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return mesh
}

func LoadTrisFromOBJatPath(path string) []FaceData {
	tris := []FaceData{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		faceDataChunks := strings.Split(line, " ")
		if strings.HasPrefix(faceDataChunks[0], "f") {
			currentFace := FaceData{}
			//TODO for now we only handle vert position data
			for _, faceDataChunk := range faceDataChunks[1:] {

				dataIndicies := strings.Split(faceDataChunk, "/")
				for dataTypeIndex, chunk := range dataIndicies {
					actualIndexValue, _ := strconv.ParseInt(chunk, 0, 64)

					switch dataTypeIndex {
					case 0:
						currentFace.VertIndicies = append(currentFace.VertIndicies, int(actualIndexValue))
					case 2:
						currentFace.NormalIndicies = append(currentFace.NormalIndicies, int(actualIndexValue))
					case 1:
						currentFace.TexCoordIndicies = append(currentFace.TexCoordIndicies, int(actualIndexValue))
					}
				}

			}
			//there are 4 verts for this face or more - this is a tri strip
			vertCount := len(faceDataChunks[1:])
			if vertCount > 3 {
				for i := 3; i < vertCount; i++ {
					newFace := FaceData{VertIndicies: []int{currentFace.VertIndicies[i-3], currentFace.VertIndicies[i-1], currentFace.VertIndicies[i]},
						NormalIndicies:   []int{currentFace.NormalIndicies[i-3], currentFace.NormalIndicies[i-1], currentFace.NormalIndicies[i]},
						TexCoordIndicies: []int{currentFace.TexCoordIndicies[i-3], currentFace.TexCoordIndicies[i-1], currentFace.TexCoordIndicies[i]}}

					tris = append(tris, newFace)
				}
			}
			tris = append(tris, currentFace)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tris
}
