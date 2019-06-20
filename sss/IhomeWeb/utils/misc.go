package utils

import (
	"encoding/json"
	"github.com/astaxie/beego/cache"
	"fmt"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"crypto/md5"
	"encoding/hex"
	"github.com/weilaihui/fdfs_client"
)

/* 将url加上 http://IP:PROT/  前缀 */
//http:// + 127.0.0.1 + ：+ 8080 + 请求
func AddDomain2Url(url string) (domain_url string) {
	domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url

	return domain_url
}

//链接redis函数
func RedisOpen(server_name, redis_addr, redis_port, redis_dbnum string) (bm cache.Cache, err error) {
	redis_config_map := map[string]string{
		"key":   server_name,
		"conn":  redis_addr + ":" + redis_port,
		"dbNum": redis_dbnum,
	}
	//将 配置信息的map 转化为json
	redis_config, _ := json.Marshal(redis_config_map)

	//连接redis
	bm, err = cache.NewCache("redis", string(redis_config))
	if err != nil {
		fmt.Println("连接 redis错误", err)
		return nil, err
	}
	return bm, nil
}

//加密函数
func Getmd5string(s string) string {
	m := md5.New()
	e := m.Sum([]byte(s))
	return hex.EncodeToString(e)
}

//fastdfs上传函数 参数:(二进制 文件,后缀名); 返回值:(fileid,err)
func Uploadbybuf(file []byte, Extname string) (filed string, err error) {
	//读取配置文件 创建fastdfs客户端句柄
	//fdfsclient, err := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf ")
	fdfsclient, err := fdfs_client.NewFdfsClient("./conf/client.conf")

	if err != nil {
		return "", err
	}

	//上传文件
	rsp, err := fdfsclient.UploadByBuffer(file, Extname)
	if err != nil {
		fmt.Println("上传错误",err)
		return "", err
	}
	fmt.Println("GroupName", rsp.GroupName, "FileId:", rsp.RemoteFileId)

	//返回filed
	return rsp.RemoteFileId, nil

}
