package mysql

import "github.com/micro-plat/hydra"

func init() {
	hydra.Installer.DB.AddSQL(tsk_system_seq, tsk_system_task)
}
