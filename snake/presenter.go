package snake

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor   = termbox.ColorGreen
)

// render 方法用于渲染游戏界面。
// 该方法首先清除终端屏幕，然后根据游戏当前状态绘制游戏标题、游戏区域、蛇、食物、得分以及退出游戏提示。
// 最后刷新终端屏幕，使绘制的内容显示出来。
// 返回值为可能发生的错误。
func (g *Game) render() error {
	termbox.Clear(defaultColor, defaultColor) // 清除终端屏幕

	// 计算屏幕中央位置以及游戏区域在屏幕上的显示范围
	var (
		w, h   = termbox.Size()                  // 获取终端窗口的宽度和高度
		midY   = h / 2                           // 屏幕中央的Y坐标
		left   = (w - g.arena.width) / 2         // 游戏区域左边界
		right  = (w + g.arena.width) / 2         // 游戏区域右边界
		top    = midY - (g.arena.height / 2)     // 游戏区域上边界
		bottom = midY + (g.arena.height / 2) + 1 // 游戏区域下边界，+1是为了包含游戏区域底部的得分显示
	)

	// 分别绘制游戏标题、游戏区域、蛇、食物、得分和退出游戏提示
	renderTitle(left, top)
	renderArena(g.arena, top, bottom, left)
	renderSnake(left, bottom, g.arena.snake)
	renderFood(left, bottom, g.arena.food)
	renderScore(left, bottom, g.score)
	renderQuitMessage(right, bottom)

	return termbox.Flush() // 刷新终端屏幕，使绘制的内容显示出来
}

// renderSnake 在指定位置渲染蛇的身体。
//
// 参数:
//
//	left - 蛇身体左上角x轴坐标。
//	bottom - 蛇身体左上角y轴坐标。
//	s - 指向snake结构体的指针，包含蛇的身体部位。
//
// 返回值:
//
//	无
func renderSnake(left, bottom int, s *snake) {
	// 遍历蛇的身体部位，并设置每个部位的终端字符及颜色。
	for _, b := range s.body {
		termbox.SetCell(left+b.x, bottom-b.y, ' ', snakeColor, snakeColor)
	}
}

// renderFood 在终端上渲染食物图标。
//
// 参数:
//
//	left - 食物图标的左边缘位置（以字符为单位）。
//	bottom - 食物图标的底边缘位置（以字符为单位）。
//	f - 指向食物结构体的指针，包含食物的位置（x, y）和表情符号。
//
// 无返回值。
func renderFood(left, bottom int, f *food) {
	// 使用终端绘图函数设置指定位置的字符和颜色为食物图标。
	termbox.SetCell(left+f.x, bottom-f.y, f.emoji, defaultColor, bgColor)
}

// renderArena 在终端上渲染一个矩形区域，表示一个竞技场的边界。
//
// 参数:
// a *arena: 表示要渲染的竞技场对象。
// top, bottom, left int: 定义竞技场边界在终端屏幕上的位置。
//
//	top 是上边界行号，bottom 是下边界行号，left 是左边界列号。
//	这些参数用来确定竞技场在终端屏幕上的显示位置。
//
// 无返回值。
func renderArena(a *arena, top, bottom, left int) {
	// 绘制竞技场的垂直边界线
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '│', defaultColor, bgColor)
		termbox.SetCell(left+a.width, i, '│', defaultColor, bgColor)
	}

	// 绘制竞技场的顶部和底部水平边界线及四个角
	termbox.SetCell(left-1, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left-1, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(left+a.width, top, '┐', defaultColor, bgColor)
	termbox.SetCell(left+a.width, bottom, '┘', defaultColor, bgColor)

	// 填充竞技场的水平边界线
	fill(left, top, a.width, 1, termbox.Cell{Ch: '─'})
	fill(left, bottom, a.width, 1, termbox.Cell{Ch: '─'})
}

// renderScore 在指定位置渲染分数信息。
//
// 参数：
// left - 分数显示的左边界位置。
// bottom - 分数显示的下边界位置。
// s - 要显示的分数值。
//
// 该函数没有返回值。
func renderScore(left, bottom, s int) {
	// 格式化分数字符串
	score := fmt.Sprintf("Score: %v", s)
	// 在指定位置打印分数
	tbprint(left, bottom+1, defaultColor, defaultColor, score)
}

// renderQuitMessage 在屏幕的指定位置渲染退出游戏的消息。
// right: 消息最右边的坐标。
// bottom: 消息最底边的坐标。
func renderQuitMessage(right, bottom int) {
	// 在指定位置打印退出提示消息
	m := "Press ESC to quit"
	tbprint(right-17, bottom+1, defaultColor, defaultColor, m)
}

// renderTitle 在指定位置渲染游戏标题。
//
// 参数:
// left - 标题文本的左边缘位置。
// top - 标题文本的上边缘位置。
//
// 该函数没有返回值。
func renderTitle(left, top int) {
	// 使用默认颜色在指定位置打印游戏标题
	tbprint(left, top-1, defaultColor, defaultColor, "Snake Game")
}

// fill函数用于在终端屏幕上用指定的单元格内容填充指定区域。
// x: 填充区域的起始x坐标
// y: 填充区域的起始y坐标
// w: 填充区域的宽度
// h: 填充区域的高度
// cell: 用于填充的单元格，包含字符、前景色和背景色
func fill(x, y, w, h int, cell termbox.Cell) {
	// 遍历填充区域的每个位置，设置对应的单元格内容和颜色
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			// 设置单元格内容和颜色
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

// tbprint 在指定位置使用给定的前景色和背景色打印字符串。
//
// 参数:
//
//	x - 起始横坐标。
//	y - 起始纵坐标。
//	fg - 前景色。
//	bg - 背景色。
//	msg - 要打印的字符串。
func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg { // 遍历字符串中的每个字符
		termbox.SetCell(x, y, c, fg, bg) // 在指定位置设置细胞（字符）
		x += runewidth.RuneWidth(c)      // 根据字符的宽度更新下一个字符的横坐标
	}
}
