package main

import(
	json2 "encoding/json"
	"fmt"
	"net/http"
	"web_01db/model"
)

//创建处理器函数
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"你发送的请求地址是：",r.URL.Path)
	fmt.Fprintln(w,"你发送的请求的请求地址后的查询字符串时：",r.URL.RawQuery)
	fmt.Fprintln(w,"请求头中的信息有：",r.Header)
	fmt.Fprintln(w,"请求头中Accept_Encoding的的信息是：",r.Header["Accept-Encoding"])
	fmt.Fprintln(w,"请求头中Accept_Encoding的的信息是：",r.Header.Get("Accept-Encoding"))
	//获取请求体内容的长度
	//len := r.ContentLength
	//创建bytes切片
	//body := make([]byte,len)
	//将请求体中的内容读到body中
	//r.Body.Read(body)
	//在浏览器中显示请求体
	//mt.Fprintln(w,"请求体中的内容有：",string(body))

	//解析表单，在调用r.Form之前必须执行此操作
	err := r.ParseForm()
	//if err != nil{
	//	fmt.Println(err)
	//}
	fmt.Println("cuowuxinxi",err)
	//获取请求参数
	//如果form表单的action属性的URL地址也有与form表单参数名相同的请求参数，
	//那么参数值都可以得到，并且form表单中的参数值在URL的参数值前面
	fmt.Fprintln(w,"返回的参数是：",r.Form)
	fmt.Fprintln(w,"POST请求的form表单中的请求参数是：",r.PostForm)
}


//测试发送Json格式
func testJsonRes(w http.ResponseWriter,r *http.Request){
	//设置响应头的类型
	w.Header().Set("Content_Type","application/json")
	user :=model.User{
		ID:1,
		Username:"admin",
		Password: "123456",
		Email:"admin@atguigui.com",
	}
	//将user转化为Json格式
	json,_ :=json2.Marshal(user)
	//将json格式数据享用到客户端
	w.Write(json)
}
//测试重定向客户端
func testRedire(w http.ResponseWriter,r *http.Request){
	//设置响应头中的Location属性
	w.Header().Set("Location","https://www.baidu.com")
	//设置响应的状态码
	w.WriteHeader(302)

}
func main(){
	http.HandleFunc("/hello",handler)
	http.HandleFunc("/testJson",testJsonRes)
	http.HandleFunc("/testRedirect",testRedire)
	http.ListenAndServe(":8082",nil)
}