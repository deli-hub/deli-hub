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
}
