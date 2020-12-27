package core

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
)

//Vec3 contains 3 components - can
//be used to represent colors, points, vectors etc.
type Vec3 struct {
	X, Y, Z float64
}

func Clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func remap(x, min, max float64) float64 {
	return x*(max-min) + min
}

func RandomVector() Vec3 {
	return NewVector3(rand.Float64(), rand.Float64(), rand.Float64())
}

func RandomVectorByRange(min, max float64) Vec3 {
	x := remap(rand.Float64(), min, max)
	y := remap(rand.Float64(), min, max)
	z := remap(rand.Float64(), min, max)

	return NewVector3(x, y, z)

}

func GetRandomVectorInUnitSphere() Vec3 {
	//while loop.
	for i := 0; i < 1; {
		randpt := RandomVectorByRange(-1, 1)
		if randpt.LengthSquared() >= 1 {
			continue

		}
		return randpt
	}
	return Vec3{}
}

//TODO - don't do this
type Col3 = Vec3
type Pt3 = Vec3

//New constructs a vec3 by components
func NewVector3(x float64, y float64, z float64) Vec3 {
	v := Vec3{}
	v.X = x
	v.Y = y
	v.Z = z
	return v
}

//Negate returns a new vector with componets negated
func (v Vec3) Negate() Vec3 {
	return NewVector3(-v.X, -v.Y, -v.Z)
}

func (v Vec3) Add(u Vec3) Vec3 {
	return NewVector3(
		v.X+u.X,
		v.Y+u.Y,
		v.Z+u.Z)
}
func (v Vec3) Subtract(u Vec3) Vec3 {
	return NewVector3(
		v.X-u.X,
		v.Y-u.Y,
		v.Z-u.Z)
}

func (v Vec3) Scale(factor float64) Vec3 {
	return NewVector3(
		v.X*factor,
		v.Y*factor,
		v.Z*factor)
}
func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}
func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) String() string {
	return fmt.Sprintf("[%f, %f, %f]", v.X, v.Y, v.Z)
}

//Multiply does componentwise multiplication - hadamard product
func Multiply(v Vec3, u Vec3) Vec3 {
	return NewVector3(
		v.X*u.X,
		v.Y*u.Y,
		v.Z*u.Z)
}

func Dot(v Vec3, u Vec3) float64 {
	return u.X*v.X +
		v.Y*u.Y +
		v.Z*u.Z
}

func Cross(v Vec3, u Vec3) Vec3 {
	return NewVector3(v.Y*u.Z-v.Z*u.Y,
		v.Z*u.X-v.X*u.Z,
		v.X*u.Y-v.Y-u.X)
}

func Normalize(v Vec3) Vec3 {
	return v.Scale(1.0 / (v.Length()))
}

func (v Vec3) ToRGBA() color.RGBA {
	col := color.RGBA{uint8(v.X), uint8(v.Y), uint8(v.Z), 255}
	return col
}

func Test() {
	println("inside core package")
}
