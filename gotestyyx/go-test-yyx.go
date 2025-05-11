package gotestyyx

import (
	"reflect"
	"runtime"
)

type GoTestYyx struct {
}

func (goTestYyx GoTestYyx) Run(funcnames []func() bool) ([]string, []bool) {
	//funcname()
	names := []string{}
	results := []bool{}
	for _, funcname := range funcnames {
		fn := runtime.FuncForPC(reflect.ValueOf(funcname).Pointer())
		fn.Name()
		names = append(names, fn.Name())
		results = append(results, funcname())
	}
	return names, results
}
