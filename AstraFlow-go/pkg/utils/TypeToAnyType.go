package typeUtils

import "github.com/spf13/cast"

func AnyToInt64(param any) (int64, bool) {
	result, err := cast.ToInt64E(param)
	return result, err == nil
}

func AnyToInt(param any) (int, bool) {
	result, err := cast.ToIntE(param)
	return result, err == nil
}

func AnyToString(param any) (string, bool) {
	result, err := cast.ToStringE(param)
	return result, err == nil
}

func AnyToFloat64(param any) (float64, bool) {
	result, err := cast.ToFloat64E(param)
	return result, err == nil
}

func AnyToBool(param any) (bool, bool) {
	result, err := cast.ToBoolE(param)
	return result, err == nil
}

func AnyToInt32(param any) (int32, bool) {
	result, err := cast.ToInt32E(param)
	return result, err == nil
}

func AnyToFloat32(param any) (float32, bool) {
	result, err := cast.ToFloat32E(param)
	return result, err == nil
}

func AnyToSliceString(param any) ([]string, bool) {
	result, err := cast.ToStringSliceE(param)
	return result, err == nil
}

func AnyToSliceInt64(param any) ([]int64, bool) {
	result, err := cast.ToInt64SliceE(param)
	return result, err == nil
}
