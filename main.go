package main

import (
	"context"
	"fmt"
	"time"
)

func goRutineFunc(ctx context.Context) {
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
	defer cancel()
	go goRutineFunc(ctx)

	time.Sleep(11 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)

}
