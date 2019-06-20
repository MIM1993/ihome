package handler

import (
	"context"
	"encoding/json"
	"net/http"
	GETAREA "sss/GetArea/proto/example"
	GETIMAGECD "sss/GetImageCd/proto/example"
	GETSMSCD "sss/GetSmscd/proto/example"
	POSTRET "sss/PostRet/proto/example"
	GETSESSION "sss/GetSession/proto/example"
	POSTLOGIN "sss/PostLogin/proto/example"
	DELETESESSION "sss/DeleteSession/proto/example"
	GETUSERINFO "sss/GetUserInfo/proto/example"
	POSTAVATAR "sss/PostAvatar/proto/example"
	PUTUSERINFO "sss/PutUserInfo/proto/example"
	POSTUSERAUTH "sss/PostUserAuth/proto/example"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	"github.com/astaxie/beego"
	"sss/IhomeWeb/models"
	"fmt"
	"image"
	"github.com/afocus/captcha"
	"image/png"
	"sss/IhomeWeb/utils"
)

//damo
func ExampleCall(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// decode the incoming request as json
	//var request map[string]interface{}
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//
	////创建grpc链接 客户端
	//cli := grpc.NewService()
	////初始化
	//cli.Init()
	//
	//// call the backend service
	//exampleClient := example.NewExampleService("go.micro.srv.template", cli.Client())
	//rsp, err := exampleClient.Call(context.TODO(), &example.Request{
	//	Name: request["name"].(string),
	//})
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//
	//// we want to augment the response
	//response := map[string]interface{}{
	//"errno": utils.RECODE_MOBILEERR,
	//	"errmsg":   utils.RecodeText(utils.RECODE_MOBILEERR),
	//}
	//
	//设置返回数据的格式
	//w.Header().Set("Content-type", "application/json")
	//// encode and write the response as json
	//if err := json.NewEncoder(w).Encode(response); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//return
}

//获取地域信息
func GetArea(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取地区请求客户端  url：api/v1.0/areas")

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// 调用f服务端函数病返回数据
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea", cli.Client())
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//整理获取到的数据
	area_list := []models.Area{}
	for _, value := range rsp.Data {
		temp := models.Area{
			Id:   int(value.Aid),
			Name: value.Name,
		}
		area_list = append(area_list, temp)
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   area_list,
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")
	//fmt.Println(w.Header()["Content-type"])

	//  将返回数据map发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

//欺骗浏览器  首页登录
func Getindex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Getindex /api/v1.0/house/index")
	// we want to augment the response
	response := map[string]interface{}{
		"errno":  "0",
		"errmsg": "ok",
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//登录检查session
func GetSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" 登陆检查  GetSession  /api/v1.0/session")

	//获取sessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		response := map[string]interface{}{
			"errno":  "4101",
			"errmsg": "用户未登录",
		}

		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := GETSESSION.NewExampleService("go.micro.srv.GetSession", cli.Client())
	rsp, err := exampleClient.GetSession(context.TODO(), &GETSESSION.Request{
		Sessionid: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//定义map容器承载data
	data := make(map[string]string)
	data["name"] = rsp.Name

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

//验证码
func GetImageCd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("获取验证码 GetImageCd  /api/v1.0/imagecode/:uuid")

	//获取url中的参数
	uuid := ps.ByName("uuid")

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := GETIMAGECD.NewExampleService("go.micro.srv.GetImageCd", cli.Client())
	rsp, err := exampleClient.GetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid: uuid,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//判断返回值  如果不等于 返回错误
	if rsp.Errno != "0" {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  rsp.Errno,
			"errmsg": rsp.Errmsg,
		}

		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//接受图片数据 拼接图片结构体 发送给前段

	//定义容器
	var img image.RGBA
	for _, value := range rsp.Pix {
		img.Pix = append(img.Pix, uint8(value))
	}

	img.Stride = int(rsp.Stride)
	//point
	//min-->x,y
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	//max-->x,y
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)

	//类型转换
	var image captcha.Image
	image.RGBA = &img

	//将图片发送给前端
	png.Encode(w, image)

	return
}

//发送短信
func GetSmscd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" 获取短信验证码 GetSmscd  /api/v1.0/smscode/:mobile")
	//接收电话号码
	mobile := ps.ByName("mobile")

	/*
		//用正则验证电话号码是否正确
		myreg := regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)
		bo := myreg.MatchString(mobile)
		if bo == false {
			response := map[string]interface{}{
				"errno":  utils.RECODE_MOBILEERR,
				"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
			}
			//设置返回时的响应头数据格式
			w.Header().Set("Content-type", "application/json")
			// 将数据反回前段
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			return
		}
	*/

	//获取图片验证码
	text := r.URL.Query()["text"][0]
	uuid := r.URL.Query()["id"][0]

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// 调用服务端的函数
	exampleClient := GETSMSCD.NewExampleService("go.micro.srv.GetSmscd", cli.Client())
	rsp, err := exampleClient.GetSmscd(context.TODO(), &GETSMSCD.Request{
		Mobile: mobile,
		Uuid:   uuid,
		Text:   text,
	})
	//判断返回值
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 返回给前端的数据
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	//设置返回时的响应头数据格式
	w.Header().Set("Content-type", "application/json")

	// 将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

//注册服务
func PostRet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" 注册服务  PostRet  /api/v1.0/users")

	//decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//判断数据
	/*    mobile: "123",
    password: "123",
    sms_code: "123"
   */
	if request["mobile"].(string) == "" || request["password"].(string) == "" || request["sms_code"].(string) == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := POSTRET.NewExampleService("go.micro.srv.PostRet", cli.Client())
	rsp, err := exampleClient.PostRet(context.TODO(), &POSTRET.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
		SmsCode:  request["sms_code"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	//设置cookie
	cookie, err := r.Cookie("ihomelogin")
	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{Name: "ihomelogin", Value: rsp.Sessionid, MaxAge: 600, Path: "/"}
		http.SetCookie(w, &cookie)
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

//登录
func PostLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("登录业务 PostLogin  api/v1.0/sessions")
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//打印手机号号密码
	fmt.Println("账号：", request["mobile"].(string))
	fmt.Println("密码：", request["password"].(string))

	//校验数据
	if request["mobile"].(string) == "" || request["password"].(string) == "" {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := POSTLOGIN.NewExampleService("go.micro.srv.PostLogin", cli.Client())
	rsp, err := exampleClient.PostLogin(context.TODO(), &POSTLOGIN.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
	})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	/*将获取到的sessionid设置到cookie*/
	//先获取cookie没有的话在设置
	cookie, err := r.Cookie("ihomelogin")
	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{Name: "ihomelogin", Value: rsp.Sessionid, MaxAge: 600, Path: "/"}
		http.SetCookie(w, &cookie)
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

//退出登录
func DeleteSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("退出登录   DeleteSession  /api/v1.0/session")

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	//获取cookie
	cookie, err := r.Cookie("ihomelogin")
	if err != nil || cookie.Value == "" {
		//cookie不存在,直接返回
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 需要先将sessionid传递到sev在设置cookie,不然数据就穿不到后台了
	exampleClient := DELETESESSION.NewExampleService("go.micro.srv.DeleteSession", cli.Client())
	rsp, err := exampleClient.DeleteSession(context.TODO(), &DELETESESSION.Request{
		Sessionid: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//为了防止从srv调函数时有延迟,导致cookie过期,所以要在读取一次
	if rsp.Errno == "0" {
		//再次读取数据
		cookie, err := r.Cookie("ihomelogin")
		//数据不为空则将数据设置副的
		if err != nil || "" == cookie.Value {
			return
		} else {
			//设置sessionid时间为负值
			delcookie := http.Cookie{Name: "ihomelogin", MaxAge: -1, Path: "/"}
			http.SetCookie(w, &delcookie)
		}
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	return
}

//获取用户信息
func GetUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("获取用户信息 GetUserInfo /api/v1.0/user")

	//获取cookie
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := GETUSERINFO.NewExampleService("go.micro.srv.GetUserInfo", cli.Client())
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &GETUSERINFO.Request{
		Sessionid: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//接受数据
	//"user_id": 1,
	//"name": "Panda",
	//"mobile": "110",
	//"real_name": "熊猫",
	//"id_card": "210112244556677",
	//"avatar_url":
	data := make(map[string]interface{})
	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		//差点忘了
		"data": data,
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

//实名认证服务
func GetUserAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("获取用户信息 GetUserInfo /api/v1.0/user")

	//获取cookie
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := GETUSERINFO.NewExampleService("go.micro.srv.GetUserInfo", cli.Client())
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &GETUSERINFO.Request{
		Sessionid: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//接受数据
	//"user_id": 1,
	//"name": "Panda",
	//"mobile": "110",
	//"real_name": "熊猫",
	//"id_card": "210112244556677",
	//"avatar_url":
	data := make(map[string]interface{})
	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		//差点忘了
		"data": data,
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

//上传用户头像服务
func PostAvatar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("上传用户头像 PostAvatar /api/v1.0/user/avatar")

	//获取前端发送过来的消息
	file, head, err := r.FormFile("avatar")
	if err != nil {
		fmt.Println("获取前端发送图片数据失败")
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//TODO  可以进行文件类型判断

	/*获取图片数据基本信息*/
	//文件大小
	filesize := head.Size
	fmt.Println("文件大小:", filesize)
	//文件名
	filename := head.Filename
	fmt.Println("文件名:", filename)

	//定义二进制容器,储存图片数据
	filebuffer := make([]byte, filesize)

	//将文件发读取到filebuffer中
	_, err = file.Read(filebuffer)
	if err != nil {
		fmt.Println("读取filebuffer中的数据失败")
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//获取cookie,从中取出sessionid ,出现错误直接返回
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		fmt.Println("获取sessionid失败")
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := POSTAVATAR.NewExampleService("go.micro.srv.PostAvatar", cli.Client())
	rsp, err := exampleClient.PostAvatar(context.TODO(), &POSTAVATAR.Request{
		Sessionid: cookie.Value,
		Filesize:  filesize,
		Filename:  filename,
		Buffer:    filebuffer,
	})
	//判断是否有错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//定义"data"数据容器
	data := make(map[string]string)
	data["avatar_url"] = utils.AddDomain2Url(rsp.Fileid)

	// 返回给前端的数据
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")

	// 将数据加密后发给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

//更新用户名
func PutUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("更新用户名  PutUserInfo  api/v1.0/user/name")

	//decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//参数校验
	if request["name"] == "" {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}

		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//获取usessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}

		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := PUTUSERINFO.NewExampleService("go.micro.srv.PutUserInfo", cli.Client())
	rsp, err := exampleClient.PutUserInfo(context.TODO(), &PUTUSERINFO.Request{
		Username:  request["name"].(string),
		Sessionid: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//重新设置 cookie 的时间  刷新时间
	newcookie := http.Cookie{Name: "ihomelogin", Value: cookie.Value, Path: "/", MaxAge: 600}
	http.SetCookie(w, &newcookie)

	//接受返回数据
	data := make(map[string]string)
	data["name"] = rsp.Username
	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

//设置实名信息
func PostUserAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" 实名认证服务  PostUserAuth   /api/v1.0/user/auth  ")

	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//数据校验
	if request["real_name"].(string) == "" || request["id_card"].(string) == "" {
		fmt.Println("获取数据不完整")
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}
		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//TODO 需要对身份证进行 正则校验

	//获取sessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil || cookie.Value == "" {
		fmt.Println("cookie 数据错误")
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := POSTUSERAUTH.NewExampleService("go.micro.srv.PostUserAuth", cli.Client())
	rsp, err := exampleClient.PostUserAuth(context.TODO(), &POSTUSERAUTH.Request{
		IdCard:    request["id_card"].(string),
		RealName:  request["real_name"].(string),
		Sessionid: cookie.Value,
	})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return

	}

	//刷新cookie时间
	newcookie := http.Cookie{Name: "ihomelogin", Value: cookie.Value, Path: "/", MaxAge: 600}
	http.SetCookie(w, &newcookie)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return

}
