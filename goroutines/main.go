package main

import (
	"fmt"
	"time"
)

func main() {
	// 함수 앞에 go를 더해줌으로써 두 가지 함수가 동시에 처리가 된다.
	// "SW" 카운트와 "Chris" 카운트가 일렬선상에 놓인 채 실행되는 것과 같다.
	// go mainItCount("SW")

	// 만약 아래 함수에서도 go를 붙여서 실행시키면 아무것도 터미널에 찍히지 않게되는데,
	// 이 이유는 goRoutine은 프로그램이 실행되는 중에서만 유효하기 때문이다.
	// main 함수가 종료되었기 때문에 다른 함수도 종료된다.
	// mainItCount("Chris")

	/* CHANNEL */
	//Channel -> goroutine이랑 메인함수 사이에 정보를 전달하기 위한 방법
	// make(chan infoType)
	c := make(chan bool)

	people := [2]string{"SW", "Chris"}
	for _, person := range people {
		// we're gonna check if the person is sexy
		go isSexy(person, c)
	}
	// receiving msg from chan
	fmt.Println(<-c)

}

/** GOROUTINE*/
/*func mainItCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		// 1초 쉬고 다시 수행된다. (모든 과정을 끝내는 데 걸리는 시간은 20초)
		time.Sleep(time.Second)
	}
}*/

/** CHANNEL example*/
func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	// c라는 채널에 true라는 메시지를 보낸다. (return으로 보내지 않음)
	c <- true
}
