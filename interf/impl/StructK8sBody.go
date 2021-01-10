package impl

import (
	"fmt"
)

//定义接收k8sconfig配置文件序列化为结构体
type K8sBodyList struct {
	APIVersion     string      `json:"apiVersion"`
	Clusters       []Clusters  `json:"clusters"`
	Contexts       []Contexts  `json:"contexts"`
	CurrentContext string      `json:"current-context"`
	Kind           string      `json:"kind"`
	Preferences    Preferences `json:"preferences"`
	Users          []Users     `json:"users"`
}

type Cluster struct {
	CertificateAuthorityData string `json:"certificate-authority-data"`
	Server                   string `json:"server"`
}
type Clusters struct {
	Cluster Cluster `json:"cluster"`
	Name    string  `json:"name"`
}
type Context struct {
	Cluster string `json:"cluster"`
	User    string `json:"user"`
}
type Contexts struct {
	Context Context `json:"context"`
	Name    string  `json:"name"`
}
type Preferences struct {
}
type User struct {
	ClientCertificateData string `json:"client-certificate-data"`
	ClientKeyData         string `json:"client-key-data"`
}
type Users struct {
	Name string `json:"name"`
	User User   `json:"user"`
}

//接口的空实现
func (k K8sBodyList) K8sDeploymentGet(namespace, control, address string) {
	fmt.Println("无实现")
}
