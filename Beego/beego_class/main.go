package main

import (
	_ "class/routers"
	_ "class/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("showprepage",prepage)
	beego.AddFuncMap("shownextpage",shownextpage)
	beego.Run()
}

//试图函数，获取上一页页码

/*
1.在试图中定义视图函数函数名    | funcName

2.一般在main.go里面实现试图函数

3.在main函数里面把实现的函数和试图函关联起来   beego.AddFuncMap()
 */
func prepage(pageindex int)(preIndex int){
	preIndex = pageindex - 1
	return
}

func shownextpage(pageindex int)(nextIndex int){
	nextIndex = pageindex + 1
	return
}