package gomap

import (
	"fmt"
)

type StrMap map[string]interface{}

func NewStrMap() StrMap {
	return make(StrMap, 0)
}

//exists
func (this StrMap) Exists(key string) bool {
	_, ok := this[key]
	return ok
}

//get
func (this StrMap) Get(key string) interface{} {
	return this[key]
}

//get
func (this StrMap) String(key string) string {
	return fmt.Sprintf("%s", this[key])
}

//get
func (this StrMap) Int(key string, defaultValue int) int {
	mValue := this[key]
	tValue, err := stringToType(toString(mValue), defaultValue)
	if err != nil {
		return defaultValue
	}
	if v, ok := tValue.(int); ok{
		return v
	}
	return defaultValue
}

func (this StrMap) Int32(key string, defaultValue int32) int32 {
	mValue := this[key]
	tValue, err := stringToType(toString(mValue), defaultValue)
	if err != nil {
		return defaultValue
	}
	if v, ok := tValue.(int32); ok{
		return v
	}
	return defaultValue
}

func (this StrMap) Int64(key string, defaultValue int64) int64 {
	mValue := this[key]
	tValue, err := stringToType(toString(mValue), defaultValue)
	if err != nil {
		return defaultValue
	}
	if v, ok := tValue.(int64); ok{
		return v
	}
	return defaultValue
}



//get
func (this StrMap) Float32(key string, defaultValue float32) float32 {
	mValue := this[key]
	tValue, err := stringToType(toString(mValue), defaultValue)
	if err != nil {
		return defaultValue
	}
	if v, ok := tValue.(float32); ok{
		return v
	}
	return defaultValue
}

func (this StrMap) Float64(key string, defaultValue float64) float64 {
	mValue := this[key]
	tValue, err := stringToType(toString(mValue), defaultValue)
	if err != nil {
		return defaultValue
	}
	if v, ok := tValue.(float64); ok{
		return v
	}
	return defaultValue
}

func (this StrMap) StrMap(key string) StrMap {
	mValue, ok := this[key]
	if !ok{
		return nil
	}
	if v, ok := mValue.(StrMap); ok{
		return v
	}
	if v, ok := mValue.(map[string]interface{}); ok{
		return StrMap(v)
	}
	if v, ok := mValue.(map[interface{}]interface{}); ok{
		mapValue := NewStrMap()
		for key, value := range v{
			mapValue[toString(key)] = value
		}
		return mapValue
	}
	return nil
}