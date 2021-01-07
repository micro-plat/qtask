package oracle

const tsk_system_seq = `create table tsk_system_seq(
	seq_id number(20)  not null ,
	name varchar2(32)  not null ,
	create_time date default sysdate not null 
	);


comment on table tsk_system_seq is '序列表';
comment on column tsk_system_seq.seq_id is '编号';	
comment on column tsk_system_seq.name is '名称';	
comment on column tsk_system_seq.create_time is '创建时间';	



alter table tsk_system_seq
add constraint pk_system_seq primary key(seq_id);

create sequence seq_system_seq
increment by 1
minvalue 1
maxvalue 99999999999
start with 1
cache 20;
`
