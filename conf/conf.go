package conf

import (
	"blog-admin-service/model"
	"github.com/joho/godotenv"
	"os"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()
	//
	//// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}
	//
	//// 连接数据库
	model.DataBase(os.Getenv("MYSQL_DSN"))
	//cache.Redis()
	//
	//// 启动定时任务
	//tasks.CronJob()
}
