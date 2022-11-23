package dao

import (
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/transfer_repo"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type Transfer struct {
	Id                           int64      // 用户id
	FromUid                      int64      // 来源用户ID
	FromNickname                 string     // 来源昵称
	FromAvatar                   string     // 来源头像
	FromWalletHash               string     // 来源钱包地址
	ToUid                        int64      // 接收人id
	ToNickname                   string     // 接收人昵称
	ToAvatar                     string     // 接收人头像
	ToWalletHash                 string     // 接收人钱包地址
	TransactionHash              string     // 交易哈希
	GoodsId                      int64      // 购买商品
	GoodsName                    string     // 商品名称
	GoodsThumbPic                string     // 商品缩略图
	GoodsImage                   string     // 商品封面图
	GoodsAr                      string     // ar 模型
	GoodsArImage                 string     // ar加载图
	GoodsTags                    string     // 标签 使用英文逗号间隔
	CopyrightId                  int64      // 版权方ID
	CopyrightName                string     // 版权方名字
	BrandId                      int64      // 品牌ID
	BrandName                    string     // 品牌方名称
	ReleaseId                    int64      // 发行方ID
	ReleaseName                  string     // 发行方名称
	BlockchainId                 int64      // 区块链ID
	BlockchainName               string     // 区块链名字
	BlockchainKey                string     // 区块链key
	BlockchainIcon               string     // 区块链ICON
	ContractTokenId              int64      // 合约 token id
	ContractMetadataId           int64      // 商品ID
	ContractMetadataName         string     // 合约metadata对应的名字
	ContractMetadataDescription  string     // 合约metadata对应的描述
	ContractMetadataImage        string     // 合约metadata对应的固定图片
	ContractMetadataAnimationUrl string     // 合约metadata对应播放媒体地址
	Type                         int32      // 下单类型 0:购买 1:赠出 2:获赠 3:空投 4:合成
	UpdatedAt                    timef.Time // 更新时间
	CreatedAt                    timef.Time // 创建时间
	DeletedAt                    timef.Time // 标记删除时间
}

// 查询会员钱包信息ByUid
func (d *Dao) CreateTransfer(dao *Transfer) (int64, error) {

	m := convert.StructAssign(dao, &transfer_repo.Transfer{}).(*transfer_repo.Transfer)

	id, err := m.Create()

	if err != nil {
		return 0, err
	}
	return id, nil
}

// 获取列表
func (d *Dao) GetTransferListCountByUID(uid int64) (int64, error) {

	return transfer_repo.NewQueryBuilder().
		WhereRaw(" from_uid = ? OR to_uid = ?", uid, uid).
		Count()

	//var c int64
	//db := transfer_repo.Connection()
	//res := db.Where("from_uid = ?", uid).Or("to_uid = ?", uid).Model(&transfer_repo.Transfer{}).Count(&c)
	//if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
	//	c = 0
	//}
	//return c, res.Error
}

// 查询我的NFT交易
func (d *Dao) GetTransferListByUID(uid int64, page int, pageSize int) ([]*Transfer, error) {

	//listm := []*transfer_repo.Transfer{}
	//db := transfer_repo.Connection()
	//
	//err := db.Where("from_uid = ?", uid).Or("to_uid = ?", uid).
	//	Offset(app.GetPageOffset(page, pageSize)).
	//	Limit(pageSize).
	//	Order("id desc").
	//	Find(&listm).Error

	listm, err := transfer_repo.NewQueryBuilder().
		WhereRaw(" from_uid = ? OR to_uid = ?", uid, uid).
		OrderById(false).
		Offset(app.GetPageOffset(page, pageSize)).
		Limit(pageSize).
		Get()
	if err != nil {
		return nil, err
	}
	var list []*Transfer
	for _, m := range listm {
		list = append(list, convert.StructAssign(m, &Transfer{}).(*Transfer))
	}
	return list, nil
}

func (d *Dao) GetTransferById(id int64) (*Transfer, error) {
	m, err := transfer_repo.NewQueryBuilder().
		WhereId(model.Eq, id).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &Transfer{}).(*Transfer), nil
}
