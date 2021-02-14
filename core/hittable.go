package core

import (
	"math"
	"math/rand"
	"sort"
)

type HitRecord struct {
	Hitpoint    Pt3
	Normal      Vec3
	t           float64
	FrontFacing bool
	HitMaterial Material
	U           float64
	V           float64
}

func (hr *HitRecord) CalculateIfHitIsFrontFacing(ray *Ray) {
	hr.FrontFacing = Dot(ray.Direction, hr.Normal) < 0
	if !hr.FrontFacing {
		hr.Normal = hr.Normal.Negate()
	}
}

type Hittable interface {
	GetMaterial() Material
	Hit(ray *Ray, tmin float64, tmax float64, record *HitRecord) bool
	BoundingBox(outAABB *AABB)
}

func GetClosestHit(scene *[]Hittable, ray *Ray, min float64, max float64, record *HitRecord) bool {
	temprecord := HitRecord{}
	anyhit := false
	closestT := max

	for _, item := range *scene {
		//fmt.Println(index, item)
		if item.Hit(ray, min, closestT, &temprecord) {
			anyhit = true
			closestT = temprecord.t
			*record = temprecord
		}
	}
	return anyhit
}

type AABB struct {
	Min Pt3
	Max Pt3
}

func (box *AABB) HitOLD(ray *Ray, tmin float64, tmax float64) bool {
	for i := 0; i < 3; i++ {
		t0 := math.Min((box.Min.Index(i)-ray.Origin.Index(i))/ray.Direction.Index(i),
			(box.Max.Index(i)-ray.Origin.Index(i))/ray.Direction.Index(i))
		t1 := math.Max((box.Min.Index(i)-ray.Origin.Index(i))/ray.Direction.Index(i),
			(box.Max.Index(i)-ray.Origin.Index(i))/ray.Direction.Index(i))

		tmin = math.Max(t0, tmin)
		tmax = math.Min(t1, tmax)
		if tmax <= tmin {
			return false
		}
	}
	return true
}

func (box *AABB) Hit(ray *Ray, tmin float64, tmax float64) bool {
	for i := 0; i < 3; i++ {
		invD := 1.0 / ray.Direction.Index(i)
		t0 := (box.Min.Index(i) - ray.Origin.Index(i)) * invD
		t1 := (box.Max.Index(i) - ray.Origin.Index(i)) * invD
		if invD < 0 {
			t3 := t1
			t1 = t0
			t0 = t3
		}
		if t0 > tmin {
			tmin = t0
		}
		if t1 < tmax {
			tmax = t1
		}
		if tmax <= tmin {
			return false
		}
	}
	return true
}

func AABBFromTwoAABBs(box1 AABB, box2 AABB) AABB {
	small := Pt3{
		math.Min(box1.Min.X, box2.Min.X),
		math.Min(box1.Min.Y, box2.Min.Y),
		math.Min(box1.Min.Z, box2.Min.Z),
	}

	large := Pt3{
		math.Max(box1.Max.X, box2.Max.X),
		math.Max(box1.Max.Y, box2.Max.Y),
		math.Max(box1.Max.Z, box2.Max.Z),
	}
	return AABB{small, large}
}

func BoundingBoxFromScene(scene *[]Hittable, outAABB *AABB) bool {
	firstBox := false
	tempbox := AABB{}

	for _, h := range *scene {
		h.BoundingBox(&tempbox)
		if firstBox {
			*outAABB = tempbox
		} else {
			*outAABB = AABBFromTwoAABBs(*outAABB, tempbox)
		}
	}
	return true
}

type BvhNode struct {
	Left  *Hittable
	Right *Hittable
	Box   AABB
}

func (node *BvhNode) GetMaterial() Material {
	return &DiffuseMaterial{&SolidTexture{Col3{0, 1, 0}}}
}

func NewBVHNode(hittables *[]Hittable, start int, end int) BvhNode {

	var left Hittable
	var right Hittable
	output := BvhNode{&left, &right, AABB{}}
	axis := rand.Intn(3)
	var compareFunc func(a *Hittable, b *Hittable) bool

	switch axis {
	case 0:
		compareFunc = boxXCompare
	case 1:
		compareFunc = boxYCompare
	case 2:
		compareFunc = boxZCompare
	}
	objectSpan := end - start
	if objectSpan == 1 {
		*output.Left = (*hittables)[start]
		*output.Right = (*hittables)[start]
	} else if objectSpan == 2 {
		if compareFunc(&(*hittables)[start], &(*hittables)[start+1]) {
			*output.Left = (*hittables)[start]
			*output.Right = (*hittables)[start+1]
		} else {
			*output.Left = (*hittables)[start+1]
			*output.Right = (*hittables)[start]
		}
	} else {
		sort.Slice((*hittables)[start:end], func(i, j int) bool {
			a := (*hittables)[start:end][i]
			b := (*hittables)[start:end][j]
			return compareFunc(&a, &b)
		})
		mid := start + objectSpan/2
		outleft := NewBVHNode(hittables, start, mid)
		outright := NewBVHNode(hittables, mid, end)
		*output.Left = &outleft
		*output.Right = &outright
	}
	box_left := AABB{}
	box_right := AABB{}
	(*output.Left).BoundingBox(&box_left)
	(*output.Right).BoundingBox(&box_right)
	output.Box = AABBFromTwoAABBs(box_left, box_right)
	return output
}

func boxCompareInternal(a *Hittable, b *Hittable, axis int) bool {
	boxa := AABB{}
	boxb := AABB{}
	(*a).BoundingBox(&boxa)
	(*b).BoundingBox(&boxb)
	return boxa.Min.Index(axis) < boxb.Min.Index(axis)
}

func boxXCompare(a *Hittable, b *Hittable) bool {
	return boxCompareInternal(a, b, 0)
}
func boxYCompare(a *Hittable, b *Hittable) bool {
	return boxCompareInternal(a, b, 1)
}
func boxZCompare(a *Hittable, b *Hittable) bool {
	return boxCompareInternal(a, b, 2)
}

func (node *BvhNode) BoundingBox(outAABB *AABB) {
	*outAABB = node.Box
}

func (node *BvhNode) Hit(ray *Ray, tmin float64, tmax float64, record *HitRecord) bool {
	if !node.Box.Hit(ray, tmin, tmax) {
		return false
	}
	hitLeft := (*node.Left).Hit(ray, tmin, tmax, record)
	rightMax := tmax
	if hitLeft {
		rightMax = record.t
	}
	hitright := (*node.Right).Hit(ray, tmin, rightMax, record)
	return hitLeft || hitright
}
