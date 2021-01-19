package advent20201220

import (
	"fmt"
	"math"
)

type Index struct {
	colIdx int
	rowIdx int
}

type Tile map[Index]int

func (t Tile) Boundaries() (minRowIdx int, maxRowIdx int, minColIdx int, maxColIdx int) {
	for idx := range t {
		minRowIdx = idx.rowIdx
		maxRowIdx = idx.rowIdx
		minColIdx = idx.colIdx
		maxColIdx = idx.colIdx
		break
	}

	for idx := range t {
		minRowIdx = int(math.Min(float64(idx.rowIdx), float64(minRowIdx)))
		maxRowIdx = int(math.Max(float64(idx.rowIdx), float64(maxRowIdx)))
		minColIdx = int(math.Min(float64(idx.colIdx), float64(minColIdx)))
		maxColIdx = int(math.Max(float64(idx.colIdx), float64(maxColIdx)))
	}
	return minRowIdx, maxRowIdx, minColIdx, maxColIdx
}

func (t Tile) ToString() string {
	minRowIdx, maxRowIdx, minColIdx, maxColIdx := t.Boundaries()
	repr := ""
	for row := minRowIdx; row <= maxRowIdx; row++ {
		for col := minColIdx; col <= maxColIdx; col++ {
			repr += fmt.Sprint(t[Index{colIdx: col, rowIdx: row}])
		}
		repr += fmt.Sprintln()
	}
	return repr
}

func (t Tile) Top() Edge {
	result := ""
	for i := 0; i < 10; i++ {
		result += fmt.Sprint(t[Index{colIdx: i, rowIdx: 0}])
	}
	return Edge(result)
}

func (t Tile) Bottom() Edge {
	result := ""
	for i := 0; i < 10; i++ {
		result += fmt.Sprint(t[Index{colIdx: i, rowIdx: 9}])
	}
	return Edge(result)
}

func (t Tile) Left() Edge {
	result := ""
	for j := 0; j < 10; j++ {
		result += fmt.Sprint(t[Index{colIdx: 0, rowIdx: j}])
	}
	return Edge(result)
}

func (t Tile) Right() Edge {
	result := ""
	for j := 0; j < 10; j++ {
		result += fmt.Sprint(t[Index{colIdx: 9, rowIdx: j}])
	}
	return Edge(result)
}

func (t Tile) CanonicalRight() Edge {
	return t.Right().AsCanonical()
}
func (t Tile) CanonicalLeft() Edge {
	return t.Left().AsCanonical()
}
func (t Tile) CanonicalTop() Edge {
	return t.Top().AsCanonical()
}
func (t Tile) CanonicalBottom() Edge {
	return t.Bottom().AsCanonical()
}

func (t Tile) CanonicalEdges() []Edge {
	return []Edge{t.CanonicalBottom(), t.CanonicalLeft(), t.CanonicalTop(), t.CanonicalRight()}
}

func (t Tile) Rotate90() Tile {
	newTile := make(Tile)
	minRowIdx, maxRowIdx, minColIdx, maxColIdx := t.Boundaries()
	for col := minColIdx; col <= maxColIdx; col++ {
		for row := minRowIdx; row <= maxRowIdx; row++ {
			newTile[Index{colIdx: maxColIdx - row, rowIdx: col}] = t[Index{rowIdx: row, colIdx: col}]
		}
	}
	return newTile
}

func (t Tile) Flip() Tile {
	newTile := make(Tile)
	minRowIdx, maxRowIdx, minColIdx, maxColIdx := t.Boundaries()
	for row := minRowIdx; row <= maxRowIdx; row++ {
		for col := minColIdx; col <= maxColIdx; col++ {
			newTile[Index{colIdx: row, rowIdx: col}] = t[Index{rowIdx: row, colIdx: col}]
		}
	}
	return newTile
}

// SeaMonsterMask gets which indices are part of a sea monster, and returns a boolean for if it found any.
func (t Tile) SeaMonsterMask() (bool, map[Index]int) {
	seaMonsterMask := make(map[Index]int)
	minRowIdx, maxRowIdx, minColIdx, maxColIdx := t.Boundaries()
	hasSeaMonsters := false
	for row := minRowIdx; row <= maxRowIdx; row++ {
		for col := minColIdx; col <= maxColIdx; col++ {

			if (t[Index{rowIdx: row, colIdx: col}] == 1 &&
				t[Index{rowIdx: row + 1, colIdx: col + 1}] == 1 &&
				t[Index{rowIdx: row + 1, colIdx: col + 4}] == 1 &&
				t[Index{rowIdx: row + 0, colIdx: col + 5}] == 1 &&
				t[Index{rowIdx: row + 0, colIdx: col + 6}] == 1 &&
				t[Index{rowIdx: row + 1, colIdx: col + 7}] == 1 &&
				t[Index{rowIdx: row + 1, colIdx: col + 10}] == 1 &&
				t[Index{rowIdx: row + 0, colIdx: col + 11}] == 1 &&
				t[Index{rowIdx: row + 0, colIdx: col + 12}] == 1 &&
				t[Index{rowIdx: row + 1, colIdx: col + 13}] == 1 &&
				t[Index{rowIdx: row + 1, colIdx: col + 16}] == 1 &&
				t[Index{rowIdx: row + 0, colIdx: col + 17}] == 1 &&
				t[Index{rowIdx: row + 0, colIdx: col + 18}] == 1 &&
				t[Index{rowIdx: row + 0, colIdx: col + 19}] == 1 &&
				t[Index{rowIdx: row - 1, colIdx: col + 18}] == 1) {
				hasSeaMonsters = true
				seaMonsterMask[Index{rowIdx: row, colIdx: col}] = 1
				seaMonsterMask[Index{rowIdx: row + 1, colIdx: col + 1}] = 1
				seaMonsterMask[Index{rowIdx: row + 1, colIdx: col + 4}] = 1
				seaMonsterMask[Index{rowIdx: row + 0, colIdx: col + 5}] = 1
				seaMonsterMask[Index{rowIdx: row + 0, colIdx: col + 6}] = 1
				seaMonsterMask[Index{rowIdx: row + 1, colIdx: col + 7}] = 1
				seaMonsterMask[Index{rowIdx: row + 1, colIdx: col + 10}] = 1
				seaMonsterMask[Index{rowIdx: row + 0, colIdx: col + 11}] = 1
				seaMonsterMask[Index{rowIdx: row + 0, colIdx: col + 12}] = 1
				seaMonsterMask[Index{rowIdx: row + 1, colIdx: col + 13}] = 1
				seaMonsterMask[Index{rowIdx: row + 1, colIdx: col + 16}] = 1
				seaMonsterMask[Index{rowIdx: row + 0, colIdx: col + 17}] = 1
				seaMonsterMask[Index{rowIdx: row + 0, colIdx: col + 18}] = 1
				seaMonsterMask[Index{rowIdx: row + 0, colIdx: col + 19}] = 1
				seaMonsterMask[Index{rowIdx: row - 1, colIdx: col + 18}] = 1
			}
		}
	}
	return hasSeaMonsters, seaMonsterMask
}

func (t Tile) WaterRoughness(mask Tile) (total int) {
	minRowIdx, maxRowIdx, minColIdx, maxColIdx := t.Boundaries()
	for row := minRowIdx; row <= maxRowIdx; row++ {
		for col := minColIdx; col <= maxColIdx; col++ {
			if mask[Index{rowIdx: row, colIdx: col}] == 0 {
				total += t[Index{rowIdx: row, colIdx: col}]
			}
		}
	}
	return
}
