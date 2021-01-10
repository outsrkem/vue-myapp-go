package impl

import (
	"encoding/json"
	"fmt"
	"menu/db"
	"time"
)

/*
	获取到k8s配置信息，转为json格式的指定结构体数据，存入数据库
*/
func (k K8sBodyList) K8sBodyAdd() {

	//创建items的结构体类对象，并赋值
	var item Items

	item.ClusterName = k.Clusters[0].Name
	item.UserName = k.Users[0].Name
	item.CurrentContext = k.CurrentContext
	item.Server = k.Clusters[0].Cluster.Server
	item.CreationTime = time.Now().Unix()
	item.Status = "success"
	item.CaCrt = k.Clusters[0].Cluster.CertificateAuthorityData
	item.ClientCrt = k.Users[0].User.ClientCertificateData
	item.ClientKey = k.Users[0].User.ClientKeyData

	//以指定的key类型存入数据库，序列化为json格式的[]byte类型
	key := []byte(k.Clusters[0].Cluster.Server)
	val, err := json.Marshal(&item)
	if err != nil {
		fmt.Println(err)
	}

	db.Add(key, val, db.K8sList)
}
