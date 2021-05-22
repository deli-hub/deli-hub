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
	c := make(chan string)

	people := [2]string{"SW", "Chris"}
	for _, person := range people {
		// we're gonna check if the person is sexy
		go isSexy(person, c)
	}

	resultOne := <-c
	resultTwo := <-c

	// receiving msg from chan
	fmt.Println("Waiting for messages")
	// getting a message from channel
	fmt.Println("Received first msg:: ", resultOne)
	fmt.Println("Got the first msg")
	// once we got the msg from upper line, we're getting the message below from channel until we get another msg.
	// SW가 들어오고 Chris가 들어온다는 등의 순서는 없다. 다만 동시수행돼서 먼저 들어오는 게 resultOne이 된다.
	// 이런 식으로 다음 기능(?)이 수행될까지 기다리게 하는 작업을 blocking operation 이라고 한다.
	fmt.Println("Received second msg:: ", resultTwo)
	fmt.Println("DONE")

	for i := 0; i < len(people); i++ {
		fmt.Println("Waiting for ", i)
		fmt.Println(<-c) // 배열 people의 인원이 추가 될 경우 resultThree, four, five..늘리지말고 loop를 사용한다.
	}

}

/** GOROUTINE*/
/*func makeItCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		// 1초 쉬고 다시 수행된다. (모든 과정을 끝내는 데 걸리는 시간은 20초)
		time.Sleep(time.Second)
	}
}*/

/** CHANNEL example*/
// c chan 후에는 chan에서 주고 받을 데이터의 형식을 표기해준다.
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	// c라는 채널에 메시지를 보낸다. (return으로 보내지 않음)
	c <- person + " is sexy"
}
