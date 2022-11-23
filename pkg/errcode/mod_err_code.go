package errcode

var (
	ErrorUploadFileFail = NewError(20001, "上传文件失败")

	ErrorUserNotExist   = NewError(21001, "用户不存在或被禁用")
	ToUserNotExist      = NewError(21002, "转赠接收方用户不存在或被禁用")
	ToUserIsNotValidate = NewError(21003, "转赠接收方用户尚未实名认证")

	ErrorGetAddressFail         = NewError(22001, "获取我的地址失败")
	ErrorCreateAddressFail      = NewError(22002, "新建我的地址失败")
	ErrorUpdateAddressFail      = NewError(22003, "更新我的地址失败")
	ErrorDeleteAddressFail      = NewError(22004, "删除我的地址失败")
	ErrorGetDefaultAddressFail  = NewError(22005, "没有设置默认地址")
	ErrorGetAddressNotExistFail = NewError(22006, "这个收货地址不存在")

	ErrorOrderCreateFail               = NewError(23001, "订单创建失败")
	ErrorOrderGoodsAddressEmptyFail    = NewError(23002, "没有设置订单收货地址")
	ErrorOrderGoodsAddressNotExistFail = NewError(23003, "订单收货地址不存在")
	ErrorOrderGoodsSoldOutFail         = NewError(23004, "下单商品已经售罄")
	ErrorOrderGoodsNumNotEnoughFail    = NewError(23005, "订单购买数量比商品剩余数量多,不能下单")
	ErrorOrderGoodsPurchasedMaxNumFail = NewError(23006, "超过单品单人最大购买数量")
	ErrorOrderFail                     = NewError(23007, "下单失败")
	ErrorOrderContractMetadataFail     = NewError(23008, "订单数据源错误,不能下单")
	ErrorOrderDetailsNotSelfFail       = NewError(23009, "非本人订单不能获取详情")
	ErrorOrderGoodsDownFail            = NewError(23010, "商品已经下架不能购买")
	ErrorOrderRedeemCodeNotExistFail   = NewError(23011, "兑换码不存在")
	ErrorOrderRedeemCodeUsedFail       = NewError(23012, "兑换券已经使用过了,不能重复兑换")
	ErrorOrderRedeemCodeNotStartFail   = NewError(23013, "兑换券还没有开启兑换")
	ErrorOrderRedeemCodeIsEndFail      = NewError(23014, "兑换券已经过期不能使用")
	ErrorOrderRedeemCodeEmptyFail      = NewError(23015, "兑换券已经会用一空")
	ErrorOrderRedeemCodeNotSelfFail    = NewError(23016, "用户绑定兑换券,其他人无权使用")

	ErrorOrderViewFail             = NewError(23021, "订单无法查看失败")
	ErrorGoodsSaleTimeNotStartFail = NewError(23022, "商品尚未开售")
	ErrorOrderNotOwnFail           = NewError(23023, "非本人订单无法进行操作")
	RelationGiveGoodsIdIsFail      = NewError(23024, "商品未关联发放商品不能进行发放操作")
	ErrorOrderGiveFail             = NewError(23025, "订单错误进行发放操作")

	ErrorPayCreateFail                = NewError(24001, "订单支付失败")
	ErrorPayOrderNotExistFail         = NewError(24002, "订单不存在不能支付")
	ErrorPayOrderNotUnpaidFail        = NewError(24003, "订单不是待支付状态,不能发起支付")
	ErrorPayOrderPaidFail             = NewError(24004, "订单已经支付过了,不能重复支付")
	ErrorPayOrderOwnFail              = NewError(24005, "非本人订单,不能支付")
	ErrorPayOrderNotNeedFail          = NewError(24006, "非用户支付订单,不需要支付")
	ErrorPayOrderPaidDuplicateACKFail = NewError(24007, "订单已经支付过了,不需要重复确认支付")
	ErrorPayOrderTotalCostZeroFail    = NewError(24008, "订单总价为零不需要付款")

	WalletGivenExistError                  = NewError(25001, "有一个转赠正在进行中,不能重复提交")
	WalletGoodsNotIsShiped                 = NewError(25002, "藏品尚未铸造成功不能转赠")
	WalletGivenPasswordError               = NewError(25003, "您的密码错误")
	ErrorWalletGoodsMetadataDataFormatFail = NewError(25004, "钱包藏品Metadata数据格式错误")

	ErrorOrderHasChargeOffFail = NewError(26001, "核销失败,订单已经核销过了,不能重复核销")
	ErrorOrderNotExistFail     = NewError(26002, "核销失败,订单不存在,无法核销")
	ErrorOrderNotPayFail       = NewError(26003, "核销失败,订单尚未付款,无法核销")
	ErrorOrderChargeOffFail    = NewError(26004, "核销失败,请重试")

	ErrorOpenServiceMemberFail             = NewError(27001, "开放服务同步用户信息失败")
	ErrorOpenServiceNotifyUrlParseFail     = NewError(27002, "异步通知地址解析失败")
	ErrorOpenServiceNotifyDomainVerifyFail = NewError(27003, "异步通知地址域名效验失败")
	ErrorOpenServiceNotifySchemeVerifyFail = NewError(27004, "异步通知地址必须为 http或https 协议")
	ErrorOpenServiceOrderGoodsVerifyFail   = NewError(27005, "下订单的商品和商户订单商品不一致,无法下单")
	ErrorOpenServiceOrderRepeatCreateFail  = NewError(27006, "商户订单已下订单,不能重复下单,请到个人中心支付订单")
	ErrorOpenServiceOrderImageTypeFail     = NewError(27007, "图片类型不正确")
	ErrorOpenServiceOrderUserVerifyFail    = NewError(27008, "登录用户和商户订单用户不一致,无法下单")
	ErrorOpenServiceImageUrlParseFail      = NewError(27009, "上链图片地址解析失败")

	ErrorRecommend = NewError(27010, "推荐相关信息获取失败")
	ErrorTopic     = NewError(27011, "专题相关信息获取失败")
	ErrorCategory  = NewError(27012, "分类相关信息获取失败")
)
