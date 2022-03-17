package main

import (
	"fmt"
	"github.com/mohith/DIY-2/pool"
)

func main() {
	p := pool.GetPool(2, 4, 100, 10)
	p.Start()
	for i := 0; i < 10; i++ {
		p.Submit(func(args ...interface{}) {
			fmt.Println(args[0])
		}, "go")
	}
	p.Submit(func(args ...interface{}) {
		fmt.Println(args[0], args[1])
	}, "go", "lang")
	p.Stop()
}
