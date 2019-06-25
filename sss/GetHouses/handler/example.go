package handler

import (
	"context"

	"fmt"

	example "sss/GetHouses/proto/example"
	"sss/IhomeWeb/utils"
	"strconv"
	"sss/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetHouses(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取（搜索）房源服务  /api/v1.0/houses  GetHouses")

	//初始化返回数据
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//获取数据 获取url上的参数
	//地区
	aid, _ := strconv.Atoi(req.Aid)
	//起始时间
	sd := req.Sd
	//结束时间
	ed := req.Ed
	//第三栏信息
	sk := req.Sk
	//页码
	page, _ := strconv.Atoi(req.P)

	fmt.Println(aid, sd, ed, sk, page)

	//查询相关地域的房屋信息
	houses := []models.House{}

	//创建orm句柄
	o := orm.NewOrm()
	//设置要查找的表
	qs := o.QueryTable("House")
	//根据地域信息查找房源
	num, err := qs.Filter("Area__Id", aid).All(&houses)
	if err != nil {
		fmt.Println("mysql数据库查找数据失败,err：", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//计算总页数
	total_page := int(num) / models.HOUSE_LIST_PAGE_CAPACITY
	//当前页数
	house_page := 1

	//定义容器，承载房屋数据
	house_list := []interface{}{}

	//循环houses ，补全数据
	for _, house := range houses {
		o.LoadRelated(&house, "User")
		o.LoadRelated(&house, "Area")
		o.LoadRelated(&house, "Facilities")
		o.LoadRelated(&house, "Images")
		house_list = append(house_list, house)
	}

	//返回数据
	rsp.Houses, err = json.Marshal(house_list)
	if err != nil {
		fmt.Println("数据转json失败,err：", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	rsp.TotalPage = int64(total_page)
	rsp.CurrentPage = int64(house_page)

	return nil
}
