package util

// ArrayStructMerge 结构体合并
func ArrayStructMerge(content interface{}, content2 interface{}) map[string]interface{} {
	var name = make(map[string]interface{})

	map1 := JSONToMap(content)
	map2 := JSONToMap(content2)
	for k, v := range map1 {
		name[k] = v
	}
	for k, v := range map2 {
		name[k] = v
	}
	return name
}
