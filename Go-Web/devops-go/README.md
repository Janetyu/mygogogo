## 概念

#### 什么是云计算

云计算跟云没有任何关系

- 虚拟化（Virtualization）是基础（计算、存储、网络等）
- 产品服务化（ laas 基础设施服务化，Paas 平台服务化 ，Saas 软件服务化 ... Xaas）
- 弹性伸缩， 没有边界

#### 云计算分类

- 公有云（AWS，阿里云，Azure等）
- 私有云（Vmware等）
- 混合云（Azure，Rackspace）

#### 公有云

- 云服务提供商对基础设施维护
- 多租户
- Pay For Use

#### 私有云

- 自己维护云基础设施
- 单租户或狭义上的多租户
- Pay For Cloud

#### 混合云（专属云）

- 云服务提供商维护自己的云设施
- 用户范围内租户隔离
- Pay For Use of Cloud

#### DevOps

什么是DevOps

- DevOps = Development + Operations
- 极速的迭代和快速的用户反馈

#### DevOps完整研发周期

Dev->CI/Build->Deploy->Ops->Feedback->Dev

#### 在阿里云上以DevOps简单实现网站搭建

- 基本业务代码实现
- 管理层面代码实现
- 代码托管打包构建
- 部署上线
- 设置监控告警
- 网站简单的自动恢复

## 网站上云

#### 阿里云简单介绍

- 云服务器 ECS
- 云监控

#### 简单的 web server 构建

- 利用github仓库
- 编写代码，并进行本地测试
- go install 打包代码，在 bin 文件夹中会生成一个可执行文件
- 因为需要部署到服务器上，因此用交叉编译打包代码
  - ` env GOOS=linux GOARCH=amd64 go build`
  - 在当前目录下生成了二进制可执行文件
- 利用github，把二进制文件上传到服务器上，git add，git status，git commit，git push origin master
- 在服务器上把仓库克隆下来，并直接启动二进制文件
  - 可以通过 `netstat -antp` 命令查看进程
  - 也可以用 `ps aux | grep name` 命令查看 “name” 的进程

#### 持续构建， 部署， 监控告警设置

- 上面操作存在的问题
  1. git pull
  2. git push -> git pull
  3. deploy
- 对这些人工操作进行自动化改造
  - 新建一个 sh 脚本，对一些部署所需要的操作进行记录
  - 在 main.go 中编写函数，通过 `cmd` 工具包对脚本进行调用启动完成自动化部署
  - 注意指定**不同的端口**进行启动，属于一个管理层面的程序
  - 设置 github 仓库中的 webhooks ，设置成 push 的事件启动后通知部署服务器执行脚本，注意填的是 **管理层面的服务器的URL地址**："xxx.xxx.xxx.xxx:5000"

#### 设置监控告警

- 使用云服务中的云监控产品服务
- 配置监控的云服务器以及站点管理