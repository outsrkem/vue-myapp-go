package impl

/*
	定义k8sconfig配置文件存储到数据库的结构体
*/
type K8sList struct {
	MetaInfo MetaInfo `json:"metaInfo"`
	Response Response `json:"response"`
}
type MetaInfo struct {
	Status      int    `json:"status"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}
type PageInfo struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
	Page     int `json:"page"`
	Total    int `json:"total"`
}
type Items struct {
	ID             int    `json:"id"`
	ClusterName    string `json:"clusterName"`
	UserName       string `json:"userName"`
	CurrentContext string `json:"currentContext"`
	Server         string `json:"server"`
	CreationTime   int64  `json:"creationTime"`
	Status         string `json:"status"`
	CaCrt          string `json:"caCrt"`
	ClientCrt      string `json:"clientCrt"`
	ClientKey      string `json:"clientKey"`
}

type Response struct {
	PageInfo PageInfo `json:"pageInfo"`
	Items    []Items  `json:"items"`
}
