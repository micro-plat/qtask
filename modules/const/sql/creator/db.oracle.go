// +build oracle

package creator

import (
	"github.com/micro-plat/lib4go/db"
)

//自定义安装程序
func CreateDB(xdb db.IDB) error {
	return db.CreateDB(xdb, "src/github.com/micro-plat/qtask/modules/const/sql/oracle")
}
