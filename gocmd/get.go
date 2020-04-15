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

var (
	// 协议名对应的协议[id, typeid]
	name2id = map[string]uint16{}
)

func init() {

	// 卡顿调度服务
	name2id[Heartbeat] = IdHeartbeat                 //
	name2id[LoginRequest] = IdLoginRequest   //
	name2id[LoginReply] = IdLoginReply                 //
	name2id[LogoutRequest] = IdLogoutRequest               //
	name2id[LogoutReply] = IdLogoutReply       //

}

// 返回[cmdid, typeid]
func GetIdByCmd(cmd string) uint16 {
	if id, ok := name2id[cmd]; ok {
		return id
	}
	return 0
}
