package advent20201220

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TileMap map[int]Tile

func (tm TileMap) ToString() (repr string) {
	repr += fmt.Sprintf("%v tiles", len(tm))
	repr += fmt.Sprintln()
	for tileID, tile := range tm {
		repr += fmt.Sprint(tileID)
		repr += fmt.Sprint("   ")
		repr += fmt.Sprint(tile.Boundaries())
		repr += fmt.Sprintln()
		repr += fmt.Sprintln(tile.ToString())
	}
	return
}

func (tm TileMap) First() Tile {
	for _, tile := range tm {
		return tile
	}
	return nil
}

func TilesFromFile(filename string) TileMap {
	tileMap := make(TileMap)
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	valueMap := make(map[Index]int)
	var rowIdx, tileID int
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		if text[:4] == "Tile" {
			if tileID != 0 {
				tileMap[tileID] = valueMap
			}
			rowIdx = 0
			tileString := text[5:9]
			tileNumber, err := strconv.ParseInt(tileString, 10, 64)
			check(err)
			tileID = int(tileNumber)

			valueMap = make(map[Index]int)
			continue
		}
		for colIdx, rn := range text {
			switch rn {
			case '.':
				valueMap[Index{colIdx, rowIdx}] = 0
			case '#':
				valueMap[Index{colIdx, rowIdx}] = 1
			}
		}
		rowIdx++
	}
	tileMap[tileID] = valueMap
	return tileMap
}

func (tm TileMap) EdgeNeighbors() map[Edge][]int {
	adjacencyMap := make(map[Edge][]int)
	for ID, tile := range tm {
		edges := tile.CanonicalEdges()
		for _, edge := range edges {
			if _, ok := adjacencyMap[edge]; !ok {
				adjacencyMap[edge] = make([]int, 0)
			}
			adjacencyMap[edge] = append(adjacencyMap[edge], ID)
		}
	}
	return adjacencyMap
}

func (tm TileMap) DifficultEdgeMatches() {
	adjacencyMap := tm.EdgeNeighbors()
	fmt.Println("Checking for problems...")
	for edge, lst := range adjacencyMap {
		if len(lst) != 2 {
			fmt.Printf("%v is difficult with %v", edge, len(lst))
			fmt.Println()
		}
	}
	fmt.Println("Done.")
}

func (tm TileMap) NeighborCountMap() (result map[int]int) {
	result = make(map[int]int)
	adjMap := tm.EdgeNeighbors()
	for ID, tile := range tm {
		for _, edge := range tile.CanonicalEdges() {
			for _, neighbor := range adjMap[edge] {
				if neighbor != ID {
					result[ID]++
				}
			}
		}
	}
	return
}

func (tm TileMap) Corners() (result []int) {
	for tile, ct := range tm.NeighborCountMap() {
		if ct == 2 {
			result = append(result, tile)
		}
	}
	return
}
