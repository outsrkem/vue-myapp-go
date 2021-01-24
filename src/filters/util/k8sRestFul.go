package util

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 查询k8s资源,接受证书及url，
// 证书为base64编码字符串
type k8sResources struct {
	CodedCa        string
	CodedClientPem string
	CodedClientKey string
	url            string
}

func NewK8sResources(ca, clientPem, clientKey, url string) *k8sResources {
	return &k8sResources{
		CodedCa:        ca,
		CodedClientPem: clientPem,
		CodedClientKey: clientKey,
		url:            url,
	}
}

func K8sResourcesGet(k *k8sResources) *[]byte {
	ca, _ := base64.StdEncoding.DecodeString(k.CodedCa)
	clientPem, _ := base64.StdEncoding.DecodeString(k.CodedClientPem)
	clientKey, _ := base64.StdEncoding.DecodeString(k.CodedClientKey)

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca)

	cliCrt, err := tls.X509KeyPair(clientPem, clientKey)
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return nil
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(k.url)
	if err != nil {
		fmt.Println("Get error:", err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return &body
}