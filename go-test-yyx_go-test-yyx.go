package main

import (
	"fmt"

	"github.com/Issei0804-ie/go-test-yyx/gotestyyx"
)

func main() {
	goTestYyx := gotestyyx.GoTestYyx{}
	calledTestNames, results := goTestYyx.Run([]func() bool{
		gotestyyx.TestAdd,
		gotestyyx.TestAdd2,
	})

	if calledTestNames[0] != "github.com/Issei0804-ie/go-test-yyx/gotestyyx.TestAdd" {
		fmt.Printf("error: %s", calledTestNames[0])
		return
	}

	if results[0] == false {
		fmt.Printf("error: %s %s", results[0], calledTestNames[0])
		return
	}

	if results[1] == true {
		fmt.Printf("error: %s %s", results[1], calledTestNames[1])
		return
	}

	fmt.Printf("pass!")
}
