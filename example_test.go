package env

import "fmt"

func ExampleSet() {
	Set("key1", "value1", DefaultNamespace)
	fmt.Println(Get("key1", DefaultNamespace))
	// 输出: value1
}

func ExampleGet() {
	Set("key1", "value1", DefaultNamespace)

	fmt.Println(Get("key1", DefaultNamespace))
	// 输出: value1
}

func ExampleGetAll() {
	Set("key1", "value1", "namespace1")
	Set("key2", "value2", "namespace2")

	fmt.Println(GetAll())
	// 输出: map[namespace1:map[key1:value1] namespace2:map[key2:value2]]
}