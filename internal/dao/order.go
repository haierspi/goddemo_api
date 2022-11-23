package dao

import (
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/order_repo"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// 订单类型
const (
	//用户支付下单
	OrderTypeUserPayOrder int32 = iota
	//免费领取
	OrderTypeFreeReceive
	//后台发放
	OrderTypeAdminSend
	//报名空投
	OrderTypeAirdrop
	//兑换码兑换
	OrderTypeRedeemCode
)

// 订单状态
const (
	//OrderStatusClosed 订单关闭
	OrderStatusClosed int32 = iota
	//OrderStatusUnpaid 订单待支付
	OrderStatusUnpaid
	//OrderStatusPaid 订单已经支付
	OrderStatusPaid
	//OrderStatusShipped 订单已发货
	OrderStatusShipped
)

// 订单核销状态
const (
	//订单尚未核销
	OrderNotChargeOff int32 = iota
	//订单已经核销
	OrderHasChargeOff
)

type Order struct {
	OrderId                 int64      `gorm:"column:order_id;primary_key;AUTO_INCREMENT" json:"orderId" form:"orderId"`                                   // 订单ID
	OrderSn                 string     `gorm:"column:order_sn;default:''" json:"orderSn" form:"orderSn"`                                                   // 订单号
	GoodsId                 int64      `gorm:"column:goods_id;default:0" json:"goodsId" form:"goodsId"`                                                    // 购买商品
	GoodsSn                 string     `gorm:"column:goods_sn;default:''" json:"goodsSn" form:"goodsSn"`                                                   // 货号
	Uid                     int64      `gorm:"column:uid;default:0" json:"uid" form:"uid"`                                                                 // 用户id
	Nickname                string     `gorm:"column:nickname;default:''" json:"nickname" form:"nickname"`                                                 // 用户昵称
	AppOrderId              int64      `gorm:"column:app_order_id;default:0" json:"appOrderId" form:"appOrderId"`                                          // APP 订单ID
	HannelsId               int64      `gorm:"column:hannels_id;default:0" json:"hannelsId" form:"hannelsId"`                                              // 渠道ID
	GoodsName               string     `gorm:"column:goods_name;default:''" json:"goodsName" form:"goodsName"`                                             // 购买商品名称
	GoodsUrl                string     `gorm:"column:goods_url;default:''" json:"goodsUrl" form:"goodsUrl"`                                                // 连接关键词(英文)
	GoodsType               int32      `gorm:"column:goods_type;default:0" json:"goodsType" form:"goodsType"`                                              // 0：实物商品，1：数字藏品，2：实物礼包， 3：数字藏品-盲盒
	GoodsTitlePic           string     `gorm:"column:goods_title_pic;default:''" json:"goodsTitlePic" form:"goodsTitlePic"`                                // 商品标题图
	GoodsThumbPic           string     `gorm:"column:goods_thumb_pic;default:''" json:"goodsThumbPic" form:"goodsThumbPic"`                                // 商品缩略图
	GoodsImage              string     `gorm:"column:goods_image;default:''" json:"goodsImage" form:"goodsImage"`                                          // 商品封面图
	GoodsBuyNum             int32      `gorm:"column:goods_buy_num;default:0" json:"goodsBuyNum" form:"goodsBuyNum"`                                       // 商品购买数量
	GoodsPrice              float64    `gorm:"column:goods_price;default:0.00" json:"goodsPrice" form:"goodsPrice"`                                        // 商品单价
	GoodsTotalCost          float64    `gorm:"column:goods_total_cost;default:0.00" json:"goodsTotalCost" form:"goodsTotalCost"`                           // 订单总价
	PaymentKey              string     `gorm:"column:payment_key;default:''" json:"paymentKey" form:"paymentKey"`                                          // 供前端暂存使用的 支付类型
	Payment                 int32      `gorm:"column:payment;default:0" json:"payment" form:"payment"`                                                     // 是否需要三方支付
	PaymentType             string     `gorm:"column:payment_type;default:''" json:"paymentType" form:"paymentType"`                                       // 三方支付方式 wxpay:微信
	PaymentCost             float64    `gorm:"column:payment_cost;default:0.00" json:"paymentCost" form:"paymentCost"`                                     // 三方支付金额
	PaymentDatetime         timef.Time `gorm:"column:payment_datetime;time;default:'0000-00-00 00:00:00'" json:"paymentDatetime" form:"paymentDatetime"`   // 三方支付时间
	PaymentPayTransactionId string     `gorm:"column:payment_pay_transaction_id;default:''" json:"paymentPayTransactionId" form:"paymentPayTransactionId"` // 三方支付订单号
	PaymentData             string     `gorm:"column:payment_data;default:''" json:"paymentData" form:"paymentData"`                                       // 详情
	AddressId               int64      `gorm:"column:address_id;default:0" json:"addressId" form:"addressId"`                                              // 地址ID
	LinkPhone               string     `gorm:"column:link_phone;default:''" json:"linkPhone" form:"linkPhone"`                                             // 收货手机
	LinkMan                 string     `gorm:"column:link_man;default:''" json:"linkMan" form:"linkMan"`                                                   // 收货联系人
	LinkAddress             string     `gorm:"column:link_address;default:''" json:"linkAddress" form:"linkAddress"`                                       // 收货地址
	ProvinceId              int32      `gorm:"column:province_id;default:0" json:"provinceId" form:"provinceId"`                                           // 省ID
	Province                string     `gorm:"column:province;default:''" json:"province" form:"province"`                                                 // 省
	CityId                  int32      `gorm:"column:city_id;default:0" json:"cityId" form:"cityId"`                                                       // 市ID
	City                    string     `gorm:"column:city;default:''" json:"city" form:"city"`                                                             // 市
	AreaId                  int32      `gorm:"column:area_id;default:0" json:"areaId" form:"areaId"`                                                       // 区县ID
	Area                    string     `gorm:"column:area;default:''" json:"area" form:"area"`                                                             // 区县
	CountyId                int32      `gorm:"column:county_id;default:0" json:"countyId" form:"countyId"`                                                 // 国家ID
	Address                 string     `gorm:"column:address;default:''" json:"address" form:"address"`                                                    // 收货人详细地址
	ContractMetadataId      int64      `gorm:"column:contract_metadata_id;default:0" json:"contractMetadataId" form:"contractMetadataId"`                  // contract_metadata_id
	WalletGoodsId           int64      `gorm:"column:wallet_goods_id;default:0" json:"walletGoodsId" form:"walletGoodsId"`                                 // 我的藏品ID(钱包商品ID)
	IsShiped                int32      `gorm:"column:is_shiped;default:0" json:"isShiped" form:"isShiped"`                                                 // nft是否发放
	IsChargeOff             int32      `gorm:"column:is_charge_off;default:0" json:"isChargeOff" form:"isChargeOff"`                                       // 订单是否核销
	ChargeOffAt             timef.Time `gorm:"column:charge_off_at;time;default:'0000-00-00 00:00:00'" json:"chargeOffAt" form:"chargeOffAt"`              // 核销时间
	OrderType               int32      `gorm:"column:order_type;default:0" json:"orderType" form:"orderType"`                                              // 下单类型 0:用户下单 1:免费领取 2:后台发放  3:报名空投 4:兑换码兑换
	Status                  int32      `gorm:"column:status;default:0" json:"status" form:"status"`                                                        // 订单状态码 0:订单关闭 1订单创建 2:订单已支付或 非支付订单初始状态  3:订单已发货
	RelationGiveGoodsIsSend int32      `gorm:"column:relation_give_goods_is_send;default:0" json:"relationGiveGoodsIsSend" form:"relationGiveGoodsIsSend"` // 是否已经赠送
	RelationGiveGoodsId     int64      `gorm:"column:relation_give_goods_id;default:0" json:"relationGiveGoodsId" form:"relationGiveGoodsId"`              // 关联赠送商品 值为赠送的goodsid
	ExpressCompany          string     `gorm:"column:express_company;default:''" json:"expressCompany" form:"expressCompany"`                              //
	ExpressNumber           string     `gorm:"column:express_number;default:''" json:"expressNumber" form:"expressNumber"`                                 //
	UpdatedAt               timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"`                     // 更新时间
	CreatedAt               timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"`                     // 创建时间
	DeletedAt               timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"`                     // 标记删除时间
}

type OrderPayParams struct {
	Payment                 int32      // 是否需要三方支付
	PaymentType             string     // 三方支付方式 wxpay:微信
	PaymentCost             float64    // 三方支付金额
	PaymentDatetime         timef.Time // 三方支付时间
	PaymentPayTransactionId string     // 三方支付订单号
	PaymentData             string     // 详情
}

// 获取我订单商品总计
func (d *Dao) GetOrderGoodsCountByUID(uid int64, goodsId int64) (int64, error) {
	return order_repo.NewQueryBuilder().
		WhereGoodsId(model.Eq, goodsId).
		WhereUid(model.Eq, uid).WhereStatusIn([]int32{1, 2}).
		Count()
}

// 查询用户未付款订单一个
func (d *Dao) GetUserGoodsOrderIDByUIDAndGoodsId(uid int64, goodsId int64) (int64, error) {

	m, err := order_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereGoodsId(model.Eq, goodsId).
		WhereStatus(model.Eq, 1).
		First()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}

	return m.OrderId, nil
}

// 我的订单列表数量
func (d *Dao) CountOrderListByUID(uid int64) (int64, error) {
	return order_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		Count()
}

// 查询我的订单列表
func (d *Dao) GetOrderListByUID(uid int64, page int, pageSize int) ([]*Order, error) {

	listm, err := order_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		OrderByOrderId(false).
		Offset(app.GetPageOffset(page, pageSize)).
		Limit(pageSize).
		Get()
	if err != nil {
		return nil, err
	}

	var list []*Order
	for _, m := range listm {
		list = append(list, convert.StructAssign(m, &Order{}).(*Order))
	}
	return list, nil
}

// 查询某个状态下的我的订单列表数量
func (d *Dao) CountOrderStatusListByUID(uid int64, status int32) (int64, error) {
	return order_repo.NewQueryBuilder().
		WhereStatus(model.Eq, status).
		WhereUid(model.Eq, uid).
		Count()
}

// 查询某个状态下的我的订单列表
func (d *Dao) GetOrderStatusListByUID(uid int64, status int32, page int, pageSize int) ([]*Order, error) {

	listm, err := order_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereStatus(model.Eq, status).
		OrderByOrderId(false).
		Offset(app.GetPageOffset(page, pageSize)).
		Limit(pageSize).
		Get()
	if err != nil {
		return nil, err
	}

	var list []*Order
	for _, m := range listm {
		list = append(list, convert.StructAssign(m, &Order{}).(*Order))
	}
	return list, nil
}

func (d *Dao) CreateOrder(dao *Order) (int64, error) {

	m := convert.StructAssign(dao, &order_repo.Order{}).(*order_repo.Order)

	id, err := m.Create()

	if err != nil {
		return 0, err
	}
	return id, nil
}

// 获取订单dao
func (d *Dao) GetOrderById(orderID int64) (*Order, error) {
	m, err := order_repo.NewQueryBuilder().
		WhereOrderId(model.Eq, orderID).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &Order{}).(*Order), nil
}

// 根据订单号获取订单信息
func (d *Dao) GetOrderBySn(orderSn string) (*Order, error) {
	m, err := order_repo.NewQueryBuilder().
		WhereOrderSn(model.Eq, orderSn).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &Order{}).(*Order), nil
}

// 获取订单dao
func (d *Dao) GetUserOrderById(orderID int64, uid int64) (*Order, error) {
	m, err := order_repo.NewQueryBuilder().
		WhereOrderId(model.Eq, orderID).
		WhereUid(model.Eq, uid).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &Order{}).(*Order), nil
}

// 关闭订单
func (d *Dao) SetOrderStatusClosedById(id int64) error {

	err := order_repo.NewQueryBuilder().
		WhereOrderId(model.Eq, id).
		Updates(map[string]interface{}{
			"Status": OrderStatusClosed,
		})

	if err != nil {
		return err
	}

	return nil
}

// 订单NFT铸造成功
func (d *Dao) SetOrderStatusShipedById(id int64) error {

	err := order_repo.NewQueryBuilder().
		WhereOrderId(model.Eq, id).
		Updates(map[string]interface{}{
			"IsShiped": 1,
		})

	if err != nil {
		return err
	}

	return nil
}

// 设置订单核销
func (d *Dao) SetOrderChargeOffById(id int64) error {

	err := order_repo.NewQueryBuilder().
		WhereOrderId(model.Eq, id).
		Updates(map[string]interface{}{
			"IsChargeOff": OrderHasChargeOff,
			"ChargeOffAt": timef.Now(),
		})

	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) SetOrderRelationGiveGoodsSend(id int64) error {

	err := order_repo.NewQueryBuilder().
		WhereOrderId(model.Eq, id).
		WhereRelationGiveGoodsId(model.Neq, 0).
		Updates(map[string]interface{}{
			"RelationGiveGoodsIsSend": 1,
		})

	if err != nil {
		return err
	}

	return nil
}

// 订单支付成功
func (d *Dao) SetOrderStatusPaidById(id int64, params *OrderPayParams) error {

	err := order_repo.NewQueryBuilder().
		WhereOrderId(model.Eq, id).
		Updates(map[string]interface{}{
			"Status":                  OrderStatusPaid,
			"PaymentType":             params.PaymentType,
			"PaymentCost":             params.PaymentCost,
			"PaymentDatetime":         params.PaymentDatetime,
			"PaymentPayTransactionId": params.PaymentPayTransactionId,
			"PaymentData":             params.PaymentData,
		})

	if err != nil {
		return err
	}

	return nil
}

// 订单设置关联藏品
func (d *Dao) SetOrderWalletGoodsIdById(id int64, walletGoodsId int64) error {
	err := order_repo.NewQueryBuilder().
		WhereOrderId(model.Eq, id).
		Updates(map[string]interface{}{
			"WalletGoodsId": walletGoodsId,
		})

	if err != nil {
		return err
	}
	return nil

}
