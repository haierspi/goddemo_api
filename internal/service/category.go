package service

import (
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
)

type GoodsCategory struct {
	CategoryId     int64  `json:"categoryID"`     // 分类ID
	CategoryName   string `json:"categoryName"`   // 分类名称
	CategoryNameEn string `json:"categoryNameEn"` // 英文分类名
	CategoryURL    string `json:"categoryUrl"`    // 分类URL
}

type GoodsCategoryDetailsRequest struct {
	CategoryId int64 `form:"categoryID" binding:"required,gte=0"`
}

type GoodsCategoryURLDetailsRequest struct {
	CategoryUrl string `form:"categoryURL" binding:"required"`
}

// GoodsCategoryList 获取分类列表
func (svc *Service) GoodsCategoryList(pager *app.Pager) ([]*GoodsCategory, int, error) {

	count, err := svc.dao.CountGoodsCategoryList()
	if err != nil {
		return nil, 0, err
	}
	daoList, err := svc.dao.GetGoodsCategoryList(pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var list []*GoodsCategory
	for _, order := range daoList {
		list = append(list, convert.StructAssign(order, &GoodsCategory{}).(*GoodsCategory))
	}

	return list, int(count), nil
}

// GoodsCategoryDetails 根据ID获取分类详情
func (svc *Service) GoodsCategoryDetails(param *GoodsCategoryDetailsRequest) (*GoodsCategory, error) {

	dao, err := svc.dao.GetGoodsCategoryById(param.CategoryId)
	if err != nil {
		return nil, err
	}
	return convert.StructAssign(dao, &GoodsCategory{}).(*GoodsCategory), nil
}

// GoodsCategoryURLDetails 根据URL获取分类详情
func (svc *Service) GoodsCategoryURLDetails(param *GoodsCategoryURLDetailsRequest) (*GoodsCategory, error) {

	dao, err := svc.dao.GetGoodsCategoryByURL(param.CategoryUrl)
	if err != nil {
		return nil, err
	}
	return convert.StructAssign(dao, &GoodsCategory{}).(*GoodsCategory), nil
}
