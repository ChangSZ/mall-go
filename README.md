## 关于

`go-mall` 是基于 [gin-boilerplate](https://github.com/ChangSZ/gin-boilerplate) 框架实现的一套电商系统的后台管理系统，包含商品管理、订单管理、会员管理、促销管理、运营管理、内容管理、统计报表、财务管理、权限管理、设置等模块。


> - **快速体验admin** --> [在线访问地址](http://mall.water-melon.top/admin) 
> - **快速体验app**   --> [在线访问地址](http://mall.water-melon.top/app)

</br>
! 本项目暂未未包含任何效果图, 可以去源项目查看, 链接如下:

本项目对[macrozheng/mall](https://github.com/macrozheng/mall)商城项目后端代码的重构 ==> [进度](./note.md)

mall_admin前端 --> [mall-admin-web](https://github.com/ChangSZ/mall-admin-web) 

mall_portal(app)前端 --> [mall-app-web](https://github.com/ChangSZ/mall-app-web)
<hr/>

## 友情提示
铁子们, 当前mall_admin、mall_portal基本功能均已OK, 使用中暂未发现问题, 我还未进行覆盖性测试, 你们可以搭建前后端自己先玩着. 有问题麻烦狠狠issue

工程相关文档后期会逐步补充完善

推荐先mark吧🤩, 靠谱楼主, 会持续更新的~

## 快速开始
### 拉取代码
```bash
$ git clone https://github.com/ChangSZ/mall-go.git
$ cd mall-go
```

### 环境准备
```bash
# 启动mysql、redis等中间件(也可以选择其他方式)
$ docker-compose -f deploy/docker-compose-env.yml up -d
# sql文件拷贝进mysql镜像中
$ docker cp internal/proposal/tablesqls/mall.sql mysql:/
# 进入mysql容器
$ docker exec -it mysql bash
# 登录数据库
$ mysql -u root -proot
# 创建数据库并退出登录
$ CREATE DATABASE api;  # 框架使用
$ CREATE DATABASE mall; 
$ exit;
# 数据导入
$ mysql -u root -p mall < mall.sql
```

### 开始运行
```bash
# 运行GO框架
$ go run main.go -env fat  
# 运行mall_admin
$ go run cmd/mall_admin/main.go -env fat
# 运行mall_portal
$ go run cmd/mall_portal/main.go -env fat
# -env 表示设置哪个环境，主要是区分使用哪个配置文件，默认为 fat
```

## Go框架
### 安装界面
首次启动程序之后，会在浏览器中自动打开安装界面，[链接地址](http://127.0.0.1:8080/render/install)

重新启动程序，会在浏览器中自动打开登录界面，[链接地址](http://127.0.0.1:8080)

输入默认账号 admin，密码 admin 即可登录成功

如果想重新安装，删除INSTALL.lock文件即可。该文件存在即认为无需安装。

### 格式化代码
```bash
  go run cmd/mfmt/main.go
```

### 重点使用
- 代码生成器
  - 生成数据表CURD - 选择对应的mysql数据表即可
  - 生成控制器方法 - 输入相对于/internal/api的目录路径即可
- 查询小助手
- 其他的可以用来学习娱乐

### 详细使用介绍
参见[go框架: gin-boilerplate](https://water-melon.top/detail/10)