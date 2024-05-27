package snake

import "github.com/nsf/termbox-go"

// keyboardEventType 表示键盘事件的类型
type keyboardEventType int

// 定义各种键盘事件类型的常量
const (
	MOVE  keyboardEventType = 1 + iota // 移动事件
	RETRY                              // 重试事件
	END                                // 结束事件
)

// keyboardEvent 结构体表示一个键盘事件
type keyboardEvent struct {
	eventType keyboardEventType // 事件类型
	key       termbox.Key       // 按键
}

// keyToDirection 将键盘按键转换为方向
//
// 参数:
// k - termbox.Key 类型，表示键盘按键。
//
// 返回值:
// direction - 表示蛇移动的方向。
// 如果按键是方向键之一（上、下、左、右），则返回相应的方向常量。
// 如果按键不是方向键，则返回 0。
func keyToDirection(k termbox.Key) direction {
	switch k { // 根据按键转换为方向
	case termbox.KeyArrowLeft: // 如果按键是左箭头键，返回 LEFT 方向
		return LEFT
	case termbox.KeyArrowDown: // 如果按键是下箭头键，返回 DOWN 方向
		return DOWN // 如果按键是下箭头键，返回 DOWN 方向
	case termbox.KeyArrowRight: // 如果按键是右箭头键，返回 RIGHT 方向
		return RIGHT // 如果按键是右箭头键，返回 RIGHT 方向
	case termbox.KeyArrowUp: // 如果按键是上箭头键，返回 UP 方向
		return UP // 如果按键是上箭头键，返回 UP 方向
	default: // 如果按键不是方向键，返回 0
		return 0
	}
}

// listenToKeyboard 监听键盘事件并发送到事件通道
//
// 参数:
// evChan - chan keyboardEvent 类型，表示用于传递键盘事件的通道。
//
// 功能:
// 该函数会不断监听键盘输入，根据不同的按键生成相应的键盘事件，并将事件发送到事件通道。
// 支持的按键包括方向键（上、下、左、右）、ESC 键和 'r' 键。
func listenToKeyboard(evChan chan keyboardEvent) {
	// 设置输入模式为 ESC 模式，以便可以捕获 ESC 键事件
	termbox.SetInputMode(termbox.InputEsc)

	for {
		// PollEvent 从事件队列中获取一个事件
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			// 如果事件类型是键盘事件，则处理该事件
			switch ev.Key {
			case termbox.KeyArrowLeft:
				// 如果按下左箭头键，创建一个 MOVE 类型的键盘事件并发送到事件通道
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowDown:
				// 如果按下下箭头键，创建一个 MOVE 类型的键盘事件并发送到事件通道
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowRight:
				// 如果按下右箭头键，创建一个 MOVE 类型的键盘事件并发送到事件通道
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowUp:
				// 如果按下上箭头键，创建一个 MOVE 类型的键盘事件并发送到事件通道
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyEsc:
				// 如果按下 ESC 键，创建一个 END 类型的键盘事件并发送到事件通道
				evChan <- keyboardEvent{eventType: END, key: ev.Key}
			default:
				// 如果按下的不是方向键或 ESC 键，检查是否按下了 'r' 键
				if ev.Ch == 'r' {
					// 如果按下 'r' 键，创建一个 RETRY 类型的键盘事件并发送到事件通道
					evChan <- keyboardEvent{eventType: RETRY, key: ev.Key}
				}
			}
		case termbox.EventError:
			// 如果事件类型是错误事件，触发 panic 并显示错误信息
			panic(ev.Err)
		}
	}
}
