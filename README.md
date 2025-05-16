# loc - Locate and Align Points and Rectangles

`loc` is a Go library providing types and functions for 2D geometry, focusing on points (`Point[S]`) and rectangles (`Rect[S]`).

## Overview

The library offers:
- `Point[S]`: Represents an X, Y coordinate.
- `Rect[S]`: Represents a rectangle defined by `Min` and `Max` points.
- Utility functions for creating points and rectangles (e.g., `Xy`, `Xyxy`, `Xywh`).
- Methods for geometric operations:
    - Point arithmetic (`Add`, `Sub`, `Mul`, `Div`).
    - Rectangle manipulation (`Add`, `Sub`, `Inset`, `Intersect`, `Union`).
    - Alignment (`Point.Align`, `Point.AlignCenter`).
    - Cutting and Splitting (`Rect.CutX`, `Rect.CutY`, `Rect.CutXByRate`, `Rect.CutYByRate`, `Rect.SplitX`, `Rect.SplitY`).
    - Repeating (`Rect.RepeatX`, `Rect.RepeatY`).
    - Anchoring (`Rect.Anchor`).

## Examples

### Basic Point and Rectangle Creation

```go
package main

import (
	"fmt"
	"github.com/eihigh/loc"
)

func main() {
	// Create points
	p1 := loc.Xy(10, 20)
	p2 := loc.Xy(30, 40)
	fmt.Println("Point 1:", p1)
	fmt.Println("Point 2:", p2)

	// Create a rectangle from two points
	rect1 := loc.Xyxy(p1.X, p1.Y, p2.X, p2.Y) // (10,20)-(30,40)
	fmt.Println("Rectangle 1:", rect1)

	// Create a rectangle from position and size
	rect2 := loc.Xywh(0, 0, 100, 50) // (0,0)-(100,50)
	fmt.Println("Rectangle 2:", rect2)
}
```

### Aligning a Rectangle within Another

This example demonstrates aligning a smaller `modal` rectangle to the center of a larger `screen` rectangle.

```go
package main

import (
	"fmt"
	"github.com/eihigh/loc"
)

func main() {
	screen := loc.Xyxy(0, 0, 800, 600)
	modal := loc.Xyxy(0, 0, 400, 300)

	// Align 'modal' to the center of 'screen'
	// screen.Center() returns the center Point of the screen.
	// AlignCenter aligns the center of 'modal' to this point.
	alignedModal := screen.Center().AlignCenter(modal)
	fmt.Println("Screen:", screen)
	fmt.Println("Modal (original):", modal)
	fmt.Println("Modal (aligned):", alignedModal)
	// Output: Modal (aligned): (200,150)-(600,450)
}
```

### Cutting a Rectangle

This example shows how to cut a rectangle horizontally by an absolute value and by a rate.

```go
package main

import (
	"fmt"
	"github.com/eihigh/loc"
)

func main() {
	rect := loc.Xyxy(0, 0, 100, 50)

	// Cut by absolute value
	cutAbs, restAbs := rect.CutX(30)
	fmt.Printf("Cut by 30: %s, Rest: %s\n", cutAbs, restAbs)
	// Output: Cut by 30: (0,0)-(30,50), Rest: (30,0)-(100,50)

	// Cut by rate (30% of width)
	cutRate, restRate := rect.CutXByRate(0.3)
	fmt.Printf("Cut by 0.3 rate: %s, Rest: %s\n", cutRate, restRate)
	// Output: Cut by 0.3 rate: (0,0)-(30,50), Rest: (30,0)-(100,50)
}
```

### Repeating a Rectangle

This example demonstrates repeating a rectangle horizontally.

```go
package main

import (
	"fmt"
	"github.com/eihigh/loc"
)

func main() {
	baseRect := loc.Xyxy(0, 0, 20, 30)

	// Repeat 'baseRect' 3 times in X direction with a gap of 5
	repeatedRects, overallRect := baseRect.RepeatX(3, 5)

	fmt.Println("Overall Bounding Box:", overallRect)
	for i, r := range repeatedRects {
		fmt.Printf("Repeated Rect %d: %s\n", i, r)
	}
	// Output:
	// Overall Bounding Box: (0,0)-(70,30)
	// Repeated Rect 0: (0,0)-(20,30)
	// Repeated Rect 1: (25,0)-(45,30)
	// Repeated Rect 2: (50,0)-(70,30)
}
```

For more detailed examples, please refer to the `example_test.go` file.
