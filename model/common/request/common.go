package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	PageNumber  int    `json:"pageNumber" form:"pageNumber"`   // 页码
	PageSize    int    `json:"pageSize" form:"pageSize"`       // 每页大小
	OrderStatus string `json:"orderStatus" form:"orderStatus"` //
	OrderNo     string `json:"orderNo" form:"orderNo"`         //
	LoginName   string `json:"loginName" form:"loginName"`     //
}

// GetById Find by id structure
type GetById struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

//type PageBean struct {
//	PagInfo     PageInfo `json:"params" form:"params"`
//	OrderStatus string   `json:"orderStatus" form:"orderStatus"`
//}
