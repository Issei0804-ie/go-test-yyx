package main

import (
	"fmt"

	"github.com/Issei0804-ie/go-test-yyx/gotestyyx"
)

func main() {
	goTestYyx := gotestyyx.GoTestYyx{}
	calledTestNames, results, testsResult := goTestYyx.Run([]func() bool{
		gotestyyx.TestAdd,
		gotestyyx.TestAdd2,
	})

	if calledTestNames[0] != "github.com/Issei0804-ie/go-test-yyx/gotestyyx.TestAdd" {
		fmt.Printf("error: %s", calledTestNames[0])
		return
	}

	//called tests の中に[name, results] 入ってて欲しい
	if (testsResult.calledTest[0].name != "github.com/Issei0804-ie/go-test-yyx/gotestyyx.TestAdd") || (testsResult.calledTest[0].result != true) || (testsResult.calledTest[1].name != "github.com/Issei0804-ie/go-test-yyx/gotestyyx.TestAdd2") || (testsResult.calledTest[1].result != false) {
		fmt.Printf(
			"error: %s %b, %s %b",
			testsResult.calledTest[0].name,
			testsResult.calledTest[0].result,
			testsResult.calledTest[1].name,
			testsResult.calledTest[1].result,
		)
	}

	if results[0] == false {
		fmt.Printf("error: %s %s", results[0], calledTestNames[0])
		// loggerを入れる
		// Type Resultが.Runで返ってくる
		return
	}

	if results[1] == true {
		fmt.Printf("error: %s %s", results[1], calledTestNames[1])
		return
	}

	fmt.Printf("pass!")
}
