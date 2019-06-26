# qtask
管理实时任务和延时任务，并提供安装，配置功能。

实时任务：将任务存入DB,并放入消息队列。业务系统订阅队列消息，处理逻辑，根据需要结束任务。未结束的任务超时后自动放入队列，继续处理。

延时任务：将任务存入DB,超时后再放入消息队列。业务系统订阅队列消息，处理逻辑，根据需要结束任务。未结束的任务超时后自动放入队列，继续处理。

√　实时任务
√　延时任务
√　自动重发
√　DB存储
√　过期清理




## 示例:

#### 1. 创建任务

```go
//业务逻辑

//创建实时任务，将任务保存到数据库并发送消息队列
qtask.Create(c,"订单绑定任务",map[string]interface{}{
    "order_no":"8973097380045"
},3600,"GCR:ORDER:BIND")


//创建延时任务，将任务保存到数据库,超时后放入消息队列
qtask.Delay(c,"订单绑定任务",map[string]interface{}{
    "order_no":"8973097380045"
},3600,"GCR:ORDER:BIND")
```


#### 2. 编写MQC服务，该服务处理 `GCR:ORDER:BIND`消息队列数据

```go

func OrderBind(ctx *context.Context) (r interface{}) {
    //检查输入参数...
    
    //将任务修改为正在处理中,可以不调用
    qtask.Processing(ctx,ctx.Request.GetInt64("task_id"))


    //处理业务逻辑...


    //成功处理，结束任务
    qtask.Finish(ctx,ctx.Request.GetInt64("task_id"))
}

```

#### 3. 定时任务

a. 注册服务

```go

app.MQC("/order/bind",OrderBind) //消息处理任务

qtask.Bind(app,3)　//绑定扫描任务和定时删除过期任务

```

#### 4. 其它处理

1. 自定义数据库名，队列名
```go

qtask.Config("order_db","rds_queue") //配置数据库名，队列名

```

2. 创建数据库表
   
```go

qtask.CreateDB(ctx) //创建数据库结构

```


3. 使用不同的数据库
   
使用mysql数据库
```sh
 go install

```
或
```sh
 go install -tags "mysql"

```
使用oracle数据库
```sh
 go install -tags "oci" 

```