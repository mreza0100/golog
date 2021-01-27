package golog

func combine(in ...[]interface{}) (result []interface{}) {
	result = make([]interface{}, 0)

	for _, i := range in {
		result = append(result, i...)
	}

	return
}
