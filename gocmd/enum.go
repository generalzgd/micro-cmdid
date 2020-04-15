/**
* @version: 1.0.0
* @author: zhangguodong:general_zgd
* @license: LGPL v3
* @contact: general_zgd@163.com
* @site: github.com/generalzgd
* @software: GoLand
* @time: 2020-03-02
 */

package gocmd

const (

	Heartbeat        = "heartbeat"        // 普通心跳，比较特殊，需要特定的type值，任意body
	LoginRequest              = "LoginRequest"              // 登录请求
	LoginReply                = "LoginReply"                // 登录响应
	LogoutRequest             = "LogoutRequest"             //
	LogoutReply               = "LogoutReply"               //

)

const (

	IdHeartbeat      = 1 // (1<<8)|1 普通心跳，比较特殊，需要特定的type值，任意body
	IdLoginRequest  = 2 // (1<<8)|2
	IdLoginReply    = 3 // (1<<8)|3
	IdLogoutRequest = 4 // (1<<8)|4
	IdLogoutReply   = 5 // (1<<8)|5

)
