package impl

type K8sNamespace struct {
	MetaInfo MetaInfoN `json:"metaInfo"`
	Response ResponseN `json:"response"`
}

type MetaInfoN struct {
	Status      int    `json:"status"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}
type PageInfoN struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
	Page     int `json:"page"`
	Total    int `json:"total"`
}
type ItemsN struct {
	Ns string `json:"ns"`
}
type ResponseN struct {
	PageInfo PageInfoN `json:"pageInfo"`
	Items    []ItemsN  `json:"items"`
}
