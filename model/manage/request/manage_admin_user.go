package request

type MallAdminLoginParam struct {
	UserName    string `json:"userName"`
	PasswordMd5 string `json:"passwordMd5"`
}

type MallAdminCreateParam struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	AgentId  string `json:"agentId"`
}
type MallAdminParam struct {
	LoginUserName string `json:"loginUserName"`
	LoginPassword string `json:"loginPassword"`
	NickName      string `json:"nickName"`
}

type MallUpdateNameParam struct {
	LoginUserName string `json:"loginUserName"`
	NickName      string `json:"nickName"`
}

type MallUpdatePswParam struct {
	PasswordMd5 string `json:"passwordMd5"`
	UserId      int    `json:"userId"`
}

type OrderStatusParam struct {
	OrderNo    string `json:"orderNo"`
	OrderMoney string `json:"orderMoney"`
}

type MallUpdateMoneyLevelParam struct {
	UserMoney     int `json:"userMoney"`
	UserLevel     int `json:"userLevel"`
	UserId        int `json:"userId"`
	RechargeMoney int `json:"rechargeMoney"`
	//LoginUserName int `json:"loginUserName"`

}

type MallUpdateWithdrawalParam struct {
	DealFlag   int `json:"dealFlag"`
	WithdrawId int `json:"withdrawId"`
}

type MallUpdateChatParam struct {
	ChatValue string `json:"ChatValue" `
	IsDeleted int    `json:"isDeleted"`
	ChatId    int    `json:"ChatId" "`
}

type MallUpdatePasswordParam struct {
	OriginalPassword string `json:"originalPassword"`
	NewPassword      string `json:"newPassword"`
}

type MallUserParam struct {
	UserMoney int    `json:"userMoney"`
	UserLevel int    `json:"userLevel"`
	UserId    string `json:"userId"`
	LoginName string `json:"loginName"`
}

type MallChatParam struct {
	ChatId string `json:"chatId"`
}

type BankParam struct {
	BankNumber string `json:"BankNumber" `
	UserId     int    `json:"userId"`
}

type BankUpdateParam struct {
	BankNumber string `json:"BankNumber" `
	BankId     int    `json:"BankId"`
}
type MsgParam struct {
	MsgTxt string `json:"MsgTxt" `
	UserId int    `json:"userId"`
}
