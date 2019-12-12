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

func sgn(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

// Sgn returns vector signed on every axis
func (v Vector3D) Sgn() Vector3D {
	return Vector3D{X: sgn(v.X), Y: sgn(v.Y), Z: sgn(v.Z)}
}

// L1Norm returns norm in taxicab geometry (Manhattan distance)
func (v Vector3D) L1Norm() int {
	return Abs(v.X) + Abs(v.Y) + Abs(v.Z)
}

// Eq returns if vectors are equal
func (v Vector3D) Eq(other Vector3D) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}
