package dao

import (
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/goods_repo"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

const (
	//GoodsTypeReal 实物+数字商品
	GoodsTypeReal int32 = iota
	//GoodsTypeNFT 数字藏品
	GoodsTypeNFT
	//GoodsTypeRealGiftBox 实物礼包
	GoodsTypeRealGiftBox
	//报名空投
	GoodsType
	//GoodsTypeNFTRandomBox 数字藏品-盲盒
	GoodsTypeNFTRandomBox
)

const (
	//树图
	BlockchainConfluxKey string = "Conflux"
)

type Goods struct {
	GoodsId                   int64      `gorm:"column:goods_id;primary_key;AUTO_INCREMENT" json:"goodsId" form:"goodsId"`                                         // 商品ID
	GoodsSn                   string     `gorm:"column:goods_sn;default:''" json:"goodsSn" form:"goodsSn"`                                                         // 货号
	Auid                      int64      `gorm:"column:auid;default:0" json:"auid" form:"auid"`                                                                    // 后端发布用户ID
	CategoryId                int64      `gorm:"column:category_id;default:0" json:"categoryId" form:"categoryId"`                                                 // 一级分类
	TopicIds                  string     `gorm:"column:topic_ids;default:''" json:"topicIds" form:"topicIds"`                                                      // 跳转地址
	CategoryName              string     `gorm:"column:category_name;default:''" json:"categoryName" form:"categoryName"`                                          // 一级分类
	CopyrightId               int64      `gorm:"column:copyright_id;default:0" json:"copyrightId" form:"copyrightId"`                                              // 版权方ID
	CopyrightName             string     `gorm:"column:copyright_name;default:''" json:"copyrightName" form:"copyrightName"`                                       // 版权方名字
	CopyrightImage            string     `gorm:"column:copyright_image;default:''" json:"copyrightImage" form:"copyrightImage"`                                    // 版权方图片
	BrandId                   int64      `gorm:"column:brand_id;default:0" json:"brandId" form:"brandId"`                                                          // 品牌ID
	BrandName                 string     `gorm:"column:brand_name;default:''" json:"brandName" form:"brandName"`                                                   // 品牌方名称
	ReleaseId                 int64      `gorm:"column:release_id;default:0" json:"releaseId" form:"releaseId"`                                                    // 发行方ID
	ReleaseName               string     `gorm:"column:release_name;default:''" json:"releaseName" form:"releaseName"`                                             // 发行方名称
	GoodsType                 int32      `gorm:"column:goods_type;default:0" json:"goodsType" form:"goodsType"`                                                    // 0：藏品带实物商品，1：数字藏品，2：实物礼包， 3：数字藏品-盲盒
	GoodsName                 string     `gorm:"column:goods_name;default:''" json:"goodsName" form:"goodsName"`                                                   // 商品名称
	GoodsPrice                float64    `gorm:"column:goods_price;default:0.00" json:"goodsPrice" form:"goodsPrice"`                                              // 商品价格
	GoodsMarketPrice          float64    `gorm:"column:goods_market_price;default:0.00" json:"goodsMarketPrice" form:"goodsMarketPrice"`                           // 市场价
	GoodsExpressType          int32      `gorm:"column:goods_express_type;default:0" json:"goodsExpressType" form:"goodsExpressType"`                              // 运费类型 0:免运费，1:全国运费一个价, 2:根据地区和货物重量单独计算）
	GoodsExpressPrice         float64    `gorm:"column:goods_express_price;default:0.00" json:"goodsExpressPrice" form:"goodsExpressPrice"`                        // 商品运费 0 免运费
	BuyNumLimit               int64      `gorm:"column:buy_num_limit;default:1" json:"buyNumLimit" form:"buyNumLimit"`                                             // 限购数量
	GoodsUrl                  string     `gorm:"column:goods_url;default:''" json:"goodsUrl" form:"goodsUrl"`                                                      // 连接关键词(英文)
	GoodsWeight               int64      `gorm:"column:goods_weight;default:0" json:"goodsWeight" form:"goodsWeight"`                                              // 商品单位重量（g）
	GoodsStock                int64      `gorm:"column:goods_stock;default:0" json:"goodsStock" form:"goodsStock"`                                                 // 当前库存
	GoodsTotalStock           int64      `gorm:"column:goods_total_stock;default:0" json:"goodsTotalStock" form:"goodsTotalStock"`                                 // 历史总库存
	GoodsTitlePic             string     `gorm:"column:goods_title_pic;default:''" json:"goodsTitlePic" form:"goodsTitlePic"`                                      // 商品标题图
	GoodsThumbPic             string     `gorm:"column:goods_thumb_pic;default:''" json:"goodsThumbPic" form:"goodsThumbPic"`                                      // 商品缩略图
	GoodsImage                string     `gorm:"column:goods_image;default:''" json:"goodsImage" form:"goodsImage"`                                                // 商品封面图
	GoodsAr                   string     `gorm:"column:goods_ar;default:''" json:"goodsAr" form:"goodsAr"`                                                         // ar 模型
	GoodsArImage              string     `gorm:"column:goods_ar_image;default:''" json:"goodsArImage" form:"goodsArImage"`                                         // ar加载图
	GoodsTags                 string     `gorm:"column:goods_tags;default:''" json:"goodsTags" form:"goodsTags"`                                                   // 标签 使用英文逗号间隔
	GoodsBody                 string     `gorm:"column:goods_body;default:''" json:"goodsBody" form:"goodsBody"`                                                   // 商品内容
	GoodsBodyMobile           string     `gorm:"column:goods_body_mobile;default:''" json:"goodsBodyMobile" form:"goodsBodyMobile"`                                // 商品内容移动版
	ContractTemplateId        int64      `gorm:"column:contract_template_id;default:0" json:"contractTemplateId" form:"contractTemplateId"`                        // 合约类型模板ID
	BlockchainId              int64      `gorm:"column:blockchain_id;default:0" json:"blockchainId" form:"blockchainId"`                                           // 区块链类型
	BlockchainName            string     `gorm:"column:blockchain_name;default:''" json:"blockchainName" form:"blockchainName"`                                    // 区块链名字
	BlockchainKey             string     `gorm:"column:blockchain_key;default:''" json:"blockchainKey" form:"blockchainKey"`                                       // 区块链key
	BlockchainIcon            string     `gorm:"column:blockchain_icon;default:''" json:"blockchainIcon" form:"blockchainIcon"`                                    // 区块链ICON
	BlockchainAddress         string     `gorm:"column:blockchain_address;default:''" json:"blockchainAddress" form:"blockchainAddress"`                           // 区块链地址
	ContractMetadataId        string     `gorm:"column:contract_metadata_id;default:''" json:"contractMetadataId" form:"contractMetadataId"`                       // 合约metadata类型ID
	ContractType              string     `gorm:"column:contract_type;default:''" json:"contractType" form:"contractType"`                                          // 合约类型
	ContractNetwork           string     `gorm:"column:contract_network;default:''" json:"contractNetwork" form:"contractNetwork"`                                 // 所在的网络(树图直接网络ID)
	ContractTokenuriUrlDomain string     `gorm:"column:contract_tokenuri_url_domain;default:''" json:"contractTokenuriUrlDomain" form:"contractTokenuriUrlDomain"` // tokenuri 访问地址域名
	ContractTokenuriUrlPre    string     `gorm:"column:contract_tokenuri_url_pre;default:''" json:"contractTokenuriUrlPre" form:"contractTokenuriUrlPre"`          // tokenuri 访问地址前缀
	ContractKeystorePath      string     `gorm:"column:contract_keystore_path;default:''" json:"contractKeystorePath" form:"contractKeystorePath"`                 // keystore地址
	IsHide                    int32      `gorm:"column:is_hide;default:0" json:"isHide" form:"isHide"`                                                             // 是否隐藏显示
	IsDeleted                 int32      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                                                    // 是否删除
	IsSkipNopayCheck          int32      `gorm:"column:is_skip_nopay_check;default:0" json:"isSkipNopayCheck" form:"isSkipNopayCheck"`                             // 是否跳过未付款重复检测
	Status                    int32      `gorm:"column:status;default:0" json:"status" form:"status"`                                                              // 商品状态（0下架,1上架）
	RelationGiveGoodsId       int64      `gorm:"column:relation_give_goods_id;default:0" json:"relationGiveGoodsId" form:"relationGiveGoodsId"`                    // 关联赠送商品 值为赠送的goodsid
	Weight                    int64      `gorm:"column:weight;default:0" json:"weight" form:"weight"`                                                              // 排序权重
	ReleasedTime              int64      `gorm:"column:released_time;default:0" json:"releasedTime" form:"releasedTime"`                                           // 发行时间(时间戳)
	SaleTime                  int64      `gorm:"column:sale_time;default:0" json:"saleTime" form:"saleTime"`                                                       // 销售时间(时间戳)
	ReleasedAt                timef.Time `gorm:"column:released_at;time;default:'0000-00-00 00:00:00'" json:"releasedAt" form:"releasedAt"`                        // 发行时间
	SaleAt                    timef.Time `gorm:"column:sale_at;time;default:'0000-00-00 00:00:00'" json:"saleAt" form:"saleAt"`                                    // 销售时间
	CreatedAt                 timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"`                           // 创建时间
	UpdatedAt                 timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"`                           // 更新时间
	DeletedAt                 timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"`                           // 标记删除时间
}

// GetGoodsById 根据商品Id获取商品信息
func (d *Dao) GetGoodsById(goodsId int64) (*Goods, error) {
	m, err := goods_repo.NewQueryBuilder().
		WhereGoodsId(model.Eq, goodsId).
		WhereStatus(model.Eq, 1).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &Goods{}).(*Goods), nil
}

// GetGoodsByURLKey 根据商品URLKey获取商品信息
func (d *Dao) GetGoodsByURLKey(goodsURL string) (*Goods, error) {
	m, err := goods_repo.NewQueryBuilder().
		WhereGoodsUrl(model.Eq, goodsURL).
		WhereStatus(model.Eq, 1).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}

	dao := &Goods{}
	convert.StructAssign(m, dao)
	return dao, nil
}

// GetGoodsList 获取商品列表
func (d *Dao) GetGoodsList(categoryId int64, page int, pageSize int) ([]*Goods, error) {

	queryBuilder := goods_repo.NewQueryBuilder()
	if categoryId != 0 {
		queryBuilder = queryBuilder.WhereCategoryId(model.Eq, categoryId)
	}

	listm, err := queryBuilder.
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

// CountGoodsList 根据商品URLKey获取商品数量
func (d *Dao) CountGoodsList(categoryId int64) (int64, error) {

	queryBuilder := goods_repo.NewQueryBuilder()
	if categoryId != 0 {
		queryBuilder = queryBuilder.WhereCategoryId(model.Eq, categoryId)
	}

	return queryBuilder.
		WhereIsHide(model.Eq, 0).
		WhereStatus(model.Eq, 1).
		WhereIsDeleted(model.Eq, 0).
		Count()
}

// IncGoodsStockById 根据商品Id增加库存
func (d *Dao) IncGoodsStockById(goodsId int64, number int64) error {
	err := goods_repo.NewQueryBuilder().
		WhereGoodsId(model.Eq, goodsId).
		WhereStatus(model.Eq, 1).
		WhereIsDeleted(model.Eq, 0).
		Increment("goods_stock", number)

	if err != nil {
		return err
	}
	return nil
}

// DecGoodsStockById 根据商品Id扣减库存
func (d *Dao) DecGoodsStockById(goodsId int64, number int64) error {
	err := goods_repo.NewQueryBuilder().
		WhereGoodsId(model.Eq, goodsId).
		WhereStatus(model.Eq, 1).
		WhereIsDeleted(model.Eq, 0).
		Decrement("goods_stock", number)

	if err != nil {
		return err
	}

	return nil
}
