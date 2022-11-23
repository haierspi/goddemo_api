package dao

import (
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/app_order_repo"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type AppOrder struct {
	AppOrderId        int64      `gorm:"column:app_order_id;primary_key;AUTO_INCREMENT" json:"appOrderId" form:"appOrderId"`     // APP 订单ID
	OrderId           int64      `gorm:"column:order_id;default:0" json:"orderId" form:"orderId"`                                // 订单ID
	AppUid            int64      `gorm:"column:app_uid;default:0" json:"appUid" form:"appUid"`                                   // 应用UID
	OpenId            string     `gorm:"column:open_id" json:"openId" form:"openId"`                                             // 用户openID
	GoodsId           int64      `gorm:"column:goods_id;default:0" json:"goodsId" form:"goodsId"`                                // 商品ID
	Uid               int64      `gorm:"column:uid;default:0" json:"uid" form:"uid"`                                             // 对应的用户id
	AppId             int64      `gorm:"column:app_id;default:0" json:"appId" form:"appId"`                                      // appID
	OutTradeNo        string     `gorm:"column:out_trade_no" json:"outTradeNo" form:"outTradeNo"`                                // 商户订单
	MetadataImage     string     `gorm:"column:metadata_image" json:"metadataImage" form:"metadataImage"`                        // 图像元数据
	NotifyUrl         string     `gorm:"column:notify_url" json:"notifyUrl" form:"notifyUrl"`                                    // 通知URL
	IsSuccessNotifyed int32      `gorm:"column:is_success_notifyed;default:0" json:"isSuccessNotifyed" form:"isSuccessNotifyed"` // 是否已经成功通知
	IsDeleted         int32      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                          // 是否删除
	CreatedAt         timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt         timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt         timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}

func (d *Dao) CreateAppOrder(dao *AppOrder) (int64, error) {

	m := convert.StructAssign(dao, &app_order_repo.AppOrder{}).(*app_order_repo.AppOrder)

	id, err := m.Create()

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *Dao) GetAppOrderById(id int64) (*AppOrder, error) {

	m, err := app_order_repo.NewQueryBuilder().
		WhereAppOrderId(model.Eq, id).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}
	return convert.StructAssign(m, &AppOrder{}).(*AppOrder), nil
}

// SetAppOrderRelationOrderIdById 设置APP订单关联订单
func (d *Dao) SetAppOrderRelationOrderIdById(id int64, orderId int64) error {

	err := app_order_repo.NewQueryBuilder().
		WhereAppOrderId(model.Eq, id).
		Updates(map[string]interface{}{
			"OrderId": orderId,
		})

	if err != nil {
		return err
	}

	return nil
}

// SetAppOrderSuccessNotifyedById 设置APP订单成功通知
func (d *Dao) SetAppOrderSuccessNotifyedById(id int64) error {

	err := app_order_repo.NewQueryBuilder().
		WhereAppOrderId(model.Eq, id).
		Updates(map[string]interface{}{
			"IsSuccessNotifyed": 1,
		})

	if err != nil {
		return err
	}

	return nil
}
