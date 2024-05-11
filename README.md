## 关于

`go-mall` 是基于 [go-gin-api](https://github.com/xinliangnote/go-gin-api) 框架(基本已经改得面目全非了)实现的一套电商系统的后台管理系统，包含商品管理、订单管理、会员管理、促销管理、运营管理、内容管理、统计报表、财务管理、权限管理、设置等模块。

本项目对[macrozheng/mall](https://github.com/macrozheng/mall)商城项目后端代码的重构

web前端 ==> [mall-admin-web](https://github.com/ChangSZ/mall-admin-web) ==> [fork源](https://github.com/macrozheng/mall-admin-web)


## 快速开始
### 环境准备
- golang 1.21
- MySQL
  - 连接地址，例如：127.0.0.1:3306
  - 数据库名: mall，会在此数据库下初始化数据表
  - 用户名，不可为空
  - 密码，不可为空
- Redis
  - 连接地址，例如：127.0.0.1:6379
  - 密码，可为空
  - 连接DB，默认是 0

### 下载运行
```bash
$ git clone https://github.com/ChangSZ/mall-go.git
# 数据导入： internal\proposal\tablesqls\mall.sql
$ cd mall-go
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

### 使用介绍
参见[go-gin-api语雀](https://www.yuque.com/xinliangnote/go-gin-api), 前端几乎没动, 可以参考使用