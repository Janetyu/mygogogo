package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id        int
	Name      string
	Birthdate string
	Gender    bool
	Score     int
}

func GetAllStudents() []*Student {
	o := orm.NewOrm()
	o.Using("default")
	var students []*Student
	q := o.QueryTable("student")
	q.All(&students)
	return students
}

func GetStudentById(id int) Student {
	u := Student{Id: id}
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

func AddStudent(student *Student) int {
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

func DeleteStudent(id int) {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&Student{Id: id})
}

func init() {
	// register model
	orm.RegisterModel(new(Student))
}
