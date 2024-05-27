package snake

import (
	"math/rand"
	"time"
)

// arena 结构体表示游戏场地
type arena struct {
	food       *food                    // 场地中的食物
	snake      *snake                   // 场地中的蛇
	hasFood    func(*arena, coord) bool // 判断是否在食物上
	height     int                      // 场地高度
	width      int                      // 场地宽度
	pointsChan chan (int)               // 分数
}

// newArena 创建一个新的游戏场地实例。
//
// 参数:
// s - 指向snake结构体的指针，代表游戏中的蛇。
// p - 一个整数类型的通道，用于向游戏场发送点数更新。
// h - 游戏场地的高度。
// w - 游戏场地的宽度。
//
// 返回值:
// 返回一个指向arena结构体的指针，该结构体代表游戏的场地。
func newArena(s *snake, p chan (int), h, w int) *arena {
	// 使用当前时间的纳秒级种子初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 初始化arena结构体
	a := &arena{
		snake:      s,
		height:     h,
		width:      w,
		pointsChan: p,
		hasFood:    hasFood, // hasFood应该是一个之前定义的布尔值，表明场地中是否有食物
	}

	// 在场地中放置食物
	a.placeFood()

	return a
}

// moveSnake 是一个用于移动蛇并处理游戏逻辑的方法。
// 如果蛇移动过程中发生错误，或者蛇离开游戏区域，则方法会返回相应的错误。
// 如果蛇成功移动且没有发生上述错误，方法会检查是否有食物被吃掉，如果有则更新分数和蛇的长度，并重新放置食物。
// 参数:
// - a *arena: 表示游戏的竞技场实例。
// 返回值:
// - error: 如果移动过程中发生错误，或者蛇离开游戏区域，则返回相应的错误；否则返回nil。
func (a *arena) moveSnake() error {
	// 尝试移动蛇，如果移动过程中发生错误，则返回该错误。
	if err := a.snake.move(); err != nil {
		return err
	}

	// 检查蛇是否离开游戏区域，如果是，则让蛇死亡并返回相应的错误。
	if a.snakeLeftArena() {
		return a.snake.die()
	}

	// 检查蛇的头部是否与食物位置重合，如果是，则进行相应的分数增加和蛇身长度增加的处理，并重新放置食物。
	if a.hasFood(a, a.snake.head()) {
		go a.addPoints(a.food.points) // 通过协程增加分数，不影响主游戏逻辑。
		a.snake.length++              // 蛇身长度增加。
		a.placeFood()                 // 重新随机放置食物。
	}

	return nil
}

// snakeLeftArena 检查蛇是否已经碰到了竞技场的左侧边界。
// 该函数不接受参数。
// 返回值为 bool 类型，如果蛇碰到了竞技场的左侧边界，则返回 true；否则返回 false。
func (a *arena) snakeLeftArena() bool {
	// 获取蛇头的位置
	h := a.snake.head()
	// 判断蛇头是否位于竞技场的左侧边界之外
	return h.x > a.width || h.y > a.height || h.x < 0 || h.y < 0
}

// addPoints向竞技场的积分通道中添加指定数量的积分。
// 参数：
//
//	p int: 需要添加的积分数量。
//
// 说明：
//
//	此函数用于异步地向竞技场的积分池中添加积分，不返回任何值。
func (a *arena) addPoints(p int) {
	a.pointsChan <- p // 将指定积分p放入积分通道，用于后续处理。
}

// placeFood 在竞技场中随机放置食物。
// 该函数不接受参数，也不返回值。
// 它首先随机生成一个位置，然后检查该位置是否已被占用。
// 如果该位置未被占用，则将食物放置在该位置，并更新竞技场的状态。
func (a *arena) placeFood() {
	var x, y int

	// 随机生成一个不被占用的位置来放置食物
	for {
		x = rand.Intn(a.width)  // 随机生成x坐标
		y = rand.Intn(a.height) // 随机生成y坐标

		// 检查生成的位置是否空闲，如果是则退出循环
		if !a.isOccupied(coord{x: x, y: y}) {
			break
		}
	}

	// 在找到的空闲位置上创建并放置食物
	a.food = newFood(x, y)
}

// hasFood 检查指定位置是否有食物。
//
// 参数:
// a *arena: 表示游戏区域的指针，其中包含食物的位置。
// c coord: 表示要检查的食物位置。
//
// 返回值:
// bool: 如果指定位置有食物，则返回 true；否则返回 false。
func hasFood(a *arena, c coord) bool {
	// 比较给定位置和食物位置的坐标是否相同
	return c.x == a.food.x && c.y == a.food.y
}

// isOccupied 判断指定坐标是否被占据
// 参数：
//
//	c coord - 要检查的坐标
//
// 返回值：
//
//	bool - 如果坐标被蛇占据则返回true，否则返回false
func (a *arena) isOccupied(c coord) bool {
	// 利用snake对象判断坐标是否在蛇的位置上
	return a.snake.isOnPosition(c)
}
