package handler

import (
	"context"
	"fmt"
	example "sss/GetHouseInfo/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"sss/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetHouseInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取房源详细信息 GetHouseInfo  api/v1.0/houses/:id ")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//获取sessionid
	sessionid := req.Sessionid

	//拼接key
	key := sessionid + "user_id"

	//链接redis，获取userid
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("链接redis失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//在redis中获取userid
	value_id := bm.Get(key)
	value_id_string, _ := redis.String(value_id, nil)
	user_id, _ := strconv.Atoi(value_id_string)
	fmt.Println("user_id:", user_id)

	//获取房源id
	houseid, _ := strconv.Atoi(req.HouseId)

	/*先从缓存中获取信息*/
	//拼接houseid ——> key
	house_info_key := fmt.Sprintf("house_info_%s", houseid)
	//从redis中获取信息
	house_info_value := bm.Get(house_info_key)
	if house_info_value != nil {
		rsp.UserId = int64(user_id)
		rsp.Housedata = house_info_value.([]byte)
		return nil
	}

	//查询数据库当前房屋信息
	//创建容器 house
	house := models.House{Id: houseid}

	//创建mysql数据库句柄
	o := orm.NewOrm()
	//查询数据==>惰性查询
	if err = o.Read(&house); err != nil {
		fmt.Println("数据库查询失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//关联查询==>将其它表的数据关联到house表
	o.LoadRelated(&house, "Area")
	o.LoadRelated(&house, "User")
	o.LoadRelated(&house, "Images")
	o.LoadRelated(&house, "Facilities")
	o.LoadRelated(&house, "Orders")

	//调用函数，将房子的信息整理到map中
	house_one := house.To_one_house_desc()
	house_map := house_one.(map[string]interface{})
	fmt.Println("房屋数据：", house_map)

	//json打包
	house_data, err := json.Marshal(house_map)
	if err != nil {
		fmt.Println("打包数据错误 ,err :", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//将查询到的房屋信息储存到缓存中
	err = bm.Put(house_info_key, house_map, time.Second*3600)
	if err != nil {
		fmt.Println("房屋信息存入缓存失败,err:", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//返还数据给web==> 房屋信息，用户id
	rsp.Housedata = house_data
	rsp.UserId = int64(user_id)

	fmt.Println("============操作完成====================")

	return nil
}
