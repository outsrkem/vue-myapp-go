package impl

/*
	定义返回对象，K8s集群 Deployment信息的结构体
*/
type K8sDeploymentList struct {
	MetaInfo MetaInfoDpm `json:"metaInfo"`
	Response ResponseDpm `json:"response"`
}
type MetaInfoDpm struct {
	Status      int    `json:"status"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}
type PageInfoDpm struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
	Page     int `json:"page"`
	Total    int `json:"total"`
}
type ContainersDpm struct {
	Image        string `json:"image"`
	ImageVersion string `json:"imageVersion"`
	Name         string `json:"name"`
}
type ItemsDpm struct {
	ID                int             `json:"id"`
	Name              string          `json:"name"`
	Namespace         string          `json:"namespace"`
	Replicas          int64           `json:"replicas"`
	AvailableReplicas int64           `json:"availableReplicas"`
	CreationTimestamp string          `json:"creationTimestamp"`
	LastUpdateTime    string          `json:"lastUpdateTime"`
	Containers        []ContainersDpm `json:"containers"`
}
type ResponseDpm struct {
	PageInfo PageInfoDpm `json:"pageInfo"`
	Items    []ItemsDpm  `json:"items"`
}
