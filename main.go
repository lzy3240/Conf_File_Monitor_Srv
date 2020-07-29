package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"zproject/Conf_File_Monitor_Srv/conf"

	"github.com/gin-gonic/gin"
	"github.com/lzy3240/mlog"
	"github.com/lzy3240/msql"
	"github.com/lzy3240/mtools"
	"gopkg.in/ini.v1"
)

type fileConf struct {
	Filename string `json:"filename"`
	Stage    string `json:"stage"`
	Flag     string `json:"flag"`
}

type fileResult struct {
	Filename string `json:"filename"`
	Flag     string `json:"flag"`
	Result   string `json:"result"`
}

type florderConf struct {
	Florder string `json:"florder"`
	Flag    string `json:"flag"`
	Stage   string `json:"stage"`
}

type florderResult struct {
	Florder string `json:"florder"`
	//Flag    string `json:"flag"`
	Result string `json:"result"`
}

var (
	cfg    = new(conf.AppConf)
	log    *mlog.Logger
	ms     msql.Msql
	confs  map[string][]*fileConf
	fconfs map[string][]*florderConf
)

//checkErr 检查错误
func checkErr(str string, err error) {
	if err != nil {
		log.Error(str, err)
	}
}

func getConf() {
	//1.读取每个IP的配置项
	confs = make(map[string][]*fileConf)
	ipSQLStr := "select distinct(host) from file_monitor"
	ips := ms.Queryby(ms.Db, ipSQLStr)
	if len(*ips) > 0 {
		for _, v := range *ips {
			tmpslice := make([]*fileConf, 0)        //针对每个IP创建配置
			ip, err := mtools.DecideType(v["host"]) //类型断言
			checkErr("data host type error:", err)
			tmpip := mtools.Convert2uft("gbk", "utf-8", ip) //转换
			//2.1 读取文件path
			pointSQLStr := "select filename,stage,flag from file_monitor where host ='" + tmpip + "'"
			point := ms.Queryby(ms.Db, pointSQLStr)
			if len(*point) > 0 {
				for _, k := range *point {
					filename, err := mtools.DecideType(k["filename"])
					checkErr("data filename type error:", err)
					stage, err := mtools.DecideType(k["stage"])
					checkErr("data stage type error:", err)
					flag, err := mtools.DecideType(k["flag"])
					checkErr("data flag type error:", err)
					// tfilename := mtools.Convert2uft("gbk", "utf-8", filename)
					// tstage := mtools.Convert2uft("gbk", "utf-8", stage)
					// tflag := mtools.Convert2uft("gbk", "utf-8", flag)
					//fmt.Println(filename, stage, flag)
					//2.保存到全局
					fileConfObj := new(fileConf)
					fileConfObj.Filename = filename
					fileConfObj.Stage = stage
					fileConfObj.Flag = flag
					tmpslice = append(tmpslice, fileConfObj)
					confs[tmpip] = tmpslice
				}
			}
		}
	}
	log.Info("flush file confs success")
}

func getFConf() {
	//1.读取每个IP的配置项
	fconfs = make(map[string][]*florderConf)
	ipSQLStr := "select distinct(host) from florder_monitor"
	ips := ms.Queryby(ms.Db, ipSQLStr)
	if len(*ips) > 0 {
		for _, v := range *ips {
			tmpslice := make([]*florderConf, 0)     //针对每个IP创建配置
			ip, err := mtools.DecideType(v["host"]) //类型断言
			checkErr("data host type error:", err)
			tmpip := mtools.Convert2uft("gbk", "utf-8", ip) //转换
			//2.1 读取文件path以及topic
			pointSQLStr := "select florder,flag,stage from florder_monitor where host ='" + tmpip + "'"
			point := ms.Queryby(ms.Db, pointSQLStr)
			if len(*point) > 0 {
				for _, k := range *point {
					florder, err := mtools.DecideType(k["florder"])
					checkErr("data florder type error:", err)
					flag, err := mtools.DecideType(k["flag"])
					checkErr("data flag type error:", err)
					stage, err := mtools.DecideType(k["stage"])
					checkErr("data stage type error:", err)
					// tflorder := mtools.Convert2uft("gbk", "utf-8", florder)
					// tflag := mtools.Convert2uft("gbk", "utf-8", flag)
					// tstage := mtools.Convert2uft("gbk", "utf-8", stage)
					//2.保存到全局
					florderConfObj := new(florderConf)
					florderConfObj.Florder = florder
					florderConfObj.Flag = flag
					florderConfObj.Stage = stage
					tmpslice = append(tmpslice, florderConfObj)
					fconfs[tmpip] = tmpslice
				}
			}
		}
	}
	log.Info("flush florder confs success")
}

func httpserver(str string) {
	r := gin.Default()
	//获取file相关配置信息
	r.GET("fileconf", func(c *gin.Context) {
		ip := c.ClientIP() //c.Query("query")
		var confslice []string
		confdata := confs[ip]
		for _, v := range confdata {
			jsonByte, err := json.Marshal(v)
			if err != nil {
				log.Error("marshal json faild,err:%v", err)
			}
			confslice = append(confslice, string(jsonByte))
		}

		//fmt.Println(strings.Join(confslice, "|"))
		c.JSON(http.StatusOK, gin.H{
			"fileconf": strings.Join(confslice, "|"),
		})
		log.Info("response %v,data:%v", ip, strings.Join(confslice, "|"))
	})

	//获取florder相关配置信息
	r.GET("florderconf", func(c *gin.Context) {
		ip := c.ClientIP() //c.Query("query")
		var confslice []string
		confdata := fconfs[ip]
		for _, v := range confdata {
			jsonByte, err := json.Marshal(v)
			if err != nil {
				log.Error("marshal json faild,err:%v", err)
			}
			confslice = append(confslice, string(jsonByte))
		}

		//fmt.Println(strings.Join(confslice, "|"))
		c.JSON(http.StatusOK, gin.H{
			"florderconf": strings.Join(confslice, "|"),
		})
		log.Info("response %v,data:%v", ip, strings.Join(confslice, "|"))
	})

	//接收file结果信息
	r.POST("fileresult", func(c *gin.Context) {
		ip := c.ClientIP()
		//读取result值
		date, _ := ioutil.ReadAll(c.Request.Body)
		//向客户端返回消息
		c.JSON(http.StatusOK, gin.H{
			"msg": "receive success",
		})
		var flres fileResult
		err := json.Unmarshal([]byte(date), &flres)
		if err != nil {
			fmt.Printf("err:%v\n", err)
		}
		//写库
		sqlstr := "insert into file_monitor_res set host=?,filename=?,flag=?,variation=?,varytime=CURRENT_TIMESTAMP"
		ms.Modifyby(ms.Db, sqlstr, ip, flres.Filename, flres.Flag, flres.Result)
		log.Info("insert success,%s:%s:%v", ip, flres.Filename, flres.Result)
	})

	//接收florder结果信息
	r.POST("florderresult", func(c *gin.Context) {
		ip := c.ClientIP()
		//读取result值
		date, _ := ioutil.ReadAll(c.Request.Body)
		//向客户端返回消息
		c.JSON(http.StatusOK, gin.H{
			"msg": "receive success",
		})
		var flres florderResult
		err := json.Unmarshal([]byte(date), &flres)
		if err != nil {
			fmt.Printf("err:%v\n", err)
		}
		//写库
		sqlstr := "update florder_monitor set num=? where host=? and florder=?"
		ms.Modifyby(ms.Db, sqlstr, flres.Result, ip, flres.Florder)
		log.Info("update success,%s:%s:%v", ip, flres.Florder, flres.Result)
	})
	//运行gin
	r.Run(str)
}

func init() {
	//1.1加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("init config faild,err:%v", err)
		return
	}
	//1.2初始化日志
	tmp := strings.Split(cfg.LogConf.CutParameter, "|")
	s := strings.Join(tmp, "=")
	log = mlog.Newlog(cfg.LogConf.Level, cfg.LogConf.Florder, cfg.LogConf.Perfix, s)

	//1.3初始化数据库
	ms = msql.NewMsql(cfg.MysqlConf.User, cfg.MysqlConf.Passwd, cfg.MysqlConf.Host, cfg.MysqlConf.Port, cfg.MysqlConf.Database, cfg.MysqlConf.Charset)
	//defer ms.Close()
	log.Info("Init mysql success")
}

func main() {
	//开启进程监听配置表变化
	go func() {
		for {
			getConf()
			getFConf()
			time.Sleep(time.Second * time.Duration(cfg.Server.Time))
		}
	}()
	//开启线程启动http服务
	httpserver(cfg.Server.Address)
}
