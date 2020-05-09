# acgfate
acgfate.com网站后端项目源码（未上线）

# 项目介绍

## 框架
1. [Gin](https://github.com/gin-gonic/gin): Go Web框架
2. [GORM](http://gorm.io/docs/index.html): ORM工具
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具
6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件

## 模块
1. api文件夹是MVC框架的controller，负责协调各部件完成任务
2. model文件夹负责存储数据库模型和数据库操作相关的代码
3. service负责处理比较复杂的业务
4. serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象
5. cache负责redis缓存相关的代码
6. auth权限控制文件夹
7. util一些通用的小工具
8. conf文件夹放静态存放的配置文件
9. validator文件夹放验证器

## 实现
- 用户模型
- 文章模型
- 分区投稿

# 运行项目
1. 创建.env文件，按照.env.example文件配置
2. 运行
```
go build
./acgfate.exe
```