package kubernetes

import (
	"github.com/tidwall/gjson"
	"mana/src/filters/util"
	"mana/src/models"
	"strings"
)

//

// 获取工作负载结构体
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

// 获取工作负载
func GetWorkingLoad(cid, k8sLink string) *map[string]interface{} {

	// 调用函数，获取负载信息
	k, _ := models.FindByK8sConfigParticulars(cid)
	// 获取名称空间下 deploy
	url := k.Server + k8sLink
	// 传递证书及url
	k8s := util.NewK8sResources(k.CertificateAuthorityData, k.ClientCertificateData, k.ClientKeyData, url)
	// 获取信息
	body := util.K8sResourcesGet(k8s)

	/**
	遍历json数据中item数组，使用gjson模块
	*/
	////定义索引i，gjson遍历数组，没有索引，需要自己设置
	//i := 0
	//
	////计算数组长度
	//n := len(gjson.Get(string(*body), "items").Array())
	////初始化一个ItemsDpm结构体数组，用于存放遍历数据
	//item := make([]map[string]string, n)

	//foreach遍历item数组，如果遍历完成，返回false
	i := 0
	items := gjson.Get(string(*body), "items")
	n := len(items.Array())
	itemDpm := make([]ItemsDpm, n)
	//foreach遍历item数组，如果遍历完成，返回false
	items.ForEach(func(_, value gjson.Result) bool {
		itemDpm[i].ID = i + 1
		itemDpm[i].Name = value.Map()["metadata"].Map()["name"].Str
		itemDpm[i].Namespace = "namespace"

		//k8s时间转为csr时区
		CreationTimestamp := value.Map()["metadata"].Map()["creationTimestamp"]
		itemDpm[i].CreationTimestamp = CreationTimestamp.Time().UTC().Local().Format("2006-01-02 15:04:05")
		LastUpdateTime := value.Map()["status"].Map()["conditions"].Array()[0].Map()["lastUpdateTime"]
		itemDpm[i].LastUpdateTime = LastUpdateTime.Time().UTC().Local().Format("2006-01-02 15:04:05")

		replicas := value.Map()["spec"].Map()["replicas"].Int()
		itemDpm[i].Replicas = replicas
		//因为replicas参数如果为0，没有AvailableReplicas和unavailableReplicas参数
		if replicas == 0 {
			itemDpm[i].AvailableReplicas = 0
		} else if value.Map()["status"].Map()["availableReplicas"].Int() != 0 {
			itemDpm[i].AvailableReplicas = value.Map()["status"].Map()["availableReplicas"].Int()
		} else {
			itemDpm[i].AvailableReplicas = replicas - value.Map()["status"].Map()["unavailableReplicas"].Int()
		}

		//如果deployment有多容器，遍历多容器数组
		c := value.Map()["spec"].Map()["template"].Map()["spec"].Map()["containers"]
		l := len(c.Array())
		containersDpm := make([]ContainersDpm, l)
		y := 0
		c.ForEach(func(_, v gjson.Result) bool {
			containersDpm[y].Name = v.Map()["name"].Str
			containersDpm[y].Image = v.Map()["image"].Str
			version := v.Map()["image"].Str
			s := version[strings.LastIndex(version, ":")+1:]
			containersDpm[y].ImageVersion = s
			y++
			return true
		})

		//把ContainersDpm类型切片赋值哥返回对象
		itemDpm[i].Containers = containersDpm
		i++
		return true
	})

	// 返回数据
	pageInfo := models.NewPageInfo(1, 10, 1, n)
	msg := models.NewResMessage("200", "successfully")
	response := models.NewResponse(itemDpm, pageInfo)
	returns := models.NewReturns(response, msg)
	return &returns

}
