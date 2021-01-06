package mysql

const tsk_system_task = `
DROP TABLE IF EXISTS tsk_system_task;
CREATE TABLE IF NOT EXISTS tsk_system_task (
		task_id BIGINT(20)  NOT NULL AUTO_INCREMENT COMMENT '编号' ,
		NAME VARCHAR(32)  NOT NULL  COMMENT '名称' ,
		create_time DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL  COMMENT '创建时间' ,
		last_execute_time DATETIME    COMMENT '上次执行时间' ,
		next_execute_time DATETIME  NOT NULL  COMMENT '下次执行时间' ,
		max_execute_time DATETIME  NOT NULL  COMMENT '执行期限(此时间前的任务可以被执行)' ,
		next_interval BIGINT(10)  NOT NULL  COMMENT '时间间隔,秒数' ,
		delete_interval BIGINT(10)  COMMENT '删除间隔,秒数' ,
		delete_time DATETIME   COMMENT '删除期限' ,
		COUNT BIGINT(10) DEFAULT 0 NOT NULL  COMMENT '执行次数' ,
	    max_count BIGINT(10) DEFAULT 100 NOT NULL  COMMENT '最大执行次数' ,
		order_no VARCHAR(32)  COMMENT '外部业务单号' ,
		STATUS TINYINT(2)  NOT NULL  COMMENT '状态(20 等待，30 正在,0 已处理,90 处理失败)' ,
		batch_id BIGINT(20)    COMMENT '执行批次号' ,
		queue_name VARCHAR(64)  NOT NULL  COMMENT '消息队列' ,
		msg_content VARCHAR(256)    COMMENT '消息内容' ,
		PRIMARY KEY (task_id),
		KEY key_tsk_system_task_delete_time (delete_time),
		KEY QTASK_INFO_BATCH_ID (batch_id,next_execute_time),
		KEY qtask_max_execute_time (max_execute_time,next_execute_time)
		) ENGINE=INNODB  DEFAULT CHARSET=utf8 COMMENT='任务表';`
