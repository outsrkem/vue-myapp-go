package impl

/*
	返回Linux数据结构体
*/
type LinuxList struct {
	MetaInfo MetaInfoL `json:"metaInfo"`
	Response ResponseL `json:"response"`
}

type MetaInfoL struct {
	Status      int    `json:"status"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}

type ResponseL struct {
	PageInfo PageInfoL `json:"pageInfo"`
	Items    []ItemsL  `json:"items"`
}

type PageInfoL struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
	Page     int `json:"page"`
	Total    int `json:"total"`
}

type ItemsL struct {
	Id       int    `json:"id"`
	Hostname string `json:"hostname"`
	HostIp   string `json:"hostip"`
	Memory   Memory `json:"memory"`
	Cpu      Cpu    `json:"cpu"`
	Disk     Disk   `json:"disk"`
}

type Memory struct {
	Total      int `json:"total"`
	Used       int `json:"used"`
	Section    int `json:"section"`
	Proportion int `json:"proportion"`
}

type Cpu struct {
	Total      int `json:"total"`
	Used       int `json:"used"`
	Section    int `json:"section"`
	Proportion int `json:"proportion"`
}

type Disk struct {
	Total      int `json:"total"`
	Used       int `json:"used"`
	Section    int `json:"section"`
	Proportion int `json:"proportion"`
}
