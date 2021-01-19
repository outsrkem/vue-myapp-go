package tablees

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

type KubeConfig interface {
	NewKubeConfig()
}
