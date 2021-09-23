package print_2D_clockwise

import "fmt"

// 方向枚举
type Dire int

const (
	Right Dire = iota
	Down
	Left
	Up
)

type Pos struct {
	X, Y int
}

func PrintSquare(s [][]int) {
	if len(s) == 0 { // 一行也没有，不用打印
		return
	}
	if len(s[0]) == 0 { // 一列也没有，不用打印
		return
	}
	count := 0 // 用于记录已经遍历的次数
	record := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		record[i] = make([]int, len(s[0]))
	}
	// 起始位置
	pos := Pos{0, 0}
	dire := Right
	for count < len(s)*len(s[0]) { // 当遍历次数打到元素个数时结束
		x := s[pos.Y][pos.X]
		fmt.Printf("%d ", x) // 打印
		pos, record, dire = newPos(pos, record, dire)
		count++
	}
	fmt.Printf("\n")
	// end
}

func newPos(pos Pos, record [][]int, dire Dire) (newPos Pos, newRecord [][]int, newDire Dire) {
	// 更新标记
	record[pos.Y][pos.X] = 1
	newRecord = record
	// 更新方向
	newDire = dire
	// 尝试移动
	switch dire {
	case Right:
		pos.X += 1
		if pos.X >= len(record[0]) || record[pos.Y][pos.X] == 1 {
			pos.X -= 1 // 恢复
			pos.Y += 1
			newDire = Down
		}
	case Down:
		pos.Y += 1
		if pos.Y >= len(record) || record[pos.Y][pos.X] == 1 {
			pos.X -= 1
			pos.Y -= 1
			newDire = Left
		}

	case Left:
		pos.X -= 1
		if pos.X < 0 || record[pos.Y][pos.X] == 1 {
			pos.X += 1
			pos.Y -= 1
			newDire = Up
		}

	case Up:
		pos.Y -= 1
		if pos.X < 0 || record[pos.Y][pos.X] == 1 {
			pos.X += 1
			pos.Y += 1
			newDire = Right
		}
	}
	newPos = pos
	return
}
