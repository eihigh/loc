package loc_test

import (
	"fmt"

	"github.com/eihigh/loc"
)

func ExamplePoint_AlignCenter() {
	screen := loc.Xyxy(0, 0, 800, 600)
	modal := loc.Xyxy(0, 0, 400, 300)
	modal = screen.Center().AlignCenter(modal)
	fmt.Println(modal.String())

	// Output:
	// (200,150)-(600,450)
}

func ExampleRect_Anchor() {
	rect := loc.Xyxy(10, 20, 110, 120) // Dx=100, Dy=100
	p1 := rect.Anchor(0, 0)
	p2 := rect.Anchor(0.5, 0.5)
	p3 := rect.Anchor(1, 1)
	p4 := rect.Anchor(0.25, 0.75)
	fmt.Println(p1, p2, p3, p4)

	// Output:
	// (10,20) (60,70) (110,120) (35,95)
}

func ExampleRect_Center() {
	rect := loc.Xyxy(10, 20, 110, 220) // Dx=100, Dy=200
	center := rect.Center()
	fmt.Println(center)

	// Output:
	// (60,120)
}

func ExamplePoint_Align() {
	screen := loc.Xyxy(0, 0, 800, 600)
	box := loc.Xyxy(0, 0, 100, 100)

	fmt.Println("top left:", screen.Anchor(0, 0).Align(box, 0, 0))
	fmt.Println("top right:", screen.Anchor(1, 0).Align(box, 1, 0))
	fmt.Println("bottom left:", screen.Anchor(0, 1).Align(box, 0, 1))
	fmt.Println("bottom right:", screen.Anchor(1, 1).Align(box, 1, 1))

	// Output:
	// top left: (0,0)-(100,100)
	// top right: (700,0)-(800,100)
	// bottom left: (0,500)-(100,600)
	// bottom right: (700,500)-(800,600)
}

func ExampleRect_CutX() {
	rect := loc.Xyxy(0, 0, 100, 50)
	cut1, rest1 := rect.CutX(30)
	fmt.Printf("Cut1: %s, Rest1: %s\n", cut1, rest1)

	// Output:
	// Cut1: (0,0)-(30,50), Rest1: (30,0)-(100,50)
}

func ExampleRect_CutY() {
	rect := loc.Xyxy(0, 0, 100, 50)
	cut1, rest1 := rect.CutY(20)
	fmt.Printf("Cut1: %s, Rest1: %s\n", cut1, rest1)

	// Output:
	// Cut1: (0,0)-(100,20), Rest1: (0,20)-(100,50)
}

func ExampleRect_CutXRate() {
	rect := loc.Xyxy(0, 0, 100, 50)
	cut, rest := rect.CutXRate(0.3) // Cut 30% of width
	fmt.Printf("Cut: %s, Rest: %s\n", cut, rest)

	cutNegative, restNegative := rect.CutXRate(-0.1) // Cut -10% (should be 0)
	fmt.Printf("CutNegative: %s, RestNegative: %s\n", cutNegative, restNegative)

	cutOver, restOver := rect.CutXRate(1.2) // Cut 120% (should be 100%)
	fmt.Printf("CutOver: %s, RestOver: %s\n", cutOver, restOver)
	// Output:
	// Cut: (0,0)-(30,50), Rest: (30,0)-(100,50)
	// CutNegative: (0,0)-(0,50), RestNegative: (0,0)-(100,50)
	// CutOver: (0,0)-(100,50), RestOver: (100,0)-(100,50)
}

func ExampleRect_CutYRate() {
	rect := loc.Xyxy(0, 0, 100, 50)
	cut, rest := rect.CutYRate(0.4) // Cut 40% of height (20)
	fmt.Printf("Cut: %s, Rest: %s\n", cut, rest)

	cutNegative, restNegative := rect.CutYRate(-0.2) // Cut -20% (should be 0)
	fmt.Printf("CutNegative: %s, RestNegative: %s\n", cutNegative, restNegative)

	cutOver, restOver := rect.CutYRate(1.5) // Cut 150% (should be 100%)
	fmt.Printf("CutOver: %s, RestOver: %s\n", cutOver, restOver)
	// Output:
	// Cut: (0,0)-(100,20), Rest: (0,20)-(100,50)
	// CutNegative: (0,0)-(100,0), RestNegative: (0,0)-(100,50)
	// CutOver: (0,0)-(100,50), RestOver: (0,50)-(100,50)
}

func ExampleRect_SplitX() {
	rect := loc.Xyxy(0, 0, 100, 50)

	// Gap = 0
	parts3NoGap := rect.SplitX(3, 0)
	for i, p := range parts3NoGap {
		fmt.Printf("3 parts, no gap, part %d: %s\n", i, p)
	}

	parts4NoGap := rect.SplitX(4, 0) // 100 / 4 = 25
	for i, p := range parts4NoGap {
		fmt.Printf("4 parts, no gap, part %d: %s\n", i, p)
	}

	parts1NoGap := rect.SplitX(1, 0)
	fmt.Printf("1 part, no gap, part 0: %s\n", parts1NoGap[0])

	// With positive gap
	parts3WithGap5 := rect.SplitX(3, 5) // 100, n=3, gap=5. totalItemWidth = 100 - 2*5 = 90. singleItemWidth = 90/3 = 30.
	for i, p := range parts3WithGap5 {
		fmt.Printf("3 parts, gap 5, part %d: %s\n", i, p)
	}

	// Output:
	// 3 parts, no gap, part 0: (0,0)-(33,50)
	// 3 parts, no gap, part 1: (33,0)-(66,50)
	// 3 parts, no gap, part 2: (66,0)-(100,50)
	// 4 parts, no gap, part 0: (0,0)-(25,50)
	// 4 parts, no gap, part 1: (25,0)-(50,50)
	// 4 parts, no gap, part 2: (50,0)-(75,50)
	// 4 parts, no gap, part 3: (75,0)-(100,50)
	// 1 part, no gap, part 0: (0,0)-(100,50)
	// 3 parts, gap 5, part 0: (0,0)-(30,50)
	// 3 parts, gap 5, part 1: (35,0)-(65,50)
	// 3 parts, gap 5, part 2: (70,0)-(100,50)
}

func ExampleRect_SplitY() {
	rect := loc.Xyxy(0, 0, 50, 100)

	// Gap = 0
	parts3NoGap := rect.SplitY(3, 0)
	for i, p := range parts3NoGap {
		fmt.Printf("3 parts, no gap, part %d: %s\n", i, p)
	}

	parts4NoGap := rect.SplitY(4, 0) // 100 / 4 = 25
	for i, p := range parts4NoGap {
		fmt.Printf("4 parts, no gap, part %d: %s\n", i, p)
	}

	parts1NoGap := rect.SplitY(1, 0)
	fmt.Printf("1 part, no gap, part 0: %s\n", parts1NoGap[0])

	// With positive gap
	parts3WithGap5 := rect.SplitY(3, 5) // 100, n=3, gap=5. totalItemHeight = 100 - 2*5 = 90. singleItemHeight = 90/3 = 30.
	for i, p := range parts3WithGap5 {
		fmt.Printf("3 parts, gap 5, part %d: %s\n", i, p)
	}

	// Output:
	// 3 parts, no gap, part 0: (0,0)-(50,33)
	// 3 parts, no gap, part 1: (0,33)-(50,66)
	// 3 parts, no gap, part 2: (0,66)-(50,100)
	// 4 parts, no gap, part 0: (0,0)-(50,25)
	// 4 parts, no gap, part 1: (0,25)-(50,50)
	// 4 parts, no gap, part 2: (0,50)-(50,75)
	// 4 parts, no gap, part 3: (0,75)-(50,100)
	// 1 part, no gap, part 0: (0,0)-(50,100)
	// 3 parts, gap 5, part 0: (0,0)-(50,30)
	// 3 parts, gap 5, part 1: (0,35)-(50,65)
	// 3 parts, gap 5, part 2: (0,70)-(50,100)
}

func ExampleRect_RepeatX() {
	rect := loc.Xyxy(0, 0, 20, 30)
	repeatedRects, overallRect := rect.RepeatX(3, 5) // Repeat 3 times with gap 5

	fmt.Println("Overall:", overallRect)
	for i, r := range repeatedRects {
		fmt.Printf("Part %d: %s\n", i, r)
	}

	// Output:
	// Overall: (0,0)-(70,30)
	// Part 0: (0,0)-(20,30)
	// Part 1: (25,0)-(45,30)
	// Part 2: (50,0)-(70,30)
}

func ExampleRect_RepeatY() {
	rect := loc.Xyxy(0, 0, 20, 30)
	repeatedRects, overallRect := rect.RepeatY(2, 10) // Repeat 2 times with gap 10

	fmt.Println("Overall:", overallRect)
	for i, r := range repeatedRects {
		fmt.Printf("Part %d: %s\n", i, r)
	}

	// Output:
	// Overall: (0,0)-(20,70)
	// Part 0: (0,0)-(20,30)
	// Part 1: (0,40)-(20,70)
}
