package model

import (
	"fmt"
	"testing"
)

//TestMain函数可以再测试函数执行之前做一些其他操作
//如果不屑m.Run()就只会运行TestMain（）函数
func TestMain(m *testing.M){
	fmt.Println("测试开始")
	//通过m.Run()来执行测试函数
	m.Run()
}

func TestUser(t *testing.T){
	fmt.Println("开始测试User里的相关方法")
	//t.Run("测试添加用户",testAddUser)
	//t.Run("测试查找单个用户",testUser_GetUserById)
	t.Run("测试查找所有用户",testUser_GetUsers)
}
//如果函数名不是以Test开头，那么该函数默认不执行，可以设置为一个子测试函数
func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户")
	user := &User{}
	//fmt.Printf("%v",user)
	//添加调用用户的方法
	user.AddUser()
	user.AddUser2()
}
func testUser_GetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录")
	user := User{
		ID : 1,
	}
	u ,_ := user.GetUserById()
	fmt.Println("得到的User信息是",u)
 }

 func testUser_GetUsers(t *testing.T) {
	 fmt.Println("测试查找所有用户")
	 user := &User{}
	 //调用获取所有用户信息的方法
	 us ,_ := user.GetUsers()
	  for k,v := range(us){
	  	fmt.Printf("第%d个用户的信息是:%v \n",k+1,v)
	  }
 }