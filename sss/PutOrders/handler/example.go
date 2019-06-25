package handler

import (
	"context"

	"fmt"

	example "sss/PutOrders/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PutOrders(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("更新房东同意/拒绝订单	PUT	api/v1.0/orders/:id/status	PutOrders")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//sessionid
	sessionid := req.Sessionid

	//拼接key
	key := sessionid + "user_id"

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("链接redis错误,err :", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	//获取userid
	value_id := bm.Get(key)
	userid, _ := redis.Int(value_id, nil)
	fmt.Println("userid :", userid)

	//接收web传来的订单id数据
	orderid, _ := strconv.Atoi(req.Orderid)

	//查找订单，并检验数据，确定订单数据正确，链接mysql
	o := orm.NewOrm()

	//创建容器，承载订单数据
	order := models.OrderHouse{}

	//查询数据库订单表
	qs := o.QueryTable("OrderHouse").Filter("Id", orderid)
	if err = qs.Filter("Status", models.ORDER_STATUS_WAIT_ACCEPT).One(&order); err != nil {
		fmt.Println("查询mysql数据库错误,err :", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	//关联查询House表,同时将房屋拥有者的id与当前用户id进行对比，必须相同才能继续，因为只有房东才有接单或拒单的权利
	if _, err := o.LoadRelated(&order, "House"); err != nil {
		fmt.Println("关联查询mysql数据库错误,err :", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return err
	}

	house := order.House

	//校验该订单的user_id是否是当前用户的user_id
	if house.User.Id != userid {
		fmt.Println("订单用户不匹配,操作无效")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = "订单用户不匹配,操作无效"
		return nil
	}

	//接受web传来的action数据，同时判断是否接单
	action := req.Action

	//判断是接单还是拒单
	if action == "accept" {
		//TODO：业务

		//接单后将订单状态改为待评价
		order.Status = models.ORDER_STATUS_WAIT_COMMENT

		fmt.Println("action = accept ! 接单成功")
	} else if action == "reject" {
		//如果拒单,尝试获得拒单原因，并保存拒单原因
		//更换订单状态，改为拒绝状态
		order.Status = models.ORDER_STATUS_REJECTED

		//添加评论
		reason := req.Action

		order.Comment = reason

		fmt.Println("action = reject!, reason is ", reason)
	}

	//更新数据到数据库
	if _, err := o.Insert(&order); err != nil {
		fmt.Println("更新订单表数据失败,err :", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	return nil
}
