package handler

import (
	"context"

	example "sss/PutUserInfo/proto/example"
	"fmt"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PutUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("更新用户名  PutUserInfo")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//拼接key
	key := req.Sessionid + "user_id"

	//查询user_id
	value := bm.Get(key)
	value_int, _ := redis.Int(value, nil)

	//链接mysql
	o := orm.NewOrm()
	user := models.User{Id: value_int, Name: req.Username}

	//跟新用户名
	_, err = o.Update(&user, "Name")
	if err != nil {
		fmt.Println("用户名更新失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//更新session信息
	bm.Put(req.Sessionid+"name", user.Name, time.Second*600)
	bm.Put(req.Sessionid+"user_id", user.Id, time.Second*600)
	bm.Put(req.Sessionid+"mobile", user.Mobile, time.Second*600)

	//返回数据
	rsp.Username = req.Username

	return nil
}
