package helpers

import (
	"os/exec"
	"reflect"
	"strings"
	"unsafe"
)

func Combine(in ...[]interface{}) (result []interface{}) {
	result = make([]interface{}, 0, len(in)*2)

	for _, i := range in {
		result = append(result, i...)
		result = append(result, " ")
	}

	return result
}

// Just don't touch this func for god's sake
// https://stackoverflow.com/questions/42664837/how-to-access-unexported-struct-fields/43918797#43918797
func ParseVal(x interface{}) interface{} {
	val := reflect.ValueOf(x)

	if val.Type().Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return x
	}

	result := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		rs2 := reflect.New(val.Type()).Elem()
		rs2.Set(val)
		rf := rs2.Field(i)
		valOfFiled := reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
		result[val.Type().Field(i).Name] = valOfFiled.Addr().Elem().Interface()
	}

	return result
}
func CreateDir(logPath string) (didCreate bool) {
	splited := strings.Split(logPath, "/")

	dirPath := strings.Join(splited[:len(splited)-1], "/")

	cmd := exec.Command("mkdir", "-p", dirPath)
	if _, err := cmd.Output(); err != nil {
		return false
	}

	return true
}

func Unshift(new interface{}, vals []interface{}) []interface{} {
	result := make([]interface{}, 0, len(vals)+1)

	result = append(result, new)
	result = append(result, vals...)

	return result
}
