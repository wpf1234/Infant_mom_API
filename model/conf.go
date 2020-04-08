package model

type MysqlConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       string `json:"db"`
}

type RedisConf struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	Pool     int    `json:"pool"`
}

type LogConf struct {
	LogPath string `json:"log_path"`
	LogFile string `json:"log_file"`
}
