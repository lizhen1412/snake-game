package snake

import "errors"

// 允许的蛇移动方向常量
const (
	RIGHT direction = 1 + iota // 向右移动
	LEFT                       // 向左移动
	UP                         // 向上移动
	DOWN                       // 向下移动
)

// direction 类型表示蛇的移动方向
type direction int

// snake 结构体表示游戏中的蛇
type snake struct {
	body      []coord   // 蛇身体
	direction direction // 蛇移动方向
	length    int       // 蛇的长度
}

// newSnake 创建一个新的snake实例。
//
// 参数:
// d - 表示蛇的初始方向。
// b - 表示蛇身体的初始坐标序列。
//
// 返回值:
// 返回一个初始化好的snake指针。
func newSnake(d direction, b []coord) *snake {
	// 创建并返回一个新的snake实例，包含长度、身体和方向。
	return &snake{
		length:    len(b), // 蛇的初始长度设置为身体坐标段的数量。
		body:      b,      // 蛇的身体初始化为传入的坐标序列。
		direction: d,      // 设置蛇的初始方向。
	}
}

// changeDirection 更新蛇的移动方向。
// 如果新方向不是当前方向的相反方向，则更新蛇的方向。
// 参数:
//
//	d - 想要改变成的新方向
//
// 返回值:
//
//	无
func (s *snake) changeDirection(d direction) {
	// 定义方向的相反方向映射
	opposites := map[direction]direction{
		RIGHT: LEFT,  // 如果新方向是RIGHT，则相反方向是LEFT
		LEFT:  RIGHT, // 如果新方向是LEFT，则相反方向是RIGHT
		UP:    DOWN,  // 如果新方向是UP，则相反方向是DOWN
		DOWN:  UP,    // 如果新方向是DOWN，则相反方向是UP
	}

	// 检查新方向是否是当前方向的相反方向，如果不是，则更新方向
	if o := opposites[d]; o != 0 && o != s.direction {
		s.direction = d // 更新方向为新方向
	}
}

// head 返回蛇的头部坐标。
//
// 参数:
// s *snake: 表示蛇的结构体实例，其中包含蛇的身体坐标。
//
// 返回值:
// coord: 蛇头部的坐标，是一个结构体，具体结构取决于coord的定义。
func (s *snake) head() coord {
	// 返回蛇身体最后一个元素的坐标，即蛇的头部位置。
	return s.body[len(s.body)-1]
}

// die方法用于表示蛇的死亡状态。
// 此方法不接受参数，但会返回一个错误信息，表明蛇已经死亡。
func (s *snake) die() error {
	// 创建并返回一个表示死亡错误的新实例
	return errors.New("Died")
}

// move 移动蛇
//
// 功能:
// 根据当前方向移动蛇，并检查是否撞到自己。
// 如果蛇撞到自己，则返回错误；否则更新蛇的身体坐标。
//
// 返回值:
// error 类型，如果蛇撞到自己，则返回错误；否则返回 nil。
func (s *snake) move() error {
	// 获取蛇头的当前坐标
	h := s.head()
	// 创建一个新的坐标，初始值为蛇头的当前坐标
	c := coord{x: h.x, y: h.y}

	switch s.direction { // 根据当前方向移动蛇
	case RIGHT:
		c.x++ // 向右移动 x坐标加1
	case LEFT:
		c.x-- // 向左移动 x坐标减1
	case UP:
		c.y++ // 向上移动 y坐标加1
	case DOWN:
		c.y-- // 向下移动 y坐标减1
	}

	// 如果新位置与蛇身体重叠，标记蛇为死亡并返回错误
	if s.isOnPosition(c) {
		return s.die() // 蛇死亡
	}

	// 根据蛇的长度更新身体坐标
	if s.length > len(s.body) {
		// 如果蛇的长度大于当前身体坐标的长度，增加新位置到身体末尾
		s.body = append(s.body, c)
	} else {
		// 如果蛇的长度不大于当前身体坐标的长度，删除身体的第一个位置，并增加新位置到身体末尾
		s.body = append(s.body[1:], c)
	}

	// 返回 nil 表示没有错误
	return nil
}

// isOnPosition 检查蛇是否在指定的位置
//
// 参数:
//
//	c coord - 要检查的位置，包含 x 和 y 坐标
//
// 返回值:
//
//	bool - 如果蛇的身体在指定位置上，则返回 true；否则返回 false
func (s *snake) isOnPosition(c coord) bool {
	// 遍历蛇的身体部位
	for _, b := range s.body {
		// 检查当前部位的坐标是否与指定位置匹配
		if b.x == c.x && b.y == c.y {
			return true
		}
	}

	return false
}
