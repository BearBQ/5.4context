package main

import (
	"context"
	"fmt"
	"time"
)

func goRoutineFunc(ctx context.Context) {
	i := 0
	for {

		select {
		case <-ctx.Done():
			fmt.Println("goRutine is canceled")
			fmt.Println("количество циклов го рутины", i)

			return
		case <-time.After(1 * time.Second):
			fmt.Println("1 second was missing")
			i++
		}
	}

}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx2, cancel2 := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()
	defer cancel2()
	go goRoutineFunc(ctx)
	go goRoutineWithWait((ctx2))
	time.Sleep(11 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)

}

func goRoutineWithWait(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine2 is finished")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("goroutine2 is working")
		}
	}
}
