package main

import "fmt"

func Homework2() {
	const lim = 12307
	var x int
	flag := false
	for {
		fmt.Printf("Введите число меньше 12307: ")
		fmt.Scanf("%d", &x)
		if x < lim {
			break
		}
	}

	for x < lim {
		switch {
		case x < 0:
			x = -x
		case x%7 == 0:
			x *= 39
		case x%9 == 0:
			x *= 13
			x++
		default:
			x += 2
			x *= 3
		}

		if x%9 == 0 && x%13 == 0 {
			fmt.Printf("service error winth x = %d\n", x)
			flag = true
			break
		}
	}
	if !flag {
		fmt.Printf("получившееся число: %d\n", x)
	}
}
