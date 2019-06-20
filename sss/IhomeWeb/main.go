package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"sss/IhomeWeb/handler"
	"github.com/julienschmidt/httprouter"
	_ "sss/IhomeWeb/models"
	"net/http"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.IhomeWeb"),
		web.Version("latest"),
		web.Address(":23333"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	//创建路由 REST设计模式 不是micro框架中的元素,只是一个插件
	rou := httprouter.New()
	//文件服务器 映射  静态页面
	rou.NotFound = http.FileServer(http.Dir("html"))
	//将路由与对应的业务绑定(模板)
	//rou.GET("/example/call", handler.ExampleCall)

	//获取地区信息
	rou.GET("/api/v1.0/areas", handler.GetArea)


	//获取图片验证吗服务
	rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCd)

	//短信验证码
	rou.GET("/api/v1.0/smscode/:mobile", handler.GetSmscd)

	//注册业务
	rou.POST("/api/v1.0/users", handler.PostRet)

	//session发先业务
	rou.GET("/api/v1.0/session", handler.GetSession)

	//登录业务
	rou.POST("/api/v1.0/sessions", handler.PostLogin)

	//退出登录
	rou.DELETE("/api/v1.0/session", handler.DeleteSession)

	//获取用户信息
	rou.GET("/api/v1.0/user", handler.GetUserInfo)

	//获取用户实名信息  实名认证服务
	rou.GET("/api/v1.0/user/auth", handler.GetUserAuth)

	//更新实名认证信息
	rou.POST("/api/v1.0/user/auth", handler.PostUserAuth)

	//上传用户头像
	rou.POST("/api/v1.0/user/avatar",handler.PostAvatar)

	//更新用户名
	rou.PUT("/api/v1.0/user/name",handler.PutUserInfo)

	//首页
	rou.GET("/api/v1.0/house/index", handler.Getindex)


	//将router 注册到访服务
	service.Handle("/", rou)
	//http.Handle("/",rou)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}