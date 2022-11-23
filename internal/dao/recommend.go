package dao

import (
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/recommend_repo"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type Recommend struct {
	Id           int64      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id" form:"id"`                               // 推荐ID
	DisplayType  int32      `gorm:"column:display_type;default:0" json:"displayType" form:"displayType"`                    // 显示类型   1:banner只显示在某一个分类下 2:banner显示在所有列表内(包括全部商品列表,专题商品列表)
	CategoryId   int64      `gorm:"column:category_id;default:0" json:"categoryId" form:"categoryId"`                       // 所属分类ID,如果为0 则显示在
	Title        string     `gorm:"column:title;default:''" json:"title" form:"title"`                                      // 标题
	TitleShort   string     `gorm:"column:title_short;default:''" json:"titleShort" form:"titleShort"`                      // 短标题
	TitlePic     string     `gorm:"column:title_pic;default:NULL" json:"titlePic" form:"titlePic"`                          // 标题图
	Description  string     `gorm:"column:description;default:''" json:"description" form:"description"`                    // 介绍 非html 代码
	ItemType     int32      `gorm:"column:item_type;default:0" json:"itemType" form:"itemType"`                             // 类型 0 网址链接 1:goods 商品ID 2:topic 专题ID 3: category分类ID
	ItemTypeData string     `gorm:"column:item_type_data;default:''" json:"itemTypeData" form:"itemTypeData"`               // 类型关联数据(链接 或者 ID)
	StartAt      timef.Time `gorm:"column:start_at;time;default:'0000-00-00 00:00:00'" json:"startAt" form:"startAt"`       // 开始时间 - 每日零时
	EndAt        timef.Time `gorm:"column:end_at;time;default:'0000-00-00 00:00:00'" json:"endAt" form:"endAt"`             // 结束时间- 当日某个时刻时间
	Weight       int32      `gorm:"column:weight;default:0" json:"weight" form:"weight"`                                    // 排序
	IsDeleted    int32      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                          // 是否删除
	CreatedAt    timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt    timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt    timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}

// CountRecommendList 获取推荐数量
func (d *Dao) CountRecommendList() (int64, error) {

	return recommend_repo.NewQueryBuilder().
		WhereDisplayType(model.Eq, 2).
		WhereIsDeleted(model.Eq, 0).
		Count()
}

// CountCategoryRecommendList 获取分类推荐的数量
func (d *Dao) CountCategoryRecommendList(categoryId int64) (int64, error) {

	return recommend_repo.NewQueryBuilder().
		WhereDisplayType(model.Eq, 1).
		WhereCategoryId(model.Eq, categoryId).
		WhereIsDeleted(model.Eq, 0).
		Count()
}

// GetRecommendList 获取推荐列表
func (d *Dao) GetRecommendList(page int, pageSize int) ([]*Recommend, error) {

	listm, err := recommend_repo.NewQueryBuilder().
		WhereDisplayType(model.Eq, 2).
		WhereIsDeleted(model.Eq, 0).
		OrderByWeight(false).
		Offset(app.GetPageOffset(page, pageSize)).
		Limit(pageSize).
		Get()

	if err != nil {
		return nil, err
	}

	var list []*Recommend
	for _, m := range listm {
		list = append(list, convert.StructAssign(m, &Recommend{}).(*Recommend))
	}
	return list, nil
}

// GetCategoryRecommendList 获取分类推荐列表
func (d *Dao) GetCategoryRecommendList(categoryId int64, page int, pageSize int) ([]*Recommend, error) {

	listm, err := recommend_repo.NewQueryBuilder().
		WhereDisplayType(model.Eq, 1).
		WhereCategoryId(model.Eq, categoryId).
		WhereIsDeleted(model.Eq, 0).
		OrderByWeight(false).
		Offset(app.GetPageOffset(page, pageSize)).
		Limit(pageSize).
		Get()

	if err != nil {
		return nil, err
	}

	var list []*Recommend
	for _, m := range listm {
		list = append(list, convert.StructAssign(m, &Recommend{}).(*Recommend))
	}
	return list, nil
}
