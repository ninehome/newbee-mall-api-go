package mall

// 如果含有time.Time 请自行import time包
type MallUserChat struct {
	ChatId    int    `json:"ChatId" form:"ChatId" gorm:"primarykey;AUTO_INCREMENT"`
	ChatName  string `json:"ChatName" form:"ChatName" gorm:"column:chat_name;comment:chat名称;type:varchar(30);"`
	ChatValue string `json:"ChatValue" form:"ChatValue" gorm:"column:chat_value;comment:注册名称;type:varchar(50);"`
	IsDeleted int    `json:"isDeleted" form:"isDeleted" gorm:"column:is_delete;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`
	Type      int    `json:"Type" form:"Type" gorm:"column:type;comment:删除标识字段(1-tg 2- ws  3-vb);type:tinyint"`
}

// TableName MallUserAddress 表名
func (MallUserChat) TableName() string {
	return "tb_newbee_mall_service"
}
