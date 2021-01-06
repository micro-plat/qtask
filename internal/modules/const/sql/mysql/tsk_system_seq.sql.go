package mysql

const tsk_system_seq = `
DROP TABLE IF EXISTS tsk_system_seq;
create table if not exists tsk_system_seq
(
	seq_id bigint not null PRIMARY KEY AUTO_INCREMENT comment '编号',
	name varchar(32)  not null    comment '名称',
	create_time datetime default current_timestamp not null    comment '创建时间',
 	KEY idx_create_time (create_time),
	UNIQUE KEY unq_name (name)
  )COMMENT='### 序列表';`
