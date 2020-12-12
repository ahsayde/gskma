package gskma

import (
	"fmt"
	"strconv"
)

func converToInt(v interface{}, bitSize int) (int64, error) {
	msg := "value is not of type boolean"

	switch val := v.(type) {
	case string:
		value, err := strconv.ParseInt(val, 10, bitSize)
		if err != nil {
			return 0, fmt.Errorf(msg)
		}
		return value, nil
	case int:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case int64:
		return val, nil
	default:
		return 0, fmt.Errorf(msg)
	}
}

func converToInt64(v interface{}) (int64, error) {
	return converToInt(v, 64)
}

func converToInt32(v interface{}) (int32, error) {
	val, err := converToInt(v, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

func converToFloat(v interface{}, bitSize int) (float64, error) {
	tname := "float"
	if bitSize == 64 {
		tname = "double"
	}
	msg := fmt.Sprintf("value is not of type %s", tname)

	switch val := v.(type) {
	case string:
		value, err := strconv.ParseFloat(val, bitSize)
		if err != nil {
			return 0, fmt.Errorf(msg)
		}
		return value, nil
	case float32:
		return float64(val), nil
	case float64:
		return val, nil
	default:
		return 0, fmt.Errorf(msg)
	}
}

func converToFloat64(v interface{}) (float64, error) {
	return converToFloat(v, 64)
}

func converToFloat32(v interface{}) (float32, error) {
	val, err := converToFloat(v, 32)
	if err != nil {
		return 0, err
	}
	return float32(val), nil
}

func convertToBool(value interface{}) (bool, error) {
	msg := "value is not of type boolean"
	switch v := value.(type) {
	case string:
		val, err := strconv.ParseBool(v)
		if err != nil {
			return false, fmt.Errorf(msg)
		}
		return val, nil
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf(msg)
	}
}
