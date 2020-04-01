package common

type ResponseError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Ret   int           `json:"ret"`
	Data  interface{}   `json:"data,omitempty"`
	Error ResponseError `json:"error,omitempty"`
}

// 公共日志字段
type CommonLogFields struct {
	IP      string `json:"ip"`
	Method  string `json:"method"`
	Path    string `json:"path"`
	TraceID string `json:"trace_id"`
	Header  map[string][]string
}
