package gotestyyx

import "fmt"

func TestAdd() {
	if add(1, 2) != 3 {
		fmt.Errorf("error!")
	}
}
