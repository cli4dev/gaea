# gaea

用于创建，监控基于 hydra 的项目

### 一、安装

```sh
  go get github.com/micro-plat/gaea
```

### 二、创建项目

#### 1. 生成项目

`gaea new project -n [项目名称]`

```sh
➜  gaea new project -n myproject/apiserver

生成文件: /home/colin/work/src/myproject/apiserver/main.go
生成文件: /home/colin/work/src/myproject/apiserver/bind.go

项目生成完成
```

### 2. 生成指定服务类型的项目

`gaea new project -n [项目名称] -s [服务器类型]`

```sh
➜  gaea new project -n myproject/apiserver -s api-cron

生成文件: /home/colin/work/src/myproject/myserver/main.go
生成文件: /home/colin/work/src/myproject/myserver/bind.go

项目生成完成
```

### 3. 给项目添加指定配置

`gaea new project  -s [服务器类型] -a`

```sh
myproject/apiserver ➜ gaea new project  -s cron -a

生成文件: /home/colin/work/src/myproject/myserver/main.go
生成文件: /home/colin/work/src/myproject/myserver/bind.go

项目生成完成
```

### 4. 根据数据表生成模块代码，sql语句

`gaea new module -c -r -u -d -add -t ./db.md -f trd_order_info -o modules`

```sh
// 1 使用默认的md文件生成到默认文件夹，多个操作加或追加 -add
➜ gaea new module -c -r -u -d -add

// 2 使用指定的 md 文件或指定生成目录 
➜ gaea new module -c -r -u -d -add -t ./db.md -o modules

// 3 使用过滤器生成指定数据表对应的模型代码
➜ gaea new module -c -r -u -d -add -f trd_order_info

// 4 覆盖原有代码并生成新模型代码
➜ gaea new module -c -r -u -d -cover

```

### 5. 生成服务代码

`gaea new service [项目名称] -s ["服务1 服务2"]`

```sh
gaea new service myproject/apiserver -s order/request
生成文件: /home/colin/work/src/myproject/apiserver/services/order/request.go

项目生成完成
```

### 6. 生成实体类

`gaea new struct [md文件路径] [输出路径] -f ["表名"]`

```sh
gaea new struct ../src/coupon/coupon/coupon.md ../src/coupon/coupon -f "sys_product_info"
生成文件: ../src/coupon/coupon/sys.product.info.go

struct生成完成,共生成1个文件
```

### 7. 生成数据库表创建语句

`gaea new sql [md文件路径] [输出路径]`

```sh
gaea new struct ../src/coupon/coupon/coupon.md ../src/coupon/coupon
[2018/09/04 09:28:15.739798][i][bf948ee43]生成文件: ../src/coupon/coupon/trd_order_info.sql
[2018/09/04 09:28:15.739893][i][bf948ee43]生成文件: ../src/coupon/coupon/prd_repository_time.sql
[2018/09/04 09:28:15.739941][i][bf948ee43]生成文件: ../src/coupon/coupon/prd_coupon_temp_pool.sql
[2018/09/04 09:28:15.739979][i][bf948ee43]生成文件: ../src/coupon/coupon/trd_order_refund.sql
[2018/09/04 09:28:15.740018][i][bf948ee43]生成文件: ../src/coupon/coupon/sys_supplier_record.sql
....
SQL生成完成,共生成35个文件
```
