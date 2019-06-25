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
	GETUSERHOUSES "sss/GetUserHouses/proto/example"
	POSTHOUSES "sss/PostHouses/proto/example"
	POSTHOUSESIMAGE "sss/PostHousesImage/proto/example"
	GETHOUSEINFO "sss/GetHouseInfo/proto/example"
	GETINDEX "sss/GetIndex/proto/example"
	GETHOUSES "sss/GetHouses/proto/example"
	POSTORDERS "sss/PostOrders/proto/example"
	GETUSERORDERS "sss/GetUserOrder/proto/example"
	PUTORDERS "sss/PutOrders/proto/example"
	PUTCOMMENT "sss/PutComment/proto/example"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	"github.com/astaxie/beego"
	"sss/IhomeWeb/models"
	"fmt"
	"image"
	"github.com/afocus/captcha"
	"image/png"
	"sss/IhomeWeb/utils"
	"io/ioutil"
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
	//判断错误
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

//获取首页轮播  首页登录
func GetIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("获取首页轮播 GetIndex /api/v1.0/house/index")

	//创建grpc
	cil := grpc.NewService()
	//初始化
	cil.Init()

	//创建句柄，调用函数
	exampleClient := GETINDEX.NewExampleService("go.micro.srv.GetIndex", cil.Client())
	rsp, err := exampleClient.GetIndex(context.TODO(), &GETINDEX.Request{})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	//创建接受数据容器
	index := []interface{}{}

	//将数据解码到data容器中
	if err := json.Unmarshal(rsp.Data, &index); err != nil {
		fmt.Println("解码数据错误")
		// we want to augment the response
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

	//创建返回前端的数据容器
	//data := make(map[string]interface{})
	//data["houses"] = index

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   index,
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	return
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

//获取用户已发布房源信息服务
func GetUserHouses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("获取用户已发布房源信息服务 api/v1.0/user/houses  GetUserHouses")

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	//获取cookie从中取出sessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		fmt.Println("Cookie获取错误")
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

	// 通过句柄调用函数，得到返回值
	exampleClient := GETUSERHOUSES.NewExampleService("go.micro.srv.GetUserHouses", cli.Client())
	rsp, err := exampleClient.GetUserHouses(context.TODO(), &GETUSERHOUSES.Request{
		Sessionid: cookie.Value,
	})
	//判断错误信息
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	/*定义容器返回返房屋信息*/
	//房屋切片
	house_list := []models.House{}

	//json解码，将返回的数据加载进 房屋切片
	if err := json.Unmarshal(rsp.Data, &house_list); err != nil {
		fmt.Println("返回数据解码错误, err:", err)
		// we want to augment the response
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

	//json.Unmarshal(rsp.Data, &house_list)

	//定义返回给前段的数据容器 ===>切片
	var houses []interface{}

	//遍历house_list，取出数据，将数据装进houses切片中，切片中是一个个map[string]interface{}
	for _, value := range house_list {
		fmt.Printf("house.user = %+v\n", value.Id)
		fmt.Printf("house.area = %+v\n", value.Area)
		houses = append(houses, value.To_house_info())
	}

	fmt.Println("houses:", houses)

	//返回的数据Data也是个map，key=“houses”  value=houses（切片[]interface{}）
	data_map := make(map[string]interface{})
	data_map["houses"] = houses

	// 准备返回结构体
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data_map,
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

//发送（发布）房源信息服务
func PostHouses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("发送（发布）房源信息服务 PostHouses   /api/v1.0/houses")

	//不接受数据，直接将数据存放在二进制数据容器中,获取post请求包体中的数据
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("获取数据失败")
		// we want to augment the response
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

	//获取sessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		fmt.Println("session获取失败")
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
	exampleClient := POSTHOUSES.NewExampleService("go.micro.srv.PostHouses", cli.Client())
	rsp, err := exampleClient.PostHouses(context.TODO(), &POSTHOUSES.Request{
		Sessionid: cookie.Value,
		Data:      body,
	})
	//判断错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//定义容器，接受返回的数据
	data := make(map[string]interface{})
	data["house_id"] = int(rsp.HouseId)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}

	//设置返回数据的格式
	w.Header().Set("Content-type", "application/json")

	// 将返回数据转化为json格式，传输给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

//发送（上传）房屋图片服务
func PostHousesImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("上传房屋图片服务  PostHousesImage   /api/v1.0/houses/:id/images")

	//获取前端发送的参数
	houseid := ps.ByName("id")

	//获取sessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		fmt.Println("获取cookie错误")
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

	//获取前端发送的图片数据
	file, head, err := r.FormFile("house_image")
	if err != nil {
		fmt.Println("获取图片数据错误")
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

	fmt.Println("===========================================")
	fmt.Println("文件名：", head.Filename)
	fmt.Println("文件大小：", head.Size)
	fmt.Println("===========================================")

	//定义容器储存图片数据
	filebuffer := make([]byte, head.Size)
	//将文件读取进入容器中
	_, err = file.Read(filebuffer)
	if err != nil {
		fmt.Println("图片数据读取失败：", err)
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

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := POSTHOUSESIMAGE.NewExampleService("go.micro.srv.PostHousesImage", cli.Client())
	rsp, err := exampleClient.PostHousesImage(context.TODO(), &POSTHOUSESIMAGE.Request{
		Sessionid: cookie.Value,
		Image:     filebuffer,
		HouseId:   houseid,
		Filesize:  head.Size,
		Filename:  head.Filename,
	})
	//判断错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//定义返回数据容器
	data := make(map[string]interface{})
	data["url"] = utils.AddDomain2Url(rsp.Url)

	//返回数据map
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

//获取房屋详细信息服务
func GetHouseInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("获取房源详细信息 GetHouseInfo  api/v1.0/houses/:id ")

	//获取房屋id
	houseid := ps.ByName("id")

	//获取sessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		fmt.Println("获取sessionid错误")
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
	exampleClient := GETHOUSEINFO.NewExampleService("go.micro.srv.GetHouseInfo", cli.Client())
	rsp, err := exampleClient.GetHouseInfo(context.TODO(), &GETHOUSEINFO.Request{
		Sessionid: cookie.Value,
		HouseId:   houseid,
	})
	//判断错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//定义承载房屋信息的数据容器
	house_data := make(map[string]interface{})
	//将srv返回的数据解密装在容器中
	if err := json.Unmarshal(rsp.Housedata, &house_data); err != nil {
		fmt.Println("解密房屋数据失败")
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

	//定义返回前端的数据容器
	data_map := make(map[string]interface{})
	//房屋详细信息
	data_map["house"] = house_data
	//用户id
	data_map["user_id"] = int(rsp.UserId)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data_map,
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

//获取（搜索）房源服务
func GetHouses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("获取（搜索）房源服务  /api/v1.0/houses  GetHouses")

	//获取url中的参数
	aid := r.URL.Query()["aid"][0] //aid=5   地区编号
	sd := r.URL.Query()["sd"][0]   //sd=2017-11-1   开始世界
	ed := r.URL.Query()["ed"][0]   //ed=2017-11-3   结束世界
	sk := r.URL.Query()["sk"][0]   //sk=new    第三栏条件
	p := r.URL.Query()["p"][0]     //tp=1   页数

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := GETHOUSES.NewExampleService("go.micro.srv.GetHouses", cli.Client())
	rsp, err := exampleClient.GetHouses(context.TODO(), &GETHOUSES.Request{
		Aid: aid,
		Sd:  sd,
		Ed:  ed,
		Sk:  sk,
		P:   p,
	})
	//判断错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//定义接受houses数据容器
	houses := []interface{}{}
	if err := json.Unmarshal(rsp.Houses, houses); err != nil {
		fmt.Println("解码houses数据错误")
		// we want to augment the response
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

	//定义返回数据容器
	data := make(map[string]interface{})
	data["current_page"] = rsp.CurrentPage
	data["houses"] = houses
	data["total_page"] = rsp.TotalPage

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

//-------------------------------未校验------------------------------------------------------------

//发送（发布）订单服务
func PostOrders(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("发送（发布）订单服务	PostOrders	api/v1.0/orders	PostOrders")

	//接收前段发送过来数据，转化为二进制
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("接受前端数据错误，err：", err)
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

	//获取sessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		fmt.Println("获取sessionid错误，err：", err)
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
	exampleClient := POSTORDERS.NewExampleService("go.micro.srv.PostOrders", cli.Client())
	rsp, err := exampleClient.PostOrders(context.TODO(), &POSTORDERS.Request{
		Sessionid: cookie.Value,
		Body:      body,
	})
	//判断错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//定义返回数据容器
	houseid_map := make(map[string]interface{})
	houseid_map["order_id"] = int(rsp.OrderId)

	// 发送给前段的数据
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   houseid_map,
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

//获取房东/租户订单信息服务
func GetUserOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("获取房东/租户订单信息服务	GET	api/v1.0/user/orders	GetUserOrder")

	//获取前端发送过来的数据 role{custom / landlord}
	role := r.URL.Query()["role"][0]

	//获取sessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		fmt.Println("获取sessionid错误，err ：", err)
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
	exampleClient := GETUSERORDERS.NewExampleService("go.micro.srv.GetUserOrder", cli.Client())
	rsp, err := exampleClient.GetUserOrder(context.TODO(), &GETUSERORDERS.Request{
		Sessionid: cookie.Value,
		Role:      role,
	})
	//判断错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//定义返回前端数据容器
	order_map := map[string]interface{}{}
	orders := []interface{}{}

	//将数据放进容器中
	if err := json.Unmarshal(rsp.Data, &orders); err != nil {
		fmt.Println("解码json数据错误,err ：", err)
		// we want to augment the response
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

	//将数据装进map中
	order_map["orders"] = orders

	// 返回前端数据
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   order_map,
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

//更新房东同意/拒绝订单
func PutOrders(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("更新房东同意/拒绝订单	PUT	api/v1.0/orders/:id/status	PutOrders")

	//接收请求携带的数据    将数据用json处理函数，转换为map格式数据
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	//是否接受
	action := request["action"].(string)

	//获取前端数据（订单id）
	orderid := ps.ByName("id")

	//获取sessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		fmt.Println("获取sessionid错误,err :", err)
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
	exampleClient := PUTORDERS.NewExampleService("go.micro.srv.PutOrders", cli.Client())
	rsp, err := exampleClient.PutOrders(context.TODO(), &PUTORDERS.Request{
		Sessionid: cookie.Value,
		Action:    action,
		Orderid:   orderid,
	})
	//判断错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
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

//更新用户评价订单信息
func PutComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("更新用户评价订单信息	PUT	api/v1.0/orders/:id/comment   	PutComment")

	// 获取前端传来的参数，进行数据转换
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//获取订单编号
	orderid := ps.ByName("id")

	//获取sessionid
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		fmt.Println("获取sessionid错误,err :", err)
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

	//获取评论内容
	comment := request["comment"].(string)

	//创建grpc链接 客户端
	cli := grpc.NewService()
	//初始化
	cli.Init()

	// call the backend service
	exampleClient := PUTCOMMENT.NewExampleService("go.micro.srv.PutComment", cli.Client())
	rsp, err := exampleClient.PutComment(context.TODO(), &PUTCOMMENT.Request{
		Sessionid: cookie.Value,
		OrderId:   orderid,
		Comment:   comment,
	})
	//判断错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 返回的数据map
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
