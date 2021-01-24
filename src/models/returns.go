package models

var returns map[string]interface{}

func NewReturns(result interface{}, msg interface{}) map[string]interface{} {
	returns = make(map[string]interface{})

	returns["response"] = result
	returns["metaInfo"] = msg

	return returns
}

// response 中的页码信息
type pageInfo struct {
	Page     int `json:"page"`
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

func NewPageInfo(page, pageSize, pageNum, total int) *pageInfo {
	return &pageInfo{
		Page:     page,
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
	}
}
