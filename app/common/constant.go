package common

// 环境
const (
	ENV_DEV   = "dev"
	ENV_STAGE = "stage"
	ENV_PROD  = "prod"
)

// response
const (
	RET_SUCCESS = 1
	RET_ERROR   = 0
)

// context 的公用key
const (
	COMMON_LOG_FIELD_KEY = "clf_key"
	COMMON_START_TIME    = "cst"
)

// time format
const (
	FORMAT_YYYY_MM_DD_HH_II_SS_DETAIL = "2006-01-02 15:04:05.000000"
)

const (
	HTTP_HEADER_MJ_TRACE_ID = "orange-trace-id"
)

var REQUEST_IN_LOG_BLACKLIST = []string{
	"/health/check",
	"/upload/index",
}

var NOT_AUTH_URIS = []string{
	"/health/check",
}
