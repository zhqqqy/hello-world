package util

import "fmt"

/**
* @author zhaohq
* @date 2020/3/20 10:00 上午
 */

func Add(a, b int) {
	c := a + b
	switch c {
	case 5:
		fmt.Println("yes")
	case 6:
		fmt.Println("big")
	}
}
