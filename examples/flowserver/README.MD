# 使用方法　


### 1. 编译项目

使用oracle数据库
```go
 go install -tags "oci" 

```
使用mysql数据库
```go
 go install

```


### 2. 安装服务
```sh
 flowserver install -r zk://192.168.0.109 -c test
```

### 3. 运行服务
```sh
 flowserver run -r zk://192.168.0.109 -c test
```


### 4. 创建立即执行任务
```sh
 curl  http://192.168.4.121:9090/order/request/
```

### 5. 创建延迟执行任务
```sh
 curl  http://192.168.4.121:9090/order/request/
```