package interf

import "menu/interf/impl"

//定义接口k8s类型接口，如果要新增接口，需要增加对应的方法，所有该接口的结构体实现都要增加方法
type K8sInterface interface {
	//添加 k8sconfig 配置信息件到十几块
	K8sBodyAdd()
	//查询数据库所有k8s集群信息
	K8sBodyGetAll() *impl.K8sList
	//更加请求参数，查询对应deployment信息
	K8sDeploymentGet(namespace, control, address string)
	//删除指定k8s集群数据
	K8sBodyDel(address string)
}
