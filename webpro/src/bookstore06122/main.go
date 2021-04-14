package main

import (
	"net/http"
	"text/template"
)

//去首页
func IndexHandler(w http.ResponseWriter,r *http.Request){
	//解析模板,Must用于保护里面，会在err非nil时panic,一般用于变量初始化
//D:/gecode/webpro/src/bookstore0612/views/index.html
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行,func (*Template) Execute
	//func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
	//Execute方法将解析好的模板应用到data上，并将输出写入wr。如果执行时出现错误，
	//会停止执行，但有可能已经写入wr部分数据。模板可以安全的并发执行。
	t.Execute(w,"")
}

func main(){
	http.HandleFunc("/testindex",IndexHandler)
	http.ListenAndServe(":8082",nil)
}
