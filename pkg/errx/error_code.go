package errx

var errMsgMap = map[int64]string{
	CODE_SUCCESS:           "success",
	CODE_UNDEFINED:         "未定义的错误",
	CODE_INVALID_PARAMS:    "参数错误",
	CODE_DATA_NOT_FOUND:    "数据不存在",
	CODE_DATA_EXISTS:       "数据已存在",
	CODE_NOT_FOUND_HANDLER: "不存在的路由",
}

const CODE_SUCCESS = 0
const CODE_UNDEFINED = 1000
const CODE_INVALID_PARAMS = 1001
const CODE_DATA_NOT_FOUND = 1002
const CODE_DATA_EXISTS = 1003
const CODE_UNAUTHORIZED = 1004
const CODE_NOT_FOUND_HANDLER = 1005

func init() {
	// errMsgMap = make(map[int64]string, 0)
	// errMsgMap[CODE_SUCCESS] = "success"
	// errMsgMap[CODE_UNDEFINED] = "未定义的错误"

	// errMsgMap[CODE_INVALID_PARAMS] = "参数错误"
	// errMsgMap[CODE_DATA_NOT_FOUND] = "数据不存在"
	// errMsgMap[CODE_DATA_EXISTS] = "数据已存在"
}

func getMsgByCode(code int64) string {
	msg, ok := errMsgMap[code]
	if !ok {
		return "未知的错误"
	} else {
		return msg
	}
}
