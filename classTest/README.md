## 前后端完全分离登录注册

### 前端

​	使用原生html页面和Ajax来调用后端接口，传递参数。

### 后端

​	使用Go语言中的gin框架提供的router等功能实现，其中跨域访问是通过gin中的cors中间件来实现的。

### 搭建方法

1. 在classTest/webserver下执行go run main.go命令，启动后端服务器，监听8081端口。
2. 打开classTest下的server.html，在其中可以测试注册、登录功能。