# 日志收集的agent开发

## 1. 配置文件版logagent

### ini配置文件解析
```go
cfg, err := ini.Load("./conf/config.ini")
	if err != nil {
		logrus.Error("load config failed,err:%v", err)
		return
	}
	kafkaAddr := cfg.Section("kafka").Key("address").String()
	fmt.Println(kafkaAddr)
```

## 2. 介绍etcd

类似于zookeeper,etcd\consul


## 3. logagent使用etcd管理收集项

