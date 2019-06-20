package handler

import (
	"context"

	example "sss/PostUserAuth/proto/example"
	"fmt"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostUserAuth(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println(" 实名认证服务  PostUserAuth   /api/v1.0/user/auth  ")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//TODO  关联到有关部门平台--》身份证号，姓名的验证

	//来接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//拼接key
	key := req.Sessionid + "user_id"

	//查询user——id
	value := bm.Get(key)
	value_int, _ := redis.Int(value, nil)
	fmt.Println("user_id:", value_int)

	//链接mysql
	o := orm.NewOrm()
	user := models.User{Id: value_int, Real_name: req.RealName, Id_card: req.IdCard}

	//更新数据
	_, err = o.Update(&user, "Real_name", "Id_card")
	if err != nil {
		fmt.Println("实名认证数据 更新失败 ", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*更新session信息时间*/
	bm.Put(req.Sessionid+"user_id", user.Id, time.Second*600)
	bm.Put(req.Sessionid+"mobile", user.Mobile, time.Second*600)
	bm.Put(req.Sessionid+"name", user.Name, time.Second*600)

	return nil
}
