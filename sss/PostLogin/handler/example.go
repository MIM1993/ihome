package handler

import (
	"context"

	example "sss/PostLogin/proto/example"
	"sss/IhomeWeb/utils"
	"sss/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostLogin(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println(" 登陆服务 PostLogin  /api/v1.0/sessions")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//打印接收到的数据
	fmt.Println("mobile:", req.Mobile)
	fmt.Println("password:", req.Password)

	//获取数据
	var user models.User
	//user.Name = req.Mobile

	//查询数据
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("mobile", req.Mobile).One(&user)
	if err != nil {
		fmt.Println("用户名查询失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//判断密码
	if utils.Getmd5string(req.Password) != user.Password_hash {
		fmt.Println("密码错误", utils.Getmd5string(req.Password), user.Password_hash)
		rsp.Errno = utils.RECODE_PWDERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//生成sessionid
	sessionid := utils.Getmd5string(req.Mobile + req.Password + fmt.Sprintln(time.Now()))

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("链接redis失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//将用户数据存入redis
	bm.Put(sessionid+"name", user.Name, time.Second*600)
	bm.Put(sessionid+"mobile", user.Mobile, time.Second*600)
	bm.Put(sessionid+"user_id", user.Id, time.Second*600)

	//返回sessionid
	rsp.Sessionid = sessionid

	return nil
}
