package snake

import (
	"time"

	"github.com/nsf/termbox-go"
)

var (
	pointsChan         = make(chan int)           // 用于传递得分的通道
	keyboardEventsChan = make(chan keyboardEvent) // 用于传递键盘事件的通道
)

// Game 定义游戏的结构体
type Game struct {
	arena  *arena // 游戏场地
	score  int    // 游戏得分
	isOver bool   // 游戏是否结束
}

// initialSnake 创建并返回一个新的snake实例。
// 初始方向为RIGHT，蛇的身体由四个coord元素组成，初始位置分别为 {x: 1, y: 1}, {x: 1, y: 2}, {x: 1, y: 3}, {x: 1, y: 4}。
// 返回值 *snake 是指向snake实例的指针。
func initialSnake() *snake {
	// 使用newSnake函数创建一个新的snake实例，初始方向为RIGHT，身体坐标列表初始化为给定的coord数组。
	return newSnake(RIGHT, []coord{
		coord{x: 1, y: 1}, // 蛇头
		coord{x: 1, y: 2}, // 身体
		coord{x: 1, y: 3}, // 身体
		coord{x: 1, y: 4}, // 身体
	})
}

// initialScore 初始化游戏得分。
// 返回值 int 是初始得分，固定为0。
func initialScore() int {
	return 0
}

// initialArena 创建并返回一个新的arena实例。
// 使用initialSnake函数生成的snake实例初始化arena。
// arena的宽度为50，高度为20，得分通过pointsChan通道传递。
// 返回值 *arena 是指向arena实例的指针。
func initialArena() *arena {
	return newArena(initialSnake(), pointsChan, 20, 50) // 创建新的arena实例
}

// NewGame 创建并返回一个新的Game实例。
// 返回值 *Game 是指向Game实例的指针。
func (g *Game) end() {
	g.isOver = true // 设置游戏结束标志为true
}

// moveInterval 计算并返回游戏对象移动的间隔时间。
// 随着游戏得分增加，移动间隔时间会减少，使得游戏节奏变快。
// 返回值为时间间隔，单位为毫秒。
func (g *Game) moveInterval() time.Duration {
	// 根据当前得分计算移动间隔时间，得分每增加10点，移动间隔减少10毫秒。
	ms := 100 - (g.score / 10)
	return time.Duration(ms) * time.Millisecond // 返回时间间隔
}

// retry 重置游戏状态，使游戏可以重新开始。
// 该方法不接受参数，也不返回任何值。
func (g *Game) retry() {
	// 重置竞技场状态为初始状态
	g.arena = initialArena()
	// 重置分数为初始分数
	g.score = initialScore()
	// 重置游戏结束状态为未结束
	g.isOver = false
}

// addPoints 函数为游戏增加分数。
//
// 参数:
// p int - 要增加的分数值。
//
// 无返回值。
func (g *Game) addPoints(p int) {
	g.score += p // 总结：增加指定的分数到当前游戏得分。
}

// NewGame 创建一个新的游戏实例。
//
// 返回值:
// *Game: 返回一个初始化好的游戏实例指针。
func NewGame() *Game {
	// 初始化游戏场和分数
	return &Game{arena: initialArena(), score: initialScore()}
}

// Start 开始游戏。
// 该函数初始化游戏环境，启动游戏主循环，处理游戏事件（如键盘输入）和游戏逻辑（如蛇的移动、得分计算等）。
func (g *Game) Start() {
	// 初始化termbox库，如果初始化失败则触发panic
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close() // 确保在函数结束时关闭termbox

	// 启动一个协程监听键盘事件，并将事件发送到keyboardEventsChan通道
	go listenToKeyboard(keyboardEventsChan)

	// 初始渲染游戏场景，如果渲染失败则触发panic
	if err := g.render(); err != nil {
		panic(err)
	}

mainloop: // 定义主循环标签
	for {
		select {
		// 从得分通道接收得分，并增加到当前游戏得
		case p := <-pointsChan:
			g.addPoints(p)
		// 从键盘事件通道接收事件，并处理不同类型的事件
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE: // 处理移动事件
				d := keyToDirection(e.key)       // 将按键转换为方向
				g.arena.snake.changeDirection(d) // 改变蛇的方向
			case RETRY: // 处理重试事件
				g.retry() // 重置游戏
			case END: // 处理结束事件
				break mainloop // 退出主循环，结束游戏
			}
		default: // 默认处理，用于处理游戏的持续逻辑
			if !g.isOver { // 如果游戏未结束
				// 移动蛇并检测是否碰到自己或边界
				if err := g.arena.moveSnake(); err != nil {
					// 如果发生碰撞，则结束游戏
					g.end()
				}
			}

			// 渲染游戏场景，如果渲染失败则触发panic
			if err := g.render(); err != nil {
				panic(err)
			}

			// 控制蛇的移动间隔，动态调整蛇的移动速度
			time.Sleep(g.moveInterval())
		}
	}
}
