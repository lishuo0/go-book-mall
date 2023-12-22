package httputils

var (
	OK            = response{Code: 0, Message: "成功"}
	InterNalError = response{Code: 500, Message: "服务异常"}

	PARAM_MISSING = response{Code: 101, Message: "参数校验缺失"}
	PARAM_ERROR   = response{Code: 102, Message: "参数校验错误"}

	UserNotLogin        = response{Code: 401, Message: "用户未登录"}
	UserCheckLoginError = response{Code: 402, Message: "登录态校验失败"}
	UserLoginError      = response{Code: 403, Message: "登录失败"}
	UserNotExists       = response{Code: 404, Message: "用户不存在"}
	UserLogoutError     = response{Code: 405, Message: "退出登录失败"}

	GoodsNotExists = response{Code: 501, Message: "商品不存在"}
	GoodsNotEnough = response{Code: 502, Message: "商品库存不足"}

	OrderNotExists  = response{Code: 601, Message: "订单不存在"}
	CreateOrderFail = response{Code: 602, Message: "创建订单失败"}
)
