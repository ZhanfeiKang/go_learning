# 常用命令

**1. 启动zookeeper**

```
bin\windows\zookeeper-server-start.bat config\zookeeper.properties
```

**2. 启动kafka**

```
bin\windows\kafka-server-start.bat config\server.properties
```

**3. 启用消费者读取信息**

```
bin\windows\kafka-console-consumer.bat --bootstrap-server 127.0.0.1:9092 --topic shopping --from-beginning
```

**4. 启动es**

```bash
bin\elasticsearch.bat
```

**5. 启动kibana**

```bash
bin\kibana.bat
```



# etcd命令

`key-value`存储

**1. 放入值**
```
etcdctl.exe --endpoints=http://127.0.0.1:2379 put kkite "csy"
```

**2. 取值**
```
etcdctl.exe --endpoints=http://127.0.0.1:2379 get kkite
```

**3. 删除值**

```
etcdctl.exe --endpoints=http://127.0.0.1:2379 del kkite
```

# 内容回顾

## 1. context

控制goroutine,跨goroutine追踪
### 两个根节点：

- `context.Background()` 
- `context.TODO()`

### 四个方法：
- `context.WithCancel()`
- `context.WithTimeOut()`
- `context.WithDeadline()`
- `context.WithValue()`，注意事项:key推荐使用自己的类型

### pprof
性能调优的两个方向：CPU和内存
`go tool pprof cpu.pprof`
`go tool pprof mem.pprof`
调用图
火焰图

## 2. kafka

### kafka和nsq有什么区别？

`nsq`: 更多的是用来做消息队列。
`kafka`: 比较重量级的兼顾存储和消息队列。

### kafka的原理

见图

### Go语言往kafka中发数据

`sarama`:
windows平台要使用v1.19之前的版本