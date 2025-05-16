package loc

import "github.com/eihigh/ng"

// rel returns a relative portion of length.
func rel[S ng.Scalar](length S, r float64) S {
	return S(r * float64(length))
}

// Anchor returns a point within r, scaled by rx and ry.
// rx=0, ry=0 is r.Min; rx=1, ry=1 is r.Max.
func (r Rect[S]) Anchor(rx, ry float64) Point[S] {
	return Point[S]{
		X: r.Min.X + rel(r.Dx(), rx),
		Y: r.Min.Y + rel(r.Dy(), ry),
	}
}

// Center returns the center point of r.
func (r Rect[S]) Center() Point[S] {
	return r.Anchor(0.5, 0.5)
}

// Align returns a new rectangle with the same size as r,
// where the point p is at the relative position (rx, ry) within the new rectangle.
func (p Point[S]) Align(r Rect[S], rx, ry float64) Rect[S] {
	return Rect[S]{
		Min: Point[S]{
			X: p.X - rel(r.Dx(), rx),
			Y: p.Y - rel(r.Dy(), ry),
		},
		Max: Point[S]{
			X: p.X + rel(r.Dx(), 1-rx),
			Y: p.Y + rel(r.Dy(), 1-ry),
		},
	}
}

// AlignCenter returns a new rectangle with the same size as r,
// where the point p is at the center of the new rectangle.
func (p Point[S]) AlignCenter(r Rect[S]) Rect[S] {
	return p.Align(r, 0.5, 0.5)
}

// CutX cuts r into two rectangles at x = r.Min.X + w.
// It returns the left part (got) and the right part (rest).
// If w is negative, got is empty and rest is r.
// If w is larger than r.Dx(), got is r and rest is empty.
func (r Rect[S]) CutX(w S) (got, rest Rect[S]) {
	w = min(max(w, 0), r.Dx())
	return Rect[S]{
			Min: r.Min,
			Max: Point[S]{X: r.Min.X + w, Y: r.Max.Y},
		}, Rect[S]{
			Min: Point[S]{X: r.Min.X + w, Y: r.Min.Y},
			Max: r.Max,
		}
}

// CutY cuts r into two rectangles at y = r.Min.Y + h.
// It returns the top part (got) and the bottom part (rest).
// If h is negative, got is empty and rest is r.
// If h is larger than r.Dy(), got is r and rest is empty.
func (r Rect[S]) CutY(h S) (got, rest Rect[S]) {
	h = min(max(h, 0), r.Dy())
	return Rect[S]{
			Min: r.Min,
			Max: Point[S]{X: r.Max.X, Y: r.Min.Y + h},
		}, Rect[S]{
			Min: Point[S]{X: r.Min.X, Y: r.Min.Y + h},
			Max: r.Max,
		}
}

// CutXRate cuts the rectangle by a rate of its width.
// It returns two rectangles: the cut part and the rest.
// If rate < 0, the cut part has zero width.
// If rate > 1, the cut part has the original width.
func (r Rect[S]) CutXRate(rate float64) (Rect[S], Rect[S]) {
	return r.CutX(rel(r.Dx(), rate))
}

// CutYRate cuts the rectangle by a rate of its height.
// It returns two rectangles: the cut part and the rest.
// If rate < 0, the cut part has zero height.
// If rate > 1, the cut part has the original height.
func (r Rect[S]) CutYRate(rate float64) (Rect[S], Rect[S]) {
	return r.CutY(rel(r.Dy(), rate))
}

// SplitX splits r into n rectangles of (mostly) equal width, with a specified gap between them.
// If n <= 0, returns nil. If n == 1, returns r.
// Gap is space between items. Assumed to be non-negative; negative gap leads to overlap.
func (r Rect[S]) SplitX(n int, gap S) []Rect[S] {
	if n <= 0 {
		return nil
	}

	rects := make([]Rect[S], n)
	if n == 1 {
		rects[0] = r
		return rects
	}

	if gap < 0 { // Treat negative gap as no gap for this calculation, or could be an error.
		gap = 0
	}

	// Effective width available for items (after subtracting all gaps)
	totalItemWidth := r.Dx() - (S(n-1) * gap)
	if totalItemWidth < 0 {
		totalItemWidth = 0 // Items will have zero width if gaps are too large
	}

	// Width for each of the n items, if distributed equally
	// This is the width we'll assign to the first n-1 items
	singleItemWidth := totalItemWidth / S(n)
	if singleItemWidth < 0 { // Should be prevented by totalItemWidth >= 0
		singleItemWidth = 0
	}

	currentPosX := r.Min.X
	for i := range n - 1 {
		// Create rectangle for the current item
		rects[i] = Xywh(currentPosX, r.Min.Y, singleItemWidth, r.Dy())
		// Move cursor to the start of the next item
		currentPosX += singleItemWidth + gap
	}
	// Last rectangle takes all remaining space from currentPosX to r.Max.X
	lastItemWidth := r.Max.X - currentPosX
	if lastItemWidth < 0 {
		lastItemWidth = 0 // Ensure the last item's width is not negative
	}
	rects[n-1] = Xywh(currentPosX, r.Min.Y, lastItemWidth, r.Dy())

	return rects
}

// SplitY splits r into n rectangles of (mostly) equal height, with a specified gap between them.
// If n <= 0, returns nil. If n == 1, returns r.
// Gap is space between items. Assumed to be non-negative; negative gap leads to overlap.
func (r Rect[S]) SplitY(n int, gap S) []Rect[S] {
	if n <= 0 {
		return nil
	}

	rects := make([]Rect[S], n)
	if n == 1 {
		rects[0] = r
		return rects
	}

	if gap < 0 { // Treat negative gap as no gap for this calculation
		gap = 0
	}

	// Effective height available for items (after subtracting all gaps)
	totalItemHeight := r.Dy() - (S(n-1) * gap)
	if totalItemHeight < 0 {
		totalItemHeight = 0 // Items will have zero height if gaps are too large
	}

	// Height for each of the n items, if distributed equally
	// This is the height we'll assign to the first n-1 items
	singleItemHeight := totalItemHeight / S(n)
	if singleItemHeight < 0 { // Should be prevented by totalItemHeight >= 0
		singleItemHeight = 0
	}

	currentPosY := r.Min.Y
	for i := range n - 1 {
		// Create rectangle for the current item
		rects[i] = Xywh(r.Min.X, currentPosY, r.Dx(), singleItemHeight)
		// Move cursor to the start of the next item
		currentPosY += singleItemHeight + gap
	}
	// Last rectangle takes all remaining space from currentPosY to r.Max.Y
	lastItemHeight := r.Max.Y - currentPosY
	if lastItemHeight < 0 {
		lastItemHeight = 0 // Ensure the last item's height is not negative
	}
	rects[n-1] = Xywh(r.Min.X, currentPosY, r.Dx(), lastItemHeight)

	return rects
}

// RepeatX repeats the rectangle n times in the X direction with a given gap.
// It returns a slice of the repeated rectangles and the bounding box of all repeated rectangles.
// If n <= 0, it returns nil and an empty rectangle.
func (r Rect[S]) RepeatX(n int, gap S) ([]Rect[S], Rect[S]) {
	if n <= 0 {
		return nil, Rect[S]{}
	}
	rects := make([]Rect[S], n)
	dx := r.Dx()
	currentX := r.Min.X
	for i := range n {
		rects[i] = Xywh(currentX, r.Min.Y, dx, r.Dy())
		currentX += dx + gap
	}
	overallRect := Xyxy(r.Min.X, r.Min.Y, currentX-gap, r.Max.Y)
	return rects, overallRect
}

// RepeatY repeats the rectangle n times in the Y direction with a given gap.
// It returns a slice of the repeated rectangles and the bounding box of all repeated rectangles.
// If n <= 0, it returns nil and an empty rectangle.
func (r Rect[S]) RepeatY(n int, gap S) ([]Rect[S], Rect[S]) {
	if n <= 0 {
		return nil, Rect[S]{}
	}
	rects := make([]Rect[S], n)
	dy := r.Dy()
	currentY := r.Min.Y
	for i := range n {
		rects[i] = Xywh(r.Min.X, currentY, r.Dx(), dy)
		currentY += dy + gap
	}
	overallRect := Xyxy(r.Min.X, r.Min.Y, r.Max.X, currentY-gap)
	return rects, overallRect
}
