package core

import (
	"math"
)

type HitRecord struct {
	Hitpoint    Pt3
	Normal      Vec3
	t           float64
	FrontFacing bool
}

func (hr *HitRecord) CalculateIfHitIsFrontFacing(ray *Ray) {
	hr.FrontFacing = Dot(ray.Direction, hr.Normal) < 0
}

type Hittable interface {
	Hit(ray *Ray, tmin float64, tmax float64, record *HitRecord) bool
}

type Sphere struct {
	Center Pt3
	Radius float64
}

func (s *Sphere) Hit(ray *Ray, tmin float64, tmax float64, record *HitRecord) bool {
	oc := ray.Origin.Subtract(s.Center)
	a := Dot(ray.Direction, ray.Direction)
	b := 2.0 * Dot(oc, ray.Direction)
	c := Dot(oc, oc) - s.Radius*s.Radius
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
	return true

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
