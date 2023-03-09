# go-ws

使用gin，viper，websocket，etcd，grpc实现的分布式可扩展的websocket im系统

## 一、单机使用

1、下载

```shell
git clone https://github.com/lackone/go-ws.git
```

2、编译

```shell
// 编译适用于本机的版本
go build

// 编译Linux版本
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

// 编译Windows位版本
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

// 编译MacOS版本
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
```

3、修改配置 ./configs/app.yaml 中 ws.IsCluster 设置为 false。

4、运行

```shell
./go-ws -conf ./configs/app.yaml
```

如果不指定 -conf，则默认会读取 ./configs/app.yaml 文件。

## 二、分布式使用

1、修改 configs/app.yaml 中 ws.IsCluster 为 true，并配置 ws.AesKey 的key。

保证多个节点中 ws.AesKey 是相同的，不然无法正确解析。

2、修改 configs/app.yaml 中 snowflake.NodeId 的 ID，保证在节点中唯一，不然雪花算法生成有可能重复。

3、安装etcd

```shell
docker run -d --name etcd-server \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
```

当然，etcd 你可以安装多个，我这里测试只安装了一个。

修改 config/app.yaml 中的 etcd.Endpoints 添加上你配置的 etcd 节点。

4、安装nginx进行反向代理，比如我在 192.168.1.4 主机上安装 nginx。

然后反向代理主机 111，112，113 上的 ws 和 http 服务。

```shell
upstream ws_cluster {
    server 192.168.1.111:8882;
    server 192.168.1.112:8882;
    server 192.168.1.113:8882;
}

upstream http_cluster {
    server 192.168.1.111:8881;
    server 192.168.1.112:8881;
    server 192.168.1.113:8881;
}

server {
    listen  9992;
    server_name 0.0.0.0;
    
    location /ws {
        proxy_pass http://ws_cluster; # 代理转发地址
        proxy_read_timeout 60s; # 超时设置
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        # 启用支持websocket连接
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}

server {
    listen 9991;
    server_name 0.0.0.0;

    location / {
        proxy_pass http://http_cluster;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

5、然后在 111，112，113 主机上分别启动 ./go-ws，注意关闭防火墙或开放端。

6、这时就可以访问 ws://192.168.1.4:9992/ws 进行 websocket 连接了，连接成功会返回 client_id。

## 三、http接口

http的接口地址为 http://192.168.1.4:9991 这与你在上面 nginx 配置的端口一致。

1、[http接口文档](docs/api.md)


## 四、实现的功能

- [x] 分布式
- [x] 节点与节点间相互独立
- [x] 一对一或一对多个客户端发送消息
- [x] 一对一或一对多个组发送消息
- [x] 一对一或一对多个主机发送消息
- [x] 全局广播
- [x] 客户端添加组
- [x] 客户端删除组
- [x] 在线列表
- [x] 在线组列表
- [x] 在线主机列表
- [x] 服务端针对客户端的请求回调