package routers

import (
	"time"

	_ "starfission_go_api/docs"
	"starfission_go_api/global"
	"starfission_go_api/internal/middleware"
	"starfission_go_api/internal/routers/api"
	"starfission_go_api/internal/routers/api/v1"
	"starfission_go_api/pkg/limiter"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.AppInfo())

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	//对404 的处理
	r.NoRoute(middleware.NoFound())

	//	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.Use(middleware.Cors())

	r.GET("/debug/vars", api.Expvar)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//upload := api.NewUpload()
	//r.POST("/upload/file", upload.UploadFile)

	//r.StaticFS("/static", httpclient.Dir(global.AppSetting.UploadSavePath))

	// //middleware.JWT() middleware.Token()
	APIV1Token := r.Group("/api/v1").Use(middleware.Token())

	APIV1 := r.Group("/api/v1")

	//member
	member := v1.NewMember()
	APIV1Token.Use()
	{
		APIV1Token.GET("/member/profile", member.Profile)

		//我的地址
		APIV1Token.GET("/member/address/default", member.AddressDefault)
		APIV1Token.GET("/member/address", member.AddressList)
		APIV1Token.POST("/member/address", member.AddressCreate)
		APIV1Token.PUT("/member/address/:id", member.AddressUpdate)
		APIV1Token.DELETE("/member/address/:id", member.AddressDelete)
		//我的钱包信息
		APIV1Token.GET("/member/wallet", member.Wallet)

	}

	goods := v1.NewGoods()
	APIV1.Use()
	{
		//分类
		APIV1.GET("/category/list", goods.CategoryList)
		APIV1.GET("/category/details", goods.CategoryDetails)
		APIV1.GET("/category/url_details", goods.CategoryURLDetails)
		//专题
		APIV1.GET("/topic/list", goods.TopicList)
		APIV1.GET("/topic/url_list", goods.CategoryURLTopicList)
		APIV1.GET("/topic/details", goods.TopicDetails)
		APIV1.GET("/topic/url_details", goods.TopicURLDetails)
		APIV1.GET("/topic/goods_list", goods.TopicGoodsList)
		APIV1.GET("/topic/url_goods_list", goods.TopicURLGoodsList)
		//商品
		APIV1.GET("/goods/list", goods.List)
		APIV1.GET("/goods/details", goods.Details)
		APIV1.GET("/goods/app_details", goods.AppOrderDetails)
		APIV1.GET("/goods/url_details", goods.URLDetails)
		//推荐
		APIV1.GET("/recommend/list", goods.RecommendList)
		APIV1.GET("/recommend/category_list", goods.RecommendCategoryList)
	}

	order := v1.NewOrder()
	APIV1Token.Use()
	{
		APIV1Token.POST("/order", order.Create)

		APIV1Token.POST("/order/free_receive", order.OrderFreeReceiveCreate)
		APIV1Token.POST("/order/redeem_code", order.OrderRedeemCodeCreate)
		APIV1Token.POST("/order/nft_metadata", order.OrderWalletGoodsMetadataCreate)
		APIV1Token.GET("/order/nft_metadata", order.OrderWalletGoodsMetadataList)
		APIV1Token.GET("/order", order.List)
		APIV1Token.GET("/order/give", order.Give)
		APIV1Token.GET("/order/:id", order.URLDetails)
		APIV1Token.POST("/order/details", order.Details)
		APIV1Token.GET("/order/details", order.Details)
	}
	//订单核销
	APIV1.POST("/order/charge_off", order.OrderChargeOff)

	pay := v1.NewPay()
	APIV1Token.Use()
	{
		APIV1Token.POST("/pay/wechat_h5", pay.OrderWechatH5PayCreate)
		APIV1Token.POST("/pay/wechat_jsapi", pay.OrderWechatJSAPIPayCreate)
		APIV1Token.POST("/pay/alipay_h5", pay.OrderAlipayH5PayCreate)
	}

	APIV1.POST("/pay/wechat_h5_notify", pay.OrderWechatH5PayNotify)
	APIV1.POST("/pay/wechat_jsapi_notify", pay.OrderWechatJSSDKPayNotify)
	APIV1.POST("/pay/alipay_h5_notify", pay.OrderAlipayH5PayNotify)

	transfer := v1.NewTransfer()
	APIV1Token.Use()
	{
		APIV1Token.GET("/transfer", transfer.List)
		APIV1Token.GET("/transfer/:id", transfer.Details)
	}
	wallet := v1.NewWallet()
	APIV1Token.Use()
	{
		APIV1Token.GET("/wallet", wallet.List)
		APIV1Token.GET("/wallet/:id", wallet.Details)
		APIV1Token.POST("/wallet_given", wallet.Given)
	}

	client := v1.NewClient()
	APIV1.GET("/client/version", client.Version)

	return r
}
