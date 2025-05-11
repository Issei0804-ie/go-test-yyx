package main

import "testing"

func TestMain(t *testing.T) {
	goTestYyx := GoTestYyx{}
	calledTestNames := goTestYyx.Run()

	if calledTestNames[0] != "サンプルテスト" {
		t.Errorf("calledTestNames[0] is not サンプルテスト")
	}
}
