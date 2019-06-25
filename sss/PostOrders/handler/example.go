package handler

import (
	"context"
	"fmt"
	example "sss/PostOrders/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"time"
	"sss/IhomeWeb/models"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostOrders(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("发送（发布）订单服务	PostOrders	api/v1.0/orders	PostOrders")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//获取sessionid
	sessionid := req.Sessionid
	//拼接key
	key := sessionid + "user_id"

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("链接redis错误，err：", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	//获取userid
	value_id := bm.Get(key)
	userid, err := redis.Int(value_id, nil)
	if err != nil {
		fmt.Println("redis数据错误，err：", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}
	fmt.Println("user_id :", userid)

	/*获取前端发送的数据*/
	//定义容器
	RequestMap := make(map[string]interface{})

	//将数据放进容器中
	err = json.Unmarshal(req.Body, &RequestMap)
	if err != nil {
		fmt.Println("转化json数据错误，err：", err)
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}
	fmt.Println("RequestMap data :", RequestMap)

	//校验数据
	if RequestMap["house_id"] == "" || RequestMap["start_date"] == "" || RequestMap["end_date"] == "" {
		fmt.Println("前端发送数据为空，err：", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	//确定end_data 在 start_date 之后
	//格式化时间
	start_date_time, _ := time.Parse("2006-01-02 15:04:05", RequestMap["start_date"].(string)+" 00:00:00")
	end_date_time, _ := time.Parse("2006-01-02 15:04:05", RequestMap["end_date"].(string)+" 00:00:00")

	//得到一共入住的天数
	fmt.Println("start_date_time :", start_date_time)
	fmt.Println("send_date_time :", end_date_time)
	days := end_date_time.Sub(start_date_time).Hours()/24 + 1
	fmt.Println("入住的天数为：", days)

	//根据house_id 得到房源信息
	house_id, _ := strconv.Atoi(RequestMap["house_id"].(string))
	house := models.House{Id: house_id}

	//创建orm对象，查询房源信息
	o := orm.NewOrm()
	if err := o.Read(&house); err != nil {
		fmt.Println("查询房源信息错误，err", err)
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	//关联查询房屋的用户信息，房屋主人
	o.LoadRelated(&house, "User")

	//将当前下单用户与房屋主人用户id进行比较，判断是否是同一人，以防刷单情况出现
	if userid == house.User.Id {
		fmt.Println("刷单行为不被容许，请从新提交订单")
		rsp.Errno = utils.RECODE_ROLEERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	//确保用户选择的房屋未被预定,日期没有冲突
	if end_date_time.Before(start_date_time) {
		fmt.Println("订单结束时间在订单开始之间之前，信息错误")
		rsp.Errno = utils.RECODE_ROLEERR
		rsp.Errmsg = "订单结束时间在订单开始之间之前，信息错误"
		return err
	}

	//TODO :添加征信查询业务

	//封装order订单
	order := models.OrderHouse{}
	order.House = &house
	user := models.User{Id: userid}
	order.User = &user
	order.Begin_date = start_date_time
	order.End_date = end_date_time
	order.Days = int(days)
	order.House_price = house.Price
	order.Status = models.ORDER_STATUS_WAIT_ACCEPT
	amount := days * float64(house.Price)
	order.Amount = int(amount)
	//征信
	order.Credit = false

	//将订单信息存入数据库mysql中
	if _, err := o.Insert(&order); err != nil {
		fmt.Println("订单信息插入数据库错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	//将订单信息存入缓存
	bm.Put(key, string(userid), time.Second*3600)

	//返回订单id
	rsp.OrderId = int64(order.Id)

	return nil
}
