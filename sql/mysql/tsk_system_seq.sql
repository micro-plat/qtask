
drop table tsk_system_seq;

create table tsk_system_seq
(
	seq_id bigint not null PRIMARY KEY AUTO_INCREMENT comment
	'编号' ,
	name varchar
	(32)  not null    comment '名称' 
				
  )COMMENT='### 序列表';

 




