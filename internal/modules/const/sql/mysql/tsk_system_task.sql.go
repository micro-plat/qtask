package mysql
 
//tsk_system_task 任务表
const tsk_system_task=`
	DROP TABLE IF EXISTS tsk_system_task;
	CREATE TABLE IF NOT EXISTS tsk_system_task (
		task_id bigint  not null auto_increment comment '编号' ,
		name varchar(32)  not null  comment '名称' ,
		plat_name varchar(32)  not null  comment '平台名称' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		last_execute_time datetime    comment '上次执行时间' ,
		next_execute_time datetime  not null  comment '下次执行时间' ,
		max_execute_time datetime  not null  comment '执行期限(此时间前的任务可以被执行)' ,
		next_interval int  not null  comment '时间间隔,秒数' ,
		delete_interval int    comment '删除间隔,秒数' ,
		delete_time datetime    comment '删除期限' ,
		count int default 0 not null  comment '执行次数' ,
		max_count int default 0 not null  comment '最大执行次数' ,
		order_no varchar(32)    comment '外部业务单号' ,
		status tinyint  not null  comment '状态(20 等待，30 正在,0 已处理,90 处理失败)' ,
		batch_id bigint    comment '执行批次号' ,
		queue_name varchar(64)  not null  comment '消息队列' ,
		msg_content varchar(256)    comment '消息内容' 
		,index qtask_max_execute_time(max_execute_time)
		,index qtask_order_no(order_no)
		,primary key (task_id)
		,index qtask_info_batch_id(next_execute_time,batch_id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='任务表'`