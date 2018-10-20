## 关于Beego框架Restful风格开发笔记



### 指令生成api项目

```go
bee api [appname]
```



## 运行项目

进入新建项目下 `cd $GOPATH/src/[appname]`

```go
bee run -gendoc=true -downdoc=true
```

启动命令运行后会自动下载swagger,通过http://localhost:8080/swagger/ 即可打开api接口网址进行测试

指令解释：

- -gendoc=true表示每次会自动化的build文档
- -downdoc=true表示会自动下载swagger文档查看器



## 生成的文件目录结构

```yaml
卷 项目代码 的文件夹 PATH 列表
卷序列号为 22C6-ECFB
D:.
│  imc.exe
│  lastupdate.tmp
│  main.go
│  swagger.zip
│  
├─conf
│      app.conf
│      
├─controllers
│      object.go
│      user.go
│      
├─models
│      object.go
│      user.go
│      
├─routers
│      commentsRouter_controllers.go
│      router.go
│      
├─swagger
│      favicon-16x16.png
│      favicon-32x32.png
│      index.html
│      oauth2-redirect.html
│      swagger-ui-bundle.js
│      swagger-ui-bundle.js.map
│      swagger-ui-standalone-preset.js
│      swagger-ui-standalone-preset.js.map
│      swagger-ui.css
│      swagger-ui.css.map
│      swagger-ui.js
│      swagger-ui.js.map
│      swagger.json
│      swagger.yml
│      
└─tests
        default_test.go
        

```



## 自动生成文件（代码）分析

1. app.conf是beego的默认配置文件

   ```go
   appname = imc
   httpport = 8080
   runmode = dev
   autorender = false
   copyrequestbody = true
   EnableDocs = true
   ```



   - `appname`： 应用名称，默认是beego。通过bee new或bee api创建，则是项目名称。beego.BConfig.AppName = "imc"。

   - `httpport`：应用监听端口，默认是8080。beego.BConfig.Listen.HTTPPort = 8080。

   - `runmode`：应用的运行模式，可选值为prod、dev或test。默认是dev开发模式。beego.BConfig.RunMode = "dev"。

   - `autorender`：是否模板自动渲染，默认值为true，对于API类型的应用，需要把该项设为false。beego.BConfig.WebConfig.AutoRender = true。
   - `copyrequestbody`：是否允许在HTTP请求时，返回原始请求体数据字节，默认为false（GET or HEAD or 上传文件请求除外）。beego.BConfig.CopyRequestBody = false。

   - `EnableDocs`：是否开启文档内置功能，默认为false。beego.BConfig.WebConfig.EnableDocs = true。

2. main.go

   程序入口，通过beego.Run()将beego的框架运行起来。当以dev开发者模式启动时，设置一些参数。

   `_ "mygithub/Beego/imc/routers"`是对router.go中的init函数运行进行初始化

   Go允许一个package中多个文件包含init函数，**执行顺序为文件名字母顺序**。

3. imc/routers/*.go

   `commentsRouter_controller.go`：

   - 根据router.go和controllers/*.go自动生成。

   - 从beego 1.3版本开始支持注解路由功能，用户无需在router中注册路由，只需要Include相应的controller，然后在controller的method方法上面写上router注释(// @router)即可。

   router.go:

   ```go
   // @APIVersion 1.0.0
   // @Title beego Test API
   // @Description beego has a very cool tools to autogenerate documents for your API
   // @Contact astaxie@gmail.com
   // @TermsOfServiceUrl http://beego.me/
   // @License Apache 2.0
   // @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
   package routers
   
   import (
   	"mygithub/Beego/imc/controllers"
   
   	"github.com/astaxie/beego"
   )
   
   func init() {
   	ns := beego.NewNamespace("/v1",
   		beego.NSNamespace("/object",
   			beego.NSInclude(
   				&controllers.ObjectController{},
   			),
   		),
   		beego.NSNamespace("/user",
   			beego.NSInclude(
   				&controllers.UserController{},
   			),
   		),
   	)
   	beego.AddNamespace(ns)
   }
   
   ```

   添加/v1/object/对应ObjectController中的注解路由。

   添加/v1/user/对应UserControoler中的注解路由。

4. imc/controllers/*.go

   `object.go`:

   ```go
   // @router / [post]
   
   该注解路由表示支持/v1/object/的POST方法(此处url为/，结合router.go中的/v1/object)
   
   // @router /:objectId [get]  
   
   该注解路由表示支持/v1/object/{objectId}的GET方法
   
   // @router / [get]
   
   该注解路由表示支持/v1/object的GET方法
   
   // @router /:objectId [put]
   
   该注解路由表示支持/v1/object/{objectId}的PUT方法
   
   // @router /:objectId [delete]
   
   该注解路由表示支持/v1/object/{objectId}的DELETE方法
   ```



   `user.go`:

   ```go
   // @router / [post]
   
   该注解路由表示支持/v1/user/的POST方法
   
   // @router / [get]
   
   该注解路由表示支持/v1/user/的GET方法
   
   // @router /:uid [get]
   
   该注解路由表示支持/v1/user/{uid}的GET方法
   
   // @router /:uid [put]
   
   该注解路由表示支持/v1/user/{uid}的PUT方法
   
   // @router /:uid [delete]
   
   该注解路由表示支持/v1/user/{uid}的DELETE方法
   
   // @router /login [get]
   
   该注解路由表示支持/v1/user/login的GET方法
   
   // @router /logout [get]
   
   该注解路由表示支持/v1/user/logout的GET方法
   ```

5. imc/models/*.go

object.go和user.go，代码不再罗列。model层主要数据与数据库相关的操作，自动生成的例子不含数据库操作，后续ORM的部分放在该目录下。



6. ORM使用

1. [安装MySQL 8](http://www.cnblogs.com/xiaochuizi/p/9274642.html)

2. 下载Go语言的驱动

go get github.com/go-sql-driver/mysql

3. 使用orm访问数据库

```go
func init() {
    // set default database
    orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8", 30)
 
    // register model
    orm.RegisterModel(new(models.Role))
 
    // create table
    orm.RunSyncdb("default", false, true)
}
```

注册模型之后，调用RunSyncdb，若对应模型的表不存在会自动创建。以上代码考虑放在main.go中。



4. 修改代码，实现ORMapping

   1. 新建Student model和对应的表

      - 在mysql中创建student表

        - ```mysql
          CREATE TABLE `student` ( 
            `Id` int(11) NOT NULL, 
            `Name` varchar(10), 
            `Birthdate` date , 
            `Gender` tinyint(1) , 
            `Score` int(11), 
            PRIMARY KEY (`Id`) 
          )
          ```

        - 

      - 在model文件夹下新建student.go，添加student对象

        - ```go
          type Student struct {
                 Id int
                 Name string
                 Birthdate string
                 Gender bool
                 Score int
          }
          ```

        - 

   2. 初始化ORM模块

      - 我们要通过ORM来操作对象和数据库，但是ORM需要初始化才能使用，我们需要在`main.go`文件中增加以下内容：

      - ```go
        import (
               "github.com/astaxie/beego/orm"
               _ "github.com/go-sql-driver/mysql"
        )
        func init() {
               orm.RegisterDriver("mysql", orm.DRMySQL)
               orm.RegisterDataBase("default", "mysql", "zengyi:123@tcp(127.0.0.1:3306)/testdb?charset=utf8")
        }
        ```

      - 这里需要注意的是数据库连接字符串和普通的写法不一样，要写成如下格式：
        用户名:密码@tcp(MySQL服务器地址:端口)/数据库名字?charset=utf8

   3. 提供数据库查询Student的方法

      - 接下来就是数据库访问方法了。我们可以仿照user.go一样，把方法都写在Student.go文件里面。这是完整的Student.go文件：

      - ```go
        package models
        
        import (
               "github.com/astaxie/beego/orm"
               "fmt"
               "time"
        )
        
        type Student struct {
               Id int
               Name string
               Birthdate string
               Gender bool
               Score int
        }
        
        func GetAllStudents() []*Student {
               o := orm.NewOrm()
               o.Using("default")
               var students []*Student
               q:= o.QueryTable("student")
               q.All(&students)
               return students
        
        }
        func GetStudentById(id int) Student{
               u:=Student{Id:id}
               o := orm.NewOrm()
               o.Using("default")
               err := o.Read(&u)
               if err == orm.ErrNoRows {
                      fmt.Println("查询不到")
               } else if err == orm.ErrMissPK {
                      fmt.Println("找不到主键")
               }
               return u
        }
        func AddStudent(student *Student) int{
               o := orm.NewOrm()
               o.Using("default")
               o.Insert(student)
               return student.Id
        }
        func UpdateStudent(student *Student) {
               o := orm.NewOrm()
               o.Using("default")
               o.Update(student)
        }
        
        func DeleteStudent(id int){
               o := orm.NewOrm()
               o.Using("default")
               o.Delete(&Student{Id:id})
        }
        
        func init() {
               // 需要在init中注册定义的model
               orm.RegisterModel(new(Student))
        }
        ```

   4. 创建StudentController对外提供Student的增加、删除、修改、查询一个或所有的方法

      - 这里我们也可以仿照usercontroller，直接改写成我们需要的StudentController.go。这是内容：

      - ```go
        package controllers
        
        import "github.com/astaxie/beego"
        import (
               "testApi/models"
               "encoding/json"
        )
        
        type StudentController struct {
               beego.Controller
        }
        // @Title 获得所有学生
        // @Description 返回所有的学生数据
        // @Success 200 {object} models.Student
        // @router / [get]
        func (u *StudentController) GetAll() {
               ss := models.GetAllStudents()
               u.Data["json"] = ss
               u.ServeJSON()
        }
        // @Title 获得一个学生
        // @Description 返回某学生数据
        // @Param      id            path   int    true          "The key for staticblock"
        // @Success 200 {object} models.Student
        // @router /:id [get]
        func (u *StudentController) GetById() {
               id ,_:= u.GetInt(":id")
               s := models.GetStudentById(id)
               u.Data["json"] = s
               u.ServeJSON()
        }
        // @Title 创建用户
        // @Description 创建用户的描述
        // @Param      body          body   models.Student true          "body for user content"
        // @Success 200 {int} models.Student.Id
        // @Failure 403 body is empty
        // @router / [post]
        func (u *StudentController) Post() {
               var s models.Student
               json.Unmarshal(u.Ctx.Input.RequestBody, &s)
               uid := models.AddStudent(&s)
               u.Data["json"] = uid
               u.ServeJSON()
        }
        // @Title 修改用户
        // @Description 修改用户的内容
        // @Param      body          body   models.Student true          "body for user content"
        // @Success 200 {int} models.Student
        // @Failure 403 body is empty
        // @router / [put]
        func (u *StudentController) Update() {
               var s models.Student
               json.Unmarshal(u.Ctx.Input.RequestBody, &s)
               models.UpdateStudent(&s)
               u.Data["json"] = s
               u.ServeJSON()
        }
        // @Title 删除一个学生
        // @Description 删除某学生数据
        // @Param      id            path   int    true          "The key for staticblock"
        // @Success 200 {object} models.Student
        // @router /:id [delete]
        func (u *StudentController) Delete() {
               id ,_:= u.GetInt(":id")
               models.DeleteStudent(id)
               u.Data["json"] = true
               u.ServeJSON()
        }
        ```

      - 这里需要注意的是，函数上面的注释是很重要的，有一定的格式要求，Swagger就是根据这些注释来展示的，所以必须写正确。

   5. 将StudentController注册进路由

      - 现在大部分工作已经完成，我们只需要把新的StudentController注册进路由即可，打开router.go，增加以下内容：

      - ```go
        beego.NSNamespace("/student", 
        beego.NSInclude( 
            &controllers.StudentController{}, 
        ), 
        ),
        ```

      - 当然对于系统默认的user和object，如果我们不需要，可以注释掉。

   6. 运行并通过Swagger测试