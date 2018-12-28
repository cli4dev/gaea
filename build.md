### 1.1 生成后端项目

```
//在GOPAT下

//下面的命令会自动添加 jwt,cros 配置，登录模块和菜单模块 默认端口 9090，
//-login 参数由三部分组成，以"|"分隔，第一个参数为sso的host，第二个为签名的key,第三个为系统英文标识
//-n  项目的名字
//-t md文件路径
gaea api create -login "http://192.168.5.93:6688|B128F779D5741E701923346F7FA9F95C|convoy" -n ./apiserver

cd apiserver

//配置db
gaea api cover -db "mysql:convoy:MsqlDb4567$%^&@tcp(192.168.0.36)/convoy"

```

### 1.2 生成 module,sql 和 services

```
cd admin-api

//生成modules和sql
gaea module -crud -t ../convoy.md

//生成services,此命令会自动创建路由
gaea service -crud -t ../convoy.md

//导入依赖
godep save

//调整冲突的包名后，编译即可运行
```

### 2.1 生成前端 VUE 项目

```

gaea vue create -n mgrweb -t ./convoy.md

cd mgrweb

npm install

//生产环境 端口 8089
sh build.sh

//开发环境 端口 8085
npm run serve
```
