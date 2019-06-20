package handler

import (
	"context"
	"fmt"
	example "sss/GetSmscd/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
	"reflect"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
	"strconv"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetSmscd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println(" 获取短信验证码 GetSmscd  /api/v1.0/smscode/:mobile")

	// 1 初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	// 2 验证手机号是否在数据库中
	o := orm.NewOrm()
	user := models.User{Mobile: req.Mobile}
	err := o.Read(&user, "Mobile")
	if err == nil {
		fmt.Println("该用户已经注册")
		rsp.Errno = utils.RECODE_USERERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 3 链接redis,获取图片验证码
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("redis 链接错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 4 获取图片验证码
	auth_code := bm.Get(req.Uuid)

	//打印一下从redis中取出的数据的类型
	fmt.Println(reflect.TypeOf(auth_code), auth_code)

	//将从redis中取出来的数据进行转化,转化为基础类型---->助手函数,err不是自定义的,是获取redis数据时得到的
	value_string, _ := redis.String(auth_code, nil)

	// 5 数据对比
	if value_string != req.Text {
		fmt.Println("图片验证码错误")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 6 生成随机数--->四位数
	t := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := t.Intn(8999) + 1000

	//打印生成的随机数
	fmt.Println("验证码：", size)

	/*
		// 7 调用短信接口,发送短信
		//发送短信的配置信息
		messageconfig := make(map[string]string)
		//预先创建好的appid
		messageconfig["appid"] = "29672"
		//预先获得的app的key
		messageconfig["appkey"] = "89d90165cbea8cae80137d7584179bdb"
		//加密方式默认
		messageconfig["signtype"] = "md5"

		//messagexsend
		//创建短信发送的句柄
		messagexsend := submail.CreateMessageXSend()
		//短信发送的手机号
		submail.MessageXSendAddTo(messagexsend, req.Mobile)
		//短信发送的模板
		submail.MessageXSendSetProject(messagexsend, "NQ1J94")
		//验证码
		submail.MessageXSendAddVar(messagexsend, "code", strconv.Itoa(size))
		//发送短信请求
		send := submail.MessageXSendRun(submail.MessageXSendBuildRequest(messagexsend), messageconfig)
		//发送短信的请求
		fmt.Println("MessageXSend ", send)

		//验证短信信息是否发送成功
		bo := strings.Contains(send, "success")
		if bo != true {
			fmt.Println("图片验证码错误")
			rsp.Errno = utils.RECODE_DATAERR
			rsp.Errmsg = utils.RecodeText(rsp.Errno)
			return nil
		}
	*/

	//将随机数与手机号存入redis
	err = bm.Put(req.Mobile, strconv.Itoa(size), time.Second*300)
	if err != nil {
		fmt.Println("随机数存储失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//返回错误
	return nil
}
