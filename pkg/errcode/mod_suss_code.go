package errcode

var (
	SuccessOrderGoodsUnpaidRepeatCreate = NewSuss(1, "当前商品存在一个未支付订单,不能再创建订单")
	SuccessOrderChargeOff               = NewSuss(1, "订单核销成功")
)
