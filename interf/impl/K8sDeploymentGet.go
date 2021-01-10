package impl

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"menu/db"
	"net/http"
	"strings"
	"time"
)

/*
	根据指定参数，连接对应集群，获取deployment信息
*/
func (k *K8sDeploymentList) K8sDeploymentGet(namespace, control, address string) {
	//声明Items结构体，用来映射数据库获取到的k8s配置信息，如连接证书等
	var item Items
	//查询k8s数据库信息
	dpm := db.Get([]byte(address), db.K8sList)
	//数据库[]byte数据反序列化为Items结构体
	err := json.Unmarshal(dpm.Value, &item)
	if err != nil {
		fmt.Println("错误是", err)
	}

	//解码k8s证书
	ca, _ := base64.StdEncoding.DecodeString(item.CaCrt)
	cl, _ := base64.StdEncoding.DecodeString(item.ClientCrt)
	clkey, _ := base64.StdEncoding.DecodeString(item.ClientKey)

	//创建证书校验程序
	pool := x509.NewCertPool()
	// 配置ca证书
	pool.AppendCertsFromPEM(ca)

	// 配置用户证书和私钥
	cliCrt, err := tls.X509KeyPair(cl, clkey)
	if err != nil {
		fmt.Println("x509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
			//关闭ssl校验
			InsecureSkipVerify: true,
		},
	}

	//创建http请求，请求机制为Transport
	client := &http.Client{Transport: tr}

	//http请求路径
	resp, err := client.Get(address + "/apis/apps/v1/namespaces/" + namespace + "/" + control)
	if err != nil {
		fmt.Println("Get error:", err)
	}
	//函数执行完，关闭函数
	defer resp.Body.Close()

	//读取body中的json数据，如果错误，返回错误请求
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		k.MetaInfo.Status = 502
		k.MetaInfo.Msg = "读取k8s返回体错误error"
		k.MetaInfo.RequestTime = time.Now().Unix()
		return
	}

	//执行到这里，返回正确请求
	k.MetaInfo.Status = 200
	k.MetaInfo.Msg = "success"
	k.MetaInfo.RequestTime = time.Now().Unix()

	/*
		遍历json数据中item数组，使用gjson模块
	*/

	//定义索引i，gjson遍历数组，没有索引，需要自己设置
	i := 0
	//读取items数组
	items := gjson.Get(string(body), "items")
	//计算数组长度
	n := len(items.Array())
	//初始化一个ItemsDpm结构体数组，用于存放遍历数据
	itemDpm := make([]ItemsDpm, n)
	//foreach遍历item数组，如果遍历完成，返回false
	items.ForEach(func(_, value gjson.Result) bool {
		itemDpm[i].ID = i + 1
		itemDpm[i].Name = value.Map()["metadata"].Map()["name"].Str
		itemDpm[i].Namespace = namespace

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

	//把ItemsDpm对象切片赋值哥返回结构体对象
	k.Response.Items = itemDpm

}

//接口的空实现
func (k K8sDeploymentList) K8sBodyGetAll() *K8sList {
	fmt.Println("无实现")
	return nil
}

//接口的空实现
func (k K8sDeploymentList) K8sBodyAdd() {
	fmt.Println("无实现")
}

func (k K8sDeploymentList) K8sBodyDel(address string) {
	fmt.Println("无实现")
}
