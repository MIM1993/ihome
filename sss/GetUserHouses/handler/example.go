package handler

import (
	"context"

	"fmt"

	example "sss/GetUserHouses/proto/example"
	"sss/IhomeWeb/utils"
	"reflect"
	"github.com/gomodule/redigo/redis"
	"sss/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetUserHouses(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取用户已发布房源信息服务 api/v1.0/user/houses  GetUserHouses")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("链接redis数据库错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//获取sessionid
	sessionid := req.Sessionid

	//拼接key
	key := sessionid + "user_id"

	//查询userid
	value_id := bm.Get(key)
	fmt.Println("从redis中获取信息：", value_id, reflect.TypeOf(value_id))

	//助手函数转化类型
	value_id_int, _ := redis.Int(value_id, nil)
	fmt.Println("从redis中获取信息，转化为int类型：", value_id_int, reflect.TypeOf(value_id_int))

	/*查询mysql数据库中的返房源信息*/
	house_list := []models.House{}

	//创建数据库句柄
	o := orm.NewOrm()
	qs := o.QueryTable("House")

	//查询数据
	num, err := qs.Filter("User__Id", value_id_int).All(&house_list)
	fmt.Println("从mysql数据库中获取的数据条数：", num)
	if err != nil {
		fmt.Println("mysql数据库查询失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		fmt.Println("mysql err : ", err)
		return nil
	}
	if num == 0 {
		fmt.Println("mysql数据库中无数据")
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//把是所有的房源信息打包成json二进制
	data, err := json.Marshal(house_list)
	if err != nil {
		fmt.Println("打包json数据错误")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//返回数据
	rsp.Data = data

	fmt.Println("========================操作完成===========================")

	return nil
}
