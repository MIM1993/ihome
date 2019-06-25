package handler

import (
	"context"
	"fmt"
	example "sss/PostHousesImage/proto/example"
	"sss/IhomeWeb/utils"
	"path"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
	"strconv"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostHousesImage(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("上传房屋图片服务  PostHousesImage   /api/v1.0/houses/:id/images")

	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//获取文件名后缀
	filename := req.Filename
	fileExt := path.Ext(filename)

	//将文件存入fastDFS
	fileid, err := utils.Uploadbybuf(req.Image, fileExt[1:])
	if err != nil {
		fmt.Println("图片数据存入fastdfs失败")
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	fmt.Println("图片ID：", fileid)

	//获取house对象
	houseid, _ := strconv.Atoi(req.HouseId)
	house := models.House{Id: houseid}

	//链接数据库mysql
	o := orm.NewOrm()
	//获取house全部数据
	err = o.Read(&house)
	if err != nil {
		fmt.Println("数据库查询错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		fmt.Println("err:", err)
		return nil
	}

	//判断house表中首页图片是否为空
	if house.Index_image_url == "" {
		house.Index_image_url = fileid
	}

	//将图片加入房屋图片表中   House不能为空
	house_image := models.HouseImage{House: &house, Url: fileid}

	//将图片加入house表中的Images字段中
	house.Images = append(house.Images, &house_image)

	//插入数据库操作
	_, err = o.Insert(&house_image)
	if err != nil {
		fmt.Println("数据库插入错误==>house_image,err:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//更新house表
	_, err = o.Update(&house)
	if err != nil {
		fmt.Println("数据库更新错误==>house_image")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//返回正确的数据回显给前端
	rsp.Url = fileid

	return nil
}
