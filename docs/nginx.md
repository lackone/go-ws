### docker安装nginx

```
docker run -d -p 80:80 -p 9991:9991 -p 9992:9992 --name nginx \
 -v $PWD/html:/usr/share/nginx/html \
 -v $PWD/nginx.conf:/etc/nginx/nginx.conf \
 -v $PWD/conf.d:/etc/nginx/conf.d \
 -v $PWD/logs:/var/log/nginx \
 nginx
```

### nginx反向代理

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