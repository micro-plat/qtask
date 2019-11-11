// +build !oracle

package creator

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/lib4go/db"
)

//CreateDB 自定义安装程序
func CreateDB(xdb db.IDBExecuter) error {
	return db.CreateDB(xdb, "src/github.com/micro-plat/qtask/modules/const/sql/mysql")
}
