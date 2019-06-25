package handler

import (
	"context"
	"fmt"
	example "sss/PutComment/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PutComment(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("更新用户评价订单信息	PUT	api/v1.0/orders/:id/comment   	PutComment")

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
		return nil
	}

	//获取userid
	value_id := bm.Get(key)
	userid, _ := redis.Int(value_id, nil)
	fmt.Println("userid :", userid)

	//获取订单id
	orderid, _ := strconv.Atoi(req.OrderId)

	//获取评论数据
	comment := req.Comment

	//校验数据，判断comment是否为空
	if comment == "" {
		fmt.Println("评论内容为空，不合法！")
		rsp.Errno = utils.RECODE_PARAMERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//根据orderid找到所关联的房源信息
	//查询数据库，订单必须存在，订单状态必须为WAIT_COMMENT待评价状态
	//创建mysql操作句柄
	o := orm.NewOrm()

	//定义承载订单数据的容器
	order := models.OrderHouse{}

	//查询订单表，确认订单表存在
	qs := o.QueryTable("OrderHouse").Filter("Id", orderid)
	if err := qs.Filter("Status", models.ORDER_STATUS_WAIT_COMMENT).One(&order); err != nil {
		fmt.Println("订单数据查询失败,err:", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//关联查询用户表
	if _, err = o.LoadRelated(&order, "User"); err != nil {
		fmt.Println("订单表关联用户表查询失败,err:", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//将用户id与当前用户id进行对比，确认评论人与下单人是同一人
	if order.User.Id != userid {
		fmt.Println("订单用户id与当前用户id不同，数据错误！")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//更新订单表数据信息
	order.Comment = comment
	order.Status = models.ORDER_STATUS_COMPLETE

	//关联查询order订单所关联的House信息
	if _, err = o.LoadRelated(&order, "House"); err != nil {
		fmt.Println("订单表关联房屋表查询失败,err:", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//房屋信息更新
	house := order.House
	//房屋订单成交量加一
	house.Order_count++

	//将order和house表数据更新到mysql数据库
	if _, err = o.Update(&order, "Comment", "Status"); err != nil {
		fmt.Println("更新订单表数据库失败,err:", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	if _, err = o.Update(&house, "Order_count"); err != nil {
		fmt.Println("更新房屋表数据库失败,err:", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//将缓存中的房屋信息删掉，已经不正确了
	house_info_key := "house_info_" + strconv.Itoa(house.Id)
	if err = bm.Delete(house_info_key); err != nil {
		fmt.Println("删除redis数据库中的房屋信息缓存失败,err:", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	return nil
}
