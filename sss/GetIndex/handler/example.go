package handler

import (
	"context"

	"fmt"

	example "sss/GetIndex/proto/example"
	"sss/IhomeWeb/utils"
	"sss/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetIndex(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取首页轮播 GetIndex /api/v1.0/house/index")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//链接redis，获取缓存
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("链接redis错误，err：", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//拼接key
	key := "home_page_data"

	//从redis中获取数据
	index_info := bm.Get(key)
	if index_info != nil {
		fmt.Println("======= get house page info  from CACHE!!! ========")
		rsp.Data = index_info.([]byte)
		fmt.Println("rsp.Data:", rsp.Data)
		return nil
	}

	//实例化容器 house
	houses := []models.House{}

	//从mysql数据中获取固定数据
	o := orm.NewOrm()
	_, err = o.QueryTable("House").Limit(models.HOME_PAGE_MAX_HOUSES).All(&houses)
	if err != nil {
		fmt.Println("mysql数据库查询错误，err：", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//创建返回数据容器data
	data := []interface{}{}

	//遍历houses，关连查询
	for _, house := range houses {
		o.LoadRelated(&house, "Area")
		o.LoadRelated(&house, "User")
		o.LoadRelated(&house, "Facilities")
		o.LoadRelated(&house, "Images")
		data = append(data, house.To_house_info())
	}

	//先转换数据为json
	house_page_value, err := json.Marshal(data)
	if err != nil {
		fmt.Println("转换数据为json错误，err：", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//然后将数据存入缓存
	err = bm.Put(key, house_page_value, 3600*time.Second)
	if err != nil {
		fmt.Println("redis存储错误，err：", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//返回数据
	rsp.Data = house_page_value

	//打印数据
	//fmt.Println("发送数据：", house_page_value)

	return nil
}
