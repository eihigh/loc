package loc_test

import (
	"testing"

	"github.com/eihigh/loc"
)

func TestRect_CutX_Over(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	gotCut, gotRest := rect.CutX(120)
	wantCut := loc.Xyxy(0, 0, 100, 50)
	wantRest := loc.Xyxy(100, 0, 100, 50)
	if !wantCut.Eq(gotCut) {
		t.Errorf("CutX(120) cut mismatch, want %v, got %v", wantCut, gotCut)
	}
	if !wantRest.Eq(gotRest) {
		t.Errorf("CutX(120) rest mismatch, want %v, got %v", wantRest, gotRest)
	}
}

func TestRect_CutX_Negative(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	gotCut, gotRest := rect.CutX(-10)
	wantCut := loc.Xyxy(0, 0, 0, 50)
	wantRest := loc.Xyxy(0, 0, 100, 50)
	if !wantCut.Eq(gotCut) {
		t.Errorf("CutX(-10) cut mismatch, want %v, got %v", wantCut, gotCut)
	}
	if !wantRest.Eq(gotRest) {
		t.Errorf("CutX(-10) rest mismatch, want %v, got %v", wantRest, gotRest)
	}
}

func TestRect_CutY_Over(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	gotCut, gotRest := rect.CutY(70)
	wantCut := loc.Xyxy(0, 0, 100, 50)
	wantRest := loc.Xyxy(0, 50, 100, 50)
	if !wantCut.Eq(gotCut) {
		t.Errorf("CutY(70) cut mismatch, want %v, got %v", wantCut, gotCut)
	}
	if !wantRest.Eq(gotRest) {
		t.Errorf("CutY(70) rest mismatch, want %v, got %v", wantRest, gotRest)
	}
	if !gotRest.Empty() {
		t.Errorf("CutY(70) rest should be empty, but got %s", gotRest.String())
	}
}

func TestRect_CutY_Negative(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	gotCut, gotRest := rect.CutY(-10)
	wantCut := loc.Xyxy(0, 0, 100, 0)
	wantRest := loc.Xyxy(0, 0, 100, 50)
	if !wantCut.Eq(gotCut) {
		t.Errorf("CutY(-10) cut mismatch, want %v, got %v", wantCut, gotCut)
	}
	if !wantRest.Eq(gotRest) {
		t.Errorf("CutY(-10) rest mismatch, want %v, got %v", wantRest, gotRest)
	}
}

func TestRect_SplitX_Zero(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	got := rect.SplitX(0, 0)
	if got != nil {
		t.Errorf("SplitX(0, 0) should return nil, but got %v", got)
	}
}

func TestRect_SplitX_Negative(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	got := rect.SplitX(-2, 0)
	if got != nil {
		t.Errorf("SplitX(-2, 0) should return nil, but got %v", got)
	}
}

func TestRect_SplitX_LargeGap(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	got := rect.SplitX(3, 60) // 100, n=3, gap=60. totalItemWidth = 100 - 2*60 = -20 -> 0. singleItemWidth = 0.
	want := []loc.Rect[int]{
		loc.Xyxy(0, 0, 0, 50),
		loc.Xyxy(60, 0, 60, 50),
		loc.Xyxy(120, 0, 120, 50),
	}
	if len(want) != len(got) {
		t.Errorf("SplitX(3, 60) length mismatch, want %d, got %d", len(want), len(got))
	} else {
		for i := range want {
			if !want[i].Eq(got[i]) {
				t.Errorf("SplitX(3, 60) element %d mismatch, want %v, got %v", i, want[i], got[i])
			}
		}
	}
	// For coverage, print the parts (original example output)
	for i, p := range got {
		t.Logf("SplitX large gap part %d: %s\n", i, p)
	}
}

func TestRect_CutXRate_Normal(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	gotCut, gotRest := rect.CutXRate(0.3)
	wantCut := loc.Xyxy(0, 0, 30, 50)
	wantRest := loc.Xyxy(30, 0, 100, 50)
	if !wantCut.Eq(gotCut) {
		t.Errorf("CutXRate(0.3) cut mismatch, want %v, got %v", wantCut, gotCut)
	}
	if !wantRest.Eq(gotRest) {
		t.Errorf("CutXRate(0.3) rest mismatch, want %v, got %v", wantRest, gotRest)
	}
}

func TestRect_CutXRate_Negative(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	gotCut, gotRest := rect.CutXRate(-0.1)
	wantCut := loc.Xyxy(0, 0, 0, 50)
	wantRest := loc.Xyxy(0, 0, 100, 50)
	if !wantCut.Eq(gotCut) {
		t.Errorf("CutXRate(-0.1) cut mismatch, want %v, got %v", wantCut, gotCut)
	}
	if !wantRest.Eq(gotRest) {
		t.Errorf("CutXRate(-0.1) rest mismatch, want %v, got %v", wantRest, gotRest)
	}
}

func TestRect_CutXRate_Over(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	gotCut, gotRest := rect.CutXRate(1.2)
	wantCut := loc.Xyxy(0, 0, 100, 50)
	wantRest := loc.Xyxy(100, 0, 100, 50)
	if !wantCut.Eq(gotCut) {
		t.Errorf("CutXRate(1.2) cut mismatch, want %v, got %v", wantCut, gotCut)
	}
	if !wantRest.Eq(gotRest) {
		t.Errorf("CutXRate(1.2) rest mismatch, want %v, got %v", wantRest, gotRest)
	}
}

func TestRect_CutYRate_Normal(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	gotCut, gotRest := rect.CutYRate(0.4)
	wantCut := loc.Xyxy(0, 0, 100, 20)
	wantRest := loc.Xyxy(0, 20, 100, 50)
	if !wantCut.Eq(gotCut) {
		t.Errorf("CutYRate(0.4) cut mismatch, want %v, got %v", wantCut, gotCut)
	}
	if !wantRest.Eq(gotRest) {
		t.Errorf("CutYRate(0.4) rest mismatch, want %v, got %v", wantRest, gotRest)
	}
}

func TestRect_CutYRate_Negative(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	gotCut, gotRest := rect.CutYRate(-0.2)
	wantCut := loc.Xyxy(0, 0, 100, 0)
	wantRest := loc.Xyxy(0, 0, 100, 50)
	if !wantCut.Eq(gotCut) {
		t.Errorf("CutYRate(-0.2) cut mismatch, want %v, got %v", wantCut, gotCut)
	}
	if !wantRest.Eq(gotRest) {
		t.Errorf("CutYRate(-0.2) rest mismatch, want %v, got %v", wantRest, gotRest)
	}
}

func TestRect_CutYRate_Over(t *testing.T) {
	rect := loc.Xyxy(0, 0, 100, 50)
	gotCut, gotRest := rect.CutYRate(1.5)
	wantCut := loc.Xyxy(0, 0, 100, 50)
	wantRest := loc.Xyxy(0, 50, 100, 50)
	if !wantCut.Eq(gotCut) {
		t.Errorf("CutYRate(1.5) cut mismatch, want %v, got %v", wantCut, gotCut)
	}
	if !wantRest.Eq(gotRest) {
		t.Errorf("CutYRate(1.5) rest mismatch, want %v, got %v", wantRest, gotRest)
	}
}

func TestRect_SplitY_Zero(t *testing.T) {
	rect := loc.Xyxy(0, 0, 50, 100)
	got := rect.SplitY(0, 0)
	if got != nil {
		t.Errorf("SplitY(0, 0) should return nil, but got %v", got)
	}
}

func TestRect_SplitY_Negative(t *testing.T) {
	rect := loc.Xyxy(0, 0, 50, 100)
	got := rect.SplitY(-2, 0)
	if got != nil {
		t.Errorf("SplitY(-2, 0) should return nil, but got %v", got)
	}
}

func TestRect_SplitY_LargeGap(t *testing.T) {
	rect := loc.Xyxy(0, 0, 50, 100)
	got := rect.SplitY(3, 60) // 100, n=3, gap=60. totalItemHeight = 100 - 2*60 = -20 -> 0. singleItemHeight = 0.
	want := []loc.Rect[int]{
		loc.Xyxy(0, 0, 50, 0),
		loc.Xyxy(0, 60, 50, 60),
		loc.Xyxy(0, 120, 50, 120),
	}
	if len(want) != len(got) {
		t.Errorf("SplitY(3, 60) length mismatch, want %d, got %d", len(want), len(got))
	} else {
		for i := range want {
			if !want[i].Eq(got[i]) {
				t.Errorf("SplitY(3, 60) element %d mismatch, want %v, got %v", i, want[i], got[i])
			}
		}
	}
	// For coverage, print the parts (original example output)
	for i, p := range got {
		t.Logf("SplitY large gap part %d: %s\n", i, p)
	}
}

func TestRect_RepeatX_Normal(t *testing.T) {
	rect := loc.Xyxy(0, 0, 20, 30)
	gotRects, gotOverall := rect.RepeatX(3, 5)
	wantRects := []loc.Rect[int]{
		loc.Xyxy(0, 0, 20, 30),
		loc.Xyxy(25, 0, 45, 30),
		loc.Xyxy(50, 0, 70, 30),
	}
	wantOverall := loc.Xyxy(0, 0, 70, 30)

	if len(wantRects) != len(gotRects) {
		t.Errorf("RepeatX(3, 5) rects length mismatch, want %d, got %d", len(wantRects), len(gotRects))
	} else {
		for i := range wantRects {
			if !wantRects[i].Eq(gotRects[i]) {
				t.Errorf("RepeatX(3, 5) rect %d mismatch, want %v, got %v", i, wantRects[i], gotRects[i])
			}
		}
	}
	if !wantOverall.Eq(gotOverall) {
		t.Errorf("RepeatX(3, 5) overall mismatch, want %v, got %v", wantOverall, gotOverall)
	}
}

func TestRect_RepeatX_ZeroN(t *testing.T) {
	rect := loc.Xyxy(0, 0, 20, 30)
	gotRects, gotOverall := rect.RepeatX(0, 5)
	if gotRects != nil {
		t.Errorf("RepeatX(0, 5) rects should be nil, got %v", gotRects)
	}
	if !gotOverall.Empty() {
		t.Errorf("RepeatX(0, 5) overall should be empty, got %v", gotOverall)
	}
}

func TestRect_RepeatX_NegativeN(t *testing.T) {
	rect := loc.Xyxy(0, 0, 20, 30)
	gotRects, gotOverall := rect.RepeatX(-2, 5)
	if gotRects != nil {
		t.Errorf("RepeatX(-2, 5) rects should be nil, got %v", gotRects)
	}
	if !gotOverall.Empty() {
		t.Errorf("RepeatX(-2, 5) overall should be empty, got %v", gotOverall)
	}
}

func TestRect_RepeatY_Normal(t *testing.T) {
	rect := loc.Xyxy(0, 0, 20, 30)
	gotRects, gotOverall := rect.RepeatY(2, 10)
	wantRects := []loc.Rect[int]{
		loc.Xyxy(0, 0, 20, 30),
		loc.Xyxy(0, 40, 20, 70),
	}
	wantOverall := loc.Xyxy(0, 0, 20, 70)

	if len(wantRects) != len(gotRects) {
		t.Errorf("RepeatY(2, 10) rects length mismatch, want %d, got %d", len(wantRects), len(gotRects))
	} else {
		for i := range wantRects {
			if !wantRects[i].Eq(gotRects[i]) {
				t.Errorf("RepeatY(2, 10) rect %d mismatch, want %v, got %v", i, wantRects[i], gotRects[i])
			}
		}
	}
	if !wantOverall.Eq(gotOverall) {
		t.Errorf("RepeatY(2, 10) overall mismatch, want %v, got %v", wantOverall, gotOverall)
	}
}

func TestRect_RepeatY_ZeroN(t *testing.T) {
	rect := loc.Xyxy(0, 0, 20, 30)
	gotRects, gotOverall := rect.RepeatY(0, 10)
	if gotRects != nil {
		t.Errorf("RepeatY(0, 10) rects should be nil, got %v", gotRects)
	}
	if !gotOverall.Empty() {
		t.Errorf("RepeatY(0, 10) overall should be empty, got %v", gotOverall)
	}
}

func TestRect_RepeatY_NegativeN(t *testing.T) {
	rect := loc.Xyxy(0, 0, 20, 30)
	gotRects, gotOverall := rect.RepeatY(-2, 10)
	if gotRects != nil {
		t.Errorf("RepeatY(-2, 10) rects should be nil, got %v", gotRects)
	}
	if !gotOverall.Empty() {
		t.Errorf("RepeatY(-2, 10) overall should be empty, got %v", gotOverall)
	}
}
