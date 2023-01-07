package request

type MallAdminLoginParam struct {
	UserName    string `json:"userName"`
	PasswordMd5 string `json:"passwordMd5"`
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
	OrderNo string `json:"orderNo"`
}

type MallUpdateMoneyLevelParam struct {
	UserMoney int `json:"userMoney"`
	UserLevel int `json:"userLevel"`
	UserId    int `json:"userId"`
	//LoginUserName int `json:"loginUserName"`
	//NickName      int `json:"nickName"`
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
}

type MallChatParam struct {
	ChatId string `json:"chatId"`
}
