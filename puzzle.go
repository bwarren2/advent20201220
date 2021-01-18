package advent20201220

import (
	"fmt"
	"math"
)

type Puzzle struct {
	Tiles         TileMap
	Solution      map[Index]Tile
	UsedTiles     []Tile
	UsedIDs       []int
	EdgeNeighbors map[Edge][]int
}

func (p Puzzle) SolutionBoundaries() (minRowIdx int, maxRowIdx int, minColIdx int, maxColIdx int) {
	for idx := range p.Solution {
		minRowIdx = idx.rowIdx
		maxRowIdx = idx.rowIdx
		minColIdx = idx.colIdx
		maxColIdx = idx.colIdx
		break
	}

	for idx := range p.Solution {
		minRowIdx = int(math.Min(float64(idx.rowIdx), float64(minRowIdx)))
		maxRowIdx = int(math.Max(float64(idx.rowIdx), float64(maxRowIdx)))
		minColIdx = int(math.Min(float64(idx.colIdx), float64(minColIdx)))
		maxColIdx = int(math.Max(float64(idx.colIdx), float64(maxColIdx)))
	}
	return minRowIdx, maxRowIdx, minColIdx, maxColIdx
}

func (p Puzzle) RawPrint() string {
	minRowIdx, maxRowIdx, minColIdx, maxColIdx := p.SolutionBoundaries()
	repr := ""
	for row := minRowIdx; row <= maxRowIdx; row++ {
		for i := 0; i < 10; i++ {
			for col := minColIdx; col <= maxColIdx; col++ {
				for j := 0; j < 10; j++ {
					repr += fmt.Sprint(p.Solution[Index{rowIdx: row, colIdx: col}][Index{rowIdx: i, colIdx: j}])
				}
				repr += fmt.Sprint(" ")
			}
			repr += fmt.Sprintln()
		}
		repr += fmt.Sprintln()
	}
	return repr
}

func (p Puzzle) BorderlessPrint() string {
	minRowIdx, maxRowIdx, minColIdx, maxColIdx := p.SolutionBoundaries()
	repr := ""
	for row := minRowIdx; row <= maxRowIdx; row++ {
		for i := 1; i < 9; i++ {
			for col := minColIdx; col <= maxColIdx; col++ {
				for j := 1; j < 9; j++ {
					repr += fmt.Sprint(p.Solution[Index{rowIdx: row, colIdx: col}][Index{rowIdx: i, colIdx: j}])
				}
			}
			repr += fmt.Sprintln()
		}
	}
	return repr
}

func NewPuzzle(tiles TileMap) Puzzle {
	return Puzzle{
		Tiles:     tiles,
		Solution:  make(map[Index]Tile),
		UsedTiles: make([]Tile, 0),
		UsedIDs:   make([]int, 0),
	}
}

func (p *Puzzle) Build() {
	aCornerID := p.Tiles.Corners()[0]
	for rowID := 0; rowID < p.Size(); rowID++ {
		for colID := 0; colID < p.Size(); colID++ {
			if rowID == 0 && colID == 0 {
				tile := p.DownRightOriented(p.Tiles[aCornerID])
				p.UsedTiles = append(p.UsedTiles, tile)
				p.UsedIDs = append(p.UsedIDs, aCornerID)
				p.Solution[Index{rowIdx: rowID, colIdx: colID}] = tile
			} else if colID == 0 {
				if !p.MatchUpwards(rowID, colID) {
					panic("WTH?")
				}
			} else {
				if !p.MatchLeft(rowID, colID) {
					fmt.Println(rowID, colID)
					panic("WTF?")
				}
			}
		}
	}
}

func (p Puzzle) Size() int {
	sideLength := math.Sqrt(float64(len(p.Tiles)))
	return int(sideLength)
}

func (p Puzzle) DownRightOriented(t Tile) Tile {

	for i := 0; i < 9; i++ {
		if i == 4 {
			t = t.Flip()
		}
		if len(p.EdgeNeighborList()[t.CanonicalRight()]) > 1 && len(p.EdgeNeighborList()[t.CanonicalBottom()]) > 1 {
			return t
		}
		t = t.Rotate90()
	}
	panic("Couldn't fix the orientation!")
}

func (p *Puzzle) MatchLeft(rowID, colID int) bool {
	targetSide := p.Solution[Index{rowIdx: rowID, colIdx: colID - 1}].Right()
	for _, ID := range p.EdgeNeighborList()[targetSide.AsCanonical()] {
		tile := p.Tiles[ID]
		if p.HasUsed(ID) {
			continue
		}
		for i := 0; i < 9; i++ {
			if i == 4 {
				tile = tile.Flip()
			}
			if tile.Left() == targetSide {
				p.UsedTiles = append(p.UsedTiles, tile)
				p.UsedIDs = append(p.UsedIDs, ID)
				p.Solution[Index{rowIdx: rowID, colIdx: colID}] = tile
				return true
			}
			tile = tile.Rotate90()
		}
	}
	return false
}

func (p *Puzzle) MatchUpwards(rowID, colID int) bool {
	for ID, tile := range p.Tiles {
		if p.HasUsed(ID) {
			continue
		}
		for i := 0; i < 9; i++ {
			if i == 4 {
				tile = tile.Flip()
			}
			if tile.Top() == p.Solution[Index{rowIdx: rowID - 1, colIdx: colID}].Bottom() {
				p.UsedTiles = append(p.UsedTiles, tile)
				p.UsedIDs = append(p.UsedIDs, ID)
				p.Solution[Index{rowIdx: rowID, colIdx: colID}] = tile
				return true
			}
			tile = tile.Rotate90()
		}
	}
	return false
}

func (p Puzzle) HasUsed(ID int) bool {
	for _, usedID := range p.UsedIDs {
		if ID == usedID {
			return true
		}
	}
	return false
}

func (p *Puzzle) EdgeNeighborList() map[Edge][]int {
	if p.EdgeNeighbors == nil {
		p.EdgeNeighbors = p.Tiles.EdgeNeighbors()
	}
	return p.EdgeNeighbors
}
