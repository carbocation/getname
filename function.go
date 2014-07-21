package getname

import (
	"reflect"
	"runtime"
)

// See http://stackoverflow.com/a/7053871/199475
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
