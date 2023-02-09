package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var (
	AppMode     string
	HttpPort    string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbName      string
	JwtKey      string
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadFromEnv()
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("AllYourBase")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
	QiniuServer = file.Section("qiniu").Key("QiniuServer").MustString("QiniuServer")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("Bucket")
}

func LoadFromEnv() {
	AccessKey = os.Getenv("AccessKey")
	SecretKey = os.Getenv("SecretKey")
	if AccessKey == "" || SecretKey == "" {
		log.Fatalln("请先配置环境变量")
	}
}
