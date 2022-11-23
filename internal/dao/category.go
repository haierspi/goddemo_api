package dao

import (
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/goods_category_repo"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type GoodsCategory struct {
	CategoryId     int64      `gorm:"column:category_id;primary_key;AUTO_INCREMENT" json:"categoryId" form:"categoryId"`      // 分类ID
	CategoryName   string     `gorm:"column:category_name;default:''" json:"categoryName" form:"categoryName"`                // 分类名称
	CategoryNameEn string     `gorm:"column:category_name_en;default:''" json:"categoryNameEn" form:"categoryNameEn"`         // 英文分类名
	CategoryUrl    string     `gorm:"column:category_url" json:"categoryUrl" form:"categoryUrl"`                              // 分类URL
	IsHide         int32      `gorm:"column:is_hide;default:0" json:"isHide" form:"isHide"`                                   // 是否隐藏显示
	IsDeleted      int32      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                          // 是否删除
	Order          int64      `gorm:"column:order;default:0" json:"order" form:"order"`                                       // 排序
	UpdatedAt      timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	CreatedAt      timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	DeletedAt      timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}

// GetGoodsCategoryById 根据分类ID获取商品分类
func (d *Dao) GetGoodsCategoryById(categoryId int64) (*GoodsCategory, error) {
	m, err := goods_category_repo.NewQueryBuilder().
		WhereCategoryId(model.Eq, categoryId).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}
	return convert.StructAssign(m, &GoodsCategory{}).(*GoodsCategory), nil
}

// GetGoodsCategoryByURL 根据分类URL关键字获取商品分类
func (d *Dao) GetGoodsCategoryByURL(url string) (*GoodsCategory, error) {
	m, err := goods_category_repo.NewQueryBuilder().
		WhereCategoryUrl(model.Eq, url).
		WhereIsDeleted(model.Eq, 0).
		First()
	if err != nil {
		return nil, err
	}
	return convert.StructAssign(m, &GoodsCategory{}).(*GoodsCategory), nil
}

// GetGoodsCategoryList 获取商品分类列表
func (d *Dao) GetGoodsCategoryList(page int, pageSize int) ([]*GoodsCategory, error) {

	listm, err := goods_category_repo.NewQueryBuilder().
		WhereIsHide(model.Eq, 0).
		WhereIsDeleted(model.Eq, 0).
		OrderByOrder(false).
		OrderByCategoryId(false).
		Offset(app.GetPageOffset(page, pageSize)).
		Limit(pageSize).
		Get()

	if err != nil {
		return nil, err
	}

	var list []*GoodsCategory
	for _, m := range listm {
		list = append(list, convert.StructAssign(m, &GoodsCategory{}).(*GoodsCategory))
	}
	return list, nil
}

// CountGoodsCategoryList 获取商品分类数量
func (d *Dao) CountGoodsCategoryList() (int64, error) {

	return goods_category_repo.NewQueryBuilder().
		WhereIsHide(model.Eq, 0).
		WhereIsDeleted(model.Eq, 0).
		Count()
}
