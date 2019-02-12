package constant

const (
	APIPrefix = "/api/v1"

	TimerEveryHour = "@hourly" // 每小时触发

	/****************************************** user ****************************************/

	/****************************************** wechat ****************************************/

	// 扫码关注事件

	WechatEventKeyScanCodeJoinPhsMPGroupPrefix = "join/phs-mp/group" // 加入班级事件 join/phs-mp/group/<group-code>

	/****************************************** other ****************************************/

	EventTypeSubscribe   = "subscribe"   // 关注事件, 包括点击关注和扫描二维码(公众号二维码和公众号带参数二维码)关注
	EventTypeUnsubscribe = "unsubscribe" // 取消关注事件
	EventTypeScan        = "SCAN"        // 已经关注的用户扫描带参数二维码事件
	EventTypeLocation    = "LOCATION"    // 上报地理位置事件
	EventTypeClick       = "CLICK"
	EventTypeVIEW        = "VIEW"
)
