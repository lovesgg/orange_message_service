package components

import (
	"orange_message_service/app/common/enum"
	"orange_message_service/app/components/eventdispatcher"
	"orange_message_service/app/components/mlog"
	"orange_message_service/app/components/redis"
	"orange_message_service/app/events"
)

func Init() {
	mlog.Init()
	redis.Init()

	EventInit()
}

func EventInit() {
	eventdispatcher.AddListener(enum.EVENT_NAME_RPC_CALL_WARNING, events.LogWarningRpcCallListener)
}
