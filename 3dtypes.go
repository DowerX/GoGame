package main

import "math"

//2D

type vector2f struct {
	X float64
	Y float64
}

type vector2 struct {
	X int
	Y int
}

func (v2f *vector2f) ToInt() vector2 {
	return vector2{
		X: int(math.Round(v2f.X)),
		Y: int(math.Round(v2f.Y)),
	}
}

func (v2 *vector2) ToFloat() vector2f {
	return vector2f{
		X: float64(v2.X),
		Y: float64(v2.Y),
	}
}

func (v2f *vector2f) Add(a *vector2f) vector2f {
	return vector2f{
		X: v2f.X + a.X,
		Y: v2f.Y + a.Y,
	}
}

func (v2 *vector2) Add(a *vector2) vector2 {
	return vector2{
		X: v2.X + a.X,
		Y: v2.Y + a.Y,
	}
}

func (v2f *vector2f) MultiplyFloat(a float64) vector2f {
	return vector2f{
		X: v2f.X * a,
		Y: v2f.Y * a,
	}
}

func (v2 *vector2) MultiplyInt(a int) vector2 {
	return vector2{
		X: v2.X * a,
		Y: v2.Y * a,
	}
}

func (v2f *vector2f) MultiplyV(a *vector2f) vector2f {
	return vector2f{
		X: v2f.X * a.X,
		Y: v2f.Y * a.Y,
	}
}

func (v2 *vector2) MultiplyV(a *vector2) vector2 {
	return vector2{
		X: v2.X * a.X,
		Y: v2.Y * a.Y,
	}
}

//3D

type vector3f struct {
	X float64
	Y float64
	Z float64
}

type vector3 struct {
	X int
	Y int
	Z int
}

func (v3f *vector3f) ToInt() vector3 {
	return vector3{
		X: int(math.Round(v3f.X)),
		Y: int(math.Round(v3f.Y)),
		Z: int(math.Round(v3f.Z)),
	}
}

func (v3 *vector3) ToFloat() vector3f {
	return vector3f{
		X: float64(v3.X),
		Y: float64(v3.Y),
		Z: float64(v3.Z),
	}
}

func (v3f *vector3f) Add(a *vector3f) vector3f {
	return vector3f{
		X: v3f.X + a.X,
		Y: v3f.Y + a.Y,
		Z: v3f.Z + a.Z,
	}
}

func (v3 *vector3) Add(a *vector3) vector3 {
	return vector3{
		X: v3.X + a.X,
		Y: v3.Y + a.Y,
		Z: v3.Z + a.Z,
	}
}

func (v3f *vector3f) MultiplyFloat(a float64) vector3f {
	return vector3f{
		X: v3f.X * a,
		Y: v3f.Y * a,
		Z: v3f.Z * a,
	}
}

func (v3 *vector3) MultiplyInt(a int) vector3 {
	return vector3{
		X: v3.X * a,
		Y: v3.Y * a,
		Z: v3.Z * a,
	}
}

func (v3f *vector3f) MultiplyV(a *vector3f) vector3f {
	return vector3f{
		X: v3f.X * a.X,
		Y: v3f.Y * a.Y,
		Z: v3f.Z * a.Z,
	}
}

func (v3 *vector3) MultiplyV(a *vector3) vector3 {
	return vector3{
		X: v3.X * a.X,
		Y: v3.Y * a.Y,
		Z: v3.Z * a.Z,
	}
}
