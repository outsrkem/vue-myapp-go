package config

type Config struct {
	Listen        string `json:"listen"`
	Port          string `json:"port"`
	DbUser        string `json:"db_user"`
	DbPassword    string `json:"db_password"`
	DbNetwork     string `json:"db_network"`
	DbServer      string `json:"db_server"`
	DbPort        string `json:"db_port"`
	DbName        string `json:"db_name"`
	DbMaxIdleConn int    `json:"db_MaxIdleConn"`
	DbMaxOpenConn int    `json:"db_MaxOpenConn"`
}

//var _cfg *Config = nil
// 配置解析
//func ParseConfig(path string) (*Config, error) {
//	file, err := os.Open(path)
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	render := bufio.NewReader(file)
//	decoder := json.NewDecoder(render)
//	if err = decoder.Decode(&_cfg); err != nil {
//		return nil, err
//	}
//	return _cfg, nil
//}

// 配置
func ParseConfig() *Config {
	var _cfg Config
	_cfg.Listen = "0.0.0.0"
	_cfg.Port = "9443"
	_cfg.DbUser = "abc"
	_cfg.DbPassword = "123456"
	_cfg.DbNetwork = "tcp"
	_cfg.DbServer = "10.10.10.10"
	_cfg.DbPort = "3306"
	_cfg.DbName = "mana"
	_cfg.DbMaxIdleConn = 5
	_cfg.DbMaxOpenConn = 10
	return &_cfg
}
