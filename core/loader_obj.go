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
	VertIndicies     [3]int
	NormalIndicies   [3]int
	TexCoordIndicies [3]int
}

type Mesh struct {
	Faces []FaceData
	Verts []Vec3
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
			for i, faceDataChunk := range faceDataChunks[1:] {

				dataIndicies := strings.Split(faceDataChunk, "/")
				for dataTypeIndex, chunk := range dataIndicies {
					actualIndexValue, _ := strconv.ParseInt(chunk, 0, 64)

					switch dataTypeIndex {
					case 0:
						currentFace.VertIndicies[i] = int(actualIndexValue)
					case 1:
						currentFace.NormalIndicies[i] = int(actualIndexValue)
					case 2:
						currentFace.TexCoordIndicies[i] = int(actualIndexValue)
					}
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
