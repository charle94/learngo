package design

import (
	"fmt"
)

type User interface {
	Insert(u User)
	Get() int
}
type Department interface {
	Insert(d Department)
	Get() int
}
type SqlDepartment struct {
}

func (s SqlDepartment) Insert(d Department) {
	fmt.Println("sql中插入方法")
}
func (s SqlDepartment) Get() int {
	fmt.Println("sql中get方法")
	return 0
}

type AccessDepartment struct {
}

func (a AccessDepartment) Insert(d Department) {
	fmt.Println("access中插入方法")
}
func (a AccessDepartment) Get() int {
	fmt.Println("access中get方法")
	return 0
}

type SqlsereverUser struct {
}

func (s SqlsereverUser) Insert(u User) {
	//something insert
	fmt.Println("sql中插入方法")
}
func (s SqlsereverUser) Get() int {
	//get id
	fmt.Println("sql中get方法")
	return 0
}

type AccessUser struct {
}

func (s AccessUser) Insert(u User) {
	//something insert
	fmt.Println("access中插入方法")
}
func (s AccessUser) Get() int {
	//get id
	fmt.Println("access中get方法")
	return 0
}

//创建访问user表对象的抽象的工厂方法
type Factory interface {
	CreateUser() User
	CreateDepartment() Department
}

type SqlFactory struct {
}

func (s SqlFactory) CreateUser() User {
	return SqlsereverUser{}
}
func (s SqlFactory) CreateDepartment() Department {
	return SqlDepartment{}
}

type AccessFactory struct {
}

func (a AccessFactory) CreateUser() User {
	return AccessUser{}
}
func (a AccessFactory) CreateDepartment() Department {
	return AccessDepartment{}
}

/*type DataAccess interface {
	CreateUser(which string) User
	CreateDepartment(which string) Department
}
type DataAccessInstace struct {
}

func (d DataAccessInstace) CreateUser(which interface{}) User {

}
func (d DataAccessInstace) CreateDepartment(which string) Department {

}*/
