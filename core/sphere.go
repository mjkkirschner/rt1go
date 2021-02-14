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
		var u, v, w = BaryCoords(record, tri.Verts[0], tri.Verts[1], tri.Verts[2])
		//use u v w to interpolate u v coords.
		record.U = tri.UVs[0].X*u + tri.UVs[1].X*v + tri.UVs[2].X*w
		record.V = tri.UVs[0].Y*u + tri.UVs[1].Y*v + tri.UVs[2].Y*w

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

func BaryCoords2(hit *HitRecord, a Vec3, b Vec3, c Vec3) (float64, float64, float64) {

	var p = hit.Hitpoint

	var edge1 = b.Subtract(a)
	var edge2 = c.Subtract(a)
	var edge3 = a.Subtract(c)
	var edge4 = c.Subtract(b)
	var edge5 = b.Subtract(a)

	var edgePB = p.Subtract(b)
	var edgePC = p.Subtract(c)
	var edgePA = p.Subtract(a)

	var triNorm = Cross(edge1, edge2)
	var area = triNorm.Z / 2.0

	var temp1 = Cross(edge4, edgePB)
	var temp2 = Cross(edge3, edgePC)
	var temp3 = Cross(edge5, edgePA)

	var u = (temp1.Z / 2.0) / area
	var v = (temp2.Z / 2.0) / area
	var w = (temp3.Z / 2.0) / area
	return u, v, w
}

func BaryCoords(hit *HitRecord, v0 Vec3, v1 Vec3, v2 Vec3) (float64, float64, float64) {

	var p = hit.Hitpoint
	var v0v1 = v1.Subtract(v0)
	var v0v2 = v2.Subtract(v0)
	var triN = Cross(v0v1, v0v2)
	var area = triN.Length()

	var edge1 = v2.Subtract(v1)
	var vp1 = p.Subtract(v1)
	var c = Cross(edge1, vp1)

	var u = c.Length() / area

	var edge2 = v0.Subtract(v2)
	var vp2 = p.Subtract(v2)
	var c2 = Cross(edge2, vp2)
	var v = c2.Length() / area

	var w = 1 - u - v
	return u, v, w
}
