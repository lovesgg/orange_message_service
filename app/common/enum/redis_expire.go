package enum

import "time"

const (
	REDIS_EXPIRE_NINE_SECONDS   = 9 * time.Second //9秒
	REDIS_EXPIRE_FOUR_HOURS   = 14400 * time.Second //4个小时
	REDIS_EXPIRE_TWELVE_HOURS = 43200 * time.Second //12个小时
	REDIS_EXPIRE_NINE_MINUTE   = 360 * time.Second //9分钟

)
