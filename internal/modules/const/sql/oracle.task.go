// +build oracle

package sql

import _ "github.com/micro-plat/qtask/internal/modules/const/sql/oracle"

const SQLGetSEQ = `select seq_qtask_system_task_id.nextval from dual`

const SQLGetBatch = `select seq_qtask_system_task_batch_id.nextval from dual`

const SQLCreateTask = `
insert into tsk_system_task
  (task_id,
   name,
   next_execute_time,
   max_execute_time,
   next_interval,
   delete_interval,
   status,
   max_count,
   order_no,
   queue_name,
   msg_content)
values
  (@task_id,
   @name,
   sysdate + #first_timeout / 24 / 60 / 60,
   sysdate + #max_timeout / 24 / 60 / 60,
   @next_interval,
   @delete_interval,
   20,
   @max_count,
   @order_no,
   @queue_name,
   @content)
`

const SQLProcessingTask = `
update tsk_system_task t set 
t.next_execute_time=sysdate+t.next_interval/24/60/60,
t.status=30,
t.count=t.count + 1,
t.last_execute_time=sysdate
where  t.task_id=@task_id 
and t.count < t.max_count
and t.status in(20,30)
`

const SQLFinishTask = `
update tsk_system_task t
set t.next_execute_time = to_date('2099-12-31', '%yyyy-%mm-%dd'),
    t.status            = 0,
    t.delete_time       = sysdate + t.delete_interval / 24 / 60 / 60
where t.task_id = @task_id
and t.status in (20, 30)`

const SQLUpdateTask = `
update tsk_system_task t set 
t.batch_id=@batch_id,
t.next_execute_time= sysdate+t.next_interval/24/60/60
where  t.max_execute_time > sysdate
and t.next_execute_time <= sysdate 
and t.status in(20,30) 
and t.count < t.max_count
and rownum <= 1000`

const SQLQueryWaitProcess = `
select queue_name,msg_content content 
from 
tsk_system_task t 
where t.batch_id=@batch_id
and t.next_execute_time > sysdate
and t.count < t.max_count
`

const SQLClearTask = `delete from tsk_system_task where delete_time < sysdate and status in (0, 90)`

const SQLFailedTask = `
update tsk_system_task t set 
t.delete_time = sysdate + decode(t.delete_interval, 0, 604800, t.delete_interval) / 24 / 60 / 60,
t.status = 90
where t.max_execute_time > sysdate - 7
and (t.max_execute_time < sysdate or t.count >= t.max_count) 
and t.status in (20, 30)
and rownum <= 1000
`
