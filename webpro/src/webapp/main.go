package main

import (
	"fmt"
	"net/http"
)

//创建处理器
func handler(w http.ResponseWriter,r *http.Request){
	_, _ = fmt.Fprintln(w, "Hello World!", r.URL.Path)
}

func main(){
	//这一步就是用来映射多路复用器该调哪个服务器的，并且在底层转化为处理器
	http.HandleFunc("/",handler)
	//创建路由,如果底下端口不传，那么他默认就是80
	http.ListenAndServe(":8080",nil)

}