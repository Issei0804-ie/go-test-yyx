package main

import (
	"fmt"

	"github.com/Issei0804-ie/go-test-yyx/gotestyyx"
)

func main() {
	goTestYyx := gotestyyx.GoTestYyx{}
	calledTestNames, results := goTestYyx.Run([]func(){
		gotestyyx.TestAdd,
		gotestyyx.TestAdd2,
	})

	if calledTestNames[0] != "github.com/Issei0804-ie/go-test-yyx/gotestyyx.TestAdd" {
		fmt.Printf("error: %s", calledTestNames[0])
		return
	}

	if results[0] != nil {
		fmt.Printf("error: %s", results[0])
		return
	}

	if results[1] == nil {
		fmt.Printf("error: %s", results[1])
		return
	}

	fmt.Printf("pass!")
}
