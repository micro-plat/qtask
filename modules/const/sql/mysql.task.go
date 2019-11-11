// +build !oracle

package sql

const SQLGetSEQ = `insert into tsk_system_seq (name) values (@name)`

const SQLCreateTaskID = `insert into tsk_system_task
  (task_id,
   name,
   next_execute_time,
   max_execute_time,
   next_interval,
   status,
   queue_name,
   msg_content)
values
  (@task_id,
   @name, 
   date_add(now(),interval #first_timeout second),   
   date_add(now(),interval #max_timeout second),
   @next_interval,
   20,
   @queue_name,
   @content)`

const SQLProcessingTask = `update tsk_system_task t set t.next_execute_time=date_add(now(),interval t.next_interval second),
t.status=30,t.count=t.count + 1,t.last_execute_time=now()
where t.task_id=@task_id and t.status in(20,30)`

const SQLFinishTask = `update tsk_system_task t set t.next_execute_time= STR_TO_DATE('2099-12-31', '%Y-%m-%d'),
t.status=0
where t.task_id=@task_id and t.status in(20,30)`

const SQLUpdateTask = `update tsk_system_task t set 
t.batch_id=@seq_id,
t.next_execute_time= date_add(now(),interval t.next_interval second)
where t.status in(20,30) and t.next_execute_time < now() and t.max_execute_time > now()
limit 1000`

const SQLQueryWaitProcess = `select t.queue_name,t.msg_content content from tsk_system_task t
 where t.batch_id=@seq_id and t.next_execute_time > now()`

const SQLClearTask = `delete from tsk_system_task
where max_execute_time < date_add(now(),interval -#day day)`

const SQLClearSEQ = `delete from tsk_system_seq
where 1=1 &seq_id`
