package gomap

type Mapx map[string]interface{}

func NewMapx() Mapx {
	return make(Mapx, 0)
}

//exists
func (this Mapx) Exists(key string) bool {
	_, ok := this[key]
	return ok
}

//get
func (this Mapx) Get(key string) interface{} {
	return this[key]
}

//get
func (this Mapx) String(key string) string {
	return toString(this[key])
}

//get
func (this Mapx) Int(key string, defaultValue int) int {
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

func (this Mapx) Int32(key string, defaultValue int32) int32 {
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

func (this Mapx) Int64(key string, defaultValue int64) int64 {
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
func (this Mapx) Float32(key string, defaultValue float32) float32 {
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

func (this Mapx) Float64(key string, defaultValue float64) float64 {
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

func (this Mapx) Mapx(key string) Mapx {
	mValue, ok := this[key]
	if !ok{
		return nil
	}
	if v, ok := mValue.(Mapx); ok{
		return v
	}
	if v, ok := mValue.(map[string]interface{}); ok{
		return Mapx(v)
	}
	if v, ok := mValue.(map[interface{}]interface{}); ok{
		mapValue := NewMapx()
		for key, value := range v{
			mapValue[toString(key)] = value
		}
		return mapValue
	}
	return nil
}