package models

// 定义返回k8s集群配置的结构体
// test
type kubeConfigStruct struct {
	MetaInfo metaInfos `json:"metaInfo"`
	Response resPonses `json:"response"`
}
type metaInfos struct {
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
	Status      string `json:"status"`
}

type pageInfos struct {
	Page     string `json:"page"`
	PageNum  string `json:"pageNum"`
	PageSize string `json:"pageSize"`
	Total    string `json:"total"`
}
type resPonses struct {
	Items    interface{} `json:"items"`
	PageInfo pageInfos   `json:"pageInfo"`
}
func NewKubeConfigStruct() *kubeConfigStruct {
	return &kubeConfigStruct{
	}
}
