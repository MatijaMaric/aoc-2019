package utils

import (
	"math"
)

// Vector2D is a vector in 2D space
type Vector2D struct {
	X int
	Y int
}

// Add adds two vectors together
func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{v.X + other.X, v.Y + other.Y}
}

// Subtract subtracts two vectors together
func (v Vector2D) Subtract(other Vector2D) Vector2D {
	return Vector2D{v.X - other.X, v.Y - other.Y}
}

// Multiply multiplies vector by a factor of x
func (v Vector2D) Multiply(x int) Vector2D {
	return Vector2D{v.X * x, v.Y * x}
}

// Divide divides vector by a factor of x
func (v Vector2D) Divide(x int) Vector2D {
	return Vector2D{v.X / x, v.Y / x}
}

// IntNorm returns minimum normalized vector in integer space
func (v Vector2D) IntNorm() Vector2D {
	factor := Gcd(Abs(v.X), Abs(v.Y))
	return v.Divide(factor)
}

// Eq returns true if vectors are same
func (v Vector2D) Eq(other Vector2D) bool {
	return v.X == other.X && v.Y == other.Y
}

// Copy returns copy of vector
func (v Vector2D) Copy() Vector2D {
	return Vector2D{v.X, v.Y}
}

// Angle returns angle of vector
func (v Vector2D) Angle() float64 {
	return 2.0*math.Pi - (math.Atan2(float64(v.X), float64(v.Y)) + math.Pi)
}

// MinVec2 returns vector in 2d space with minimum of both axis
func MinVec2(a, b Vector2D) Vector2D {
	return Vector2D{X: Min(a.X, b.X), Y: Min(a.Y, b.Y)}
}

// MaxVec2 returns vector in 2d space with maximum of both axis
func MaxVec2(a, b Vector2D) Vector2D {
	return Vector2D{X: Max(a.X, b.X), Y: Max(a.Y, b.Y)}
}
