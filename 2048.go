package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	gridSize = 4
)

var (
	grid     [gridSize][gridSize]int
	score    int
	gameOver bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
	initGrid()
}

// 初始化网格，生成两个随机数
func initGrid() {
	grid = [gridSize][gridSize]int{}
	spawnRandom()
	spawnRandom()
	score = 0
	gameOver = false
}

// 随机生成 2 或 4
func spawnRandom() {
	emptyCells := make([][2]int, 0)
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 0 {
				emptyCells = append(emptyCells, [2]int{i, j})
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
	grid[pos[0]][pos[1]] = val
}

// 渲染终端界面
func renderUI() {
	// 清屏（跨平台兼容）
	fmt.Print("\033[H\033[2J")
	fmt.Println("====== 终端版 2048 ======")
	fmt.Printf("====== 得分: %d ======\n", score)
	fmt.Println("W(上) S(下) A(左) D(右) | Q(退出) R(重置)")
	fmt.Println("-------------------------")
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 0 {
				fmt.Print("·\t")
			} else {
				fmt.Printf("%d\t", grid[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println("-------------------------")
	if gameOver {
		fmt.Println("====== 游戏结束! ======")
	}
}

// 向左移动逻辑
func moveLeft() bool {
	moved := false
	for i := 0; i < gridSize; i++ {
		newRow := make([]int, 0, gridSize)
		merged := make(map[int]bool)
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if len(newRow) > 0 && newRow[len(newRow)-1] == grid[i][j] && !merged[len(newRow)-1] {
				newRow[len(newRow)-1] *= 2
				score += newRow[len(newRow)-1]
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

// 旋转网格
func rotateGrid() {
	newGrid := [gridSize][gridSize]int{}
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			newGrid[i][j] = grid[j][gridSize-1-i]
		}
	}
	grid = newGrid
}

// 处理不同方向移动
func handleMove(dir rune) bool {
	moved := false
	switch dir {
	case 'a':
		moved = moveLeft()
	case 'd':
		rotateGrid()
		rotateGrid()
		moved = moveLeft()
		rotateGrid()
		rotateGrid()
	case 'w':
		rotateGrid()
		rotateGrid()
		rotateGrid()
		moved = moveLeft()
		rotateGrid()
	case 's':
		rotateGrid()
		moved = moveLeft()
		rotateGrid()
		rotateGrid()
		rotateGrid()
	}
	return moved
}

// 检查游戏是否结束
func checkGameOver() {
	// 检查是否有空格
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 0 {
				return
			}
		}
	}
	// 检查横向是否可合并
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize-1; j++ {
			if grid[i][j] == grid[i][j+1] {
				return
			}
		}
	}
	// 检查纵向是否可合并
	for j := 0; j < gridSize; j++ {
		for i := 0; i < gridSize-1; i++ {
			if grid[i][j] == grid[i+1][j] {
				return
			}
		}
	}
	gameOver = true
}

func main() {
	// 初始化键盘监听
	if err := keyboard.Open(); err != nil {
		fmt.Println("键盘监听初始化失败:", err)
		return
	}
	defer keyboard.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		renderUI()
		if gameOver {
			fmt.Print("按 R 重置游戏，按 Q 退出: ")
			scanner.Scan()
			input := scanner.Text()
			if len(input) == 0 {
				continue
			}
			switch input[0] {
			case 'q', 'Q':
				return
			case 'r', 'R':
				initGrid()
			}
			continue
		}

		// 读取键盘按键
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("读取按键失败:", err)
			continue
		}

		switch key {
		case keyboard.KeyCtrlC, keyboard.KeyEsc:
			return
		default:
			switch char {
			case 'q', 'Q':
				return
			case 'r', 'R':
				initGrid()
			case 'w', 'a', 's', 'd':
				if handleMove(char) {
					spawnRandom()
					checkGameOver()
				}
			}
		}
	}
}