package main

import (
	"fmt"
	"strings"
)

// 파라미터 타입 명시
func multiply(a, b int) int {
	return a * b
}

// 다중 값 리턴 시
// func lenAndUpeer(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// naked return으로 작성
func lenAndUpeer(name string) (length int, uppercase string) {
	// defer -> 해당 함수가 끝난 후에 실행될 함수
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// 다중 파라미터 받기
func repeatMe(words ...string) {
	fmt.Println(words)
}

// LOOP
func superAdd(numbers ...int) int {
	// range가 numbers 안에서 조건에 맞는 경우 loop를 돌 수 있게 해준다.
	// Golang에서는 loop에서 for만 사용가능(for in, foreach는 없다)
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// if else
func canIDrink(age int) bool {

	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

// switch
func canIDrinkNow(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {

	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	var name string = "deli"
	// name := "deli"
	// line 6과 7은 동일한 기능이다.
	// ':=' 와 같이 표기하면 GO는 해당 변수 값에 대한 type을 찾아서 정의한다.
	// 이 표현은 func 안에서만 가능하다.

	name = "charlie"
	fmt.Println(name)

	fmt.Println(multiply(2, 2))

	totalLength, upper := lenAndUpeer("goTest")
	// 이 시점에서 위에 작성한 defer 함수가 실행된다.
	// 그리고 totalLength와 upper을 출력
	fmt.Println(totalLength, upper)

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	fmt.Println(canIDrink(16))
	fmt.Println(canIDrinkNow(16))
}
