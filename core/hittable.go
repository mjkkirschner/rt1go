package core

type HitRecord struct {
	Hitpoint    Pt3
	Normal      Vec3
	t           float64
	FrontFacing bool
	HitMaterial Material
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
