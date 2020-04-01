package common

type ErrorCode int

type ComError struct {
	Code       ErrorCode `json:"code"`
	Msg        string    `json:"msg"` // 对外输出
	InnerError string    `json:"error"`
	trace      string
}

func (e ComError) Error() string {
	return e.InnerError
}

func GenError(code ErrorCode, msg string, err error) *ComError {
	if msg == "" {
		msg = errmsgs[code]
	}

	var e string
	if err == nil {
		e = ""
	} else {
		e = err.Error()
	}
	return &ComError{
		Code:       code,
		Msg:        msg,
		InnerError: e,
	}
}

func (e *ComError) SetTrace(trace string) {
	e.trace = trace
}

const (
	ERROR_OK = -1 //正确码
	_        = iota + 100000
	ERROR_UNKNOW
	ERROR_REQUEST_UNKNOW
	ERROR_REQUEST_INTERNAL_ERROR_500
	ERROR_REQUEST_PARAMS
	ERROR_UNMARSHAL
	ERROR_MARSHAL
	ERROR_RET_IS_ERROR
	ERROR_RET_STRUCT_WRONG
	ERROR_NO_UPLOAD_FILE
	ERROR_REQUEST_OSS_FAILED
	ERROR_REDIS_KEY_NOT_EXISTS
	ERROR_REDIS_CALL_ERROR
	ERROR_REQUEST_WARNING_DURATION
	ERROR_READ_IO

)

var errmsgs = map[ErrorCode]string{
	ERROR_UNKNOW:                     "未知错误",
	ERROR_REQUEST_UNKNOW:             "请求未知错误",
	ERROR_REQUEST_INTERNAL_ERROR_500: "请求网络错误500",
	ERROR_REQUEST_PARAMS:             "请求参数错误",
	ERROR_UNMARSHAL:                  "json解析失败",
	ERROR_MARSHAL:                    "json序列化失败",
	ERROR_RET_IS_ERROR:               "请求错误",
	ERROR_RET_STRUCT_WRONG:           "结果结构不正确",
	ERROR_NO_UPLOAD_FILE:             "未选择上传文件！",
	ERROR_REDIS_KEY_NOT_EXISTS:       "redis 返回 redis.Nil",
	ERROR_REDIS_CALL_ERROR:           "redis 调用失败",
	ERROR_REQUEST_WARNING_DURATION:   "调用下游接口响应时间超出预期",
	ERROR_READ_IO:                   "读取io错误",
	ERROR_REQUEST_OSS_FAILED: "请求oss上传失败",
}
