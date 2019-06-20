package handler

import (
	"context"
	"fmt"
	example "sss/PostRet/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
	"sss/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostRet(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println(" 注册服务  PostRet  /api/v1.0/users")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("redis链接错误", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//从redis中获取当引验证码
	value := bm.Get(req.Mobile)
	value_string, _ := redis.String(value, nil)

	//验证短信验证码是否正确
	if value_string != req.SmsCode {
		fmt.Println("短信验证码错误", value_string, req.SmsCode)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*将手机号,密码,用户名储存到数据库中*/
	//定义容器
	user := models.User{}
	//加密密码
	user.Password_hash = utils.Getmd5string(req.Password)
	user.Mobile = req.Mobile
	user.Name = req.Mobile

	//插入数据库
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		fmt.Println("用户信息插入数据库错误", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//生成sessionid
	sessionid := utils.Getmd5string(req.Mobile + req.Password + fmt.Sprintln(time.Now()))
	rsp.Sessionid = sessionid

	//通过sessionid将数据存入数据库
	bm.Put(sessionid+"user_id", id, time.Second*600)
	bm.Put(sessionid+"mobile", user.Mobile, time.Second*600)
	bm.Put(sessionid+"name", user.Name, time.Second*600)

	return nil
}
