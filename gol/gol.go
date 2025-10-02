package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	newworld := make([][]byte, p.imageHeight)
	for i := range newworld {
		newworld[i] = make([]byte, p.imageWidth)
		// Copy initial state (optional)
		copy(newworld[i], world[i])
	}

	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			neighbours := 0
			for z := -1; z < 2; z++ {
				for l := -1; l < 2; l++ {
					if z == 0 && l == 0 {
						continue
					}
					nx, ny := x+z, y+l
					if nx == -1 {
						nx = p.imageWidth - 1
					}
					if ny == -1 {
						ny = p.imageHeight - 1
					}
					if nx >= p.imageWidth {
						nx = 0
					}
					if ny >= p.imageHeight {
						ny = 0
					}
					if world[ny][nx] == 255 {
						neighbours++

					}
				}
			}
			if world[y][x] == 255 && (neighbours > 3 || neighbours < 2) {
				newworld[y][x] = 0
			}
			//else{ newworld[y][x] =255}
			if world[y][x] == 0 && neighbours == 3 {
				newworld[y][x] = 255
			}
		}
	}
	//x_values := [3]int{}
	//fmt.Println(newworld)
	return newworld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var cells []cell
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] == 255 {
				cells = append(cells, cell{x: x, y: y})
			}
		}
	}

	return cells
}
