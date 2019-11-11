
create table
if not exists  tsk_system_task
(
	task_id bigint not null PRIMARY KEY AUTO_INCREMENT  comment
	'编号' ,
		name varchar
(32)  not null    comment '名称' ,
		create_time datetime default current_timestamp not null    comment '创建时间' ,
		last_execute_time datetime      comment '上次执行时间' ,
		next_execute_time datetime  not null    comment '下次执行时间' ,
		max_execute_time datetime  not null    comment '执行期限' ,
		`next_interval` int
(10)  not null    comment '时间间隔' ,
		count int
(10) default 0 not null    comment '执行次数' ,
		status int
(2)  not null    comment '状态' ,
		batch_id bigint      comment '执行批次号' ,
		queue_name varchar
(64)  not null    comment '消息队列' ,
		msg_content varchar
(256)      comment '消息内容',
	
KEY `next_execute_time`
(`next_execute_time`,`status`) COMMENT 'idx_task_next_time',
  KEY `max_execute_time`
(`max_execute_time`,`status`) COMMENT 'idx_task_max_time',
KEY `next_execute_time_batch_id`
(`next_execute_time`,`batch_id`) COMMENT 'idx_task_batch_id',
				
  )COMMENT='### 任务表';

 




