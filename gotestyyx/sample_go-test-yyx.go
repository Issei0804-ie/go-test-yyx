package gotestyyx

import "fmt"

func TestAdd() bool {
	if add(1, 2) != 3 {
		fmt.Errorf("error!")
		return false
	}
	return true
}

func TestAdd2() bool {
	if add(1, 2) != 0 {
		fmt.Errorf("error!")
		return false
	}
	return true
}
