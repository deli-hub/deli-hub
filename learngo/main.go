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
func lenAndUpeer(name string) (int, string) {
	return len(name), strings.ToUpper(name)
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

	totalLength, _ := lenAndUpeer("goTest")
	fmt.Println(totalLength)
}
