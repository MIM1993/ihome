package handler

import (
	"context"
	"fmt"
	example "sss/GetUserOrder/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
	"encoding/json"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetUserOrder(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取房东/租户订单信息服务	GET	api/v1.0/user/orders	GetUserOrder")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//查询用户id
	sessionid := req.Sessionid
	//拼接key
	key := sessionid + "user_id"

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("链接redis错误,err:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//查询用户id
	value_id := bm.Get(key)
	userid, _ := redis.Int(value_id, nil)
	fmt.Println("userid :", userid)

	//获取角色
	role := req.Role
	fmt.Println("Role :", role)

	//链接mysql数据库查询信息，根据角色的不同，查询的方式也不同
	o := orm.NewOrm()
	//创建容器，承载数据
	orders := []models.OrderHouse{}

	//判断角色信息
	if role == "landlord" {
		/*角色为房东，查询自己已经发布的房屋的订单信息*/
		//通过userid获取自己房屋的id
		landLordHouses := []models.House{}

		_, err := o.QueryTable("House").Filter("User__Id", userid).All(&landLordHouses)
		if err != nil {
			fmt.Println("查询房东订单信息错误,err:", err)
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(rsp.Errno)
			return nil
		}

		//获取所有的房屋的房屋id，用于进行订单的查询
		houseids := []int{}
		for _, house := range landLordHouses {
			houseids = append(houseids, house.Id)
		}

		//最后一步在订单表中查询房屋id为以上id的订单信息
		o.QueryTable("OrderHouse").Filter("House__Id__in", houseids).OrderBy("Ctime").All(&orders)

		fmt.Println("查询房东订单信息结束")
	} else {
		/*角色为租客，查询已经下过的订单*/
		//通过用户id直接查订单即可
		_, err := o.QueryTable("OrderHouse").Filter("User__Id", userid).OrderBy("Ctime").All(&orders)
		if err != nil {
			fmt.Println("查询租客订单信息错误,err:", err)
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(rsp.Errno)
			return nil
		}
	}

	//定义返回前端数据容器
	order_list := []interface{}{}

	//循环orders，将数据装进order_list中
	for _, order := range orders {
		//先关联查询
		o.LoadRelated(&order, "User")
		o.LoadRelated(&order, "House")
		//追加进入切片中
		order_list = append(order_list, order)
	}

	//将数据进行json数据转换
	data, err := json.Marshal(order_list)
	if err != nil {
		fmt.Println("json数据转换错误,err:", err)
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//将数据返回web
	rsp.Data = data

	return nil
}
