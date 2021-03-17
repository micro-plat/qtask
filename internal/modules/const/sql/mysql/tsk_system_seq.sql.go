package mysql
 
//tsk_system_seq 序列表
const tsk_system_seq=`
	DROP TABLE IF EXISTS tsk_system_seq;
	CREATE TABLE IF NOT EXISTS tsk_system_seq (
		seq_id bigint  not null auto_increment comment '编号' ,
		name varchar(32)  not null  comment '名称' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (seq_id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='序列表'`