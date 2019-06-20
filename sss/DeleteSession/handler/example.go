package handler

import (
	"context"
	"fmt"
	example "sss/DeleteSession/proto/example"
	"sss/IhomeWeb/utils"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) DeleteSession(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("退出登录   DeleteSession  /api/v1.0/session")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//获取sessionid
	sessionid := req.Sessionid

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("redis链接错误", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	//拼接key
	sessionid_name := sessionid + "name"
	sessionid_mobile := sessionid + "mobile"
	sessionid_user_id := sessionid + "user_id"

	//将redis中的相关用户数据删除
	bm.Delete(sessionid_name)
	bm.Delete(sessionid_mobile)
	bm.Delete(sessionid_user_id)

	return nil
}
