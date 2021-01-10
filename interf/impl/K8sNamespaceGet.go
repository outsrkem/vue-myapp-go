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
	"time"
)

func (k *K8sNamespace) K8sNamespaceGet(address string) {
	//声明Items结构体，用来映射数据库获取到的k8s配置信息，如连接证书等
	var item Items

	//查询k8s数据库信息
	dpm := db.Get([]byte(address), db.K8sList)
	if dpm == nil {
		fmt.Println("K8sNamespaceGet查询数据库错误")
		return
	}

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
	resp, err := client.Get(address + "/api/v1/namespaces")
	if err != nil {
		fmt.Println("Get error:", err)
	}
	//函数执行完，关闭函数
	defer resp.Body.Close()

	//读取body中的json数据，如果错误，返回错误请求
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	k.MetaInfo.Status = 200
	k.MetaInfo.RequestTime = time.Now().Unix()
	k.MetaInfo.Msg = "success"

	k.Response.PageInfo.Total = 1
	k.Response.PageInfo.PageSize = 1
	k.Response.PageInfo.Page = 1
	k.Response.PageInfo.PageNum = 1

	array := gjson.Get(string(body), "items.#.metadata.name").Array()
	num := len(array)
	namespaceItem := make([]ItemsN, num)
	for i := 0; i < num; i++ {
		namespaceItem[i].Ns = array[i].Str
	}

	k.Response.Items = namespaceItem
}
