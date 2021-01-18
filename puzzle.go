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
	for ID, tile := range p.Tiles {
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
