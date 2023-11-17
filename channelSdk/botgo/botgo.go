// Package botgo 是一个QQ频道机器人 channelSdk 的 golang 实现
package botgo

import (
	"channelSdk/errs"
	"channelSdk/local"
	"channelSdk/log"
	"channelSdk/openapi"
	v1 "channelSdk/openapi/v1"
	"channelSdk/token"
	"channelSdk/websocket/client"
)

func init() {
	v1.Setup()     // 注册 v1 接口
	client.Setup() // 注册 websocket client 实现
}

// defaultSessionManager 默认实现的 session manager 为单机版本
// 如果业务要自行实现分布式的 session 管理，则实现 SessionManger 后替换掉 defaultSessionManager
var defaultSessionManager SessionManager = local.New()

// NewSessionManager 获得 session manager 实例
func NewSessionManager() SessionManager {
	return defaultSessionManager
}

// SelectOpenAPIVersion 指定使用哪个版本的 api 实现，如果不指定，sdk将默认使用第一个 setup 的 api 实现
func SelectOpenAPIVersion(version openapi.APIVersion) error {
	if _, ok := openapi.VersionMapping[version]; !ok {
		log.Errorf("version %v openapi not found or setup", version)
		return errs.ErrNotFoundOpenAPI
	}
	openapi.DefaultImpl = openapi.VersionMapping[version]
	return nil
}

// NewOpenAPI 创建新的 openapi 实例，会返回当前的 openapi 实现的实例
// 如果需要使用其他版本的实现，需要在调用这个方法之前调用 SelectOpenAPIVersion 方法
func NewOpenAPI(token *token.Token) openapi.OpenAPI {
	return openapi.DefaultImpl.Setup(token, false)
}

// NewSandboxOpenAPI 创建测试环境的 openapi 实例
func NewSandboxOpenAPI(token *token.Token) openapi.OpenAPI {
	return openapi.DefaultImpl.Setup(token, true)
}
