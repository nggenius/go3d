package vec2

import (
	"fmt"
	"math"

	"github.com/barnex/fmath"
	"github.com/ungerik/go3d/generic"
)

var (
	Zero = T{}

	UnitX = T{1, 0}
	UnitY = T{0, 1}

	MinVal = T{-math.MaxFloat32, -math.MaxFloat32}
	MaxVal = T{+math.MaxFloat32, +math.MaxFloat32}
)

type T [2]float32

func From(other generic.T) T {
	return T{other.Get(0, 0), other.Get(0, 1)}
}

func Parse(s string) (r T, err error) {
	_, err = fmt.Sscanf(s, "%f %f", &r[0], &r[1])
	return r, err
}

func (self *T) String() string {
	return fmt.Sprintf("%f %f", self[0], self[1])
}

func (self *T) Rows() int {
	return 2
}

func (self *T) Cols() int {
	return 1
}

func (self *T) Size() int {
	return 2
}

func (self *T) Slice() []float32 {
	return []float32{self[0], self[1]}
}

func (self *T) Get(col, row int) float32 {
	return self[row]
}

func (self *T) IsZero() bool {
	return self[0] == 0 && self[1] == 0
}

func (self *T) Length() float32 {
	return float32(fmath.Sqrt(self.LengthSqr()))
}

func (self *T) LengthSqr() float32 {
	return self[0]*self[0] + self[1]*self[1]
}

func (self *T) Scale(f float32) {
	self[0] *= f
	self[1] *= f
}

func (self *T) Invert() {
	self[0] = -self[0]
	self[1] = -self[1]
}

func (self *T) Inverted() T {
	return T{-self[0], -self[1]}
}

func (self *T) Normalize() {
	sl := self.LengthSqr()
	if sl == 0 || sl == 1 {
		return
	}
	self.Scale(1 / fmath.Sqrt(sl))
}

func (self *T) Normalized() T {
	v := *self
	v.Normalize()
	return v
}

func (self *T) Add(v *T) {
	self[0] += v[0]
	self[1] += v[1]
}

func (self *T) Sub(v *T) {
	self[0] -= v[0]
	self[1] -= v[1]
}

func (self *T) Mul(v *T) {
	self[0] *= v[0]
	self[1] *= v[1]
}

func Add(a, b *T) T {
	return T{a[0] + b[0], a[1] + b[1]}
}

func Sub(a, b *T) T {
	return T{a[0] - b[0], a[1] - b[1]}
}

func Mul(a, b *T) T {
	return T{a[0] * b[0], a[1] * b[1]}
}

func Dot(a, b *T) float32 {
	return a[0]*b[0] + a[1]*b[1]
}

func Cross(a, b *T) T {
	return T{
		a[1]*b[0] - a[0]*b[1],
		a[0]*b[1] - a[1]*b[0],
	}
}

func Angle(a, b *T) float32 {
	return fmath.Acos(Dot(a, b))
}

func Min(a, b *T) T {
	min := *a
	if b[0] < min[0] {
		min[0] = b[0]
	}
	if b[1] < min[1] {
		min[1] = b[1]
	}
	return min
}

func Max(a, b *T) T {
	max := *a
	if b[0] > max[0] {
		max[0] = b[0]
	}
	if b[1] > max[1] {
		max[1] = b[1]
	}
	return max
}
