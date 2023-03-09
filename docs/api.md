### 接口

#### 一、给多个客户端发消息

**请求地址：** /send_clients

**请求方式：** GET

**URL参数**

| 字段   | 类型     | 是否必须 | 说明             |
|------|--------| -------- |----------------|
| from | string | 是       | 发送的客户端ID       |
| tos  | []string | 是       | 接收的客户端ID(可以多个) |
| msg  | string | 是       | 消息内容           |

**响应示例：**

```json
{
  "code": 200,
  "msg": "success",
  "data": {}
}
```

#### 二、给组发消息

**请求地址：** /send_groups

**请求方式：** GET

**URL参数**

| 字段   | 类型     | 是否必须 | 说明          |
|------|--------| -------- |-------------|
| from | string | 是       | 发送的客户端ID    |
| groups  | []string | 是       | 接收的组名(可以多个) |
| msg  | string | 是       | 消息内容        |

**响应示例：**

```json
{
  "code": 200,
  "msg": "success",
  "data": {}
}
```

#### 三、给系统发消息

**请求地址：** /send_machines

**请求方式：** GET

**URL参数**

| 字段   | 类型     | 是否必须 | 说明            |
|------|--------| -------- |---------------|
| from | string | 是       | 发送的客户端ID      |
| ips  | []string | 是       | 接收的机器IP(可以多个) |
| msg  | string | 是       | 消息内容          |

**响应示例：**

```json
{
  "code": 200,
  "msg": "success",
  "data": {}
}
```

#### 四、全局广播

**请求地址：** /broadcast

**请求方式：** GET

**URL参数**

| 字段   | 类型     | 是否必须 | 说明             |
|------|--------| -------- |----------------|
| from | string | 是       | 发送的客户端ID       |
| msg  | string | 是       | 消息内容           |

**响应示例：**

```json
{
  "code": 200,
  "msg": "success",
  "data": {}
}
```

#### 五、加入组

**请求地址：** /add_group

**请求方式：** GET

**URL参数**

| 字段   | 类型     | 是否必须 | 说明             |
|------|--------| -------- |----------------|
| client_id | string | 是       | 客户端ID |
| groups  | []string | 是       | 多个组名  |

**响应示例：**

```json
{
  "code": 200,
  "msg": "success",
  "data": {}
}
```

#### 六、退出组

**请求地址：** /del_group

**请求方式：** GET

**URL参数**

| 字段   | 类型     | 是否必须 | 说明    |
|------|--------| -------- |-------|
| client_id | string | 是       | 客户端ID |
| groups  | []string | 是       | 多个组名  |

**响应示例：**

```json
{
  "code": 200,
  "msg": "success",
  "data": {}
}
```

#### 七、在线列表

**请求地址：** /online_list

**请求方式：** GET

**响应示例：**

```json
{
  "code": 200,
  "msg": "success",
  "data": {
    "hpDLOYW37Hd6sD4C1bdXjinoVDlNFUA0lkDoYf9xFZGtMi7UK8E1YeXHrkmDYvAb": {
      "id": "hpDLOYW37Hd6sD4C1bdXjinoVDlNFUA0lkDoYf9xFZGtMi7UK8E1YeXHrkmDYvAb",
      "ip": "192.168.1.4"
    }
  }
}
```

#### 八、组列表

**请求地址：** /group_list

**请求方式：** GET

**URL参数**

| 字段   | 类型     | 是否必须 | 说明                |
|------|--------| -------- |-------------------|
| client_id | string | 是       | 客户端ID(如果不传，则是所有组) |

**响应示例：**

```json
{
  "code": 200,
  "msg": "success",
  "data": [
    "test"
  ]
}
```

#### 九、机器列表

**请求地址：** /machine_list

**请求方式：** GET

**响应示例：**

```json
{
  "code": 200,
  "msg": "success",
  "data": [
    "192.168.1.4"
  ]
}
```