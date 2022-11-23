package dao

import (
	"time"

	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/member_address_repo"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"

	"gorm.io/gorm"
)

type Address struct {
	Id          int64      //
	Uid         int64      // 用户ID
	LinkPhone   string     // 收货手机
	LinkMan     string     // 收货联系人
	LinkAddress string     // 收货人详细地址
	ZipCode     string     // 邮编
	ProvinceId  int32      // 省ID
	Province    string     // 省
	CityId      int32      // 市ID
	City        string     // 市
	AreaId      int32      // 区县ID
	Area        string     // 区县
	CountyId    int32      // 国家ID
	Weight      int32      // 排序
	IsDefault   int32      // 是否默认
	IsDeleted   int32      // 是否删除
	UpdatedAt   timef.Time `gorm:"time"` // 更新时间
	CreatedAt   timef.Time `gorm:"time"` // 创建时间
	DeletedAt   timef.Time `gorm:"time"` // 标记删除时间
}

// 获取我的地址数量
func (d *Dao) CountAddressListByUID(uid int64) (int64, error) {
	return member_address_repo.NewQueryBuilder().
		WhereIsDeleted(model.Eq, 0).
		WhereUid(model.Eq, uid).
		Count()
}

// 查询会员钱包信息ByUid
func (d *Dao) GetAddressListByUID(uid int64) ([]*Address, error) {

	listm, err := member_address_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereIsDeleted(model.Eq, 0).
		OrderById(false).
		Get()

	if err != nil {
		return nil, err
	}

	var list []*Address
	for _, m := range listm {
		dao := &Address{}
		convert.StructAssign(m, dao)
		list = append(list, dao)
	}
	return list, nil

}

func (d *Dao) CreateAddressByUID(uid int64, param *Address) (int64, error) {

	addressModel := &member_address_repo.MemberAddress{}
	convert.StructAssign(param, addressModel)
	addressModel.Uid = uid
	id, err := addressModel.Create()

	if err != nil {
		return 0, err
	}
	return id, nil

}

func (d *Dao) UpdateAddressByUidId(uid int64, id int64, param *Address) error {

	modelQuery := member_address_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereId(model.Eq, id).
		WhereIsDeleted(model.Eq, 0)

	addressModel, err := modelQuery.First()
	if err != nil {
		return err
	}
	convert.StructAssign(param, addressModel)

	addressModel.Save()

	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetAddress(uid int64, id int64) (*Address, error) {
	addressModel, err := member_address_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereId(model.Eq, id).
		WhereIsDeleted(model.Eq, 0).First()

	if err != nil {
		return &Address{}, err
	}

	addressDao := &Address{}
	convert.StructAssign(addressModel, addressDao)
	return addressDao, nil
}

func (d *Dao) GetAddressDefault(uid int64) (*Address, error) {
	addressModel, err := member_address_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereIsDefault(model.Eq, 1).
		WhereIsDeleted(model.Eq, 0).First()

	if err != nil {
		return &Address{}, err
	}

	addressDao := &Address{}
	convert.StructAssign(addressModel, addressDao)
	return addressDao, nil
}

func (d *Dao) ClearAddressDefault(uid int64, id int64) error {

	err := member_address_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereId(model.Neq, id).
		WhereIsDeleted(model.Eq, 0).
		Updates(map[string]interface{}{
			"IsDefault": 0,
		})

	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) DeleteAddress(uid int64, id int64) error {
	modelQuery := member_address_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereId(model.Eq, id).
		WhereIsDeleted(model.Eq, 0)

	count, err := modelQuery.Count()

	if err != nil {
		return err
	}

	if count <= 0 {
		return gorm.ErrRecordNotFound
	}

	err = modelQuery.Updates(map[string]interface{}{
		"IsDeleted": 1,
		"DeletedAt": time.Now(),
	})

	if err != nil {
		return err
	}
	return nil

}
