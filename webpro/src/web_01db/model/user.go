package model

import (
	"fmt"
	"math/big"
	"web_01db/utils"
)
//结构体
type User struct {
	ID int
	Username string
	Password string
	Email string
}
//AddUser添加方法
func (user *User)AddUser()error {
	//1.写sql语句
	sqlStr := "insert into users(user_name,password,email) values(?,?,?)"
	//2.预编译
	inStmt , err := utils.Db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("预编译出现异常！")
		return err
	}
	//3.执行
	_ , err2 := inStmt.Exec("admin","123456","admin@atguigu.com")
	if err2 != nil{
		fmt.Println("执行出现异常！")
		return err
	}
	return nil
}


func (user *User)AddUser2()error {
	//1.写sql语句
	sqlStr := "insert into users(user_name,password,email) values(?,?,?)"
	//2.执行
	_ , err := utils.Db.Exec(sqlStr,"admin2","666666","admin2@sina.com")
	if err != nil{
		fmt.Println("执行出现异常！",err)
		return err
	}
	return nil
}

//查询
func (user *User)GetUserById()(*User,error){
	sqlStr := "select id,user_name,password,email from users where id = ?"
	row := utils.Db.QueryRow(sqlStr,user.ID)
	var id int
	var username string
	var password string
	var email string
	//Scan将该行查询结果各列分别保存进dest参数指定的值中。
	//如果该查询匹配多行，Scan会使用第一行结果并丢弃其余各行。
	//如果没有匹配查询的行，Scan会返回ErrNoRows。
	err := row.Scan(&id,&username,&password,&email)
	if err != nil{
		return nil,err
	}
	u := &User{
		ID:id,
		Username:username,
		Password:password,
		Email:email,
	}
	return u,nil
}

//GetUsers 获取数据库中所有的记录
func (user *User)GetUsers()([]*User,error){
	sqlStr := "select id,user_name,password,email from users"
	rows,err := utils.Db.Query(sqlStr)
	if err != nil{
		fmt.Println("未成功查询所有记录！")
		return nil,err
	}
	//创建User切片
	var users []*User
	for rows.Next(){
		var id int
		var username string
		var password string
		var email string
		//Scan将该行查询结果各列分别保存进dest参数指定的值中。
		//如果该查询匹配多行，Scan会使用第一行结果并丢弃其余各行。
		//如果没有匹配查询的行，Scan会返回ErrNoRows。
		err := rows.Scan(&id,&username,&password,&email)
		if err != nil{
			return nil,err
		}
		u := &User{
			ID:id,
			Username:username,
			Password:password,
			Email:email,
		}
		users = append(users,u)
	}
	return users,nil
	big.Int{}
}