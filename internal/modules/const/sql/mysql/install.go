package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
)

func init() {
	//注册服务包
	hydra.OnReadying(func() error {
		hydra.Installer.DB.AddSQL(
			tsk_system_task,
			tsk_system_seq,
		)
		return nil
	})
}
