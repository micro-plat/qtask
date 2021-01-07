package oracle

import "github.com/micro-plat/hydra"

func init() {
	hydra.Installer.DB.AddSQL(tsk_system_task, tsk_system_seq)
}
