package main

import (
	"os"

	_ "github.com/micro-plat/hydra/hydra"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/logger"
	"github.com/micro-plat/qtask/modules/const/sql/creator"
)

func main() {

	//处理输入参数
	defer logger.Close()
	logger := logger.New("qtask")
	if len(os.Args) < 3 {
		logger.Error("命令错误，请使用 ’qtask [数据库类型] [连接串名称]‘ 数据库类型:mysql,oracle，连接串信息:[用户名]:[密码]@[tns名称] 或 [用户名]:[密码]@[数据库名]/数据库ip")
		return
	}

	//构建数据库对象
	c, err := db.ParseConnectString(os.Args[1], os.Args[2])
	if err != nil {
		logger.Error(err)
		return
	}
	xdb, err := db.NewDB(os.Args[1], c, 1, 0, 600)
	if err != nil {
		logger.Error(err)
		return
	}

	//创建数据库
	err = creator.CreateDB(xdb)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("数据表创建成功")
}
