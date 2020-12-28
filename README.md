# qtask

提供稳定，可靠的任务管理服务。支持两种任务：

* 实时任务
将任务放入消息队列，同时备份到DB。发生异常后从DB恢复到队列。


* 延时任务
将任务放入DB，达到延时时间后从DB拉取放入消息队列。发生异常后从DB恢复到队列。




特性:

√ 　实时任务

√ 　延时任务

√ 　自动备份

√ 　自动恢复

√ 　过期清理

√ 　一键安装

√ 　代码集成

√ 　基于 hydra 构建


## 一. 创建任务

任务分为实时任务和延时任务

### 1.实时任务  

```go

import "github.com/micro-plat/qtask"

// 创建实时任务，将任务放入指定的消息队列，并备份到DB
// taskID:任务编号
// invoke:执行函数句柄
taskID, invoke, err:=qtask.Create(c,"订单绑定任务",map[string]interface{}{
    "order_no":"8973097380045"
},3600,"ORDER:BIND")

//立即执行
err:=invoke(c)
```


### 2.延迟任务

```go

import "github.com/micro-plat/qtask"


//创建延时任务，将任务保存到数据库(状态为等待处理),超时后放入消息队列
// taskID:任务编号
taskID, err:=qtask.Delay(c,"订单绑定任务",map[string]interface{}{
    "order_no":"8973097380045"
},60,3600,"ORDER:BIND", qtask.WithDeadline(86400))
```

### 3.参数说明
* qtask.WithDeadline  秒数，设置任务执行截止时间,默认为604800(7天)
* qtask.WithDeleteDeadline 秒数，设置任务删除截止时间,默认为0.如果没有设置该参数,当任务执行成功,删除时间为当前时间,如果未执行成功,删除时间为当前时间加604800(7天),
* qtask.WithMaxCount 次数，设置任务最多执行次数,默认为100
* qtask.WithOrderNO 外部业务单号，设置外部业务单号,便于查询同一业务单号任务,默认空


## 二. 处理任务

### 1. 消息队列配置

```go
    hydra.Conf.Vars().Queue().Redis("queue", "192.168.0.111:6379"))
```

###  2. 添加、注册任务

```go
	hydra.Conf.MQC(mqc.WithRedis("queue")).Queue(queue.NewQueue("ORDER:BIND", "/order/pay"))
```

### 3. 编写处理代码
```go

import "github.com/micro-plat/qtask"


func OrderBind(ctx hydra.IContext) (r interface{}) {
    //检查输入参数...

    //业务处理前调用，修改任务状态为处理中(超时前未调用qtask.Finish，任务会被重新放入队列)
    qtask.Processing(ctx,ctx.Request.GetInt64("task_id"))


    //处理业务逻辑...


    //业务处理成功，修改任务状态为完成(任务不再放入队列),并修改删除截止时间
    qtask.Finish(ctx,ctx.Request.GetInt64("task_id"))
}

```

## 三、其它

### 1. 配置

```go
qtask.Config(qtask.WithDBName("oracle"), //数据库配置
qtask.WithQueueName("redis"), //队列配置
qtask.WithScanInterval(2)) //DB拉取时间
```

### 2. 数据库

* 使用 mysql 数据库

```sh
 go install
```

* 使用 oracle 数据库

```sh
 go install -tags "oracle"
```

[完整示例](https://github.com/micro-plat/qtask/tree/master/examples/flowserver)
