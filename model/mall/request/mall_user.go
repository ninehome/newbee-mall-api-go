package request

// 用户注册
type RegisterUserParam struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
	AgentId   string `json:"agentId"`
}

// 更新用户信息
type UpdateUserInfoParam struct {
	NickName      string `json:"nickName"`
	PasswordMd5   string `json:"passwordMd5"`
	IntroduceSign string `json:"introduceSign"`
}

type UserLoginParam struct {
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
	UserIpAddr  string `json:"userIpAddr"`
	AgentId     string `json:"agentId"`
}

// 用户提款
type WithdrawalParam struct {
	WithdrawMoney int `json:"withdrawMoney"`
	BankId        int `json:"bankId"`
}

type UserInfoParam struct {
	UserId string `json:"userId"`
}
