package request

type GetTagRequest struct {
}

type TagListRequest struct {
	Name      string `json:"name"`
	State     uint8  `json:"state"`
	Page      int32  `json:"page"`
	Page_size int32  `json:"pageSize"`
}

// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
type CreateTagRequest struct {
	Name     string `json:"name"`
	State    uint8  `json:"state"`
	CreateBy string `json:"createBy"`
}

type UpdataTagRequest struct {
	Name       string `json:"name"`
	State      uint8  `json:"state"`
	ModifiedBy string `json:"modifieBy"`
}

type DeleteTagRequest struct {
}
