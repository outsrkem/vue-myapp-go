package models

import (
	"fmt"
	"mana/src/config"
	"mana/src/connections/database/mysql"
	"strconv"
	"time"
)

type kubeConfig struct {
	ID                         string `json:"id"`                         // id
	USERID                     string `json:"userid"`                     // 用户id
	CLUSTER_ALIAS              string `json:"cluster_alias"`              // 集群别名
	CLUSTER_USER               string `json:"cluster_user"`               // 集群权限用户
	CURRENT_CONTEXT            string `json:"current_context"`            // 上下文
	SERVER                     string `json:"server"`                     // 集群地址
	CREATION_TIME              string `json:"creation_time"`              // 创建时间
	STATUS                     string `json:"status"`                     // 配置文件配置状态
	CERTIFICATE_AUTHORITY_DATA string `json:"certificate_authority_data"` // CA证书
	CLIENT_CERTIFICATE_DATA    string `json:"client_certificate_data"`    // 用户证书
	CLIENT_KEY_DATA            string `json:"client_key_data"`            // 用户证书私钥
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
	sqlStr1 := `INSERT INTO kube_config (USERID, CLUSTER_ALIAS, CLUSTER_USER, CURRENT_CONTEXT,
                         SERVER, CREATION_TIME, STATUS, CERTIFICATE_AUTHORITY_DATA, CLIENT_CERTIFICATE_DATA, CLIENT_KEY_DATA)
                         VALUES (?,?,?,?,?,?,?,?,?,?);`
	ret, err := mysql.DB.Exec(sqlStr1, k.USERID, k.CLUSTER_ALIAS, k.CLUSTER_USER, k.CURRENT_CONTEXT,
		k.SERVER, k.CREATION_TIME, k.STATUS, k.CERTIFICATE_AUTHORITY_DATA, k.CLIENT_CERTIFICATE_DATA, k.CLIENT_KEY_DATA)

	if err != nil {
		log.Error("insert kubeConfig failed,", err)
		return "", err
	}

	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		log.Error("get last Insert ID failed, err:", err)
		return "", err
	}

	return strconv.FormatInt(theID, 10), err
}

// 返回集群配置信息
// /api/v1/common/kubernetes/cluster
// ?page_size=9&page=1
func FindByKubeConfigs(uid string, pageSize int, page int) *kubeConfigStruct {
	var k = NewKubeConfig()
	var k8sJson = NewKubeConfigStruct()
	n := (page - 1) * pageSize
	m := pageSize
	sqlStr := `SELECT ID, USERID, CLUSTER_ALIAS, CLUSTER_USER, CURRENT_CONTEXT,
                         SERVER, CREATION_TIME, STATUS, CERTIFICATE_AUTHORITY_DATA, CLIENT_CERTIFICATE_DATA,
                         CLIENT_KEY_DATA FROM  kube_config WHERE USERID = ? ORDER BY USERID LIMIT ?, ?;`

	rows, err := mysql.DB.Query(sqlStr, uid, n, m)
	if err != nil {
		log.Error("exec  query failed, err", sqlStr, err)
	}

	defer rows.Close()

	// items里面的map
	var items []map[string]string
	items = make([]map[string]string, 0)

	for rows.Next() {
		err := rows.Scan(&k.ID, &k.USERID, &k.CLUSTER_ALIAS, &k.CLUSTER_USER, &k.CURRENT_CONTEXT,
			&k.SERVER, &k.CREATION_TIME, &k.STATUS, &k.CERTIFICATE_AUTHORITY_DATA, &k.CLIENT_CERTIFICATE_DATA, &k.CLIENT_KEY_DATA)
		if err != nil {

			fmt.Println("error")
		}

		item := make(map[string]string)
		item["id"] = k.ID
		item["creationTime"] = k.CREATION_TIME
		item["clusterName"] = k.CLUSTER_ALIAS
		item["userName"] = k.CLUSTER_USER
		item["server"] = k.SERVER
		item["currentContext"] = k.CURRENT_CONTEXT
		item["status"] = k.STATUS
		items = append(items, item)
	}
	// 查询记录总数,用于PageInfo信息
	var total int
	var pageNum int
	// 查询总条数
	totalRow, err := mysql.DB.Query("SELECT COUNT(*) FROM kube_config WHERE USERID = ?", uid)
	if err != nil {
		log.Error("GetKnowledgePointListTotal error", err)
		return nil
	}

	for totalRow.Next() {
		err := totalRow.Scan(
			&total,
		)
		if err != nil {
			log.Error("GetKnowledgePointListTotal error", err)
			continue
		}
	}
	// 计算页数
	if total%pageSize == 0 {
		//
		pageNum = total / pageSize
	} else {
		pageNum = total/pageSize + 1
	}

	// 页码相关信息
	k8sJson.Response.Items = items
	k8sJson.Response.PageInfo.Page = page
	k8sJson.Response.PageInfo.PageSize = pageSize
	k8sJson.Response.PageInfo.Total = total
	k8sJson.Response.PageInfo.PageNum = pageNum

	return k8sJson
}
