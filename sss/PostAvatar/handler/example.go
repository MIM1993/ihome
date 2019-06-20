package handler

import (
	"context"

	example "sss/PostAvatar/proto/example"
	"fmt"
	"sss/IhomeWeb/utils"
	"path"
	"sss/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostAvatar(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("上传用户头像 PostAvatar /api/v1.0/user/avatar")

	//初始化按返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//数据对比
	fileSize := len(req.Buffer)
	if int64(fileSize) != req.Filesize {
		fmt.Println("接收到的数据不完整")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//获取后缀名  带点 ==> .jpg
	fileExt := path.Ext(req.Filename)

	//上传得到文件fileid
	fileId, err := utils.Uploadbybuf(req.Buffer, fileExt[1:])
	if err != nil {
		fmt.Println("utils.Uploadbybuf 上传图片数据错误,err:", err)
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("链接redis错误,err:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//拼接key
	key := req.Sessionid + "user_id"

	//通过sessionid获取userid
	value_id := bm.Get(key)
	value_id_string, _ := redis.Int(value_id, nil)
	//value_id_string, _ := redis.String(value_id, nil)
	//value_id_string_int, _ := strconv.Atoi(value_id_string)

	//链接mysql数据库
	//创建表对象
	user := models.User{Id: value_id_string, Avatar_url: fileId}
	o := orm.NewOrm()

	//更新数据 == 上传数据
	_, err = o.Update(&user, "Avatar_url")
	if err != nil {
		fmt.Println("将数据存入数据库失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//返回给前端
	rsp.Fileid = fileId

	return nil
}
