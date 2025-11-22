package typeUtils

func AnyToInt64(param any) (resultNum int64, err bool) {
	// 将接口类型转换为int64
	switch v := param.(type) {
	case int64:
		resultNum = v
		return resultNum, true
	default:
		return 0, false
	}
}
