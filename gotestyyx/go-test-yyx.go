package gotestyyx

import (
	"reflect"
	"runtime"
)

type GoTestYyx struct {
}

func (goTestYyx GoTestYyx) Run(funcname func()) []string {
	funcname()
	fn := runtime.FuncForPC(reflect.ValueOf(funcname).Pointer())
	return []string{
		fn.Name(),
	}
}
