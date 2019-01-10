# gaea

用于创建，监控基于 hydra 的项目

### 一、安装

```sh
  go get github.com/micro-plat/gaea
```

### 二、创建项目

#### 1. 生成api项目

`gaea [项目类型] create -n apiserver`

```sh
//创建基础的 api 项目
➜ gaea api create -n ./apiserver

//创建带登录和菜单组件的api项目（常用）
➜ gaea api create -login "http://192.168.5.93:6688|B128F779D5741E701923346F7FA9F95C|convoy" -n ./apiserver
```

`可选命令和参数`

```
-n value              项目名称 (default: "./")
-p value              指定服务端口号 (default: ":9090")
-jwt                  是否启用jwt(不输入此参数默认为false)
-db value             指定数据库类型和数据库链接串(ora:test/123456@orcl136)
-cros                 是否启用跨域设置，默认不启用
-metric               启用metric配置
-cache                启用cache配置
-queue                启用queue配置
-appconf              启用appconf配置
-login value          启用 login 模块,参数例如: "http://0.0.0.0:6688|B128F7797FA9F95C|coupon"
```

### 2. 生成其他服务类型的项目

`2.1 生成cron项目`

```sh
➜  gaea cron create -n ./cronserver

可选的参数：
   -n value      项目名称(default: "./")
   -appconf      是否启用cron app配置
   -task         是否启用cron task配置
   -metric       是否启用cron metric 配置
   -db value     指定数据库类型和数据库链接串(ora:test/123456@orcl136)
   -cache        是否启用cache配置
   -queue        是否启用queue配置

```

`2.2 生成web项目`

```sh
➜  gaea web create -n ./webserver

可选的参数：
  -n value       项目名称 (default: "./")
  -p value       web port配置 (default: ":8090")
  -static        是否启用web static 配置
  -metric        是否启用web metric 配置
  -db value      指定数据库类型和数据库链接串(ora:test/123456@orcl136)
  -cache         是否启用cache配置
  -queue         是否启用queue配置
```

`2.3 生成mqc项目`

```sh
➜  gaea mqc create -n ./mqcserver

可选的参数：
  -n value       项目名称 (default: "./")
  -server        是否启用mqc server配置
  -queue         是否启用mqc queue 配置
  -metric        是否启用mqc metric 配置
  -db value      指定数据库类型和数据库链接串(ora:test/123456@orcl136)
  -cache         是否启用启用cache配置
```

`2.4 生成rpc项目`

```sh
➜  gaea rpc create -n ./rpcserver

可选的参数：
   -n value             项目名称 (default: "./")
   -p value             rpc port端口配置
   -metric              是否启用rpc metric 配置
   -db value            指定数据库类型和数据库链接串(ora:test/123456@orcl136)
   -cache               是否启用cache配置
   -queue               是否启用queue配置
```

`2.5 生成ws项目`

```sh
➜  gaea ws create -n ./wsserver

可选的参数：
   -n value       项目名称 (default: "./")
   -appconf       是否启用ws app配置
   -jwt           是否启用 ws jwt 配置
   -metric        是否启用ws metric 配置
   -db value      指定数据库类型和数据库链接串(ora:test/123456@orcl136)
   -cache         是否启用cache配置
   -queue         是否启用queue配置
```


`2.6 生成VUE项目`

```sh
➜  gaea vue create -n ./wsserver

可选的参数：
   -n value  vue项目名称 (default: "./")
   -f value  过滤器，指定表明或关键字
   -t value  指定数据表 md 文件
说明：
   该命令会根据数据表md文件生成前端vue项目，包括页面和路由，登录和菜单相关组件   
```

### 3. 添加，修改，移除项目配置

`3.1 添加，修改项目配置`

```sh
//以修改，添加 api项目db配置为例如
➜  cd apiserver
➜  gaea api cover -db "mysql:convoy:MsqlDb4567$%^&@tcp(192.168.0.36)/convoy?charset=utf8"

说明：
    1.添加和修改共用一个 cover 命令
    2.其他类型项目修改，添加配置与此类似
```

`3.2 移除项目配置`

```sh
//以移除 api项目db配置为例如
➜  cd apiserver
➜  gaea api remove -db

说明：
    1.其他类型项目修改，添加配置与此类似
```

### 4. 根据数据表生成模块代码，sql语句

`4.1 生成 module 和 sql`

```sh
➜ cd apiserver
➜ gaea module -crud -t ../convoy.md

可选的命令：
   -c          根据表结构生成 insert 函数,insert SQL语句，输入参数实体等文件
   -r          根据表结构生成 select 函数,select SQL语句，输入参数实体等文件
   -u          根据表结构生成 update 函数,update SQL语句，输入参数实体等文件
   -d          根据表结构生成 delete 函数,delet SQL语句，输入参数实体等文件
   --crud      根据表结构生成 insert select update delete 函数,SQL语句，输入参数实体等文件
   --cru       根据表结构生成 insert select update 函数,SQL语句，输入参数实体等文件
   -f value    过滤器，指定表明或关键字
   -t value    指定数据表 md 文件
   --db value  指定生成 mysql 或 oracle 的函数和sql,默认为 mysql (default: "mysql")
   -o value    生成的文件输出路径
   -a          是否执行追加crud函数和sql语句操作
   --cover     是否执行覆盖crud函数和sql语句操作
   -m          是否生成一个基础的module
   -n value    生成基础module的名字

说明：
    1.sql语句会生成到 modules/sql/const目录下
    2.module代码会根据表名自动创建文件夹和路径
```



### 5. 生成服务代码

`5.1 生成service代码`

```sh
➜ cd apiserver
➜ gaea module -crud -t ../convoy.md

可选的参数：
   -c         根据表结构生成 post handle函数
   -r         根据表结构生成 get and query handle 函数
   -u         根据表结构生成 update handle函数
   -d         根据表结构生成 delete handle函数
   --crud     根据表结构生成 crud handle函数
   -f value   过滤器，指定表明或关键字
   -t value   指定数据表 md 文件
   -o value   生成的文件输出路径
   -a         是否执行追加 handle 函数
   -cover     是否执行覆盖 handle 函数操作

说明：
      1.默认生成的代码位于 services/ 目录下
```


### 6. 数据库

`6.1 生成数据库建表语句`

```sh
➜ gaea create [mysql;oracle] -t ./convoy.md

可选参数：
   -f value       过滤表名
   -c             覆盖已存在的文件
   -t value       指定文件
```

`6.2 根据已经存在的数据库生成md文件`

```sh
➜ gaea md create -db "mysql:convoy_v2:convoy_v21901@tcp(192.168.0.36)/convoy_v2"

参数：
   -db value  数据库链接串(自动区分mysql与oracle)
```

### 7. 其他

`7.1 创建空的module和service函数及对象`

```
➜ gaea micservice create -m "Query,Save" -s "Get,Post" plat/base/info

可选的参数：
   -m value    生成modules中的函数,参数示例："Query,Save"
   -s value    生成服务层的函数，参数示例："Get,Delete"

说明：
   1. 生成服务层函数时，会自动注册路由
   2. plat/base/info  为代码生成的路径
```