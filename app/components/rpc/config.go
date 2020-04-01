package rpcComponent

import (
	"time"
)

type RpcConfig struct {
	Uri            string
	Timeout        time.Duration
	WaringDuration time.Duration
}

type RPC_REQUEST_INDEX int

const (
	TEST           RPC_REQUEST_INDEX = iota //普通品信息

)

const DEFAULT_TIMEOUT = 1 * time.Second                 //默认超时时间
const DEFAULT_WARNING_DURATION = 100 * time.Millisecond //默认的提醒时间, 100ms

var RpcConfigs = map[RPC_REQUEST_INDEX]RpcConfig{
	TEST: {
		"/test",
		DEFAULT_TIMEOUT,
		DEFAULT_WARNING_DURATION,
	},

}
