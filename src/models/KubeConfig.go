package models

import (
	"mana/src/config"
	"mana/src/connections/database/mysql"
	"strconv"
	"time"
)

type kubeConfig struct {
	ID                         interface{} `json:"id"`                         // id
	USERID                     interface{} `json:"userid"`                     // 用户id
	CLUSTER_ALIAS              interface{} `json:"cluster_alias"`              // 集群别名
	CLUSTER_USER               interface{} `json:"cluster_user"`               // 集群权限用户
	CURRENT_CONTEXT            interface{} `json:"current_context"`            // 上下文
	SERVER                     interface{} `json:"server"`                     // 集群地址
	CREATION_TIME              interface{} `json:"creation_time"`              // 创建时间
	STATUS                     interface{} `json:"status"`                     // 配置文件配置状态
	CERTIFICATE_AUTHORITY_DATA interface{} `json:"certificate_authority_data"` // CA证书
	CLIENT_CERTIFICATE_DATA    interface{} `json:"client_certificate_data"`    // 用户证书
	CLIENT_KEY_DATA            interface{} `json:"client_key_data"`            // 用户证书私钥
}

func NewKubeConfig() *kubeConfig {
	return &kubeConfig{
		USERID:                     "",
		CLUSTER_ALIAS:              "",
		CLUSTER_USER:               "",
		CURRENT_CONTEXT:            "",
		SERVER:                     "",
		CREATION_TIME:              "",
		STATUS:                     "1",
		CERTIFICATE_AUTHORITY_DATA: "",
		CLIENT_CERTIFICATE_DATA:    "",
		CLIENT_KEY_DATA:            "",
	}
}

var log = config.Log()

// 插入集群用户配置
func InstKubeConfig(k *kubeConfig) (string, error) {
	creationTime := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	k.CREATION_TIME = creationTime

	// 插入用户表信息
	sqlStr1 := `INSERT INTO kube_config (USERID, CLUSTER_ALIAS, CLUSTER_USER, CURRENT_CONTEXT, SERVER, CREATION_TIME, STATUS, CERTIFICATE_AUTHORITY_DATA, CLIENT_CERTIFICATE_DATA, CLIENT_KEY_DATA) VALUES (?,?,?,?,?,?,?,?,?,?);`
	ret, err := mysql.DB.Exec(sqlStr1, k.USERID, k.CLUSTER_ALIAS, k.CLUSTER_USER, k.CURRENT_CONTEXT, k.SERVER, k.CREATION_TIME, k.STATUS, k.CERTIFICATE_AUTHORITY_DATA, k.CLIENT_CERTIFICATE_DATA, k.CLIENT_KEY_DATA)

	if err != nil {
		log.Error("insert kubeConfig failed,", err)
		return "", err
	}

	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		log.Error("get lastinsert ID failed, err:", err)
		return "", err
	}

	return strconv.FormatInt(theID, 10), err
}
