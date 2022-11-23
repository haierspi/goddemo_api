package service

import (
	"starfission_go_api/internal/dao"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type Address struct {
	Id          int64      `json:"ID"`          //ID
	LinkPhone   string     `json:"linkPhone"`   //收货手机
	LinkMan     string     `json:"linkMan"`     //收货联系人
	LinkAddress string     `json:"linkAddress"` //收货人详细地址
	ZipCode     string     `json:"zipCode"`     //邮编
	ProvinceId  int32      `json:"provinceID"`  //省ID
	Province    string     `json:"province"`    //省
	CityId      int32      `json:"cityID"`      //市ID
	City        string     `json:"city"`        //市
	AreaId      int32      `json:"areaID"`      //区县	ID
	Area        string     `json:"area"`        //区县
	CountyId    int32      `json:"countyID"`    //国家ID
	Weight      int32      `json:"weight"`      //排序
	IsDefault   int32      `json:"isDefault"`   //是否默认
	UpdatedAt   timef.Time `json:"updatedAt"`   //更新时间
	CreatedAt   timef.Time `json:"createdAt"`   //创建时间
}

type CreateAddressRequest struct {
	LinkPhone   string `form:"linkPhone"  binding:"required,min=2,max=100"`   //收货手机
	LinkMan     string `form:"linkMan"  binding:"required,min=2,max=100"`     //收货联系人
	LinkAddress string `form:"linkAddress"  binding:"required,min=2,max=100"` //收货人详细地址
	ZipCode     string `form:"zipCode"`                                       //邮编
	ProvinceId  int32  `form:"provinceID" binding:"required,gte=1"`           //省ID
	Province    string `form:"province" binding:"required,min=2,max=100"`     //省
	CityId      int32  `form:"cityID" binding:"required,gte=1"`               //市ID
	City        string `form:"city" binding:"required,min=2,max=100"`         //市
	AreaId      int32  `form:"areaID" binding:"required,gte=1"`               //区县	ID
	Area        string `form:"area" binding:"required,min=2,max=100"`         //区县
	CountyId    int32  `form:"countyID"`                                      //国家ID
	Weight      int32  `form:"weight" binding:"gte=0"`                        //排序
	IsDefault   int32  `form:"isDefault,default=0" binding:"oneof=0 1"`       //是否默认
}

type UpdateAddressRequest struct {
	Id          int64  `form:"ID" binding:"required,gte=1"`
	LinkPhone   string `form:"linkPhone"  binding:"required,min=2,max=100"`   //收货手机
	LinkMan     string `form:"linkMan"  binding:"required,min=2,max=100"`     //收货联系人
	LinkAddress string `form:"linkAddress"  binding:"required,min=2,max=100"` //收货人详细地址
	ZipCode     string `form:"zipCode"`                                       //邮编
	ProvinceId  int32  `form:"provinceID" binding:"required,gte=1"`           //省ID
	Province    string `form:"province" binding:"required,min=2,max=100"`     //省
	CityId      int32  `form:"cityID" binding:"required,gte=1"`               //市ID
	City        string `form:"city" binding:"required,min=2,max=100"`         //市
	AreaId      int32  `form:"areaID" binding:"required,gte=1"`               //区县	ID
	Area        string `form:"area" binding:"required,min=2,max=100"`         //区县
	CountyId    int32  `form:"countyID"`                                      //国家ID
	Weight      int32  `form:"weight" binding:"gte=0"`                        //排序
	IsDefault   int32  `form:"isDefault,default=0" binding:"oneof=0 1"`       //是否默认
}

type DeleteAddressRequest struct {
	Id int64 `form:"ID" binding:"required,gte=1"`
}

func (svc *Service) AddressList(uid int64) ([]*Address, int, error) {

	count, err := svc.dao.CountAddressListByUID(uid)
	if err != nil {
		return nil, 0, err
	}

	addressSDao, err := svc.dao.GetAddressListByUID(uid)
	if err != nil {
		return nil, 0, err
	}

	var svcList []*Address
	for _, address := range addressSDao {
		svcList = append(svcList, convert.StructAssign(address, &Address{}).(*Address))
	}

	return svcList, int(count), nil
}

func (svc *Service) AddressDefault(uid int64) (*Address, error) {

	addressDao, err := svc.dao.GetAddressDefault(uid)

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(addressDao, &Address{}).(*Address), nil
}

func (svc *Service) AddressCreate(uid int64, param *CreateAddressRequest) (*Address, error) {

	addressParamDao := convert.StructAssign(param, &dao.Address{}).(*dao.Address)

	id, err := svc.dao.CreateAddressByUID(uid, addressParamDao)
	if err != nil {
		return nil, err
	}

	if param.IsDefault == 1 {
		svc.dao.ClearAddressDefault(uid, id)
	}

	addressDao, err := svc.dao.GetAddress(uid, id)

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(addressDao, &Address{}).(*Address), nil
}

func (svc *Service) AddressUpdate(uid int64, param *UpdateAddressRequest) error {

	err := svc.dao.UpdateAddressByUidId(
		uid, param.Id, convert.StructAssign(param, &dao.Address{}).(*dao.Address),
	)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) AddressDelete(uid int64, param *DeleteAddressRequest) error {
	err := svc.dao.DeleteAddress(uid, param.Id)
	if err != nil {
		return err
	}

	return nil
}
