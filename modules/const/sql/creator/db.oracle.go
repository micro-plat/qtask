// +build oracle

package creator

import (
	"github.com/micro-plat/lib4go/db"
	_ "github.com/zkfy/go-oci8"
)

//自定义安装程序
func CreateDB(xdb db.IDBExecuter) error {
	return db.CreateDB(xdb, "src/github.com/micro-plat/qtask/modules/const/sql/oracle")
}
