package core

type Ray struct {
	Direction, Origin Vec3
}

func NewRay(origin Pt3, direction Vec3) Ray {
	return Ray{Direction: direction, Origin: origin}
}

func (r Ray) At(t float64) Pt3 {
	return r.Origin.Add(r.Direction.Scale(t))
}
