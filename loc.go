package loc

import (
	"fmt"
	"image"
	"iter"

	"github.com/eihigh/ng"
)

// A Point is an X, Y coordinate pair. The axes increase right and down.
type Point[S ng.Scalar] struct {
	X, Y S
}

// Xy is shorthand for Point{X, Y}.
func Xy[S ng.Scalar](x, y S) Point[S] {
	return Point[S]{X: x, Y: y}
}

func (p Point[S]) Xy() (S, S) {
	return p.X, p.Y
}

// String returns a string representation of p like "(3,4)".
func (p Point[S]) String() string {
	return "(" + fmt.Sprint(p.X) + "," + fmt.Sprint(p.Y) + ")"
}

// Add returns the vector p+q.
func (p Point[S]) Add(q Point[S]) Point[S] {
	return Point[S]{X: p.X + q.X, Y: p.Y + q.Y}
}

// Sub returns the vector p-q.
func (p Point[S]) Sub(q Point[S]) Point[S] {
	return Point[S]{X: p.X - q.X, Y: p.Y - q.Y}
}

// Mul returns the vector p*k.
func (p Point[S]) Mul(k S) Point[S] {
	return Point[S]{X: p.X * k, Y: p.Y * k}
}

// Div returns the vector p/k.
func (p Point[S]) Div(k S) Point[S] {
	return Point[S]{X: p.X / k, Y: p.Y / k}
}

func (p Point[S]) MulPoint(q Point[S]) Point[S] {
	return Point[S]{X: p.X * q.X, Y: p.Y * q.Y}
}

func (p Point[S]) DivPoint(q Point[S]) Point[S] {
	return Point[S]{X: p.X / q.X, Y: p.Y / q.Y}
}

// In reports whether p is in r.
func (p Point[S]) In(r Rect[S]) bool {
	return r.Min.X <= p.X && p.X < r.Max.X &&
		r.Min.Y <= p.Y && p.Y < r.Max.Y
}

// TODO: Point.Mod?

// Eq reports whether p and q are equal.
func (p Point[S]) Eq(q Point[S]) bool {
	return p.X == q.X && p.Y == q.Y
}

// Image returns the point as an image.Point.
func (p Point[S]) Image() image.Point {
	return image.Point{X: int(p.X), Y: int(p.Y)}
}

// Int returns the point as an int point.
func (p Point[S]) Int() Point[int] {
	return Point[int]{X: int(p.X), Y: int(p.Y)}
}

// Float64 returns the point as a float64 point.
func (p Point[S]) Float64() Point[float64] {
	return Point[float64]{X: float64(p.X), Y: float64(p.Y)}
}

// Float32 returns the point as a float32 point.
func (p Point[S]) Float32() Point[float32] {
	return Point[float32]{X: float32(p.X), Y: float32(p.Y)}
}

// AsSize converts the x, y coordinates to a rectangle with the corresponding width and height.
func (p Point[S]) AsSize() Rect[S] {
	return Rect[S]{Max: p}
}

// A Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y.
// It is well-formed if Min.X <= Max.X and likewise for Y. Points are always
// well-formed. A rectangle's methods always return well-formed outputs for
// well-formed inputs.
type Rect[S ng.Scalar] struct {
	Min, Max Point[S]
}

// Xyxy creates a rectangle from two points, (x0, y0) and (x1, y1).
func Xyxy[S ng.Scalar](x0, y0, x1, y1 S) Rect[S] {
	return Rect[S]{
		Min: Point[S]{X: x0, Y: y0},
		Max: Point[S]{X: x1, Y: y1},
	}
}

// Xywh creates a rectangle from a point (x, y) and a size (w, h).
func Xywh[S ng.Scalar](x, y, w, h S) Rect[S] {
	return Rect[S]{
		Min: Point[S]{X: x, Y: y},
		Max: Point[S]{X: x + w, Y: y + h},
	}
}

// MinMax creates a rectangle from two points, min and max.
func MinMax[S ng.Scalar, P ng.Vec2like[S]](min, max P) Rect[S] {
	return Rect[S]{
		Min: Point[S](min),
		Max: Point[S](max),
	}
}

// PosSize creates a rectangle from a position and a size.
func PosSize[S ng.Scalar, P ng.Vec2like[S]](pos, size P) Rect[S] {
	vp := ng.Vec2[S](pos)
	vs := ng.Vec2[S](size)
	return Rect[S]{
		Min: Point[S](pos),
		Max: Point[S]{X: vp.X + vs.X, Y: vp.Y + vs.Y},
	}
}

// String returns a string representation of r like "(3,4)-(6,5)".
func (r Rect[S]) String() string {
	return r.Min.String() + "-" + r.Max.String()
}

// Dx returns r's width.
func (r Rect[S]) Dx() S {
	return r.Max.X - r.Min.X
}

// Dy returns r's height.
func (r Rect[S]) Dy() S {
	return r.Max.Y - r.Min.Y
}

// Size returns r's width and height.
func (r Rect[S]) Size() Point[S] {
	return Point[S]{
		X: r.Max.X - r.Min.X,
		Y: r.Max.Y - r.Min.Y,
	}
}

// Add returns the rectangle r translated by p.
func (r Rect[S]) Add(p Point[S]) Rect[S] {
	return Rect[S]{
		Min: Point[S]{X: r.Min.X + p.X, Y: r.Min.Y + p.Y},
		Max: Point[S]{X: r.Max.X + p.X, Y: r.Max.Y + p.Y},
	}
}

// Sub returns the rectangle r translated by -p.
func (r Rect[S]) Sub(p Point[S]) Rect[S] {
	return Rect[S]{
		Min: Point[S]{X: r.Min.X - p.X, Y: r.Min.Y - p.Y},
		Max: Point[S]{X: r.Max.X - p.X, Y: r.Max.Y - p.Y},
	}
}

// Inset returns the rectangle r inset by n, which may be negative. If either
// of r's dimensions is less than 2*n then an empty rectangle near the center
// of r will be returned.
func (r Rect[S]) Inset(n S) Rect[S] {
	return r.Inset4(n, n, n, n)
}

// Inset2 returns the rectangle r inset by n in both dimensions. If either of
// r's dimensions is less than 2*n then an empty rectangle near the center of
// r will be returned.
func (r Rect[S]) Inset2(x, y S) Rect[S] {
	return r.Inset4(x, y, x, y)
}

// Inset4 returns the rectangle r inset by left, top, right, and bottom.
// If either of r's dimensions is less than left+right or top+bottom then an
// empty rectangle near the center of r will be returned.
func (r Rect[S]) Inset4(left, top, right, bottom S) Rect[S] {
	if r.Dx() < left+right {
		r.Min.X = (r.Min.X + r.Max.X) / 2
		r.Max.X = r.Min.X
	} else {
		r.Min.X += left
		r.Max.X -= right
	}
	if r.Dy() < top+bottom {
		r.Min.Y = (r.Min.Y + r.Max.Y) / 2
		r.Max.Y = r.Min.Y
	} else {
		r.Min.Y += top
		r.Max.Y -= bottom
	}
	return r
}

// Intersect returns the largest rectangle contained by both r and s. If the
// two rectangles do not overlap then the zero rectangle will be returned.
func (r Rect[S]) Intersect(s Rect[S]) Rect[S] {
	if r.Min.X < s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y < s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X > s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y > s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	// Letting r0 and s0 be the values of r and s at the time that the method
	// is called, this next line is equivalent to:
	//
	// if max(r0.Min.X, s0.Min.X) >= min(r0.Max.X, s0.Max.X) || likewiseForY { etc }
	if r.Empty() {
		return Rect[S]{}
	}
	return r
}

// Union returns the smallest rectangle that contains both r and s.
func (r Rect[S]) Union(s Rect[S]) Rect[S] {
	if r.Empty() {
		return s
	}
	if s.Empty() {
		return r
	}
	if r.Min.X > s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y > s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X < s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y < s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	return r
}

// Empty reports whether the rectangle contains no points.
func (r Rect[S]) Empty() bool {
	return r.Min.X >= r.Max.X || r.Min.Y >= r.Max.Y
}

// Eq reports whether r and s contain the same set of points. All empty
// rectangles are considered equal.
func (r Rect[S]) Eq(s Rect[S]) bool {
	return r == s || r.Empty() && s.Empty()
}

// Overlaps reports whether r and s have a non-empty intersection.
func (r Rect[S]) Overlaps(s Rect[S]) bool {
	return !r.Empty() && !s.Empty() &&
		r.Min.X < s.Max.X && s.Min.X < r.Max.X &&
		r.Min.Y < s.Max.Y && s.Min.Y < r.Max.Y
}

// In reports whether every point in r is in s.
func (r Rect[S]) In(s Rect[S]) bool {
	if r.Empty() {
		return true
	}
	// Note that r.Max is an exclusive bound for r, so that r.In(s)
	// does not require that r.Max.In(s).
	return s.Min.X <= r.Min.X && r.Max.X <= s.Max.X &&
		s.Min.Y <= r.Min.Y && r.Max.Y <= s.Max.Y
}

// Canon returns the canonical version of r. The returned rectangle has minimum
// and maximum coordinates swapped if necessary so that it is well-formed.
func (r Rect[S]) Canon() Rect[S] {
	if r.Max.X < r.Min.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}
	if r.Max.Y < r.Min.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}
	return r
}

// Image returns the rectangle as an image.Rectangle.
func (r Rect[S]) Image() image.Rectangle {
	return image.Rect(int(r.Min.X), int(r.Min.Y), int(r.Max.X), int(r.Max.Y))
}

// Points returns a sequence of points in the rectangle.
func (r Rect[S]) Points() iter.Seq[Point[S]] {
	return func(yield func(Point[S]) bool) {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				if !yield(Point[S]{X: x, Y: y}) {
					return
				}
			}
		}
	}
}
