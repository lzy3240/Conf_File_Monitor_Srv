package conf

//AppConf 配置文件
type AppConf struct {
	MysqlConf `ini:"mysql"`
	Server    `ini:"server"`
}

//MysqlConf 配置文件
type MysqlConf struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Passwd   string `ini:"passwd"`
	Database string `ini:"database"`
	Charset  string `ini:"charset"`
}

//Server 配置文件
type Server struct {
	Address string `ini:"address"`
	Time    int    `ini:"time"`
}
