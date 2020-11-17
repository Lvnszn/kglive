package core

// < request body show as blew

type (
	// XFCBarRequest 设置悬浮窗信息
	XFCBarRequest struct {
		SessionContext
		Title              string `json:"title"`              // 主标题
		SubTitle           string `json:"subTitle"`           // 副标题(内容)
		CountDownText      string `json:"countDownText"`      // 倒计时文字。如果要在文字中包含倒计时，使用占位符%@。比如 在%@秒后结束
		CountDownStartTime int    `json:"countDownStartTime"` // 倒计时开始时间。时间戳（秒）
		CountDownEndTime   int    `json:"countDownEndTime"`   // 倒计时结束时间。时间戳（秒）
		Image              string `json:"image"`              // 图片链接。只支持开放平台下的图片，开发者需要先上传图片到开放平台
	}

	// XFCRunBarRequest 设置轮播信息
	XFCRunBarRequest struct {
		SessionContext
		Interval string `json:"interval"` // 主标题
		Slides   Slide  `json:"slides"`   // 副标题(内容)
	}

	// Slide slide
	Slide struct {
		SlideID  string `json:"slideId"`  // 必填/最长20位	轮播图ID
		Title    string `json:"title"`    // 主标题
		SubTitle string `json:"subTitle"` // 副标题
		Image    string `json:"image"`    // 图片链接。只支持开放平台下的图片，开发者需要先上传图片到开放平台
	}

	// SessionContext base params for all request
	SessionContext struct {
		ID     string `json:"id"`     // 必填/最长32位	唯一标识
		AppID  string `json:"appId"`  // 必填	小程序ID
		StarID string `json:"starId"` // 必填	主播openID
		Time   int    `json:"time"`   // 必填	当前时间
	}

	// WebViewShow show or close webview
	WebViewShow struct {
		AppID   string `json:"appId"`   // 必填	小程序ID
		StarID  string `json:"starId"`  // 必填	主播openID
		Display bool   `json:"display"` // 是否显示。默认false
	}

	// ContentReq send msg to global
	ContentReq struct {
		SessionContext
		Content string `json:"content"`
	}
)

// >
// < Response struct show as below

type (
	// Response .
	Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	// Gift the gift response for star
	Gift struct {
		ClassID     int64  `json:"classId"`     // 分类id
		ClassName   string `json:"className"`   // 分类名称
		SortIndex   int    `json:"sortIndex"`   // 分类序号
		ID          int64  `json:"id"`          // 礼物id
		Name        string `json:"name"`        // 礼物名称
		Pic         string `json:"pic"`         // 礼物图标
		Price       int64  `json:"price"`       // 礼物价格
		Category    int64  `json:"category"`    // 礼物种类
		Status      int64  `json:"status"`      // 分类id
		Image       string `json:"image"`       // 礼物图片
		Type        int    `json:"type"`        // 类型
		MobileImage string `json:"mobileImage"` // 手机端礼物图片
	}

	// UserProfile 用户信息
	UserProfile struct {
		NickName string `json:"nickName"` // 昵称
		UserLogo string `json:"userLogo"` // 头像
		OpenID   string `json:"openId"`   // 用户openId
		UserID   int64  `json:"userId"`   // 用户繁星ID
	}

	// StarProfile 主播信息
	StarProfile struct {
		RoomID int  `json:"roomId"` // 主播房间号。没有则默认返回0
		IsStar bool `json:"isStar"` // 是否主播。true=主播，false=非主播
		IsLive bool `json:"isLive"` // 开播状态。true=开播，false=下播
	}
)

// ResponseCode .
type ResponseCode int32

const (
	// ParamsFault 参数错误
	ParamsFault ResponseCode = 1002
	// SendMessageFault 公屏消息发送失败
	SendMessageFault ResponseCode = 2000002
	// XFCSendMessageFault 悬浮窗消息发送失败
	XFCSendMessageFault ResponseCode = 2000003
	// XFCColseFault 悬浮窗关闭失败
	XFCColseFault ResponseCode = 2000004
	// WebViewSendMessageFault webview消息发送失败
	WebViewSendMessageFault ResponseCode = 2000010
)

// >
