package handler

import (
	"context"
	"fmt"
	example "sss/PostHouses/proto/example"
	"sss/IhomeWeb/utils"
	"encoding/json"
	"sss/IhomeWeb/models"
	"strconv"
	"github.com/gomodule/redigo/redis"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostHouses(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("发送（发布）房源信息服务 PostHouses   /api/v1.0/houses")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*获取二进制数据*/
	//定义接收数据容器map[string]interface{}
	Requestmap := make(map[string]interface{})

	//将二进制数据解码到map中, 注意第二个参数必须是指针
	if err := json.Unmarshal(req.Data, &Requestmap); err != nil {
		fmt.Println("接受数据错误,err:", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//遍历打印数据---》无意义
	for key, value := range Requestmap {
		fmt.Println("key:", key, "=== value:", value)
	}

	//实例化房屋结构体，进行数据存储
	house := models.House{}

	//"title":"上奥世纪中心"
	house.Title = Requestmap["title"].(string)

	//"price":"666"
	price, _ := strconv.Atoi(Requestmap["price"].(string))
	house.Price = price * 100

	//"address":"西三旗桥东建材城1号"
	house.Address = Requestmap["address"].(string)

	//"room_count":"2"
	house.Room_count, _ = strconv.Atoi(Requestmap["room_count"].(string))

	//"acreage":"60",
	house.Acreage, _ = strconv.Atoi(Requestmap["acreage"].(string))

	//"unit":"2室1厅",
	house.Unit = Requestmap["unit"].(string)

	//"capacity":"3",
	house.Capacity, _ = strconv.Atoi(Requestmap["capacity"].(string))

	//"beds":"双人床2张",
	house.Beds = Requestmap["beds"].(string)

	//"deposit":"200",
	house.Deposit, _ = strconv.Atoi(Requestmap["deposit"].(string))

	//"min_days":"3",
	house.Min_days, _ = strconv.Atoi(Requestmap["min_days"].(string))

	//"max_days":"0",
	house.Max_days, _ = strconv.Atoi(Requestmap["max_days"].(string))

	//"facility":["1","2","3","7","12","14","16","17","18","21","22"]
	//定义设施表指针切片，承载数据
	facility := []*models.Facility{}
	for _, f_id := range Requestmap["facility"].([]interface{}) {
		//获取设施id
		fid, _ := strconv.Atoi(f_id.(string))
		//定义设施结构体指针 将得到的fid存进去
		fac := &models.Facility{Id: fid}
		//将fac追加进入facility切片中
		facility = append(facility, fac)
	}

	//地域信息  "area_id":"5",
	area_id, _ := strconv.Atoi(Requestmap["area_id"].(string))
	area := &models.Area{Id: area_id}
	house.Area = area

	//链接redis  获取user_id
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("链接redis错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//拼接key
	sessionid := req.Sessionid
	key := sessionid + "user_id"

	//从redis数据库中获取user_id
	value_id := bm.Get(key)
	userid, _ := redis.Int(value_id, nil)
	//打印user_id
	fmt.Println("user_id:", userid)

	//添加user信息到house表
	user := &models.User{Id: userid}
	house.User = user

	//创建mysql数据库句柄
	o := orm.NewOrm()
	houseid, err := o.Insert(&house)
	if err != nil {
		fmt.Println("提交数据到mysql数据库失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//打印houseid
	fmt.Println("房屋id ：", houseid)

	//多对多关系操作 插入到房源与设施信息的多对多表中
	//创建多对多的句柄
	m2m := o.QueryM2M(&house, "Facilities")
	num, err := m2m.Add(facility)
	if err != nil {
		fmt.Println("房屋设施多对多数据插入失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if num == 0 {
		fmt.Println("房屋设施多对多数据插入失败,数据不存在")
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//将房屋号返回前端
	rsp.HouseId = houseid

	return nil
}
