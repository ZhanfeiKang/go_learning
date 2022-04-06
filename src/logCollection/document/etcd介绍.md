# etcd 介绍

## 1.etcd介绍

`etcd` 是使用Go语言开发的一个开源的、高可用的分布式key-value存储系统，可以用于配置共享和服务的注册和发现。

类似项目有 `zookeeper` 和 `consul`。

etcd具有以下特点：
- 完全复制：集群中的每个节点都可以使用完整的存档
- 高可用性：Etcd可用于避免硬件的单点故障或网络的问题
- 一致性：每次读取都会返回跨多主机的最新写入
- 简单：包括一个定义良好、面向用户的API(gRPC)
- 安全：实现了带有可选的客户端证书身份验证的自动化TLS
- 快速：每秒10000次写入的基准速度
- 可靠：使用Raft算法实现了强一致、高可用的服务存储目录

## 2.etcd应用场景

### 服务发现

服务发现要解决的也是分布式系统中最常见的问题之一，即在同一个分布式集群中的进程或服务，要如何才能找到对方并建立连接。本质上来说，服务发现就是想要了解集群中是否有进程在监听udp或tcp端口，并且通过名字就可以查找和连接。

### 配置中心

将一些配置信息放到etcd上进行集中管理。
这类场景的使用方式通常是这样：应用在启动的时候主动从etcd获取一次配置信息，同时，在etcd节点上注册一个Watcher并等待，以后每次配置有更新的时候，etcd都会实时通知订阅者，一次达到获取最新配置信息的目的。

### 分布式锁

![etcd](https://gitee.com/kkite/blogimg/raw/master/202204021116545.png)



因为etcd使用Raft算法保持了数据的强一致性，其次操作存储到集群中的值必然是全局一致的，所以很容易实现分布式锁。锁服务有两种使用方式，一是保持独占，二是控制时序。

- **保持独占即所有获取锁的用户最终只有一个可以得到。**etcd为此提供了一套实现分布式锁原子操作 `CAS(CompareAndSwap)` 的API。通过设置 `prevExist` 值，可以保证在多个节点同时去创建某个目录时，只有一个成功。而创建成功的用户就可以认为是获得了锁。
- 控制时序，即所有想要获得锁的用户都会被安排执行，但是**获得锁的顺序也是全局唯一的，同时决定了执行顺序。**etcd为此也提供了一套API（自动创建有序键），对一个目录建值时指定为 `POST` 动作，这样etcd会自动在目录下生成一个当前最大的值为键，存储这个新的值（客户端编号）。同时还可以使用API按顺序列出所有当前目录下的键值。此时这些键的值就是客户端的时序，而这些键中存储的值可以使代表客户端的编号。

## 3. 为什么用etcd而不用Zookeeper？

etcd实现的这些功能，Zookeeper都能实现。那为什么要用etcd而非直接使用Zookeeper呢？

相较之下，Zookeeper有如下缺点：
  1. 复杂：Zookeeper的**部署维护复杂**，管理员需要掌握一系列的知识和技能；而Paxos强一致性**算法也是素来以复杂难懂**而闻名于世；另外，Zookeeper的使用也比较复杂，需要安装客户端，官方只提供Java和C两种语言的接口。
  2. Java编写。这里不是对Java有偏见，而是Java本身就偏向于**重型应用**，它会引入大量的依赖。而运维人员则普遍希望保持强一致、高可用的机器集群尽可能简单，维护起来也不易出错。
  3. **发展缓慢**。Apache基金会项目特有的“Apache Way”在开源界饱受争议，其中一大原因就是由于基金会庞大的结构以及松散的管理导致项目发展缓慢。

而etcd作为一个后起之秀，其优点也很明显。 
  1. 简单。使用Go语言编写**部署简单**；使用HTTP作为接口**使用简单；使用Raft算法保证强一致性让用户易于理解**。
  2. **数据持久化**。etcd默认数据一更新就进行持久化。
  3. **安全**。etcd支持SSL客户端安全认证。

最后，etcd作为一个年轻的项目，真正告诉迭代和开发中，这既是一个优点，也是一个缺点。优点是它的未来具有无限的可能性，缺点是无法得到大项目长时间使用的检验。然而，目前CoreOS、Kubernetes和CloudFoundry等知名项目均在生产环境中使用了etcd，所以总的来说，etcd值得去尝试。

## 4. etcd架构

![etcd架构](https://gitee.com/kkite/blogimg/raw/master/202204021117653.png)

从etcd的架构图中我们可以看到，etcd主要分为四个部分。

- **HTTP Server**: 用于处理用户发送的API请求以及其他etcd节点的同步与心跳信息请求。
- **Store**: 用于处理etcd支持的各类功能的事务，包括数据索引、节点状态变更、监控与反馈、事件处理与执行等等，是etcd对用户提供的大多数API功能的具体实现。
- **Raft**: Raft强一致性算法的具体实现，是etcd的核心。
- **WAL**: Write Ahead Log（预写式日志），是etcd的数据存储方式。除了在内存中存有所有数据的状态以及节点的索引以外，etcd就通过WAL进行持久化存储。WAL中，所有的数据提交前都会事先记录日志。Snapshot是为了防止数据过多而进行的状态快照；Entry表示存储的具体日志内容。

## 5. etcd集群

etcd作为一个高可用键值存储系统，天生就是为集群化而设计的。由于Raft算法在做决策时需要多数节点的投票，所以etcd一般部署集群推荐奇数个节点，推荐的数量为3、5或者7个节点构成一个集群。

### 搭建一个3节点集群示例：

在每个etcd节点指定集群成员，为了区分不同的集群最好同时配置一个独一无二的`token`。

下面是提前定义好的集群信息，其中 `n1`、`n2`、`n3`表示3个不同的etcd节点。

```shell
TOKEN=token-01
CLUSTER_STATE=new
CLUSTER=n1=http://10.240.0.17:2380,n2=http://10.240.0.18:2380,n3=http://10.240.0.19:2380
```

在 `n1` 这台机器上执行以下命令来启动etcd：

```shell
etcd --data-dir=data.etcd --name n1 \
	--initial-advertise-peer-urls http://10.240.0.17:2380 --listen-peer-urls http://10.240.0.17:2380 \
	--advertise-client-urls http://10.240.0.17:2379 --listen-client-urls http://10.240.0.17:2379 \
	--initial-cluster ${CLUSTER} \
	--initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}
```

在 `n2` 这台机器上执行以下命令来启动etcd：

```shell
etcd --data-dir=data.etcd --name n2 \
	--initial-advertise-peer-urls http://10.240.0.18:2380 --listen-peer-urls http://10.240.0.18:2380 \
	--advertise-client-urls http://10.240.0.18:2379 --listen-client-urls http://10.240.0.18:2379 \
	--initial-cluster ${CLUSTER} \
	--initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}
```

在 `n3` 这台机器上执行以下命令来启动etcd：

```shell
etcd --data-dir=data.etcd --name n3 \
	--initial-advertise-peer-urls http://10.240.0.19:2380 --listen-peer-urls http://10.240.0.19:2380 \
	--advertise-client-urls http://10.240.0.19:2379 --listen-client-urls http://10.240.0.19:2379 \
	--initial-cluster ${CLUSTER} \
	--initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}
```



> etcd官网提供了一个可以公网访问的etcd存储地址。

到此etcd集群就搭建起来了，可以使用 `etcdctl` 来连接etcd。

```shell
export ETCDCTL_API=3
HOST_1=10.240.0.17
HOST_2=10.240.0.18
HOST_3=10.240.0.19
ENDPOINTS=$HOST_1:2379,$HOST_2:2379,$HOST_3:2379

etcdctl --endpoints=$ENDPOINTS member list
```



## 6. Go语言操作etcd

这里使用官方的etcd/clientv3包来连接etcd并进行相关操作。

安装

```shell
go get go.etcd.io/etcd/clientv3
```

**put和set**

注意：put是clientv3版本的命令！！！

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// 代码连接etcd

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v", err)
		return
	}
	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "daisy", "beautiful")
	// _, err = cli.Put(context.TODO(), "daisy", "beautiful")
	if err != nil {
		fmt.Printf("put to etcd failed,err:%v", err)
		return
	}
	cancel()

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gr, err := cli.Get(ctx, "daisy")
	if err != nil {
		fmt.Printf("get from etcd failed,err:%v", err)
		return
	}
	for _, ev := range gr.Kvs {
		fmt.Printf("key:%s  value:%s\n", ev.Key, ev.Value)
	}
	cancel()
}
```

**watch**

监控etcd中key的变化（创建\更改\删除）

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// watch

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v", err)
		return
	}
	defer cli.Close()

	// watch
	watchCh := cli.Watch(context.Background(), "daisy")

	// 通道没值会阻塞
	for wresp := range watchCh {
		for _, evt := range wresp.Events {
			fmt.Printf("type:%v key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}
}
```



## 总结

### logagent v1总结

1. 只能读取一个日志文件，不支持多个日志文件同时收集
2. 无法管理日志的topic

思路：

​	用etcd存储要收集的日志项，使用json格式数据：

```json
[
    {
        "path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/s4.log",
        "topic":"s4"
    },
    {
        "path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/web.log",
        "topic":"web_log"
    }

]
```





# logagent使用etcd管理收集项

# etcd初始化

```go
// 初始化etcd连接
err = etcd.Init([]string{configObj.EtcdConfig.Address})
if err != nil {
    logrus.Errorf("init etcd failed,err:%v", err)
    return
}
// 从etcd中拉取要收集日志的配置项
allConf, err := etcd.GetConf(configObj.EtcdConfig.CollectKey)
if err != nil {
    logrus.Errorf("get conf from etcd failed,err:%v", err)
    return
}
fmt.Println(allConf)
```

## 为每个单独的配置项启动tailtask

![image-20220403084934639](https://gitee.com/kkite/blogimg/raw/master/202204030849803.png)

## 管理日志收集项

程序启动之后**拉取了最新的配置之后**，就应该派一个小弟去监控etcd中`collect_log_conf`中key的变化

![image-20220403110446944](https://gitee.com/kkite/blogimg/raw/master/202204031104016.png)

**暂留的问题**

如果logagent停了需要记录上一次的位置，参考filebeat 



## logagent中多机器根据ip拉取配置

每台服务器上的logagent的收集项可能都不一致，我们需要让logagent去etcd中根据IP获取自己的配置

![image-20220403162744594](https://gitee.com/kkite/blogimg/raw/master/202204031627758.png)

### 如何获取本机的IP？

```go
// 获取本机ip的函数
func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}
```

### etcd中配置的key要注意使用的ip

```go
_, err = cli.Put(ctx, "collect_log_10.136.232.30_conf", str)

gr, err := cli.Get(ctx, "collect_log_10.136.232.30_conf")
```

### logagent中多机器根据ip拉取配置

```go
// 从etcd中拉取要收集日志的配置项
collectKey := fmt.Sprintf(configObj.EtcdConfig.CollectKey, ip)
allConf, err := etcd.GetConf(collectKey)
```



# influxDB

待学...（版本原因）

### 启动服务

http://127.0.0.1:8086



# grafana

展示数据的工具，监控，数据可视化

### 下载

grafana下载链接: https://grafana.com/grafana/download

### 安装

解压

### 启动服务

http://127.0.0.1:3000



# 内容回顾

## gopsutil包

- 是 `psutil` 的Go语言版本

## influxDB

- 时序数据库，开源，Go开发的
- 集群功能是收费的
- 类似项目：`OpenTSDB`

### influxDB 操作

- `tags`和`fields`的区别
- 插入数据

## Grafana

- 数据可视化工具，支持很多种数据源
- 告警功能
- 丰富图表
- 后端也是使用Go语言开发



# Elastic search(ES)

> ES版本要与kibana对应上

- 开源的搜索引擎，java开发，基于`Lucene` 
- https://www.liwenzhou.com/posts/Go/go_elasticsearch/

### 倒排索引(reversed index)

 	1. 云梦有云有梦
 	2. 云梦人杰地灵
 	3. 云梦好



- 分词
- 关键词的权重
- 关键词的频次



### Near Realtime（NRT）几乎实时

Elasticsearch是一个几乎实时的搜索平台。意思是，从索引一个文档到这个文档可以被搜索只需要一点点的延迟，这个时间一般为毫秒级。



### ES基本概念与关系型数据库的比较

| ES概念                                         | 关系型数据库       |
| ---------------------------------------------- | ------------------ |
| Index（索引）支持全文检索                      | Database（数据库） |
| Type（类型）                                   | Table（表）        |
| Document（文档），不同文档可以有不同的字段集合 | Row（数据行）      |
| Field（字段）                                  | Column（数据列）   |
| Mapping（映射）                                | Schema（模式）     |



`xpack.security.enabled: false`
