### docker安装nginx

由于需要nginx.conf文件，但我们本地没有，我们可以临时启一个nginx容器，从容器中拷贝一份到当前目录。

```
docker run -d --name tmp-nginx nginx
docker cp tmp-nginx:/usr/share/nginx/html $PWD/html
docker cp tmp-nginx:/etc/nginx/nginx.conf $PWD/nginx.conf
docker cp tmp-nginx:/etc/nginx/conf.d $PWD/conf.d
```

然后把这个临时容器删除

```
docker stop tmp-nginx
docker rm tmp-nginx
```

启一个新的nginx容器

```
docker run -d -p 80:80 -p 9991:9991 -p 9992:9992 --name nginx \
 -v $PWD/html:/usr/share/nginx/html \
 -v $PWD/nginx.conf:/etc/nginx/nginx.conf \
 -v $PWD/conf.d:/etc/nginx/conf.d \
 -v $PWD/logs:/var/log/nginx \
 nginx
```

### nginx反向代理

在conf.d目录下创建一个ws.conf文件，配置如下。

如果重启nginx容器不生效，可以把nginx容器删了，再重新启一个。

```
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