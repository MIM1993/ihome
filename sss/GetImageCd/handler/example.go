package handler

import (
	"context"
	example "sss/GetImageCd/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/afocus/captcha"
	"image/color"
	"time"
	"fmt"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetImageCd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	//生成随机数 与 图片
	//创建一个句柄
	cap := captcha.New()
	//设置字体
	if err := cap.SetFont("comic.ttf"); err != nil {
		//抛出异常
		panic(err.Error())
	}

	//设置突破大小
	cap.SetSize(90, 41)
	//设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	//设置前景色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255}, color.RGBA{0, 0, 0, 200})
	//设置背景色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	//随即生存 图片 与 验证码
	img, str := cap.Create(4, captcha.NUM)

	fmt.Println("验证码:", str)

	//获取 uuid
	uuid := req.Uuid

	//链接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//存入数据 key是uuid
	bm.Put(uuid, str, time.Second*300)

	//将图片拆分 赋值到proto
	//解引用
	a := *img
	b := *(a.RGBA)

	//pix
	for _, value := range b.Pix {
		rsp.Pix = append(rsp.Pix, int32(value))
	}

	//stride
	rsp.Stride = int64(b.Stride)

	//point
	rsp.Min = &example.ResponsePoint{X: int64(b.Rect.Min.X), Y: int64(b.Rect.Min.Y)}
	rsp.Max = &example.ResponsePoint{X: int64(b.Rect.Max.X), Y: int64(b.Rect.Max.Y)}

	return nil
}
