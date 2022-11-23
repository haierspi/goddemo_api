package dao

import (
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/goods_repo"
	"starfission_go_api/internal/model/starfission/topic_goods_repo"
	"starfission_go_api/internal/model/starfission/topic_repo"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type Topic struct {
	TopicId     int64      `gorm:"column:topic_id;primary_key;AUTO_INCREMENT" json:"topicId" form:"topicId"`               // 专题ID
	CategoryId  int64      `gorm:"column:category_id;default:0" json:"categoryId" form:"categoryId"`                       // 所属分类ID
	TopicName   string     `gorm:"column:topic_name" json:"topicName" form:"topicName"`                                    // 商品名称
	TopicUrl    string     `gorm:"column:topic_url;default:''" json:"topicUrl" form:"topicUrl"`                            // 连接关键词(英文)
	TopicImage  string     `gorm:"column:topic_image;default:''" json:"topicImage" form:"topicImage"`                      // 专题描述
	TopicBanner string     `gorm:"column:topic_banner;default:''" json:"topicBanner" form:"topicBanner"`                   // 专题Banner
	TopicDesc   string     `gorm:"column:topic_desc;default:''" json:"topicDesc" form:"topicDesc"`                         // 专题描述
	RedirectUrl string     `gorm:"column:redirect_url;default:''" json:"redirectUrl" form:"redirectUrl"`                   // 跳转地址
	IsHide      int32      `gorm:"column:is_hide;default:0" json:"isHide" form:"isHide"`                                   // 是否隐藏显示
	IsDeleted   int32      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                          // 是否删除
	Weight      int64      `gorm:"column:weight;default:0" json:"weight" form:"weight"`                                    // 排序权重
	CreatedAt   timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt   timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt   timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}

// GetTopicById 根据ID获取专题详情
func (d *Dao) GetTopicById(topicId int64) (*Topic, error) {
	m, err := topic_repo.NewQueryBuilder().
		WhereTopicId(model.Eq, topicId).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &Topic{}).(*Topic), nil
}

// GetTopicByURLKey 根据URL获取专题详情
func (d *Dao) GetTopicByURLKey(topicUrl string) (*Topic, error) {
	m, err := topic_repo.NewQueryBuilder().
		WhereTopicUrl(model.Eq, topicUrl).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &Topic{}).(*Topic), nil
}

// GetTopicList 获取商品列表
func (d *Dao) GetTopicList(categoryId int64, page int, pageSize int) ([]*Topic, error) {

	queryBuilder := topic_repo.NewQueryBuilder()
	if categoryId != 0 {
		queryBuilder = queryBuilder.WhereCategoryId(model.Eq, categoryId)
	}

	listm, err := queryBuilder.
		WhereIsHide(model.Eq, 0).
		WhereIsDeleted(model.Eq, 0).
		OrderByWeight(false).
		OrderByCategoryId(false).
		Offset(app.GetPageOffset(page, pageSize)).
		Limit(pageSize).
		Get()

	if err != nil {
		return nil, err
	}

	var list []*Topic
	for _, m := range listm {
		list = append(list, convert.StructAssign(m, &Topic{}).(*Topic))
	}
	return list, nil
}

// CountTopicList 根据分类ID获取专题列表
func (d *Dao) CountTopicList(categoryId int64) (int64, error) {

	queryBuilder := topic_repo.NewQueryBuilder()
	if categoryId != 0 {
		queryBuilder = queryBuilder.WhereCategoryId(model.Eq, categoryId)
	}

	return queryBuilder.
		WhereIsHide(model.Eq, 0).
		WhereIsDeleted(model.Eq, 0).
		Count()
}

// GetTopicGoodsList 获取商品列表
func (d *Dao) GetTopicGoodsList(topicId int64, page int, pageSize int) ([]*Goods, error) {

	topicGoodsList, err := topic_goods_repo.NewQueryBuilder().
		WhereTopicId(model.Eq, topicId).
		Get()
	if err != nil {
		return nil, err
	}

	var goodsIds []int64
	for _, one := range topicGoodsList {
		goodsIds = append(goodsIds, one.GoodsId)
	}

	//todo 需要同步下架状态
	listm, err := goods_repo.NewQueryBuilder().
		WhereGoodsIdIn(goodsIds).
		WhereIsHide(model.Eq, 0).
		WhereStatus(model.Eq, 1).
		WhereIsDeleted(model.Eq, 0).
		OrderByWeight(false).
		OrderByGoodsId(false).
		Offset(app.GetPageOffset(page, pageSize)).
		Limit(pageSize).
		Get()

	if err != nil {
		return nil, err
	}

	var list []*Goods
	for _, m := range listm {
		list = append(list, convert.StructAssign(m, &Goods{}).(*Goods))
	}
	return list, nil
}

// CountTopicGoodsList 根据分类ID获取专题列表
func (d *Dao) CountTopicGoodsList(topicId int64) (int64, error) {
	//todo 需要同步下架状态
	return topic_goods_repo.NewQueryBuilder().
		WhereTopicId(model.Eq, topicId).
		Count()
}
