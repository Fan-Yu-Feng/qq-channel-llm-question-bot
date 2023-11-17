package botgo

import (
	"channelSdk/dto"
	"channelSdk/token"
)

// SessionManager 接口，管理session
type SessionManager interface {
	// Start 启动连接，默认使用 apInfo 中的 shards 作为 shard 数量，如果有需要自己指定 shard 数，请修 apInfo 中的信息
	Start(apInfo *dto.WebsocketAP, token *token.Token, intents *dto.Intent) error
}
