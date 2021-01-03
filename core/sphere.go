package core

import (
	"fmt"
	"math"
)

type Sphere struct {
	Center   Pt3
	Radius   float64
	Material Material
}

func (s *Sphere) Hit(ray *Ray, tmin float64, tmax float64, record *HitRecord) bool {
	oc := ray.Origin.Subtract(s.Center)
	a := Dot(ray.Direction, ray.Direction)
	b := 2.0 * Dot(oc, ray.Direction)
	c := Dot(oc, oc) - (s.Radius * s.Radius)
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return false
	}
	//find nearest root in range
	root := (-b - math.Sqrt(discriminant)) / (2.0 * a)
	if root < tmin || root > tmax {
		root = (-b + math.Sqrt(discriminant)) / (2.0 * a)
		if root < tmin || root > tmax {
			return false
		}
	}
	record.t = root
	record.Hitpoint = ray.At(root)
	record.Normal = (record.Hitpoint.Subtract(s.Center)).Scale(1.0 / s.Radius)
	record.CalculateIfHitIsFrontFacing(ray)
	record.HitMaterial = *&s.Material
	return true
}

func (s *Sphere) BoundingBox(outAABB *AABB) {
	*outAABB = AABB{
		s.Center.Subtract(Vec3{s.Radius, s.Radius, s.Radius}),
		s.Center.Add(Vec3{s.Radius, s.Radius, s.Radius})}
}

func (s *Sphere) GetMaterial() Material {
	return s.Material
}

type Triangle struct {
	Verts    []Vec3
	Normals  []Vec3
	UVs      []Vec3
	Material Material
}

func (tri *Triangle) Hit(ray *Ray, tmin float64, tmax float64, record *HitRecord) bool {
	// do we hit the plane?

	//tri norm by edges:
	C := tri.Verts[0]
	B := tri.Verts[1]
	A := tri.Verts[2]

	norm := Normalize(Cross(B.Subtract(A), C.Subtract(A)))

	//ray is parallel to tri - no intersection
	normDotDir := Dot(norm, ray.Direction)
	if normDotDir == 0 {
		fmt.Println("tri is perpendicular to ray")
		return false
	}

	d := Dot(norm, A)
	t := (d - Dot(norm, ray.Origin)) / normDotDir

	if t < tmin || t > tmax {
		return false
	}

	intersectionPointQ := ray.At(t)

	//now do tri inside out testing using q.
	term1 := Dot(Normalize(Cross(B.Subtract(A), intersectionPointQ.Subtract(A))), norm) >= 0
	term2 := Dot(Normalize(Cross(C.Subtract(B), intersectionPointQ.Subtract(B))), norm) >= 0
	term3 := Dot(Normalize(Cross(A.Subtract(C), intersectionPointQ.Subtract(C))), norm) >= 0

	if term1 && term2 && term3 {
		//return data
		record.t = t
		record.Hitpoint = intersectionPointQ
		//for now... TODO smooth this using averaged vert normals.
		record.Normal = norm
		record.CalculateIfHitIsFrontFacing(ray)
		record.HitMaterial = *&tri.Material
		return true
	}
	return false
}
func (t *Triangle) GetMaterial() Material {
	return t.Material
}

func (tri *Triangle) BoundingBox(outAABB *AABB) {

	//march all coordinates and find max and min
	minx := math.Min(tri.Verts[0].X, tri.Verts[1].X)
	minx = math.Min(minx, tri.Verts[2].X)

	miny := math.Min(tri.Verts[0].Y, tri.Verts[1].Y)
	miny = math.Min(miny, tri.Verts[2].Y)

	minz := math.Min(tri.Verts[0].Z, tri.Verts[1].Z)
	minz = math.Min(minz, tri.Verts[2].Z)

	maxx := math.Max(tri.Verts[0].X, tri.Verts[1].X)
	maxx = math.Max(maxx, tri.Verts[2].X)

	maxy := math.Max(tri.Verts[0].Y, tri.Verts[1].Y)
	maxy = math.Max(maxy, tri.Verts[2].Y)

	maxz := math.Max(tri.Verts[0].Z, tri.Verts[1].Z)
	maxz = math.Max(maxz, tri.Verts[2].Z)

	offset := .00001

	*outAABB = AABB{Pt3{minx - offset, miny - offset, minz - offset}, Pt3{maxx + offset, maxy + offset, maxz + offset}}

}
