package manage

import (
	"main.go/model/common"
)

type MallUserWithdraw struct {
	WithdrawId    int             `json:"withdrawId" form:"withdrawId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId        int             `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:tinyint;"`
	BankId        int             `json:"bankId" form:"bankId" gorm:"column:bank_id;comment:银行id;type:tinyint;"`
	LoginName     string          `json:"loginName" form:"loginName" gorm:"column:login_name;comment:登陆名称(默认为手机号);type:varchar(11);"`
	UserMoney     int             `json:"userMoney" form:"userMoney" gorm:"column:user_money;comment:用户余额;type:int"`
	UserLevel     int             `json:"userLevel" form:"userLevel" gorm:"column:user_level;comment:用户等级;type:tinyint"`
	WithdrawMoney int             `json:"withdrawMoney" form:"withdrawMoney" gorm:"column:withdraw_money;comment:提款金额;type:int"`
	DealFlag      int             `json:"dealFlag" form:"dealFlag" gorm:"column:deal_flag;comment:出款处理状态(0-未处理 1-已出款 2-取消出款);type:tinyint"`
	CreateTime    common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:注册时间;type:datetime"`
	UserIpAddr    string          `json:"userIpAddr" form:"userIpAddr" gorm:"column:user_ip_addr;comment:用户IP(默认null);type:varchar(32);"`
	AgentId       string          `json:"agentId" form:"agentId" gorm:"column:agent_id;comment:代理id;type:varchar(50);"`
}

// TableName MallUserWithdraw 表名
func (MallUserWithdraw) TableName() string {
	return "tb_newbee_mall_user_withdraw"
}
