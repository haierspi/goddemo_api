package v1

import (
	"starfission_go_api/global"
	"starfission_go_api/internal/service"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Goods struct{}

func NewGoods() Goods {
	return Goods{}
}

func (t *Goods) List(c *gin.Context) {

	param := &service.GoodsListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	pager := &app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	svc := service.New(c.Request.Context())
	svcList, svcCount, err := svc.GoodsList(param, pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorPayCreateFail)
		}
		return
	}

	response.ToResponseList(svcList, svcCount)
	return
}

func (t *Goods) Details(c *gin.Context) {

	param := &service.GoodsDetailsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	svcData, err := svc.GoodsDetails(param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderViewFail)
		}
		return
	}

	response.ToResponse(svcData)
	return
}
func (t *Goods) URLDetails(c *gin.Context) {

	param := &service.GoodsDetailsURLRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	svcData, err := svc.GoodsDetailsURLKey(param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderViewFail)
		}
		return
	}

	response.ToResponse(svcData)
	return
}

func (t *Goods) AppOrderDetails(c *gin.Context) {

	param := &service.AppOrderGoodsDetailsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	svcData, err := svc.AppOrderGoodsDetails(param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderViewFail)
		}
		return
	}

	response.ToResponse(svcData)
	return
}

func (t *Goods) CategoryList(c *gin.Context) {
	param := &service.GoodsListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	pager := &app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	svc := service.New(c.Request.Context())
	svcList, svcCount, err := svc.GoodsCategoryList(pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorCategory)
		}
		return
	}

	response.ToResponseList(svcList, svcCount)
	return
}

func (t *Goods) CategoryDetails(c *gin.Context) {

	param := &service.GoodsCategoryDetailsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	svcData, err := svc.GoodsCategoryDetails(param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorCategory)
		}
		return
	}

	response.ToResponse(svcData)
	return
}
func (t *Goods) CategoryURLDetails(c *gin.Context) {

	param := &service.GoodsCategoryURLDetailsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	svcData, err := svc.GoodsCategoryURLDetails(param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorCategory)
		}
		return
	}

	response.ToResponse(svcData)
	return
}

func (t *Goods) TopicDetails(c *gin.Context) {

	param := &service.TopicDetailsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	svcData, err := svc.TopicDetails(param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorTopic)
		}
		return
	}

	response.ToResponse(svcData)
	return
}
func (t *Goods) TopicURLDetails(c *gin.Context) {

	param := &service.TopicURLDetailsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	svcData, err := svc.TopicURLDetails(param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorTopic)
		}
		return
	}

	response.ToResponse(svcData)
	return
}

func (t *Goods) TopicList(c *gin.Context) {

	param := &service.TopicListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	pager := &app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	svc := service.New(c.Request.Context())
	svcList, svcCount, err := svc.TopicList(param, pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorTopic)
		}
		return
	}

	response.ToResponseList(svcList, svcCount)
	return
}

func (t *Goods) CategoryURLTopicList(c *gin.Context) {

	param := &service.CategoryURLTopicListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	pager := &app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	svc := service.New(c.Request.Context())
	svcList, svcCount, err := svc.CategoryURLTopicList(param, pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorTopic)
		}
		return
	}

	response.ToResponseList(svcList, svcCount)
	return
}

func (t *Goods) TopicGoodsList(c *gin.Context) {

	param := &service.TopicGoodsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	pager := &app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	svc := service.New(c.Request.Context())
	svcList, svcCount, err := svc.TopicGoodsList(param, pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.TopicGoodsList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorTopic)
		}
		return
	}

	response.ToResponseList(svcList, svcCount)
	return
}

func (t *Goods) TopicURLGoodsList(c *gin.Context) {

	param := &service.TopicURLGoodsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	pager := &app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	svc := service.New(c.Request.Context())
	svcList, svcCount, err := svc.TopicURLGoodsList(param, pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorTopic)
		}
		return
	}

	response.ToResponseList(svcList, svcCount)
	return
}
func (t *Goods) RecommendList(c *gin.Context) {

	response := app.NewResponse(c)

	pager := &app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	svc := service.New(c.Request.Context())
	svcList, svcCount, err := svc.RecommendList(pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorRecommend)
		}
		return
	}

	response.ToResponseList(svcList, svcCount)
	return
}
func (t *Goods) RecommendCategoryList(c *gin.Context) {

	param := &service.CategoryRecommendListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	pager := &app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	svc := service.New(c.Request.Context())
	svcList, svcCount, err := svc.CategoryRecommendList(param, pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorRecommend)
		}
		return
	}

	response.ToResponseList(svcList, svcCount)
	return
}
