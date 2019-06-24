
drop table tsk_system_task;

create table tsk_system_task
(
	task_id bigint not null PRIMARY KEY AUTO_INCREMENT  comment
	'编号' ,
		name varchar
	(32)  not null    comment '名称' ,
		create_time datetime default current_timestamp not null    comment '创建时间' ,
		last_execute_time datetime      comment '上次执行时间' ,
		next_execute_time int  not null    comment '下次执行时间' ,
		max_execute_time datetime  not null    comment '执行期限' ,
		interval int  not null    comment '时间间隔' ,
		count int default 0 not null    comment '执行次数' ,
		status int  not null    comment '状态' ,
		batch_id bigint      comment '执行批次号' ,
		queue_name varchar
	(64)  not null    comment '消息队列' ,
		msg_content varchar
	(256)      comment '消息内容' 
				
  )COMMENT='### 任务表';

 




