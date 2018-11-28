package main

import (
	"net/http"
	"time"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// handle static assets 新建多路复用器
	mux := http.NewServeMux()
	// 通过配置文件设置静态文件路径
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// 所有的路由匹配都写在这个文件中,路由处理器则在其他文件进行声明
	// all route patterns matched here
	// route handler functions defined in other files

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// starting up the server 开启服务器
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)), //设置单位为秒
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20, // 设置首部的最大值为1m
	}
	server.ListenAndServe()
}
