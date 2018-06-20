# gaea
用于创建，监控基于hydra的项目


### 一、安装
```sh
  go get github.com/micro-plat/gaea
```

### 二、创建项目

#### 1. 简单项目
`gaea new project [项目名称]`

```sh
gaea new project myproject/apiserver
生成文件: /home/colin/work/src/myproject/apiserver/main.go
生成文件: /home/colin/work/src/myproject/apiserver/bind.go
项目生成完成
```

### 2. 生成包含有模块代码的项目
`gaea new project [项目名称] -m [模块名称]`

```sh
gaea new project myproject/apiserver -m order/request
生成文件: /home/colin/work/src/myproject/apiserver/main.go
生成文件: /home/colin/work/src/myproject/apiserver/bind.go
生成文件: /home/colin/work/src/myproject/apiserver/services/order/request.go
生成文件: /home/colin/work/src/myproject/apiserver/modules/order/request.go
生成文件: /home/colin/work/src/myproject/apiserver/modules/sql/order.go
项目生成完成
```

### 3. 生成包含多个模块代码的项目
`gaea new project [项目名称] -m ["模块1 模块2"]`

```sh
gaea new project myproject/apiserver -m "order/request order/query"
生成文件: /home/colin/work/src/myproject/apiserver/main.go
生成文件: /home/colin/work/src/myproject/apiserver/bind.go
生成文件: /home/colin/work/src/myproject/apiserver/services/order/request.go
生成文件: /home/colin/work/src/myproject/apiserver/services/order/query.go
生成文件: /home/colin/work/src/myproject/apiserver/modules/order/request.go
生成文件: /home/colin/work/src/myproject/apiserver/modules/sql/order.go
生成文件: /home/colin/work/src/myproject/apiserver/modules/order/query.go
项目生成完成
```


### 4. 生成Restful风格API代码
`gaea new project [项目名称] -r`

```sh
gaea new project myproject/apiserver -m order/request -r
生成文件: /home/colin/work/src/myproject/apiserver/main.go
生成文件: /home/colin/work/src/myproject/apiserver/bind.go
生成文件: /home/colin/work/src/myproject/apiserver/services/order/request.go
生成文件: /home/colin/work/src/myproject/apiserver/modules/order/request.go
生成文件: /home/colin/work/src/myproject/apiserver/modules/sql/order.go
项目生成完成
```


### 5. 生成指定服务类型的项目
`gaea new project [项目名称] -s[服务器类型]`

```sh
gaea new project myproject/myserver -m order/request -s api-cron
生成文件: /home/colin/work/src/myproject/myserver/main.go
生成文件: /home/colin/work/src/myproject/myserver/bind.go
生成文件: /home/colin/work/src/myproject/myserver/services/order/request.go
生成文件: /home/colin/work/src/myproject/myserver/modules/order/request.go
生成文件: /home/colin/work/src/myproject/myserver/modules/sql/order.go
项目生成完成
```


### 6. 项目中添加模块
`gaea new module [项目名称] -m ["模块1 模块2"]`

```sh
gaea new module myproject/myserver -m order/query
生成文件: /home/colin/work/src/myproject/myserver/services/order/query.go
生成文件: /home/colin/work/src/myproject/myserver/modules/order/query.go
项目生成完成
```
