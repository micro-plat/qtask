drop table tsk_system_task;
drop sequence seq_qtask_system_task_id;
drop sequence seq_qtask_system_task_batch_id;


create table tsk_system_task
(
	task_id number(20) not null ,
	name varchar2(32) not null ,
	create_time date default sysdate not null ,
	last_execute_time date             ,
	next_execute_time date not null ,
	max_execute_time date not null ,
	next_interval number(10) not null ,
	count number(10) default 0 not null ,
	status number(2) not null ,
	batch_id number(20)       ,
	queue_name varchar2(64) not null ,
	msg_content varchar2(256)
);


comment on table tsk_system_task is '### 任务表';
	comment on column tsk_system_task.task_id is '编号';	
	comment on column tsk_system_task.name is '名称';	
	comment on column tsk_system_task.create_time is '创建时间';	
	comment on column tsk_system_task.last_execute_time is '上次执行时间';	
	comment on column tsk_system_task.next_execute_time is '下次执行时间';	
	comment on column tsk_system_task.max_execute_time is '执行期限';	
	comment on column tsk_system_task.next_interval is '时间间隔';	
	comment on column tsk_system_task.count is '执行次数';	
	comment on column tsk_system_task.status is '状态';	
	comment on column tsk_system_task.batch_id is '执行批次号';	
	comment on column tsk_system_task.queue_name is '消息队列';	
	comment on column tsk_system_task.msg_content is '消息内容';



alter table tsk_system_task
	add constraint pk_system_task primary key(task_id);


create index IDX_QTASK_INFO_TIME on TSK_SYSTEM_TASK (NEXT_EXECUTE_TIME, STATUS);


create sequence seq_qtask_system_task_id
	minvalue 10000
	maxvalue 99999999999
	start with 10000
	cache 20;

create sequence seq_qtask_system_task_batch_id
	minvalue 20000
	maxvalue 99999999999
	start with 20000
	cache 20;

create index IDX_QTASK_INFO_MAX_TIME on TSK_SYSTEM_TASK (MAX_EXECUTE_TIME, STATUS);
create index IDX_QTASK_INFO_BATCH_ID on TSK_SYSTEM_TASK (NEXT_EXECUTE_TIME, BATCH_ID);
