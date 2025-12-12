package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

const (
	gridSize  = 4
	cellSize  = 100
	cellSpace = 10
)

var (
	grid     [gridSize][gridSize]int
	window   fyne.Window
	tiles    [gridSize][gridSize]*canvas.Text
	bgColors = map[int]fyne.Color{
		0:    fyne.Color{0xcd, 0xc1, 0xb4, 1},
		2:    fyne.Color{0xee, 0xe4, 0xda, 1},
		4:    fyne.Color{0xed, 0xe0, 0xc8, 1},
		8:    fyne.Color{0xf2, 0xb1, 0x79, 1},
		16:   fyne.Color{0xf5, 0x95, 0x63, 1},
		32:   fyne.Color{0xf6, 0x7c, 0x5f, 1},
		64:   fyne.Color{0xf6, 0x5e, 0x3b, 1},
		128:  fyne.Color{0xed, 0xcf, 0x72, 1},
		256:  fyne.Color{0xed, 0xcc, 0x61, 1},
		512:  fyne.Color{0xed, 0xc8, 0x50, 1},
		1024: fyne.Color{0xed, 0xc5, 0x3f, 1},
		2048: fyne.Color{0xed, 0xc2, 0x2e, 1},
	}
	textColors = map[int]fyne.Color{
		0:    fyne.Color{0x77, 0x6e, 0x65, 1},
		2:    fyne.Color{0x77, 0x6e, 0x65, 1},
		4:    fyne.Color{0x77, 0x6e, 0x65, 1},
		8:    fyne.Color{0xf9, 0xf6, 0xf2, 1},
		16:   fyne.Color{0xf9, 0xf6, 0xf2, 1},
		32:   fyne.Color{0xf9, 0xf6, 0xf2, 1},
		64:   fyne.Color{0xf9, 0xf6, 0xf2, 1},
		128:  fyne.Color{0xf9, 0xf6, 0xf2, 1},
		256:  fyne.Color{0xf9, 0xf6, 0xf2, 1},
		512:  fyne.Color{0xf9, 0xf6, 0xf2, 1},
		1024: fyne.Color{0xf9, 0xf6, 0xf2, 1},
		2048: fyne.Color{0xf9, 0xf6, 0xf2, 1},
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 初始化游戏网格
func initGrid() {
	grid = [gridSize][gridSize]int{}
	spawnRandom()
	spawnRandom()
}

// 随机生成2或4
func spawnRandom() {
	emptyCells := []fyne.Position{}
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 0 {
				emptyCells = append(emptyCells, fyne.NewPos(i, j))
			}
		}
	}
	if len(emptyCells) == 0 {
		return
	}
	pos := emptyCells[rand.Intn(len(emptyCells))]
	val := 2
	if rand.Float64() < 0.1 {
		val = 4
	}
	grid[pos.X][pos.Y] = val
}

// 更新界面显示
func updateUI() {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			val := grid[i][j]
			tile := tiles[i][j]
			if val == 0 {
				tile.Text = ""
			} else {
				tile.Text = fmt.Sprintf("%d", val)
			}
			tile.Color = textColors[val]
			tile.BackgroundColor = bgColors[val]
			canvas.Refresh(tile)
		}
	}
}

// 处理向左滑动
func moveLeft() bool {
	moved := false
	for i := 0; i < gridSize; i++ {
		newRow := []int{}
		merged := make(map[int]bool)
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if len(newRow) > 0 && newRow[len(newRow)-1] == grid[i][j] && !merged[len(newRow)-1] {
				newRow[len(newRow)-1] *= 2
				merged[len(newRow)-1] = true
				moved = true
			} else {
				newRow = append(newRow, grid[i][j])
			}
		}
		for len(newRow) < gridSize {
			newRow = append(newRow, 0)
		}
		for j := 0; j < gridSize; j++ {
			if grid[i][j] != newRow[j] {
				moved = true
			}
			grid[i][j] = newRow[j]
		}
	}
	return moved
}

// 旋转网格（用于复用向左滑动逻辑处理其他方向）
func rotateGrid() {
	newGrid := [gridSize][gridSize]int{}
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			newGrid[i][j] = grid[j][gridSize-1-i]
		}
	}
	grid = newGrid
}

// 处理键盘方向键事件
func handleKeyDown(e *fyne.KeyEvent) {
	moved := false
	switch e.Name {
	case fyne.KeyLeft:
		moved = moveLeft()
	case fyne.KeyRight:
		rotateGrid()
		rotateGrid()
		moved = moveLeft()
		rotateGrid()
		rotateGrid()
	case fyne.KeyUp:
		rotateGrid()
		rotateGrid()
		rotateGrid()
		moved = moveLeft()
		rotateGrid()
	case fyne.KeyDown:
		rotateGrid()
		moved = moveLeft()
		rotateGrid()
		rotateGrid()
		rotateGrid()
	}
	if moved {
		spawnRandom()
		updateUI()
	}
}

func main() {
	myApp := app.New()
	window = myApp.NewWindow("2048 Game")
	window.Resize(fyne.NewSize(gridSize*(cellSize+cellSpace)+cellSpace, gridSize*(cellSize+cellSpace)+cellSpace+30))

	// 初始化tiles
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			tile := canvas.NewText("", fyne.ColorBlack)
			tile.TextSize = 30
			tile.Alignment = fyne.TextAlignCenter
			tile.Resize(fyne.NewSize(cellSize, cellSize))
			tile.Move(fyne.NewPos(float32(j*(cellSize+cellSpace)+cellSpace), float32(i*(cellSize+cellSpace)+cellSpace)))
			tiles[i][j] = tile
		}
	}

	// 创建网格容器
	gridContainer := container.NewWithoutLayout()
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			gridContainer.Add(tiles[i][j])
		}
	}

	// 重置按钮
	resetBtn := widget.NewButton("Reset", func() {
		initGrid()
		updateUI()
	})

	mainContainer := container.NewVBox(gridContainer, resetBtn)
	window.SetContent(mainContainer)

	// 绑定键盘事件
	if deskWin, ok := window.(desktop.Window); ok {
		deskWin.SetOnKeyDown(handleKeyDown)
	}

	initGrid()
	updateUI()
	window.ShowAndRun()
}