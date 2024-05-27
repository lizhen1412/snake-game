package main

import "github.com/lizhen1412/snake-game/snake"

// main 是程序的入口函数
// 它创建一个新的蛇游戏实例并启动游戏
func main() {
	// 创建一个新的蛇游戏实例，并启动游戏
	snake.NewGame().Start()
}
