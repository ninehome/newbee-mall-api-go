package response

type MallUserDetailResponse struct {
	NickName      string `json:"nickName"`
	LoginName     string `json:"loginName"`
	IntroduceSign string `json:"introduceSign"`
	UserMoney     int    `json:"userMoney"`
	UserLevel     int    `json:"userLevel"`
	LockedFlag    int    `json:"lockedFlag"`
}
