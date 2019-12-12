package utils

// Vector3D is a vector in 3D space
type Vector3D struct {
	X int
	Y int
	Z int
}

// Add adds two vectors together
func (v Vector3D) Add(other Vector3D) Vector3D {
	return Vector3D{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

// Subtract subtracts two vectors together
func (v Vector3D) Subtract(other Vector3D) Vector3D {
	return Vector3D{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

// Sgn returns vector signed on every axis
func (v Vector3D) Sgn() Vector3D {
	return Vector3D{X: Sgn(v.X), Y: Sgn(v.Y), Z: Sgn(v.Z)}
}

// L1Norm returns norm in taxicab geometry (Manhattan distance)
func (v Vector3D) L1Norm() int {
	return Abs(v.X) + Abs(v.Y) + Abs(v.Z)
}

// Eq returns if vectors are equal
func (v Vector3D) Eq(other Vector3D) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}

// JustX returns copy of vector with only the X axis
func (v Vector3D) JustX() Vector3D {
	return Vector3D{X: v.X}
}

// JustY returns copy of vector with only the Y axis
func (v Vector3D) JustY() Vector3D {
	return Vector3D{Y: v.Y}
}

// JustZ returns copy of vector with only the Z axis
func (v Vector3D) JustZ() Vector3D {
	return Vector3D{Z: v.Z}
}
