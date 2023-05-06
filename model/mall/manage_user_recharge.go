package mall

import (
	"main.go/model/common"
)

type Recharge struct {
	RechargeId int             `json:"rechargeId" form:"rechargeId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId     int             `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:tinyint;"`
	UserName   string          `json:"userName" form:"userName" gorm:"column:user_name;comment:登陆名称(默认为手机号);type:varchar(11);"`
	Money      int             `json:"userMoney" form:"userMoney" gorm:"column:money;comment:用户余额;type:int"`
	CreateTime common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:注册时间;type:datetime"`
	AgentId    string          `json:"agentId" form:"agentId" gorm:"column:agent_id;comment:代理id;type:varchar(50);"`
}

// TableName MallUserWithdraw 表名
func (Recharge) TableName() string {
	return "tb_newbee_mall_user_recharge"
}
