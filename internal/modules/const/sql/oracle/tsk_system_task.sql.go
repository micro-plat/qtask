package oracle

const tsk_system_task = `create table tsk_system_task(
		task_id number(20)  not null ,
		name varchar2(32)  not null ,
		create_time date default sysdate not null ,
		last_execute_time date   ,
		next_execute_time date  not null ,
		max_execute_time date  not null ,
		next_interval number(10)  not null ,
		delete_interval number(10) ,
		delete_time date ,
		count number(10) default 0 not null ,
		max_count number(10) default 100 not null ,
		order_no varchar2(32),
		status number(2)  not null ,
		batch_id number(20)   ,
		queue_name varchar2(64)  not null ,
		msg_content varchar2(256) 
		);
	

	comment on table tsk_system_task is '任务表';
	comment on column tsk_system_task.task_id is '编号';	
	comment on column tsk_system_task.name is '名称';	
	comment on column tsk_system_task.create_time is '创建时间';	
	comment on column tsk_system_task.last_execute_time is '上次执行时间';	
	comment on column tsk_system_task.next_execute_time is '下次执行时间';	
	comment on column tsk_system_task.max_execute_time is '执行期限(此时间前的任务可以被执行)';	
	comment on column tsk_system_task.next_interval is '时间间隔,秒数';	
	comment on column tsk_system_task.delete_interval is '删除间隔,秒数';	
	comment on column tsk_system_task.delete_time is '删除期限';	
	comment on column tsk_system_task.count is '执行次数';	
	comment on column tsk_system_task.max_count is '最大执行次数';	
	comment on column tsk_system_task.order_no is '外部业务单号';	
	comment on column tsk_system_task.status is '状态(20 等待，30 正在,0 已处理,90 处理失败)';	
	comment on column tsk_system_task.batch_id is '执行批次号';	
	comment on column tsk_system_task.queue_name is '消息队列';	
	comment on column tsk_system_task.msg_content is '消息内容';	

	

 
	alter table tsk_system_task
	add constraint pk_tsk_system_task primary key(task_id);
	create index qtask_delete_time on tsk_system_task(delete_time);
	create index qtask_max_execute_time on tsk_system_task(max_execute_time,next_execute_time);
	create index QTASK_INFO_BATCH_ID on tsk_system_task(batch_id,next_execute_time);
	
	create sequence seq_qtask_system_task_id
	increment by 1
	minvalue 1
	maxvalue 99999999999
	start with 1
	cache 20;

	create sequence seq_qtask_system_task_batch_id
	increment by 1
	minvalue 1
	maxvalue 99999999999
	start with 1
	cache 20;`
