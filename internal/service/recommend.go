package service

import (
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type Recommend struct {
	Id           int64      `json:"ID"`           // 推荐ID
	Title        string     `json:"title"`        // 标题
	TitleShort   string     `json:"titleShort"`   // 短标题
	TitlePic     string     `json:"titlePic"`     // 标题图
	Description  string     `json:"description"`  // 介绍 非html 代码
	ItemType     int32      `json:"itemType"`     // 类型 0 网址链接 1:goods 商品ID 2:topic 专题ID 3: category分类ID
	ItemTypeData string     `json:"itemTypeData"` // 类型关联数据(链接 或者 ID)
	CreatedAt    timef.Time `json:"createdAt"`    // 创建时间
	Data         any        `json:"data"`         // 格式化数据
}

type URL struct {
	Url string `json:"url"` // 介绍 非html 代码
}

// CategoryRecommendListRequest 专题分类请求
type CategoryRecommendListRequest struct {
	CategoryId int64 `form:"categoryID" binding:"required,gte=1"` // 所属分类ID 当displayType 设置为1 时必须提交  设置为2时
}

// RecommendList 专题分类列表
func (svc *Service) RecommendList(pager *app.Pager) ([]*Recommend, int, error) {

	count, err := svc.dao.CountRecommendList()
	if err != nil {
		return nil, 0, err
	}
	daoList, err := svc.dao.GetRecommendList(pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var list []*Recommend
	for _, order := range daoList {
		d := convert.StructAssign(order, &Recommend{}).(*Recommend)
		err := svc.recommendContent(d)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, d)
	}

	return list, int(count), nil
}

// CategoryRecommendList 专题分类列表
func (svc *Service) CategoryRecommendList(param *CategoryRecommendListRequest, pager *app.Pager) ([]*Recommend, int, error) {

	count, err := svc.dao.CountCategoryRecommendList(param.CategoryId)
	if err != nil {
		return nil, 0, err
	}
	daoList, err := svc.dao.GetCategoryRecommendList(param.CategoryId, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var list []*Recommend
	for _, order := range daoList {
		d := convert.StructAssign(order, &Recommend{}).(*Recommend)
		err := svc.recommendContent(d)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, d)
	}

	return list, int(count), nil
}

// 0 网址链接 1:goods 商品ID 2:topic 专题ID 3: category分类ID
// recommendContent 推荐内容
func (svc *Service) recommendContent(d *Recommend) error {

	if d.ItemType == 0 {
		//url
		d.Data = &URL{Url: d.ItemTypeData}
	} else if d.ItemType == 1 {
		//goods
		goodsId, _ := convert.StrTo(d.ItemTypeData).Int64()
		daod, _ := svc.dao.GetGoodsById(goodsId)
		if daod != nil {
			goods := convert.StructAssign(daod, &Goods{}).(*Goods)
			d.Data = goods
		}
	} else if d.ItemType == 2 {
		//topic
		topicId, _ := convert.StrTo(d.ItemTypeData).Int64()
		daod, _ := svc.dao.GetTopicById(topicId)
		if daod != nil {
			topic := convert.StructAssign(daod, &Topic{}).(*Topic)
			d.Data = topic
		}

	} else if d.ItemType == 3 {
		//category分类ID
		categoryId, _ := convert.StrTo(d.ItemTypeData).Int64()
		daod, _ := svc.dao.GetGoodsCategoryById(categoryId)
		if daod != nil {
			goodsCategory := convert.StructAssign(daod, &GoodsCategory{}).(*GoodsCategory)
			d.Data = goodsCategory
		}
	}

	return nil
}
