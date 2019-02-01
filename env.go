// Package env is a go port of the ruby dotenv library (https://github.com/bkeepers/dotenv)
//
// Examples/readme can be found on the github page at https://github.com/joho/godotenv
//
// The TL;DR is that you make a .env file that looks something like
//
// 		SOME_ENV_VAR=somevalue
//
// and then in your go code you can call
//
// 		godotenv.Load()
//
// and all the env vars declared in .env will be available through os.Getenv("SOME_ENV_VAR")package env
package env
const DefaultNamespace = "DefaultNamespace"

var variables = make(map[string]map[string]string)

// Set will set key,value in global env variables
func Set(key , value , namespace string) {
	namespaceMap, exist := variables[namespace]
	if (!exist) {
		namespaceMap = make(map[string]string)
	}
	namespaceMap[key] = value
	variables[namespace] = namespaceMap
}


// Get 获取指定命令空间下的键
func Get(key, namespace string) (string, bool) {
	namespaceMap, exist := GetNamespace(namespace)
	if (!exist) {
		namespaceMap = make(map[string]string)
	}
	value, exist := namespaceMap[key]
	return value, exist
}
// GetNamespace 获取命名空间下的变量列表
func GetNamespace(namespace string) (map[string]string, bool) {
	namespaceMap, exist := variables[namespace]
	return namespaceMap, exist
}
// GetAll
func GetAll() (map[string]map[string]string) {
	return variables
}
