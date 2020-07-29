package conf

//AppConf 配置文件
type AppConf struct {
	MysqlConf `ini:"mysql"`
	Server    `ini:"server"`
	LogConf   `ini:"log"`
}

//LogConf 配置文件
type LogConf struct {
	Level        string `ini:"level"`
	Florder      string `ini:"florder"`
	Perfix       string `ini:"perfix"`
	CutParameter string `ini:"cutparameter"`
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
