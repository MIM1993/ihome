package handler

import (
	"context"

	//"github.com/micro/go-log"
	"fmt"
	example "sss/GetArea/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"encoding/json"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取地域信息服务   GetArea  /api/v1.0/areas")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("链接redis错误,err:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//redis key
	key := "area_info"

	//获取数据
	area_info_value := bm.Get(key)
	/*接受数据*/
	//定义接受数据容器
	var areas []models.Area
	if area_info_value != nil {
		fmt.Println("从redis中获取数据发送到web")
		//解码
		err = json.Unmarshal(area_info_value.([]byte), &areas)

		//循环将数据发送给前端
		for _, value := range areas {
			//定义proto结构体容器
			area := example.ResponseAddress{Aid: int32(value.Id), Name: value.Name}
			//将area结构体取地添加到rsp.Data中
			rsp.Data = append(rsp.Data, &area)
		}

		return nil
	}

	/*查询数据库*/
	//创建映射条件
	o := orm.NewOrm()

	//查询条件
	qs := o.QueryTable("Area")
	//查询结果
	num, err := qs.All(&areas)
	if err != nil {
		fmt.Println("传数据库错误", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if num == 0 {
		fmt.Println("无数据", err)
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//存入数据
	//json编码
	area_info_json, _ := json.Marshal(areas)
	//存入redis
	err = bm.Put(key, area_info_json, time.Second*7200)
	if err != nil {
		fmt.Println("redis存入数据错误,err:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//将查询结果赋值到返回数据容器中
	for _, value := range areas {
		//定义proto结构体容器
		area := example.ResponseAddress{Aid: int32(value.Id), Name: value.Name}
		//将area结构体取地添加到rsp.Data中
		rsp.Data = append(rsp.Data, &area)
	}

	return nil
}
