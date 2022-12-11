package mall

// 如果含有time.Time 请自行import time包
type MallUserBank struct {
	BankId     int    `json:"bankId" form:"bankId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId     int    `json:"userId" form:"userId" gorm:"column:user_id;type:bigint"`
	BankName   string `json:"bankName" form:"bankName" gorm:"column:bank_name;comment:银行名称;type:varchar(30);"`
	UserName   string `json:"userName" form:"userName" gorm:"column:user_name;comment:注册名称;type:varchar(11);"`
	BankNumber string `json:"bankNumber" form:"bankNumber" gorm:"column:bank_number;comment:银行号码;type:varchar(32);"`
	Default    int    `json:"default" form:"default" gorm:"column:default;comment:区;type:varchar(32);"`
	IsDeleted  int    `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`
	//CreateTime common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:添加时间;type:datetime"`
	//UpdateTime common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"`
}

// TableName MallUserAddress 表名
func (MallUserBank) TableName() string {
	return "tb_newbee_mall_user_bank"
}
