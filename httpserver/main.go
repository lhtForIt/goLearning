package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/golang/glog"
)

func main() {
	glog.V(2).Infoln("服务器开始启动...................")
	http.HandleFunc("/", rootHandler)
	//绑定8012端口并获取返回值 go默认端口好像是80，和java里tomcat默认8080类似。所以你在浏览器访问不需要输端口也能访问
	err := http.ListenAndServe(":8012", nil)
	if err != nil {
		glog.V(2).Infoln(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	glog.V(2).Infoln("进入 root handler")
	// env := os.Environ()
	// for j := range env {
	// 	fmt.Println("env value", j)
	// }
	version := os.Getenv("GOVERSION")
	statusCode := ""
	if r.URL.Path == "/healthz" {
		statusCode = "200"
		w.WriteHeader(200)
	} else {
		statusCode = "500"
		w.WriteHeader(500)
	}

	//header设置version
	w.Header().Add("version", version)
	for k, v := range r.Header {
		fmt.Printf("k %s=>v %s", k, strings.Join(v, ""))
		fmt.Printf("k %s=>v %s", k, toString(v))
		w.Header().Add(k, strings.Join(v, ""))
	}

	glog.V(2).Infoln("客户端访问地址为：" + r.URL.Path)
	glog.V(2).Infoln("客户端IP为：" + r.RemoteAddr)
	glog.V(2).Infoln("客户端返回码为：" + statusCode)

	io.WriteString(w, "你访问的地址是： "+r.URL.Path+"\n")
	io.WriteString(w, "你访问的客户端IP是： "+r.RemoteAddr+"\n")
	io.WriteString(w, "你的返回码是： "+statusCode+"\n")

}

func toString(v []string) string {
	var result string
	for _, i := range v {
		result += i
	}
	return result
}
