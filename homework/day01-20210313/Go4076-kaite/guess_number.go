package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 猜数字游戏
	rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(100)) //随机生成0-100数字
	var target int = rand.Intn(100)
	// 设置输入次数，最多5次
	for i := 1; i <= 5; i++ {
		var guess int
		fmt.Print("请输入0-100间一个整数: ")
		fmt.Scan(&guess)
		// if语句进行判断用户输入与随机数的大小
		if guess > target {
			fmt.Println("太大了")
		} else if guess < target {
			fmt.Println("太小了")
		} else if guess == target {
			fmt.Println("猜对了")
		} else {
			fmt.Println()
		}
		fmt.Printf("您还剩 %d 次猜的机会!\n", 5-i)
	}
}
