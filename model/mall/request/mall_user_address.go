package request

type AddAddressParam struct {
	UserName string `json:"userName"`

	UserPhone string `json:"userPhone"`

	DefaultFlag byte `json:"defaultFlag"` // 0-不是 1-是

	ProvinceName  string `json:"provinceName"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
}

type UpdateAddressParam struct {
	AddressId     string `json:"addressId"`
	UserId        int    `json:"userId"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	DefaultFlag   byte   `json:"defaultFlag"` // 0-不是 1-是
	ProvinceName  string `json:"provinceName"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
}

type BankParam struct {
	//BankId     int    `json:"bankId" form:"bankId" gorm:"primarykey;AUTO_INCREMENT"`
	//UserId     int    `json:"userId" form:"userId" gorm:"column:user_id;type:bigint"`
	BankName   string `json:"bankName"`
	UserName   string `json:"userName"`
	BankNumber string `json:"bankNumber"`
	//Default    string `json:"default" form:"default" gorm:"column:default;comment:区;type:varchar(32);"`
	//IsDeleted  int    `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`

}
