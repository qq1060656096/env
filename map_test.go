package env

import (
	"testing"
)
// TestMap 测试map
func TestMap(t *testing.T) {
	var mapStr = make(map[string]string)
	mapStr["key1"] = "value1" // map[key1:value1]

	mapStr = make(map[string]string)
	mapStr["key2"] = "value2"// map[key2:value2]
	// 删除map key2键
	delete(mapStr, "key2")// map[]
}