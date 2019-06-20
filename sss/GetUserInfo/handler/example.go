package handler

import (
	"context"

	example "sss/GetUserInfo/proto/example"
	"fmt"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
	"strconv"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取用户信息 GetUserInfo /api/v1.0/user")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//获取sessionid
	sessionid := req.Sessionid

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("redis链接错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	//拼接key 查询 user_id
	value := bm.Get(sessionid + "user_id")
	value_string, _ := redis.String(value, nil)
	value_int, _ := strconv.Atoi(value_string)
	//value_int, _ := redis.Int(value, nil)
	fmt.Println("用户Id：", value_int)

	//链接数据库
	o := orm.NewOrm()
	user := models.User{Id: value_int}

	//查询数据
	err = o.Read(&user)
	if err != nil {
		fmt.Println("redis链接错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	//返回数据
	rsp.Name = user.Name
	rsp.Mobile = user.Mobile
	rsp.UserId = int64(user.Id)
	rsp.IdCard = user.Id_card

	rsp.RealName = user.Real_name
	rsp.AvatarUrl = user.Avatar_url

	fmt.Println("用户信息：---------------------------")
	fmt.Println("user.Name", user.Name)
	fmt.Println("user.Mobile", user.Mobile)
	fmt.Println("user.Id", int64(user.Id))
	fmt.Println("user.Id_card", user.Id_card)
	fmt.Println("user.Real_name", user.Real_name)
	fmt.Println("user.Avatar_url", user.Avatar_url)
	fmt.Println("-------------------------------------")

	return nil
}
