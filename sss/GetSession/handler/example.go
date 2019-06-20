package handler

import (
	"context"

	"fmt"

	example "sss/GetSession/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetSession(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println(" 登陆检查  GetSession  /api/v1.0/session")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*2 连接redis*/
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("连接redis失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//获取sessionid
	sessionid := req.Sessionid

	//拼接key
	key := sessionid + "name"

	//查询name
	value := bm.Get(key)
	value_string, _ := redis.String(value, nil)

	//打印从redis中提取出来的name
	fmt.Println("用户名：", value_string)

	//将name返回
	rsp.Name = value_string

	return nil
}
