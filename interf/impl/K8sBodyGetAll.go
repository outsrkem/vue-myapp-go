package impl

import (
	"encoding/json"
	"fmt"
	"menu/db"
	"time"
)

/*
	从数据库获取所有k8s集群配置信息
	是接口的实现方法
*/
func (k K8sBodyList) K8sBodyGetAll() *K8sList {

	//调用数据库函数获取索引和总数据
	index := db.GetIndex(db.K8sList)
	dbList := db.GetAll(db.K8sList)

	//创建K8sList结构体类型对象，最终返回的数据结构
	k8sList := K8sList{
		Response: Response{
			//初始化结构体中嵌套的切片对象，否则添加数据会超出索引
			Items: make([]Items, index),
		},
	}

	k8sList.MetaInfo.Status = 200
	k8sList.MetaInfo.Msg = "success"
	k8sList.MetaInfo.RequestTime = time.Now().Unix()

	k8sList.Response.PageInfo.PageSize = 10
	k8sList.Response.PageInfo.PageNum = 1
	k8sList.Response.PageInfo.Page = 1
	k8sList.Response.PageInfo.Total = index

	var item Items
	//遍历数据库数据
	for i, k := range dbList {
		err := json.Unmarshal(k.Value, &item)
		if err != nil {
			fmt.Println(err)
		}
		//证书信息不返回，设置为空
		item.CaCrt = ""
		item.ClientCrt = ""
		item.ClientKey = ""
		k8sList.Response.Items[i] = item
	}

	return &k8sList
}
