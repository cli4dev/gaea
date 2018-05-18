# gaea
用于创建，监控基于hydra的项目


### 安装
```sh
  go get github.com/micro-plat/wechat-token-server
```

### 创建项目

#### 1. 简单项目
```sh
gaea new project myproject/apiserver
创建文件: /home/yanglei/work/src/myproject/apiserver/main.go
创建文件: /home/yanglei/work/src/myproject/apiserver/bind.go
项目生成完成
```

### 2. 生成包含有模块代码的项目

```sh
gaea new project myproject/apiserver -m order/request
创建文件: /home/yanglei/work/src/myproject/apiserver/main.go
创建文件: /home/yanglei/work/src/myproject/apiserver/bind.go
创建文件: /home/yanglei/work/src/myproject/apiserver/services/order/request.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/order/request.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/sql/order.go
项目生成完成

```