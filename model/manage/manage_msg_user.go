package manage

// MallUserMsg 结构体
type MallUserMsg struct {
	Id       int    `json:"MsgId" form:"MsgId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId   int    `json:"UserId" form:"UserId" gorm:"column:user_id;comment:用户ID;"`
	MsgId    int    `json:"MsgId" form:"MsgId" gorm:"column:msg_id;comment:消息id;type:varchar(50);"`
	MsgText  string `json:"MsgText" form:"MsgText" gorm:"column:msg_text;comment:消息内容;type:varchar(50);"`
	ShowFlag int    `json:"ShowFlag" form:"ShowFlag" gorm:"column:show_flag;comment:是否显示;type:varchar(50);"`
}

func (MallUserMsg) TableName() string {
	return "tb_newbee_mall_user_msg"
}
