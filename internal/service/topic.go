package service

import (
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
)

type Topic struct {
	TopicId     int64  `json:"topicID"`     // 专题ID
	CategoryId  int64  `json:"categoryID"`  // 所属分类ID
	TopicName   string `json:"topicName"`   // 商品名称
	TopicUrl    string `json:"topicURL"`    // 连接关键词(英文)
	TopicImage  string `json:"topicImage"`  // 专题描述
	TopicBanner string `json:"topicBanner"` // 专题Banner
	RedirectUrl string `json:"redirectURL"` // 跳转地址
	TopicDesc   string `json:"topicDesc"`   // 专题描述
	Weight      int64  `json:"weight"`      // 排序权重
}

// TopicListRequest 专题分类请求
type TopicListRequest struct {
	CategoryId int64 `form:"categoryID"`
}

type CategoryURLTopicListRequest struct {
	CategoryUrl string `form:"categoryURL" binding:"required"`
}

// TopicDetailsRequest 专题分类请求
type TopicDetailsRequest struct {
	TopicId int64 `form:"topicID"`
}

// TopicURLDetailsRequest 商品列表请求参数
type TopicURLDetailsRequest struct {
	TopicUrl string `form:"topicURL"`
}

// TopicGoodsRequest 专题分类请求
type TopicGoodsRequest struct {
	TopicId int64 `form:"topicID"`
}

// TopicURLGoodsRequest 商品列表请求参数
type TopicURLGoodsRequest struct {
	TopicUrl string `form:"topicURL"`
}

// TopicList 专题分类列表
func (svc *Service) TopicList(param *TopicListRequest, pager *app.Pager) ([]*Topic, int, error) {

	//检查分类
	goodsCategorDao, err := svc.dao.GetGoodsCategoryById(param.CategoryId)
	if err != nil {
		return nil, 0, err
	}

	count, err := svc.dao.CountTopicList(goodsCategorDao.CategoryId)
	if err != nil {
		return nil, 0, err
	}
	daoList, err := svc.dao.GetTopicList(goodsCategorDao.CategoryId, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var list []*Topic
	for _, order := range daoList {
		list = append(list, convert.StructAssign(order, &Topic{}).(*Topic))
	}

	return list, int(count), nil
}

// TopicList 专题分类列表
func (svc *Service) CategoryURLTopicList(param *CategoryURLTopicListRequest, pager *app.Pager) ([]*Topic, int, error) {

	goodsCategorDao, err := svc.dao.GetGoodsCategoryByURL(param.CategoryUrl)
	if err != nil {
		return nil, 0, err
	}

	count, err := svc.dao.CountTopicList(goodsCategorDao.CategoryId)
	if err != nil {
		return nil, 0, err
	}
	daoList, err := svc.dao.GetTopicList(goodsCategorDao.CategoryId, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var list []*Topic
	for _, order := range daoList {
		list = append(list, convert.StructAssign(order, &Topic{}).(*Topic))
	}

	return list, int(count), nil
}

// TopicDetails 根据ID获取专题详情
func (svc *Service) TopicDetails(param *TopicDetailsRequest) (*Topic, error) {

	dao, err := svc.dao.GetTopicById(param.TopicId)
	if err != nil {
		return nil, err
	}
	return convert.StructAssign(dao, &Topic{}).(*Topic), nil
}

// TopicURLDetails 根据URL获取专题详情
func (svc *Service) TopicURLDetails(param *TopicURLDetailsRequest) (*Topic, error) {

	dao, err := svc.dao.GetTopicByURLKey(param.TopicUrl)
	if err != nil {
		return nil, err
	}
	return convert.StructAssign(dao, &Topic{}).(*Topic), nil
}

// TopicGoodsList 专题分类列表
func (svc *Service) TopicGoodsList(param *TopicGoodsRequest, pager *app.Pager) ([]*Goods, int, error) {

	topicDao, err := svc.dao.GetTopicById(param.TopicId)
	if err != nil {
		return nil, 0, err
	}

	count, err := svc.dao.CountTopicGoodsList(topicDao.TopicId)
	if err != nil {
		return nil, 0, err
	}
	daoList, err := svc.dao.GetTopicGoodsList(topicDao.TopicId, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var list []*Goods
	for _, order := range daoList {
		list = append(list, convert.StructAssign(order, &Goods{}).(*Goods))
	}

	return list, int(count), nil
}

// TopicURLGoodsList 专题分类列表
func (svc *Service) TopicURLGoodsList(param *TopicURLGoodsRequest, pager *app.Pager) ([]*Goods, int, error) {

	topicDao, err := svc.dao.GetTopicByURLKey(param.TopicUrl)
	if err != nil {
		return nil, 0, err
	}

	count, err := svc.dao.CountTopicGoodsList(topicDao.TopicId)
	if err != nil {
		return nil, 0, err
	}
	daoList, err := svc.dao.GetTopicGoodsList(topicDao.TopicId, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var list []*Goods
	for _, order := range daoList {
		list = append(list, convert.StructAssign(order, &Goods{}).(*Goods))
	}

	return list, int(count), nil
}
