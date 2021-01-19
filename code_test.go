package advent20201220_test

import (
	advent "advent20201220"
	"fmt"
	"testing"
)

func TestTilesFromFile(t *testing.T) {
	sampleLen := len(advent.TilesFromFile("sample.txt"))
	if sampleLen != 9 {
		t.Errorf("Got the wrong number of tiles: %v", sampleLen)
	}
	inputLen := len(advent.TilesFromFile("input.txt"))
	if inputLen != 144 {
		t.Errorf("Got the wrong number of tiles: %v", inputLen)
	}
}

func TestCanonical(t *testing.T) {
	tcs := []struct {
		input, want advent.Edge
	}{
		{input: advent.Edge("01"), want: advent.Edge("10")},
		{input: advent.Edge("10"), want: advent.Edge("10")},
	}
	for _, tc := range tcs {
		got := tc.input.AsCanonical()
		if tc.want != got {
			t.Error("Got canonicity wrong")
		}
	}
}

func TestDifficultEdgeMatches(t *testing.T) {
	tiles := advent.TilesFromFile("sample.txt")
	tiles.DifficultEdgeMatches()
	// advent.DifficultEdgeMatches("input.txt")
	// t.Fail()
}

func TestPart1(t *testing.T) {
	tiles := advent.TilesFromFile("input.txt")
	product := 1
	for _, tile := range tiles.Corners() {
		product *= tile
	}
	if product != 66020135789767 {
		t.Fail()
	}
}

func TestRotate90(t *testing.T) {
	tile := advent.TilesFromFile("input.txt").First()

	fmt.Println(tile.ToString())
	r90 := tile.Rotate90()
	fmt.Println(r90.ToString())
	r180 := r90.Rotate90()
	fmt.Println(r180.ToString())
	// t.Fail()
}

func TestFlip(t *testing.T) {
	tile := advent.TilesFromFile("input.txt").First()

	fmt.Println(tile.ToString())
	flip := tile.Flip()
	fmt.Println(flip.ToString())
	dup := flip.Flip()
	if dup.ToString() != tile.ToString() || flip.ToString() == tile.ToString() {
		t.Fail()
	}
}

// func TestDownRight(t *testing.T) {
// 	tiles := advent.TilesFromFile("input.txt")
// 	puzzle := advent.NewPuzzle(tiles)
// 	// t.Fail()
// }

func TestBuild(t *testing.T) {
	tiles := advent.TilesFromFile("input.txt")
	puzzle := advent.NewPuzzle(tiles)
	puzzle.Build()
	fmt.Println(puzzle.RawPrint())
	// t.Fail()
}

func TestBorderlessPrint(t *testing.T) {
	tiles := advent.TilesFromFile("input.txt")
	puzzle := advent.NewPuzzle(tiles)
	puzzle.Build()
	if puzzle.BorderlessPrint() == puzzle.ReconstructedImage().ToString() {
		t.Skip()
	}
	t.Fail()
}

func TestAnswer(t *testing.T) {
	tiles := advent.TilesFromFile("input.txt")
	puzzle := advent.NewPuzzle(tiles)
	puzzle.Build()
	image := puzzle.ReconstructedImage()
	// fmt.Println(puzzle.RawPrint())
	for i := 0; i < 9; i++ {
		// fmt.Println(image.ToString())
		if i == 4 {
			image = image.Flip()
		}
		hasMonsters, mask := image.SeaMonsterMask()

		if hasMonsters {
			if image.WaterRoughness(mask) == 1537 {
				t.Skip()
			}
		}
		image = image.Rotate90()
	}
	t.Fail()
}

func TestHasMonsters(t *testing.T) {
	tile := advent.TilesFromFile("sample_monster.txt")[1111]
	image := tile
	for i := 0; i < 9; i++ {
		if i == 4 {
			image = image.Flip()
		}
		hasMonsters, _ := image.SeaMonsterMask()
		if hasMonsters {
			t.Skip()
		}
		image = image.Rotate90()
	}
	t.Fail()
}

// func TestPart2(t *testing.T) {
// 	fmt.Println(advent.Part2("sample.txt"))
// 	fmt.Println(advent.Part2("input.txt"))
// 	t.Fail()
// }
