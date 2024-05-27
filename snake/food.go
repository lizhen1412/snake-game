package snake

import (
	"math/rand"
	"os"
	"strings"
)

// food 结构体表示游戏中的食物
type food struct {
	emoji        rune // 表示食物的表情符号
	points, x, y int  // points 表示食物的点数，吃掉食物后增加的分数
	// x 表示食物在地图上的 x 坐标
	// y 表示食物在地图上的 y 坐标
}

// newFood 创建一个新的食物对象。
//
// 参数:
// x - 食物在地图上的x坐标。
// y - 食物在地图上的y坐标。
//
// 返回值:
// 返回一个指向food结构体的指针，该结构体包含食物的初始点数、表情符号、以及其在地图上的坐标。
func newFood(x, y int) *food {
	// 使用提供的坐标和默认属性创建食物对象
	return &food{
		points: 10,             // 食物初始点数
		emoji:  getFoodEmoji(), // 食物表情符号
		x:      x,              // 食物x坐标
		y:      y,              // 食物y坐标
	}
}

// getFoodEmoji 返回一个与食物相关的表情符号。
// 如果当前环境支持Unicode，将返回一个随机的食物表情符号；
// 如果不支持Unicode，则返回一个默认的 '@' 字符。
func getFoodEmoji() rune {
	// 检查当前环境是否支持Unicode
	if hasUnicodeSupport() {
		// 如果支持Unicode，返回一个随机的食物表情符号
		return randomFoodEmoji()
	}

	// 如果不支持Unicode，返回默认的 '@' 字符
	return '@'
}

// randomFoodEmoji 生成并返回一个随机的食物表情符号。
// 该函数没有参数。
// 返回值：一个随机的食物表情符号（rune类型）。
func randomFoodEmoji() rune {
	f := []rune{
		'🍒',
		'🍍',
		'🍑',
		'🍇',
		'🍏',
		'🍌',
		'🍫',
		'🍭',
		'🍕',
		'🍩',
		'🍗',
		'🍖',
		'🍬',
		'🍤',
		'🍪',
	}

	// 从列表中随机选择一个表情符号并返回
	return f[rand.Intn(len(f))]
}

// hasUnicodeSupport 检查当前环境是否支持 Unicode。
// 该函数不接受任何参数。
// 返回值 bool 表示系统环境是否支持 Unicode。
func hasUnicodeSupport() bool {
	// 通过检查环境变量 LANG 中是否包含 "UTF-8" 来判断 Unicode 支持情况。
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}
