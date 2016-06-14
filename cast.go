package gomap

import (
	"strconv"
	"errors"
	"fmt"
)

func toString(val interface{}) string {
	switch val.(type) {
	//case []byte:
	//	return fmt.Sprintf("%s", val)
	case string, []byte:
		return fmt.Sprintf("%s", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func stringToType(val string, valType interface{}) (interface{}, error) {
	switch valType.(type) {
	case bool:
		return strconv.ParseBool(val)
	case string:
		return val, nil
	case int:
		return strconv.Atoi(val)
	case int8:
		return strconv.ParseInt(val, 10, 8)
	case int16:
		return strconv.ParseInt(val, 10, 16)
	case int32:
		return strconv.ParseInt(val, 10, 32)
	case int64:
		return strconv.ParseInt(val, 10, 64)
	case uint:
		newVal, err := strconv.Atoi(val)
		return uint(newVal), err
	case uint8:
		strconv.ParseUint(val, 10, 8)
	case uint16:
		strconv.ParseUint(val, 10, 16)
	case uint32:
		strconv.ParseUint(val, 10, 32)
	case uint64:
		strconv.ParseUint(val, 10, 64)
	case float32:
		iVal, err := strconv.ParseFloat(val, 32)
		return float32(iVal), err
	case float64:
		iVal, err := strconv.ParseFloat(val, 64)
		return float64(iVal), err
	default:
		return nil, errors.New("Type not handled")
	}
	return nil, errors.New("Not handled")
}