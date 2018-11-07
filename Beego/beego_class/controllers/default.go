package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	//"class/models"
	"github.com/astaxie/beego/orm"
	"class/models"
	"path"
	"time"
	"math"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
/*
	//1.有ORM对象
	o := orm.NewOrm()
	//2.有一个要插入数据的结构体对象
	user := models.User{}
	//3.对结构体赋值
	user.Name = "111"
	user.Pwd = "222"
	//4.插入
	_,err := o.Insert(&user)
	if err != nil{
		beego.Info("插入失败",err)
		return
	}
*/

/*  查询有关代码
	//1.有ORM对象
	o := orm.NewOrm()

	//2.查询的对象
	user := models.User{}
	//3.指定查询对象字段值
	user.Name = "111"

	 user.Id = 1
	//4.查询
	err := o.Read(&user)

	err := o.Read(&user,"Name")
	if err != nil {
		beego.Info("查询失败", err)
		return
	}

	beego.Info("查询成功",user)
	*/
	/*
	//1.要有ORM对象   第一种更新
	o := orm.NewOrm()

	//2.需要更新的结构体对象
	user := models.User{}
	//3.查到需要更新的数据
	user.Id = 1
	err := o.Read(&user)

	//4.给数据重新赋值
	if err == nil{
		user.Name = "111"

		//5.更新
		_,err = o.Update(&user)
		if err != nil{
			beego.Info("更新失败",err)
			return
		}
	}
	*/
	/*

	//1.有ORM对象
	o := orm.NewOrm()
	//2.删除的对象
	user := models.User{}
	//3.指定删除哪一条数据
	user.Id = 1
	//4.删除
	_,err := o.Delete(&user)
	if err != nil{
		beego.Info("删除错误",err)
		return
	}


	c.Data["data"] = "今天中午吃饺子"
	c.TplName = "test.html"
	*/
	c.TplName = "register.html"
}
func (c *MainController) Post() {

	//1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	//2.对数据进行校验
	if userName == "" || pwd == ""{
		beego.Info("数据不能为空")
		c.Redirect("/register",302)
		return
	}
	//3.插入数据库
	o := orm.NewOrm()

	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_,err := o.Insert(&user)
	if err != nil{
		beego.Info("插入数据失败")
		c.Redirect("/register",302)
		return
	}
	//4.返回登陆界面
	c.Redirect("/login",302)
}

func (c*MainController)ShowLogin(){
	c.TplName = "login.html"
}

//登陆业务处理
func (c*MainController)HandleLogin(){
	//c.Ctx.WriteString("这是登陆的POST请求")
	//1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	//2.判断数据是否合法
	if userName == ""|| pwd ==""{
		beego.Info("输入数据不合法")
		c.TplName = "login.html"
		return
	}
	//3.查询账号密码是否正确
	o := orm.NewOrm()
	user := models.User{}

	user.Name = userName
	err := o.Read(&user,"Name")
	if err != nil {
		beego.Info("查询失败")
		c.TplName = "login.html"
		return
	}
	//4.跳转
	c.Redirect("/index",302)
}

//显示列表页面内容
func (c*MainController)ShowIndex(){
	o := orm.NewOrm()
	id,_ := c.GetInt("select")


	var articles []models.Article
	_,err := o.QueryTable("Article").All(&articles)
	if err != nil{
		beego.Info("查询所有文章信息出错")
		return
	}



	//分页处理
	//获得数据总数，总页数，当前页码
	count,err := o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__Id",id).Count()
	if err != nil {
		beego.Info("查询失败",err)
		return
	}
	pagesize := int64(2)  //每页显示数据条目

	index,err := c.GetInt("pageIndex")  //当前页码
	if err != nil{
		index = 1
	}


	pageCount := math.Ceil(float64(count) / float64(pagesize))   //总页数

	if index <=0 || index > int(pageCount){
		index = 1
	}

	start := (int64(index)  -1 ) * pagesize
// inner   left join
	var artis []models.Article
	o.QueryTable("Article").Limit(pagesize,start).RelatedSel("ArticleType").Filter("ArticleType__Id",id).All(&artis)


	//获取类型数据
	var artiTypes []models.ArticleType
	_,err = o.QueryTable("ArticleType").All(&artiTypes)
	if err != nil{
		beego.Info("获取类型错误")
		return
	}

	c.Data["articleType"] = artiTypes

	c.Data["pageCount"] = pageCount
	c.Data["count"] = count
	c.Data["articles"] = artis
	c.Data["pageIndex"] = index
	c.Data["typeid"] = id    //文章类型ID

	c.TplName = "index.html"
}

//显示添加文章界面
func (c*MainController)ShowAdd(){
	o := orm.NewOrm()
	var artiTypes []models.ArticleType
	_,err := o.QueryTable("ArticleType").All(&artiTypes)
	if err != nil{
		beego.Info("获取类型错误")
		return
	}

	c.Data["articleType"] = artiTypes

	c.TplName = "add.html"
}

//处理添加文章界面数据
func (c*MainController)HandleAdd(){
	//1.拿到数据
	artiName := c.GetString("articleName")
	artiContent := c.GetString("content")
	id,err :=c.GetInt("select")
	if err !=nil{
		beego.Info("获取类型错误")
		return
	}
	f,h,err:=c.GetFile("uploadname")
	defer f.Close()


	//1.要限定格式
	fileext := path.Ext(h.Filename)
	if fileext != ".jpg" && fileext != "png"{
		beego.Info("上传文件格式错误")
		return
	}
	//2.限制大小
	if h.Size > 50000000 {
		beego.Info("上传文件过大")
		return
	}

	//3.需要对文件重命名，防止文件名重复
	filename := time.Now().Format("2006-01-02 15:04:05") + fileext  //6-1-2 3:4:5


	if err != nil{
		beego.Info("上传文件失败")
		return
 	}else {
		c.SaveToFile("uploadname","./static/img/"+filename)
	}

	//2.判断数据是否合法
	if artiContent == "" || artiName == ""{
		beego.Info("添加文章数据错误")
		return
	}
	//3.插入数据
	o := orm.NewOrm()
	arti := models.Article{}
	arti.ArtiName = artiName
	arti.Acontent = artiContent
	arti.Aimg = "./static/img/"+filename
	//查找type对象
	artiType := models.ArticleType{Id:id}
	o.Read(&artiType)


	arti.ArticleType = &artiType


	_,err = o.Insert(&arti)
	if err != nil{
		beego.Info("插入数据库错误")
		return
	}

	//4.返回文章界面
	c.Redirect("/index",302)
}

//显示内容详情页面
func (c*MainController)ShowContent(){
	//1.获取文章ID
	id,err := c.GetInt("id")
	//beego.Info("id is ",id)
	if err != nil {
		beego.Info("获取文章ID错误",err)
		return
	}
	//2.查询数据库获取数据
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询错误",err)
		return
	}
	//3.传递数据给试图
	c.Data["article"] = arti

	c.TplName = "content.html"

}

//显示编辑界面
func (c*MainController)ShowUpdate(){
	//1.获取文章ID
	id,err := c.GetInt("id")
	//beego.Info("id is ",id)
	if err != nil {
		beego.Info("获取文章ID错误",err)
		return
	}
	//2.查询数据库获取数据
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询错误",err)
		return
	}
	//3.传递数据给试图
	c.Data["article"] = arti
	c.TplName = "update.html"
}

//处理更新业务数据
func (c*MainController)HandleUpdate(){
	//1.拿到数据
	id,_ := c.GetInt("id")
	artiName := c.GetString("articleName")
	content := c.GetString("content")
	f,h,err:=c.GetFile("uploadname")
	var filename string
	if err != nil{
		beego.Info("上传文件失败")
		return
	}else {
		defer f.Close()


		//1.要限定格式
		fileext := path.Ext(h.Filename)
		if fileext != ".jpg" && fileext != "png"{
			beego.Info("上传文件格式错误")
			return
		}
		//2.限制大小
		if h.Size > 50000000 {
			beego.Info("上传文件过大")
			return
		}

		//3.需要对文件重命名，防止文件名重复
		filename = time.Now().Format("2006-01-02 15:04:05") + fileext  //6-1-2 3:4:5
		c.SaveToFile("uploadname","./static/img/"+filename)
	}

	//2.对数据进行一个处理
	if artiName == "" || content ==""{
		beego.Info("更新数据获取失败")
		return
	}

	//3.更新操作
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询数据错误")
		return
	}
	arti.ArtiName = artiName
	arti.Acontent = content
	arti.Aimg = "./static/img/"+filename


	_,err = o.Update(&arti,"ArtiName","Acontent","Aimg")
	if err != nil{
		beego.Info("更新数据显示错误")
		return
	}
	//4.返回列表页面
	c.Redirect("/index",302)
}

//删除操作
func (c*MainController)HandleDelete(){
	//1.拿到数据
	id,err:=c.GetInt("id")
	if err != nil{
		beego.Info("获取id数据错误")
		return
	}


	//2.执行删除操作
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询错误")
		return
	}
	o.Delete(&arti)

	//3.返回列表页面
	c.Redirect("/index",302)
}

//显示添加类型界面
func (c*MainController)ShowAddType(){
	o := orm.NewOrm()
	var artiTypes []models.ArticleType
	_,err:=o.QueryTable("ArticleType").All(&artiTypes)
	if err!=nil{
		beego.Info("没有获取到类型数据")
	}

	c.Data["articleType"] = artiTypes
	c.TplName = "addType.html"
}

//处理添加类型传输的信息
func (c*MainController)HandleAddType(){
	//1.获取内容
	typeName := c.GetString("typeName")
	//2.判断数据是否合法
	if typeName == ""{
		beego.Info("获取天津爱类型信息错误")
		return
	}
	//3.写入数据
	o := orm.NewOrm()
	artiType := models.ArticleType{}
	artiType.Tname = typeName
	_,err := o.Insert(&artiType)
	if err != nil{
		beego.Info("插入类型错误")
		return
	}
	//4.返回界面
	c.Redirect("/addType",302)
}
