package env

import (
	"testing"
)
func TestGetAll(t *testing.T) {
	Set("keyAll1", "kallValue1", "default")
	Set("keyAll2", "kallValue2", "default")
	envVariables := GetAll()
	namespaceVar, exist := envVariables["default"]
	if (!exist) {
		t.Error("not found namespace name")
	}
	keyAll1, exist := namespaceVar["keyAll1"]
	if (keyAll1 != "kallValue1") {
		t.Error("namespace kayAll1 key value error")
	}

	keyAll2, exist := namespaceVar["keyAll2"]
	if (keyAll2 != "kallValue2") {
		t.Error("namespace keyAll2 key value error")
	}
}

func TestSet(t *testing.T) {
	Set("key1", "value1", "default")
	value, _ := Get("key1", "default")
	if (value != "value1") {
		t.Error("get key1 value error")
	}
}

func TestMultiSet(t *testing.T) {
	Set("key1", "value1", "default")
	Set("key2", "value2", "default")
	value, _ := Get("key1", "default")
	if (value != "value1") {
		t.Error("get key1 value error")
	}
	value2, _ := Get("key2", "default")
	if (value2 != "value2") {
		t.Error("get key2 value error")
	}
}
