## beego 学习项目 guess
- 介绍Beego的基本构成和常用模块，并通过案例演示如何快速完成web项目搭建。
  - bee new 新建项目

  - bee run 运行项目

  - bee generate 根据数据库生成代码

    - `bee generate scaffold user `

       `-fields="id:int64,name:string,gender:int,age:int" `

       `-driver=mysql `

      `-conn="root:@tcp(127.0.0.1:3306)/imooc"`

- 整体介绍实践项目的产品原型和技术架构并按照MVC架构拆分成三个部分，分别演示项目的编码实现。
  - View 视图，html
  - Controller 路由及业务代码
  - Model 数据库模型
  - Static 静态文件
  - logs 日志输出地址
  - check.go toolbox模块 监控检查
  - main.go 数据库初始化、日志设置初始化
- 讲解项目的打包与部署，并介绍使用toolbox模块对项目进行监控。
  - app.conf配置修改
    - 开发模式改为生产模式：prod
    - 设置成可监控 EnableAdmin = true   http://beego.me:8088/healthcheck
  - bee pack 项目打包